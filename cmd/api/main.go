package main

import (
	"github.com/gustavosbarreto/httptunnel"
	"github.com/labstack/echo"
	"github.com/runtimekit/runtimekit/pkg/api/endpoints"
	"github.com/runtimekit/runtimekit/pkg/with"
)

func main() {
	router := echo.New()
	router.Use(with.Database)

	tunnel := httptunnel.NewTunnel(httptunnel.DefaultConnectionURL, httptunnel.DefaultRevdialURL)

	router.Any("/", echo.WrapHandler(tunnel.Router()))
	endpoints.SetupRouter(router, tunnel)

	router.Logger.Fatal(router.Start(":1313"))
}
