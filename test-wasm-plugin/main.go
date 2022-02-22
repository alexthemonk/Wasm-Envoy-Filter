package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

const (
	tickMilliseconds uint32 = 1000
)

type (
	vmContext struct {
		// Embed the default VM context here,
		// so that we don't need to reimplement all the methods.
		types.DefaultVMContext
	}

	pluginContext struct {
		// Embed the default plugin context here,
		// so that we don't need to reimplement all the methods.
		types.DefaultPluginContext
	}

	httpContext struct {
		// Embed the default plugin context here,
		// so that we don't need to reimplement all the methods.
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

func (ctx *httpContext) OnHttpRequestHeaders(int, bool) types.Action {
	proxywasm.LogInfof("OnHttpRequestHeaders with ctx %v", ctx)
	return types.ActionContinue
}
