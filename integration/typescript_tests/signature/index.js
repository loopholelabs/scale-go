// Code generated by scale-signature 0.4.2, DO NOT EDIT.
// output: local-example-latest-guest

"use strict";
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __hasOwnProp = Object.prototype.hasOwnProperty;
var __export = (target, all) => {
  for (var name in all)
    __defProp(target, name, { get: all[name], enumerable: true });
};
var __copyProps = (to, from, except, desc) => {
  if (from && typeof from === "object" || typeof from === "function") {
    for (let key of __getOwnPropNames(from))
      if (!__hasOwnProp.call(to, key) && key !== except)
        __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
  }
  return to;
};
var __reExport = (target, mod, secondTarget) => (__copyProps(target, mod, "default"), secondTarget && __copyProps(secondTarget, mod, "default"));
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);
var stdin_exports = {};
__export(stdin_exports, {
  Error: () => Error2,
  Hash: () => Hash,
  Next: () => Next,
  Read: () => Read,
  Resize: () => Resize,
  Write: () => Write
});
module.exports = __toCommonJS(stdin_exports);
var import_scale_signature_interfaces = require("@loopholelabs/scale-signature-interfaces");
var import_polyglot = require("@loopholelabs/polyglot");
__reExport(stdin_exports, require("./types"), module.exports);
var import_types = require("./types");
global.WRITE_BUFFER = new Uint8Array().buffer;
global.READ_BUFFER = new Uint8Array().buffer;
const hash = "3a592aa345d412faa2e6285ee048ca2ab5aa64b0caa2f9ca67b2c1e0792101e5";
function Write(ctx) {
  const enc = new import_polyglot.Encoder();
  if (typeof ctx === "undefined") {
    enc.null();
  } else {
    ctx.encode(enc);
  }
  const len = enc.bytes.buffer.byteLength;
  global.WRITE_BUFFER = enc.bytes.buffer;
  const addrof = global[import_scale_signature_interfaces.TYPESCRIPT_ADDRESS_OF];
  const ptr = addrof(global.WRITE_BUFFER);
  return [ptr, len];
}
function Read() {
  const dec = new import_polyglot.Decoder(new Uint8Array(global.READ_BUFFER));
  return import_types.ModelWithAllFieldTypes.decode(dec);
}
function Error2(err) {
  const enc = new import_polyglot.Encoder();
  enc.error(err);
  const len = enc.bytes.buffer.byteLength;
  global.WRITE_BUFFER = enc.bytes.buffer;
  const addrof = global[import_scale_signature_interfaces.TYPESCRIPT_ADDRESS_OF];
  const ptr = addrof(global.WRITE_BUFFER);
  return [ptr, len];
}
function Resize(size) {
  global.READ_BUFFER = new Uint8Array(size).buffer;
  const addrof = global[import_scale_signature_interfaces.TYPESCRIPT_ADDRESS_OF];
  return addrof(global.READ_BUFFER);
}
function Hash() {
  const enc = new import_polyglot.Encoder();
  enc.string(hash);
  const len = enc.bytes.buffer.byteLength;
  global.WRITE_BUFFER = enc.bytes.buffer;
  const addrof = global[import_scale_signature_interfaces.TYPESCRIPT_ADDRESS_OF];
  const ptr = addrof(global.WRITE_BUFFER);
  return [ptr, len];
}
function Next(ctx) {
  const [ptr, len] = Write(ctx);
  const next = global[import_scale_signature_interfaces.TYPESCRIPT_NEXT];
  next([ptr, len]);
  return Read();
}
//# sourceMappingURL=index.js.map