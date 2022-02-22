package main

import (
	"math/rand"
	"time"

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
func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	rand.Seed(time.Now().UnixNano())

	proxywasm.LogInfo("OnPluginStart from Go!")
	if err := proxywasm.SetTickPeriodMilliSeconds(tickMilliseconds); err != nil {
		proxywasm.LogCriticalf("failed to set tick period: %v", err)
	}

	return types.OnPluginStartStatusOK
}

// Override types.DefaultPluginContext.
func (ctx *pluginContext) OnTick() {
	t := time.Now().UnixNano()
	proxywasm.LogInfof("It's %d: random value: %d", t, rand.Uint64())
	proxywasm.LogInfof("OnTick called")
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(uint32) types.HttpContext {
	return &httpContext{}
}

func (ctx *httpContext) OnHttpRequestHeaders(int, bool) types.Action {
	proxywasm.LogInfof("OnHttpRequestHeaders with ctx %v", ctx)
	return types.ActionContinue
}
