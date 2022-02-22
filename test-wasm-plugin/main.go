package main

import (
	"errors"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type (
	vmContext struct {
	}

	pluginContext struct {
		types.DefaultPluginContext
	}

	httpContext struct {
		types.DefaultHttpContext
	}
)

const (
	sharedDataKey = "hello_world_shared_data_key"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

// Override types.VMContext.
func (*vmContext) OnVMStart(vmConfigurationSize int) types.OnVMStartStatus {

	proxywasm.LogCritical("Inside Go OnVMStart")

	headers := [][2]string{
		{":method", "GET"},
		{":path", "/uuid"},
		{":authority", "localhost"},
		{":scheme", "http"},
	}

	if _, err := proxywasm.DispatchHttpCall("httpbin2", headers, nil, nil,
		5000, httpCallResponseCallback); err != nil {
		proxywasm.LogCriticalf("HttpBin2 Dispatch http call failed: %v", err)
	}

	return types.OnVMStartStatusOK
}

func httpCallResponseCallback(numHeaders, bodySize, numTrailers int) {
	resp, _ := proxywasm.GetHttpCallResponseBody(0, bodySize)
	response_json := string(resp)
	initialValueBuf := []byte(response_json)
	if err := proxywasm.SetSharedData(sharedDataKey, initialValueBuf, 0); err != nil {
		proxywasm.LogWarnf("Error setting shared uuid data on OnVMStart: %v", err)
	}
	proxywasm.LogInfof("Httpbin2 RESPONSE %v", response_json)
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

// Override types.DefaultPluginContext.
func (*pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpContext{}
}

// Override types.DefaultHttpContext.
func (ctx *httpContext) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	for {
		value, err := ctx.getSharedData()
		if err == nil {
			proxywasm.LogInfof("shared data value: %s", value)
		} else if errors.Is(err, types.ErrorStatusCasMismatch) {
			continue
		}
		break
	}
	return types.ActionContinue
}

func (ctx *httpContext) getSharedData() (string, error) {
	value, cas, err := proxywasm.GetSharedData(sharedDataKey)
	if err != nil {
		proxywasm.LogWarnf("Error getting shared data on OnHttpRequestHeaders with cas %d: %v ", cas, err)
		return "error", err
	}

	shared_value := string(value)

	return shared_value, err
}
