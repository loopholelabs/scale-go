// Code generated by scale-extension 0.4.7, DO NOT EDIT.
// output: local_inttest_latest_guest

/* eslint no-bitwise: off */

import { Decoder, Encoder } from "@loopholelabs/polyglot";

import * as types from "./types";

let writeBuffer = new Uint8Array().buffer;
let readBuffer = new Uint8Array().buffer;

function ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize(len: number): number {
  readBuffer = new Uint8Array(len).buffer;
  const ptr = (global as any).scale_address_of(readBuffer);
  return ptr;
}

// Register it...
function ext_init() {
  let id = BigInt(0xb30af2dd);
  // TODO: This ID needs to come from config etc
  (global as any).registerResize(id, ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize);
}

// Define any interfaces we need here...
// Also define structs we can use to hold instanceId

// Define concrete types with a hidden instanceId

class _Example {
  instanceId: number;

  constructor(id: number) {
    this.instanceId = id;
  }

  Hello(params: types.Stringval): types.Stringval {
    let e = new Encoder();
    params.encode(e);
    writeBuffer = e.bytes.buffer;
    let callID = BigInt(0x611e5c44);
    let ev = (global as any).scale_ext_mux([callID, this.instanceId, (global as any).scale_address_of(writeBuffer), writeBuffer.byteLength]);
    // Decode it and return...
    let dec = new Decoder(new Uint8Array(readBuffer));
    return new types.Stringval(dec);
  }

}

// Define any global functions here...

export function New(params: types.Stringval): types.Example {
  // First encode the params...

  // Make sure this is registered for incoming resize calls.
  ext_init();

  readBuffer = new Uint8Array(0).buffer;

  let e = new Encoder();
  params.encode(e);
  writeBuffer = e.bytes.buffer;

  let callID = BigInt(0xef500a0a);
  let ev = (global as any).scale_ext_mux([callID, 0, (global as any).scale_address_of(writeBuffer), writeBuffer.byteLength]);
  // Handle error from host... (stuff in readBuffer)
  if (readBuffer.byteLength>0) {
    let dec = new Decoder(new Uint8Array(readBuffer));
    throw dec.error();
  }

  return new _Example(ev);

}

export function World(params: types.Stringval): types.Stringval {
  // First encode the params...

  // Make sure this is registered for incoming resize calls.
  ext_init();

  readBuffer = new Uint8Array(0).buffer;

  let e = new Encoder();
  params.encode(e);
  writeBuffer = e.bytes.buffer;

  let callID = BigInt(0xd034776c);
  let ev = (global as any).scale_ext_mux([callID, 0, (global as any).scale_address_of(writeBuffer), writeBuffer.byteLength]);
  // Decode it and return...
  let dec = new Decoder(new Uint8Array(readBuffer));
  return new types.Stringval(dec);

}

