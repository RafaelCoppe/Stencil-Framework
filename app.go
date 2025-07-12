//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// Welcome page with framework information
type WelcomePage struct {
	framework.BasePage
}

func (p *WelcomePage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{
		"showDetails": false,
	}
}

func (p *WelcomePage) HandleEvent(eventName string, event js.Value) {
	switch eventName {
	case "toggleDetails":
		showDetails := framework.GetStateBool("showDetails")
		framework.SetState("showDetails", !showDetails)
	}
}

func (p *WelcomePage) Render() string {
	showDetails := framework.GetStateBool("showDetails")

	// Main content
	content := StencilUtils.Join(
		// Header section
		StencilText.Titre1("Welcome to Stencil Framework! 🚀", "text-center", "text-primary", "mb-4"),
		StencilText.Paragraphe("Build modern WebAssembly applications with Go", "text-center", "lead", "mb-5"),

		StencilUtils.Hr("my-5"),

		// Getting started section
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("🏁 Getting Started", "mb-3"),
				StencilText.Paragraphe("Follow these simple steps to create your first Stencil application:", "mb-3"),

				// Step list
				StencilPage.Div(
					StencilUtils.Join(
						StencilText.Paragraphe("1. Initialize a new project:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>stencil init my-awesome-app</code>", "bg-light", "p-2", "rounded", "mb-3"),

						StencilText.Paragraphe("2. Navigate to your project:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>cd my-awesome-app</code>", "bg-light", "p-2", "rounded", "mb-3"),

						StencilText.Paragraphe("3. Setup dependencies:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>make setup</code>", "bg-light", "p-2", "rounded", "mb-3"),

						StencilText.Paragraphe("4. Build your application:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>make build</code>", "bg-light", "p-2", "rounded", "mb-3"),

						StencilText.Paragraphe("5. Start the development server:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>make serve</code>", "bg-light", "p-2", "rounded", "mb-4"),
					),
					"",
				),
			),
			"bg-white", "p-4", "rounded", "shadow-sm", "mb-4",
		),

		// Features section
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("✨ Framework Features", "mb-3"),
				StencilPage.Div(
					StencilUtils.Join(
						StencilPage.Div(
							StencilUtils.Join(
								StencilText.Titre3("🎯 Simple State Management", "mb-2"),
								StencilText.Paragraphe("Manage your application state with easy-to-use functions", ""),
							),
							"col-md-4", "mb-3",
						),
						StencilPage.Div(
							StencilUtils.Join(
								StencilText.Titre3("🔧 Rich UI Components", "mb-2"),
								StencilText.Paragraphe("Use pre-built components for forms, layout, text, and interactions", ""),
							),
							"col-md-4", "mb-3",
						),
						StencilPage.Div(
							StencilUtils.Join(
								StencilText.Titre3("⚡ WebAssembly Performance", "mb-2"),
								StencilText.Paragraphe("Enjoy near-native performance with Go compiled to WebAssembly", ""),
							),
							"col-md-4", "mb-3",
						),
					),
					"row",
				),
			),
			"bg-light", "p-4", "rounded", "shadow-sm", "mb-4",
		),

		// Toggle button for more details
		StencilPage.Div(
			StencilInteractions.Bouton(
				func() string {
					if showDetails {
						return "Hide Framework Details"
					}
					return "Show Framework Details"
				}(),
				"toggleDetails",
				[]string{"btn", "btn-outline-primary", "btn-lg"},
			),
			"text-center", "mb-4",
		),
	)

	// Conditional details section
	if showDetails {
		details := StencilUtils.Join(
			StencilUtils.Hr("my-4"),
			StencilPage.Div(
				StencilUtils.Join(
					StencilText.Titre2("📚 Available Components", "mb-3"),

					// Component categories
					StencilPage.Div(
						StencilUtils.Join(
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("📝 Text Components", "mb-2"),
									StencilText.Paragraphe("• Titre1, Titre2, Titre3", "mb-1"),
									StencilText.Paragraphe("• Paragraphe, Span", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("🎛️ Form Components", "mb-2"),
									StencilText.Paragraphe("• InputText, Checkbox", "mb-1"),
									StencilText.Paragraphe("• Label, Form", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("🔗 Interactive Components", "mb-2"),
									StencilText.Paragraphe("• Bouton (Button)", "mb-1"),
									StencilText.Paragraphe("• Lien (Link)", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("📐 Layout Components", "mb-2"),
									StencilText.Paragraphe("• Container, Div", "mb-1"),
									StencilText.Paragraphe("• Page components", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
						),
						"row",
					),
				),
				"bg-info", "bg-opacity-10", "p-4", "rounded", "shadow-sm", "mb-4",
			),

			StencilPage.Div(
				StencilUtils.Join(
					StencilText.Titre2("🔗 Useful Links", "mb-3"),
					StencilText.Paragraphe("• Framework Repository: github.com/RafaelCoppe/Stencil-Framework", "mb-2"),
					StencilText.Paragraphe("• UI Components: github.com/RafaelCoppe/Stencil-Go", "mb-2"),
					StencilText.Paragraphe("• Documentation: Check the CLI-README.md for detailed usage", "mb-2"),
				),
				"bg-success", "bg-opacity-10", "p-4", "rounded", "shadow-sm",
			),
		)
		content = StencilUtils.Join(content, details)
	}

	return StencilPage.Container(content, "container", "my-5")
}

func main() {
	// Initialize and run the welcome page
	framework.Run(&WelcomePage{}, "app")
}
