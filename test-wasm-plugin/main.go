package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type (
	vmContext struct {
		types.DefaultVMContext
	}

	pluginContext struct {
		types.DefaultPluginContext
	}

	httpContext struct {
		types.DefaultHttpContext
	}
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(uint32) types.HttpContext {
	return &httpContext{}
}

func (*httpContext) OnHttpRequestHeaders(int, bool) types.Action {
	proxywasm.LogInfo("OnHttpRequestHeaders")
	return types.ActionContinue
}
