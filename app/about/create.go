//go:build js && wasm

package about

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
)

// AboutCreatePage represents the create page for about
type AboutCreatePage struct {
	framework.BasePage
}

func (p *AboutCreatePage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *AboutCreatePage) Render() string {
	content := StencilUtils.Join(
		StencilText.Titre1("Create New About", "text-center", "text-success", "mb-4"),
		StencilText.Paragraphe("Create a new item in about", "text-center", "lead", "mb-4"),
		
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Create Form", "mb-3"),
				StencilText.Paragraphe("Add your form components here.", "mb-3"),
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
