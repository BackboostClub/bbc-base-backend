package main

import (
	_ "BbcBase/docs"
	"BbcBase/internal/api"
	"BbcBase/internal/pkg"
)

func main() {
	server := new(pkg.Server)

	handler := new(api.Handler)

	err := server.Run("8000", handler.InitRoutes())
	if err != nil {
		panic(err)
	}
}
