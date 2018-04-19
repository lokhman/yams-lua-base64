package base64

import (
	"encoding/base64"

	"github.com/lokhman/yams-lua"
)

// Preload adds base64 to the given Lua state's package.preload table. After it
// has been preloaded, it can be loaded using require:
//
//  local base64 = require("base64")
func Preload(L *lua.LState) {
	L.PreloadModule("base64", Loader)
}

// Loader is the module loader function.
func Loader(L *lua.LState) int {
	t := L.NewTable()
	L.SetFuncs(t, api)
	L.Push(t)
	return 1
}

var api = map[string]lua.LGFunction{
	"decode": apiDecode,
	"encode": apiEncode,
}

func apiDecode(L *lua.LState) int {
	data := L.CheckString(1)
	value, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(lua.LString(value))
	return 1
}

func apiEncode(L *lua.LState) int {
	value := L.CheckString(1)
	data := base64.StdEncoding.EncodeToString([]byte(value))
	L.Push(lua.LString(data))
	return 1
}
