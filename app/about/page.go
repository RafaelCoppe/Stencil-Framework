//go:build js && wasm

package about

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
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
		StencilText.Titre1("About", "text-center", "text-primary", "mb-4"),
		StencilText.Paragraphe("Welcome to the about page!", "text-center", "lead", "mb-4"),
		
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Page Actions", "mb-3"),
				StencilInteractions.Lien("/about/create", "Create New", "btn", "btn-success", "me-2"),
				StencilInteractions.Lien("/about/edit", "Edit", "btn", "btn-warning", "me-2"),
				StencilInteractions.Lien("/", "← Back to Home", "btn", "btn-secondary"),
			),
			"text-center", "bg-light", "p-4", "rounded",
		),
	)

	return StencilPage.Container(content, "container", "my-5")
}
