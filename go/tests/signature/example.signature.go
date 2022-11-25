// Code generated by scale-signature-go v0.0.1, DO NOT EDIT.
// source: signature/example/signature.proto

package example

import (
	"errors"
	"github.com/loopholelabs/polyglot-go"
)

var (
	NilDecode = errors.New("cannot decode into a nil root struct")
)

type ExampleContext struct {
	Data string
}

func NewExampleContext() *ExampleContext {
	return &ExampleContext{}
}

func (x *ExampleContext) error(b *polyglot.Buffer, err error) {
	polyglot.Encoder(b).Error(err)
}

func (x *ExampleContext) internalEncode(b *polyglot.Buffer) {
	if x == nil {
		polyglot.Encoder(b).Nil()
	} else {
		polyglot.Encoder(b).String(x.Data)
	}
}

func (x *ExampleContext) internalDecode(b []byte) error {
	if x == nil {
		return NilDecode
	}
	d := polyglot.GetDecoder(b)
	defer d.Return()
	return x.decode(d)
}

func (x *ExampleContext) decode(d *polyglot.Decoder) error {
	if d.Nil() {
		return nil
	}

	err, _ := d.Error()
	if err != nil {
		return err
	}
	x.Data, err = d.String()
	if err != nil {
		return err
	}
	return nil
}
