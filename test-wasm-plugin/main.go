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

func (*vmContext) OnVMStart(vmConfigurationSize int) types.OnVMStartStatus {
	proxywasm.LogCritical("Inside Go OnVMStart")

	return types.OnVMStartStatusOK
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(uint32) types.PluginContext {
	return &pluginContext{}
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(uint32) types.HttpContext {
	return &httpContext{}
}

// Override types.DefaultHttpContext.
func (*httpContext) OnHttpRequestHeaders(int, bool) types.Action {
	proxywasm.LogInfo("OnHttpRequestHeaders")
	err := proxywasm.AddHttpRequestHeader("X-my-custom-header-asm", "hello alex world")
	if err != nil {
		proxywasm.LogCritical("failed to set request header")
	}
	return types.ActionContinue
}
