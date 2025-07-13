//go:build js && wasm

package main

import (
	"github.com/RafaelCoppe/Stencil-Framework/app"
	"github.com/RafaelCoppe/Stencil-Framework/core/framework"
)

func main() {
	// Register routes from app package
	app.RegisterRoutes()

	// Start the application with router
	framework.RunWithRouter("app")
}
