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

// Package scalefunc implements the Schema type, as well as any helper functions
// for interacting with Schema types
package scalefunc

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"os"
	"regexp"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/loopholelabs/polyglot"

	signatureSchema "github.com/loopholelabs/scale/signature"
)

var (
	ErrVersion  = errors.New("unknown or invalid version")
	ErrLanguage = errors.New("unknown or invalid language")
	ErrHash     = errors.New("error while verifying hash")
)

var (
	InvalidStringRegex = regexp.MustCompile(`[^A-Za-z0-9-.]`)
)

// Version is the Version of the Schema definition
type Version string

const (
	// V1Alpha is the V1 Alpha definition of a Schema
	V1Alpha Version = "v1alpha"

	// V1Beta is the V1 Beta definition of a Schema
	V1Beta Version = "v1beta"
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
	// AcceptedLanguages is an array of acceptable Languages
	AcceptedLanguages = []Language{Go, Rust, TypeScript}
)

// ValidString returns true if the string is valid for use with Scale Functions
func ValidString(str string) bool {
	return !InvalidStringRegex.MatchString(str)
}

// V1AlphaDependency is a dependency of a compiled Scale Function
type V1AlphaDependency struct {
	Name     string            `json:"name" yaml:"name"`
	Version  string            `json:"version" yaml:"version"`
	Metadata map[string]string `json:"metadata" yaml:"metadata"`
}

// V1AlphaSchema is the type used to define the requirements of a
// scale function for a Scale Runtime
type V1AlphaSchema struct {
	Name            string                  `json:"name" yaml:"name"`
	Tag             string                  `json:"tag" yaml:"tag"`
	SignatureName   string                  `json:"signature_name" yaml:"signature_name"`
	SignatureSchema *signatureSchema.Schema `json:"signature_schema" yaml:"signature_schema"`
	SignatureHash   string                  `json:"signature_hash" yaml:"signature_hash"`
	Language        Language                `json:"language" yaml:"language"`
	Stateless       bool                    `json:"stateless" yaml:"stateless"`
	Dependencies    []V1AlphaDependency     `json:"dependencies" yaml:"dependencies"`
	Function        []byte                  `json:"function" yaml:"function"`
	Size            uint32                  `json:"size" yaml:"size"`
	Hash            string                  `json:"hash" yaml:"hash"`
}

// Encode encodes the Schema into a byte array
//
// Deprecated: Use V1BetaSchema instead
func (s *V1AlphaSchema) Encode() []byte {
	b := polyglot.GetBuffer()
	defer polyglot.PutBuffer(b)
	e := polyglot.Encoder(b)
	e.String(string(V1Alpha))
	e.String(s.Name)
	e.String(s.Tag)
	e.String(s.SignatureName)

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(s.SignatureSchema, f.Body())
	e.Bytes(f.Bytes())

	e.String(s.SignatureHash)

	e.String(string(s.Language))

	e.Bool(s.Stateless)

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

	e.Uint32(size)
	e.String(hex.EncodeToString(hash.Sum(nil)))

	return b.Bytes()
}

// Decode decodes the Schema from a byte array
func (s *V1AlphaSchema) Decode(data []byte) error {
	d := polyglot.GetDecoder(data)
	defer d.Return()

	version, err := d.String()
	if err != nil {
		return err
	}

	switch Version(version) {
	case V1Alpha:
	default:
		return ErrVersion
	}

	s.Name, err = d.String()
	if err != nil {
		return err
	}

	s.Tag, err = d.String()
	if err != nil {
		return err
	}

	s.SignatureName, err = d.String()
	if err != nil {
		return err
	}

	signatureSchemaBytes, err := d.Bytes(nil)
	if err != nil {
		return err
	}

	s.SignatureSchema = new(signatureSchema.Schema)
	err = s.SignatureSchema.Decode(signatureSchemaBytes)
	if err != nil {
		return err
	}

	s.SignatureHash, err = d.String()
	if err != nil {
		return err
	}

	language, err := d.String()
	if err != nil {
		return err
	}
	s.Language = Language(language)

	invalid := true
	for _, l := range AcceptedLanguages {
		if l == s.Language {
			invalid = false
			break
		}
	}
	if invalid {
		return ErrLanguage
	}

	s.Stateless, err = d.Bool()
	if err != nil {
		s.Stateless = false
	}

	dependenciesSize, err := d.Slice(polyglot.AnyKind)
	if err != nil {
		return err
	}
	s.Dependencies = make([]V1AlphaDependency, dependenciesSize)
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

	s.Hash, err = d.String()
	if err != nil {
		return err
	}

	hash := sha256.New()
	hash.Write(data[:s.Size])

	if hex.EncodeToString(hash.Sum(nil)) != s.Hash {
		return ErrHash
	}

	return nil
}

// GetHash returns the hash of the Schema
func (s *V1AlphaSchema) GetHash() []byte {
	b := polyglot.GetBuffer()
	defer polyglot.PutBuffer(b)
	e := polyglot.Encoder(b)
	e.String(string(V1Alpha))
	e.String(s.Name)
	e.String(s.Tag)
	e.String(s.SignatureName)

	f := hclwrite.NewEmptyFile()
	gohcl.EncodeIntoBody(s.SignatureSchema, f.Body())
	e.Bytes(f.Bytes())

	e.String(s.SignatureHash)

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

	hash := sha256.New()
	hash.Write(b.Bytes())

	return hash.Sum(nil)
}

// ReadV1Alpha opens a file at the given path and returns a *WriteV1Alpha
//
// Deprecated: Use V1BetaSchema instead
func ReadV1Alpha(path string) (*V1AlphaSchema, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	scaleFunc := new(V1AlphaSchema)
	return scaleFunc, scaleFunc.Decode(data)
}

// WriteV1Alpha opens a file at the given path and writes the given WriteV1Alpha to it
//
// Deprecated: Use V1BetaSchema instead
func WriteV1Alpha(path string, scaleFunc *V1AlphaSchema) error {
	data := scaleFunc.Encode()
	return os.WriteFile(path, data, 0644)
}

func Read(path string) (*V1AlphaSchema, error) {
	return ReadV1Alpha(path)
}

func Write(path string, scaleFunc *V1AlphaSchema) error {
	return WriteV1Alpha(path, scaleFunc)
}
