/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

// Package scalefunc implements the ScaleFunc type, as well as any helper functions
// for interacting with ScaleFunc types
package scalefunc

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/loopholelabs/polyglot"
	signatureSchema "github.com/loopholelabs/scale/signature"
	"os"
	"regexp"
)

var (
	VersionErr  = errors.New("unknown or invalid version")
	LanguageErr = errors.New("unknown or invalid language")
	ChecksumErr = errors.New("error while verifying checksum")
)

var (
	InvalidStringRegex = regexp.MustCompile(`[^A-Za-z0-9-.]`)
)

// Version is the Version of the ScaleFunc definition
type Version string

const (
	// V1Alpha is the V1 Alpha definition of a ScaleFunc
	V1Alpha Version = "v1alpha"
)

// Language is the Language the Scale Function's Source Language
type Language string

const (
	// Go is the Golang Source Language for Scale Functions
	Go Language = "go"

	// Rust is the Rust Source Language for Scale Functions
	Rust Language = "rust"

	// TypeScript is the TypeScript Source Language for Scale Functions
	TypeScript Language = "ts"
)

var (
	// AcceptedVersions is an array of acceptable Versions
	AcceptedVersions = []Version{V1Alpha}

	// AcceptedLanguages is an array of acceptable Languages
	AcceptedLanguages = []Language{Go, Rust, TypeScript}
)

// ValidString returns true if the string is valid for use with Scale Functions
func ValidString(str string) bool {
	return !InvalidStringRegex.MatchString(str)
}

// Dependency is a dependency of a compiled Scale Function
type Dependency struct {
	Name     string            `json:"name" yaml:"name"`
	Version  string            `json:"version" yaml:"version"`
	Metadata map[string]string `json:"metadata" yaml:"metadata"`
}

// ScaleFunc is the type used to define the requirements of a
// scale function for a Scale Runtime
type ScaleFunc struct {
	Version         Version                `json:"version" yaml:"version"`
	Name            string                 `json:"name" yaml:"name"`
	Tag             string                 `json:"tag" yaml:"tag"`
	Signature       string                 `json:"signature" yaml:"signature"`
	SignatureSchema signatureSchema.Schema `json:"signature_schema" yaml:"signature_schema"`
	Language        Language               `json:"language" yaml:"language"`
	Dependencies    []Dependency           `json:"dependencies" yaml:"dependencies"`
	Function        []byte                 `json:"function" yaml:"function"`
	Size            uint32                 `json:"size" yaml:"size"`
	Checksum        string                 `json:"checksum" yaml:"checksum"`

	// env is used by the host at runtime to specify environment variables for the function
	env map[string]string
}

// Encode encodes the ScaleFunc into a byte array
func (s *ScaleFunc) Encode() []byte {
	b := polyglot.GetBuffer()
	defer polyglot.PutBuffer(b)
	e := polyglot.Encoder(b)
	e.String(string(s.Version))
	e.String(s.Name)
	e.String(s.Tag)
	e.String(s.Signature)

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(s.SignatureSchema, f.Body())
	e.Bytes(f.Bytes())

	e.String(string(s.Language))

	e.Slice(uint32(len(s.Dependencies)), polyglot.AnyKind)
	for _, d := range s.Dependencies {
		e.String(d.Name)
		e.String(d.Version)
		e.Map(uint32(len(d.Metadata)), polyglot.StringKind, polyglot.StringKind)
		for k, v := range d.Metadata {
			e.String(k)
			e.String(v)
		}
	}

	e.Bytes(s.Function)

	size := uint32(len(b.Bytes()))
	hash := sha256.New()
	hash.Write(b.Bytes())
	checksum := hex.EncodeToString(hash.Sum(nil))

	e.Uint32(size)
	e.String(checksum)

	return b.Bytes()
}

// Decode decodes the ScaleFunc from a byte array
func (s *ScaleFunc) Decode(data []byte) error {
	d := polyglot.GetDecoder(data)
	defer d.Return()

	version, err := d.String()
	if err != nil {
		return err
	}
	s.Version = Version(version)

	invalid := true
	for _, v := range AcceptedVersions {
		if s.Version == v {
			invalid = false
			break
		}
	}
	if invalid {
		return VersionErr
	}

	s.Name, err = d.String()
	if err != nil {
		return err
	}

	s.Tag, err = d.String()
	if err != nil {
		return err
	}

	s.Signature, err = d.String()
	if err != nil {
		return err
	}

	signatureSchemaBytes, err := d.Bytes(nil)
	if err != nil {
		return err
	}

	err = s.SignatureSchema.Decode(signatureSchemaBytes)
	if err != nil {
		return err
	}

	language, err := d.String()
	if err != nil {
		return err
	}
	s.Language = Language(language)

	invalid = true
	for _, l := range AcceptedLanguages {
		if l == s.Language {
			invalid = false
			break
		}
	}
	if invalid {
		return LanguageErr
	}

	dependenciesSize, err := d.Slice(polyglot.AnyKind)
	if err != nil {
		return err
	}
	s.Dependencies = make([]Dependency, dependenciesSize)
	for i := uint32(0); i < dependenciesSize; i++ {
		s.Dependencies[i].Name, err = d.String()
		if err != nil {
			return err
		}
		s.Dependencies[i].Version, err = d.String()
		if err != nil {
			return err
		}

		metadataSize, err := d.Map(polyglot.StringKind, polyglot.StringKind)
		if err != nil {
			return err
		}
		s.Dependencies[i].Metadata = make(map[string]string, metadataSize)
		for j := uint32(0); j < metadataSize; j++ {
			key, err := d.String()
			if err != nil {
				return err
			}
			value, err := d.String()
			if err != nil {
				return err
			}
			s.Dependencies[i].Metadata[key] = value
		}
	}

	s.Function, err = d.Bytes(nil)
	if err != nil {
		return err
	}

	s.Size, err = d.Uint32()
	if err != nil {
		return err
	}

	s.Checksum, err = d.String()
	if err != nil {
		return err
	}

	hash := sha256.New()
	hash.Write(data[:s.Size])

	if hex.EncodeToString(hash.Sum(nil)) != s.Checksum {
		return ChecksumErr
	}

	return nil
}

// SetEnv sets the environment variables for the ScaleFunc
func (s *ScaleFunc) SetEnv(env map[string]string) {
	s.env = env
}

// GetEnv returns the environment variables for the ScaleFunc
func (s *ScaleFunc) GetEnv() map[string]string {
	return s.env
}

// Read opens a file at the given path and returns a *ScaleFile
func Read(path string) (*ScaleFunc, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	scaleFunc := new(ScaleFunc)
	return scaleFunc, scaleFunc.Decode(data)
}

// Write opens a file at the given path and writes the given scalefile to it
func Write(path string, scaleFunc *ScaleFunc) error {
	data := scaleFunc.Encode()
	return os.WriteFile(path, data, 0644)
}
