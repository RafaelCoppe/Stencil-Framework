//go:build js && wasm

package main

import (
	"github.com/RafaelCoppe/Stencil-Framework/app"
	"github.com/RafaelCoppe/Stencil-Framework/core/framework"
	"github.com/RafaelCoppe/Stencil-Framework/core/http"
)

func main() {
	// ========== CONFIGURATION HTTP CLIENT ==========
	// Initialiser le client HTTP global UNE SEULE FOIS

	// Option 1: Client JSONPlaceholder pré-configuré (RECOMMANDÉ pour débuter)
	http.InitJSONPlaceholder()

	// Option 2: Client personnalisé
	// http.InitCustom("https://your-api.com", func(client *http.Client) {
	//     client.SetAuth("your-token")
	//     client.SetHeader("X-API-Version", "v1")
	//     client.SetTimeout(60 * time.Second)
	// })

	// Option 3: Client GitHub
	// http.InitGitHub("your-github-token")

	// Option 4: Client basique
	// client := http.Init("https://api.example.com")
	// client.SetContentType("application/json")
	// client.SetAuth("your-token")

	// ========== DÉMARRAGE APPLICATION ==========

	// Register routes from app package
	app.RegisterRoutes()

	// Start the application with router
	framework.RunWithRouter("app")
}
