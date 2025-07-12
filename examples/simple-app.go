package main

import (
	"fmt"
	"syscall/js"

	"github.com/RafaelCoppe/Stencil-Framework/framework"
)

// HomePage represents the main page of our application
type HomePage struct {
	*framework.Page
}

// NewHomePage creates a new home page
func NewHomePage(app *framework.App) *HomePage {
	page := &HomePage{
		Page: framework.NewPage(app, "Home"),
	}

	// Initialize state
	page.State.Set("count", 0)
	page.State.Set("text", "")

	// Register event handlers
	framework.RegisterEventHandler("increment", page.handleIncrement)
	framework.RegisterEventHandler("decrement", page.handleDecrement)
	framework.RegisterEventHandler("reset", page.handleReset)
	framework.RegisterEventHandler("updateText", page.handleUpdateText)

	return page
}

// Event handlers
func (h *HomePage) handleIncrement(eventName string, event js.Value, app *framework.App) {
	count := h.State.GetInt("count")
	h.State.Set("count", count+1)
}

func (h *HomePage) handleDecrement(eventName string, event js.Value, app *framework.App) {
	count := h.State.GetInt("count")
	h.State.Set("count", count-1)
}

func (h *HomePage) handleReset(eventName string, event js.Value, app *framework.App) {
	h.State.Set("count", 0)
}

func (h *HomePage) handleUpdateText(eventName string, event js.Value, app *framework.App) {
	target := event.Get("target")
	value := target.Get("value").String()
	h.State.Set("text", value)
}

// Render renders the home page
func (h *HomePage) Render() string {
	count := h.State.GetInt("count")
	text := h.State.GetString("text")
	if text == "" {
		text = "Tapez quelque chose..."
	}

	return framework.Join(
		framework.Container(framework.Join(
			framework.H1("My WebAssembly App", "text-3xl", "font-bold", "mb-8", "text-center"),

			// Counter section
			h.renderCounter(count),

			// Text section
			h.renderTextInput(text),
		), "max-w-4xl", "mx-auto", "p-8"),
	)
}

func (h *HomePage) renderCounter(count int) string {
	return framework.Div(framework.Join(
		framework.H2("Counter", "text-xl", "font-semibold", "mb-4"),
		framework.P(fmt.Sprintf("Current value: %d", count), "text-lg", "mb-4"),
		framework.Div(framework.Join(
			framework.Button("Increment", "increment", "bg-blue-500", "text-white", "px-4", "py-2", "rounded", "mr-2"),
			framework.Button("Decrement", "decrement", "bg-red-500", "text-white", "px-4", "py-2", "rounded", "mr-2"),
			framework.Button("Reset", "reset", "bg-gray-500", "text-white", "px-4", "py-2", "rounded"),
		), "flex", "justify-center", "gap-2"),
	), "mb-8", "p-6", "bg-gray-100", "rounded-lg")
}

func (h *HomePage) renderTextInput(text string) string {
	return framework.Div(framework.Join(
		framework.H2("Text Input", "text-xl", "font-semibold", "mb-4"),
		framework.TextInput("textInput", "Type something...", "", "updateText", "border", "p-2", "rounded", "w-full", "mb-4"),
		framework.P(fmt.Sprintf("You typed: %s", text), "text-lg", "p-4", "bg-yellow-100", "rounded"),
	), "mb-8", "p-6", "bg-gray-100", "rounded-lg")
}

func main() {
	// Create the application
	app := framework.NewApp("app")

	// Create and set the home page
	homePage := NewHomePage(app)
	app.SetRootComponent(homePage)

	// Start the application
	app.Start()
}
