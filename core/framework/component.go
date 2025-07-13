//go:build js && wasm

package framework

import (
	"syscall/js"

	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
)

// BasePage provides a base implementation that users can embed in their pages
type BasePage struct{}

// GetInitialState provides default empty initial state
func (p *BasePage) GetInitialState() map[string]interface{} {
	return make(map[string]interface{})
}

// HandleEvent provides default empty event handling
func (p *BasePage) HandleEvent(eventName string, event js.Value) {
	// Override this method in your page to handle events
}

// Render provides a default render method
func (p *BasePage) Render() string {
	return StencilText.Titre1("Override the Render method in your page")
}
