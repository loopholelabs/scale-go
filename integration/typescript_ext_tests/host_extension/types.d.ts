// Code generated by scale-signature 0.4.5, DO NOT EDIT.
// output: local-inttest-latest-host

import { Encoder, Decoder, Kind } from "@loopholelabs/polyglot"

export declare class Stringval {
  value: string;

  /**
  * @throws {Error}
  */
  constructor (decoder?: Decoder);

  /**
  * @throws {Error}
  */
  encode (encoder: Encoder);

  /**
  * @throws {Error}
  */
  static decode (decoder: Decoder): Stringval | undefined;

  /**
  * @throws {Error}
  */
  static encode_undefined (encoder: Encoder);
}
