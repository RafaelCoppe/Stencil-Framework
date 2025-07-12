//go:build js && wasm

package about

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// CreatePage represents the create page
type CreatePage struct {
	framework.BasePage
}

func (p *CreatePage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *CreatePage) Render() string {
	content := StencilUtils.Join(
		StencilText.Titre1("Create New Item", "text-center", "text-success", "mb-4"),
		StencilText.Paragraphe("This is the create page - /about/create", "text-center", "lead", "mb-4"),

		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Create Form", "mb-3"),
				StencilText.Paragraphe("Here you would typically have a form to create new items.", "mb-3"),
				StencilText.Paragraphe("This demonstrates the create.go routing pattern.", "mb-3"),
			),
			"bg-success", "bg-opacity-10", "p-4", "rounded", "mb-4",
		),

		StencilPage.Div(
			StencilUtils.Join(
				StencilInteractions.Lien("/about", "← Back to About", "btn", "btn-secondary", "me-2"),
				StencilInteractions.Lien("/", "Home", "btn", "btn-primary"),
			),
			"text-center",
		),
	)

	return StencilPage.Container(content, "container", "my-5")
}
