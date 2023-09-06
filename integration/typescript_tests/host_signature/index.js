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
  New: () => New,
  Signature: () => Signature
});
module.exports = __toCommonJS(stdin_exports);
var import_polyglot = require("@loopholelabs/polyglot");
__reExport(stdin_exports, require("./types"), module.exports);
const hash = "3a592aa345d412faa2e6285ee048ca2ab5aa64b0caa2f9ca67b2c1e0792101e5";
function New() {
  return new Signature();
}
class Signature {
  constructor() {
    this.context = new ModelWithAllFieldTypes();
  }
  // Read reads the context from the given Uint8Array and returns an error if one occurred
  //
  // This method is meant to be used by the Scale Runtime to deserialize the Signature
  Read(b) {
    const dec = new import_polyglot.Decoder(b);
    try {
      Object.assign(this.context, ModelWithAllFieldTypes.decode(dec));
    } catch (err) {
      return err;
    }
    return void 0;
  }
  // Write writes the signature into a Uint8Array and returns it
  //
  // This method is meant to be used by the Scale Runtime to serialize the Signature
  Write() {
    return this.context.encode(new Uint8Array());
  }
  // Error writes the signature into a Uint8Array and returns it
  //
  // This method is meant to be used by the Scale Runtime to return an error
  Error(err) {
    const enc = new import_polyglot.Encoder();
    enc.error(err);
    return enc.bytes;
  }
  // Hash returns the hash of the signature
  //
  // This method is meant to be used by the Scale Runtime to validate Signature and Function compatibility
  Hash() {
    return hash;
  }
}
//# sourceMappingURL=index.js.map