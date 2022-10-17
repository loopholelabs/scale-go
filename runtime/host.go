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

package runtime

import (
	"context"
	"github.com/loopholelabs/scale-go/utils"
	"github.com/tetratelabs/wazero/api"
	"strings"
)

func (r *Runtime) fd_write(int32, int32, int32, int32) int32 {
	return 0
}

func (r *Runtime) next(ctx context.Context, module api.Module, offset uint32, length uint32) uint64 {
	i := r.instances[strings.Split(module.Name(), ".")[0]]
	if i == nil {
		return 0
	}

	m := i.modules[module]
	if m == nil {
		return 0
	}

	buf, ok := m.module.Memory().Read(ctx, offset, length)
	if !ok {
		return 0
	}

	err := i.Context().Read(buf)
	if err != nil {
		return 0
	}

	if m.next == nil {
		i.ctx = i.next(i.Context())
	} else {
		err = m.next.Run(ctx)
		if err != nil {
			return 0
		}
	}

	ctxBuffer := i.Context().Write()
	ctxBufferLength := uint64(len(ctxBuffer))
	writeBuffer, err := m.malloc.Call(ctx, ctxBufferLength)
	if err != nil {
		return 0
	}

	if !module.Memory().Write(ctx, uint32(writeBuffer[0]), ctxBuffer) {
		return 0
	}

	return utils.PackUint32(uint32(writeBuffer[0]), uint32(ctxBufferLength))
}

//func (r *Runtime) debug(ctx context.Context, module api.Module, offset uint32, length uint32) {
//	buf, ok := module.Memory().Read(ctx, offset, length)
//	if !ok {
//		panic("failed to read memory")
//	}
//
//	fmt.Println(string(buf))
//}