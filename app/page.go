//go:build js && wasm

package app

import (
	"syscall/js"

	"github.com/RafaelCoppe/Stencil-Framework/app/about"
	"github.com/RafaelCoppe/Stencil-Framework/components"
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
	case "navigateToAbout":
		framework.NavigateTo("/about")
	case "navigateToCreate":
		framework.NavigateTo("/about/create")
	case "navigateToEdit":
		framework.NavigateTo("/about/edit")
	}
}

func (p *WelcomePage) Render() string {
	showDetails := framework.GetStateBool("showDetails")

	// Create the main hero section using HeroSection component
	heroSection := components.HeroSection(components.ComponentProps{
		"title":               "Stencil Framework 🚀",
		"subtitle":            "Build modern WebAssembly applications with Go",
		"primaryButtonText":   "Get Started",
		"primaryButtonHref":   "/about",
		"secondaryButtonText": "Create Project",
		"secondaryButtonHref": "/about/create",
	})

	// Navigation demo section
	navigationDemo := StencilPage.Div(
		StencilUtils.Join(
			StencilText.Titre2("🧭 Navigation Demo", "text-3xl", "font-bold", "text-gray-800", "mb-4", "flex", "items-center"),
			StencilText.Paragraphe("Test the Next.js-style routing:", "text-gray-600", "mb-6"),
			StencilPage.Div(
				StencilUtils.Join(
					StencilInteractions.Lien("/about", "About Page", "bg-blue-500", "text-white", "px-6", "py-2", "rounded-lg", "hover:bg-blue-600", "transition-all", "transform", "hover:-translate-y-1", "shadow-md", "inline-block", "text-center", "no-underline"),
					StencilInteractions.Lien("/about/create", "Create Page", "bg-green-500", "text-white", "px-6", "py-2", "rounded-lg", "hover:bg-green-600", "transition-all", "transform", "hover:-translate-y-1", "shadow-md", "inline-block", "text-center", "no-underline"),
					StencilInteractions.Lien("/about/edit", "Edit Page", "bg-yellow-500", "text-white", "px-6", "py-2", "rounded-lg", "hover:bg-yellow-600", "transition-all", "transform", "hover:-translate-y-1", "shadow-md", "inline-block", "text-center", "no-underline"),
				),
				"flex", "flex-wrap", "gap-3", "justify-center",
			),
		),
		"bg-gradient-to-r", "from-blue-50", "to-purple-50", "p-8", "border-b",
	)

	// Getting started section using StepItem components
	gettingStarted := StencilPage.Div(
		StencilUtils.Join(
			StencilText.Titre2("🏁 Getting Started", "text-3xl", "font-bold", "text-gray-800", "mb-4", "flex", "items-center"),
			StencilText.Paragraphe("Follow these simple steps to create your first Stencil application:", "text-gray-600", "mb-6"),
			StencilPage.Div(
				StencilUtils.Join(
					components.StepItem(components.ComponentProps{
						"number":  "1",
						"title":   "Initialize a new project:",
						"command": "stencil init my-awesome-app",
					}),
					components.StepItem(components.ComponentProps{
						"number":  "2",
						"title":   "Navigate to your project:",
						"command": "cd my-awesome-app",
					}),
					components.StepItem(components.ComponentProps{
						"number":  "3",
						"title":   "Setup dependencies:",
						"command": "make setup",
					}),
					components.StepItem(components.ComponentProps{
						"number":  "4",
						"title":   "Build your application:",
						"command": "make build",
					}),
					components.StepItem(components.ComponentProps{
						"number":  "5",
						"title":   "Start the development server:",
						"command": "make serve",
					}),
				),
				"space-y-4",
			),
		),
		"p-8", "border-b",
	)

	// Features section using FeatureCard components
	features := StencilPage.Div(
		StencilUtils.Join(
			StencilText.Titre2("✨ Framework Features", "text-3xl", "font-bold", "text-gray-800", "mb-6", "flex", "items-center"),
			StencilPage.Div(
				StencilUtils.Join(
					components.FeatureCard(components.ComponentProps{
						"icon":        "🎯",
						"title":       "Simple State Management",
						"description": "Manage your application state with easy-to-use functions",
					}),
					components.FeatureCard(components.ComponentProps{
						"icon":        "🔧",
						"title":       "Rich UI Components",
						"description": "Use pre-built components for forms, layout, text, and interactions",
					}),
					components.FeatureCard(components.ComponentProps{
						"icon":        "⚡",
						"title":       "WebAssembly Performance",
						"description": "Enjoy near-native performance with Go compiled to WebAssembly",
					}),
				),
				"grid", "md:grid-cols-3", "gap-6",
			),
		),
		"p-8", "border-b",
	)

	// Toggle button section
	toggleButton := StencilPage.Div(
		StencilInteractions.Bouton(
			func() string {
				if showDetails {
					return "Hide Framework Details"
				}
				return "Show Framework Details"
			}(),
			"toggleDetails",
			"bg-gradient-to-r", "from-blue-500", "to-purple-600", "text-white", "px-8", "py-3", "rounded-xl", "font-semibold", "hover:from-blue-600", "hover:to-purple-700", "transition-all", "transform", "hover:-translate-y-1", "shadow-lg",
		),
		"p-8", "text-center", "border-b",
	)

	// Details section (conditionally rendered)
	detailsSection := ""
	if showDetails {
		componentCard := func(title, items string) string {
			return StencilPage.Div(
				StencilUtils.Join(
					StencilText.Titre3(title, "text-xl", "font-semibold", "text-gray-800", "mb-3"),
					StencilPage.Ul([]string{items}, "text-gray-600", "space-y-1"),
				),
				"bg-gradient-to-br", "from-gray-50", "to-gray-100", "p-6", "rounded-xl",
			)
		}

		detailsSection = StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("📚 Available Components", "text-3xl", "font-bold", "text-gray-800", "mb-6", "flex", "items-center"),
				StencilPage.Div(
					StencilUtils.Join(
						componentCard("📝 Text Components", "• Titre1, Titre2, Titre3<br>• Paragraphe, Span"),
						componentCard("🎛️ Form Components", "• InputText, Checkbox<br>• Label, Form"),
						componentCard("🔗 Interactive Components", "• Bouton (Button)<br>• Lien (Link)"),
						componentCard("📐 Layout Components", "• Container, Div<br>• Page components"),
					),
					"grid", "md:grid-cols-2", "gap-6", "mb-8",
				),
				StencilPage.Div(
					StencilUtils.Join(
						StencilText.Titre3("🔗 Useful Links", "text-xl", "font-semibold", "text-gray-800", "mb-3"),
						StencilPage.Ul([]string{
							"• Framework Repository: github.com/RafaelCoppe/Stencil-Framework",
							"• Documentation: Check the CLI-README.md for detailed usage",
						}, "text-gray-600", "space-y-2"),
					),
					"bg-gradient-to-br", "from-green-50", "to-green-100", "p-6", "rounded-xl",
				),
			),
			"p-8",
		)
	}

	// Main content container
	mainContent := StencilPage.Div(
		StencilUtils.Join(
			navigationDemo,
			gettingStarted,
			features,
			toggleButton,
			detailsSection,
		),
		"bg-white", "rounded-2xl", "shadow-2xl", "overflow-hidden", "mb-8",
	)

	// Main page container
	return StencilPage.Div(
		StencilPage.Div(
			StencilUtils.Join(
				heroSection,
				mainContent,
			),
			"max-w-6xl", "mx-auto",
		),
		"min-h-screen", "bg-gradient-to-br", "from-blue-500", "to-purple-600", "py-8", "px-4",
	)
}

// RegisterRoutes registers all routes for this application
func RegisterRoutes() {
	// Register the welcome page route
	framework.RegisterPageRoute("/",
		func() framework.PageInterface { return &WelcomePage{} },
		nil, // create.go handler
		nil, // edit.go handler
	)

	// Register about page routes
	framework.RegisterPageRoute("/about",
		func() framework.PageInterface { return &about.AboutPage{} },
		func() framework.PageInterface { return &about.AboutCreatePage{} },
		func() framework.PageInterface { return &about.AboutEditPage{} },
	)
}
