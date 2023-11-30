// Code generated by scale-extension 0.4.5, DO NOT EDIT.
// output: local-inttest-latest-host

"use strict";
var __create = Object.create;
var __defProp = Object.defineProperty;
var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
var __getOwnPropNames = Object.getOwnPropertyNames;
var __getProtoOf = Object.getPrototypeOf;
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
var __toESM = (mod, isNodeMode, target) => (target = mod != null ? __create(__getProtoOf(mod)) : {}, __copyProps(
  // If the importer is in node compatibility mode or this is not an ESM
  // file that has been converted to a CommonJS file using a Babel-
  // compatible transform (i.e. "__esModule" has not been set), then set
  // "default" to the CommonJS "module.exports" for node compatibility.
  isNodeMode || !mod || !mod.__esModule ? __defProp(target, "default", { value: mod, enumerable: true }) : target,
  mod
));
var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);
var stdin_exports = {};
__export(stdin_exports, {
  New: () => New
});
module.exports = __toCommonJS(stdin_exports);
var import_polyglot = require("@loopholelabs/polyglot");
var types = __toESM(require("./types"));
__reExport(stdin_exports, require("./types"), module.exports);
const hash = "b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7";
function hostError(mem, resize, err) {
  const enc = new import_polyglot.Encoder();
  enc.error(err);
  const ptr = resize("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize", enc.bytes.length);
  mem.Write(ptr, enc.bytes);
}
class hostExt {
  constructor(fns, h) {
    this.functions = fns;
    this.host = h;
  }
  Init() {
    return this.functions;
  }
  Reset() {
    this.host.instances_Example = /* @__PURE__ */ new Map() < number, Example();
  }
}
function New(impl) {
  let hostWrapper = new Host(impl);
  let fns = /* @__PURE__ */ new Map();
  fns.set("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_New", hostWrapper.host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_New.bind(hostWrapper));
  fns.set("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_World", hostWrapper.host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_World.bind(hostWrapper));
  hostWrapper.instances_Example = /* @__PURE__ */ new Map();
  fns.set("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Example_Hello", hostWrapper.host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Example_Hello.bind(hostWrapper));
  return new hostExt(fns, hostWrapper);
}
class Host {
  constructor(i) {
    this.gid_Example = 0n;
    this.instances_Example = /* @__PURE__ */ new Map();
    this.impl = i;
  }
  // Global functions...
  host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_New(mem, resize, params) {
    const d = mem.Read(params[1], params[2]);
    const c = types.Stringval.decode(new import_polyglot.Decoder(d));
    const r = this.impl.New(c);
    const id = this.gid_Example++;
    this.instances_Example.set(id, r);
    params[0] = id;
    return;
  }
  host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_World(mem, resize, params) {
    const d = mem.Read(params[1], params[2]);
    const c = types.Stringval.decode(new import_polyglot.Decoder(d));
    const r = this.impl.World(c);
    const enc = new import_polyglot.Encoder();
    r.encode(enc);
    const ptr = resize("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize", enc.bytes.length);
    mem.Write(ptr, enc.bytes);
    return;
  }
  // Instance functions...
  host_ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Example_Hello(mem, resize, params) {
    const d = mem.Read(params[1], params[2]);
    const c = types.Stringval.decode(new import_polyglot.Decoder(d));
    const inst = this.instances_Example.get(params[0]);
    const r = inst.Hello(c);
    const enc = new import_polyglot.Encoder();
    r.encode(enc);
    const ptr = resize("ext_b30af2dd8561988edd7b281ad5c1b84487072727a8ad0e490a87be0a66b037d7_Resize", enc.bytes.length);
    mem.Write(ptr, enc.bytes);
    return;
  }
}
//# sourceMappingURL=index.js.map