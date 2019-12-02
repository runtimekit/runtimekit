package endpoints

import (
	"github.com/gustavosbarreto/httptunnel"
	"github.com/labstack/echo"
	"github.com/runtimekit/runtimekit/pkg/api/types"
)

func SetupRouter(router *echo.Echo, tunnel *httptunnel.Tunnel) {
	d := &deviceEndpoints{tunnel: tunnel}

	router.GET(types.GetDevicesURL, d.getDevices)
	router.GET(types.DeviceAuthURL, d.deviceAuth)
	router.GET(types.DeviceConnectionURL, d.deviceConnection)
}
