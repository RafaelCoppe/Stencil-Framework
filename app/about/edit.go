//go:build js && wasm

package about

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// EditPage represents the edit page
type EditPage struct {
	framework.BasePage
}

func (p *EditPage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *EditPage) Render() string {
	content := StencilUtils.Join(
		StencilText.Titre1("Edit Item", "text-center", "text-warning", "mb-4"),
		StencilText.Paragraphe("This is the edit page - /about/edit", "text-center", "lead", "mb-4"),

		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Edit Form", "mb-3"),
				StencilText.Paragraphe("Here you would typically have a form to edit existing items.", "mb-3"),
				StencilText.Paragraphe("This demonstrates the edit.go routing pattern.", "mb-3"),
			),
			"bg-warning", "bg-opacity-10", "p-4", "rounded", "mb-4",
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
