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

import { Signature } from "@loopholelabs/scale-signature-interfaces";
import { Func } from "./function";
import { Scale } from "./scale";
import {Instance} from "./instance";

export class ModuleInstance<T extends Signature> {
    private function: Func<T>;

    private instantiatedModule: WebAssembly.Instance;

    private run: undefined | Function;
    private resize: undefined | Function;

    private memory: undefined | WebAssembly.Memory;

    private instance: Instance<T>;

    public signature: T | undefined;

    private nextInstance: undefined | ModuleInstance<T>;

    constructor(r: Scale<T>, f: Func<T>, i: Instance<T>) {
        this.function = f;
        this.instance = i;
    }

    init(i: Instance<T>) {
        this.wasmInstance = this.runtime.InstantiateModule(this.func.id, this, i);

        this.run = this.wasmInstance.exports.run as Function;
        this.resize = this.wasmInstance.exports.resize as Function;

        if (this.run === undefined || this.resize === undefined) {
            throw new Error("failed to find run or resize implementations");
        }

        this.memory = this.wasmInstance.exports.memory as WebAssembly.Memory;
    }

}