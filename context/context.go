//go:build !tinygo

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

package context

type Context struct{}

// New creates a new empty Context that must be initialized with the Deserialize method
func New() *Context {
	return &Context{}
}

func (ctx *Context) Serialize() (uint32, uint32) {
	return 0, 0
}

func (ctx *Context) Deserialize(uint32, uint32) {}

func (ctx *Context) Next() *Context {
	return ctx
}

//func Debug(string) {}
