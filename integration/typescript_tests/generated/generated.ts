// Code generated by scale-signature v0.1.0, DO NOT EDIT.
// output: generated

import { Encoder, Decoder, Kind } from "@loopholelabs/polyglot"

export class Context {
  a: number;
  b: number;
  c: number;

  /**
  * @throws {Error}
  */
  constructor (decoder?: Decoder) {
    if (decoder) {
      this.a = decoder.int32();
      this.b = decoder.int32();
      this.c = decoder.int32();
    } else {
      this.a = 0;
      this.b = 0;
      this.c = 0;
    }
  }

  /**
  * @throws {Error}
  */
  encode (encoder: Encoder) {
    encoder.int32(this.a);
    encoder.int32(this.b);
    encoder.int32(this.c);
  }

  /**
  * @throws {Error}
  */
  static decode (decoder: Decoder): Context | undefined {
    if (decoder.null()) {
      return undefined
    }
    return new Context(decoder);
  }

  /**
  * @throws {Error}
  */
  static encode_undefined (encoder: Encoder) {
    encoder.null();
  }
}

