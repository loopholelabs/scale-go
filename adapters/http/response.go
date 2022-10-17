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

package http

import (
	"fmt"
	"github.com/loopholelabs/scale-go/runtime"
	"github.com/loopholelabs/scale-go/runtime/generated"
	"net/http"
	"strings"
)

// FromResponse serializes the *ResponseWriter object into a runtime.Context
func FromResponse(ctx *runtime.Context, w *ResponseWriter) {
	ctx.Context.Response.StatusCode = int32(w.statusCode)
	for k, v := range w.headers {
		ctx.Context.Response.Headers[k] = &generated.StringList{
			Value: v,
		}
	}
	ctx.Context.Response.Body = w.buffer.Bytes()
}

// ToResponse deserializes the runtime.Context object into the http.ResponseWriter
func ToResponse(ctx *runtime.Context, w http.ResponseWriter) error {
	for k, v := range ctx.Context.Response.Headers {
		w.Header().Set(k, strings.Join(v.Value, ","))
	}
	w.WriteHeader(int(ctx.Context.Response.StatusCode))
	_, err := w.Write(ctx.Context.Response.Body)
	if err != nil {
		return fmt.Errorf("error writing response body: %w", err)
	}
	return nil
}