package endpoints

import (
	"fmt"

	"github.com/gustavosbarreto/httptunnel"
	"github.com/labstack/echo"
	"github.com/runtimekit/runtimekit/pkg/with"
)

type deviceEndpoints struct {
	tunnel *httptunnel.Tunnel
}

func (d *deviceEndpoints) getDevices(c echo.Context) error {
	db := with.GetDatabase(c)
	fmt.Println(db)
	return nil
}

func (d *deviceEndpoints) deviceAuth(c echo.Context) error {
	return nil
}

func (d *deviceEndpoints) deviceConnection(c echo.Context) error {
	return nil
}
