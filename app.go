//go:build js && wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
)

// Simple counter page demonstrating the new framework
type CounterPage struct {
	framework.BasePage
}

func (p *CounterPage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{
		"count": 0,
	}
}

func (p *CounterPage) HandleEvent(eventName string, event js.Value) {
	count := framework.GetStateInt("count")

	switch eventName {
	case "increment":
		framework.SetState("count", count+1)
	case "decrement":
		framework.SetState("count", count-1)
	case "reset":
		framework.SetState("count", 0)
	}
}

func (p *CounterPage) Render() string {
	count := framework.GetStateInt("count")

	return StencilPage.Container(
		StencilUtils.Join(
			StencilText.Titre1("Simple Counter", "text-center", "mb-4"),
			StencilText.Paragraphe(fmt.Sprintf("Current count: %d", count), "text-center", "mb-4"),
			StencilPage.Div(
				StencilUtils.Join(
					StencilInteractions.Bouton("−", "decrement", []string{"btn", "btn-danger", "me-2"}),
					StencilInteractions.Bouton("Reset", "reset", []string{"btn", "btn-secondary", "me-2"}),
					StencilInteractions.Bouton("+", "increment", []string{"btn", "btn-success"}),
				),
				"text-center",
			),
		),
		"container", "mt-5",
	)
}

func main() {
	// That's it! Just pass your page to the framework
	framework.Run(&CounterPage{}, "app")
}
