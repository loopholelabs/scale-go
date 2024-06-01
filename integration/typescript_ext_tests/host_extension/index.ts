// Code generated by scale-extension 0.4.7, DO NOT EDIT.
// output: local-inttest-latest-host

/* eslint no-bitwise: off */

import { Extension as ExtensionInterface, ModuleMemory, Resizer } from "@loopholelabs/scale-extension-interfaces";
import { Decoder, Encoder, Kind } from "@loopholelabs/polyglot";
import * as types from "./types";

export * from "./types";

const hash = "b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7";

// Write an error to the scale function guest buffer.
function hostError(mem: ModuleMemory, resize: Resizer, err: Error) {
  const enc = new Encoder();
  enc.error(err);
  const ptr = resize("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize", enc.bytes.length);

  mem.Write(ptr, enc.bytes);
}

class hostExt {
  functions: Map<string, InstallableFunc>;
  host: Host;

  constructor(fns: Map<string, InstallableFunc>, h: Host) {
    this.functions = fns;
    this.host = h;
  }

  Init(): Map<string, InstallableFunc> {
    return this.functions;
  }

  Reset() {
    // Reset any instances that have been created.
    this.host.instances_Example = new Map<number, Example();
  }
}

export function New(impl: Interface): ExtensionInterface {
  let hostWrapper = new Host(impl);

  let fns = new Map<string, InstallableFunc>();

  // Add global functions to the runtime

  fns.set("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_New", hostWrapper.host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_New.bind(hostWrapper));

  fns.set("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_World", hostWrapper.host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_World.bind(hostWrapper));

  hostWrapper.instances_Example = new Map<number, Example>();

  fns.set("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Example_Hello", hostWrapper.host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Example_Hello.bind(hostWrapper));

  return new hostExt(fns, hostWrapper);
}

class Host {
  impl: Interface

  gid_Example: bigint = 0n;
  instances_Example: Map<bigint, Example> = new Map<bigint, Example>();

  constructor(i: Interface) {
    this.impl = i;
  }

  // Global functions...

  host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_New(mem: ModuleMemory, resize: Resizer, params: number[]) {
    const d = mem.Read(params[1], params[2]);
    const c = types.Stringval.decode(new Decoder(d));
    const r = this.impl.New(c);
    const id = this.gid_Example++;
    this.instances_Example.set(id, r);
    params[0] = id;
    return;
  }

  host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_World(mem: ModuleMemory, resize: Resizer, params: number[]) {
    const d = mem.Read(params[1], params[2]);
    const c = types.Stringval.decode(new Decoder(d));
    const r = this.impl.World(c);
    const enc = new Encoder();
    r.encode(enc);
    const ptr = resize("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize", enc.bytes.length);
    mem.Write(ptr, enc.bytes);
    return;
  }

  // Instance functions...

  host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Example_Hello(mem: ModuleMemory, resize: Resizer, params: number[]): bigint {
    const d = mem.Read(params[1], params[2]);
    const c = types.Stringval.decode(new Decoder(d));
    // Do lookup...
    const inst = this.instances_Example.get(params[0]);
    const r = inst.Hello(c);
    const enc = new Encoder();
    r.encode(enc);
    const ptr = resize("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize", enc.bytes.length);
    mem.Write(ptr, enc.bytes);
    return;
  }

}

//// //// //// //// //// //// //// //// ////

// Interface to the extension impl. This is what the implementor should create

export interface Interface {
  New(params: Stringval): Example;

  World(params: Stringval): Stringval;

}

export interface Example {
  Hello(params: Stringval): Stringval;

}

