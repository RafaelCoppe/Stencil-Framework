//go:build js && wasm

package about

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// AboutPage represents the about page
type AboutPage struct {
	framework.BasePage
}

func (p *AboutPage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *AboutPage) Render() string {
	content := StencilUtils.Join(
		StencilText.Titre1("About Stencil Framework", "text-center", "text-primary", "mb-4"),
		StencilText.Paragraphe("This is the about page using Next.js-style routing!", "text-center", "lead", "mb-4"),

		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Routing Features", "mb-3"),
				StencilText.Paragraphe("• File-based routing like Next.js", "mb-2"),
				StencilText.Paragraphe("• Support for page.go, create.go, and edit.go", "mb-2"),
				StencilText.Paragraphe("• Nested routes based on folder structure", "mb-2"),
				StencilText.Paragraphe("• Client-side navigation", "mb-4"),
			),
			"bg-light", "p-4", "rounded", "mb-4",
		),

		StencilPage.Div(
			StencilUtils.Join(
				StencilInteractions.Lien("/", "← Back to Home", "btn", "btn-primary", "me-2"),
				StencilInteractions.Lien("/about/create", "Create Something", "btn", "btn-success", "me-2"),
				StencilInteractions.Lien("/about/edit", "Edit Something", "btn", "btn-warning"),
			),
			"text-center",
		),
	)

	return StencilPage.Container(content, "container", "my-5")
}
