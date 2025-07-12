package framework

import (
	"fmt"
	"syscall/js"
)

// Component represents a reusable UI component
type Component interface {
	Render() string
}

// App represents the main application framework
type App struct {
	container     js.Value
	state         map[string]interface{}
	rootComponent Component
}

// NewApp creates a new application instance
func NewApp(containerId string) *App {
	container := js.Global().Get("document").Call("getElementById", containerId)
	if container.IsNull() {
		panic(fmt.Sprintf("Element with ID '%s' not found", containerId))
	}

	return &App{
		container: container,
		state:     make(map[string]interface{}),
	}
}

// SetRootComponent sets the root component to render
func (a *App) SetRootComponent(component Component) {
	a.rootComponent = component
}

// SetState updates the application state
func (a *App) SetState(key string, value interface{}) {
	a.state[key] = value
}

// GetState retrieves a value from the state
func (a *App) GetState(key string) interface{} {
	return a.state[key]
}

// Update re-renders the application
func (a *App) Update() {
	if a.rootComponent != nil {
		html := a.rootComponent.Render()
		a.render(html)
	}
}

// render updates the DOM with the generated HTML
func (a *App) render(html string) {
	a.container.Set("innerHTML", html)
	a.attachEventListeners()
}

// attachEventListeners attaches event listeners after rendering
func (a *App) attachEventListeners() {
	// Attach click events for buttons with data-onclick
	buttons := a.container.Call("querySelectorAll", "[data-onclick]")
	for i := 0; i < buttons.Length(); i++ {
		button := buttons.Index(i)
		onclick := button.Get("dataset").Get("onclick").String()

		eventFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			a.handleEvent(onclick, args[0])
			return nil
		})

		button.Call("addEventListener", "click", eventFunc)
	}

	// Attach change events for inputs
	inputs := a.container.Call("querySelectorAll", "input[data-onchange]")
	for i := 0; i < inputs.Length(); i++ {
		input := inputs.Index(i)
		onchange := input.Get("dataset").Get("onchange").String()

		eventFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			a.handleEvent(onchange, args[0])
			return nil
		})

		input.Call("addEventListener", "change", eventFunc)
	}
}

// EventHandler represents a function that handles events
type EventHandler func(eventName string, event js.Value, app *App)

var eventHandlers = make(map[string]EventHandler)

// RegisterEventHandler registers a custom event handler
func RegisterEventHandler(eventName string, handler EventHandler) {
	eventHandlers[eventName] = handler
}

// handleEvent handles custom events
func (a *App) handleEvent(eventName string, event js.Value) {
	if handler, exists := eventHandlers[eventName]; exists {
		handler(eventName, event, a)
	}
}

// Start initializes and starts the application
func (a *App) Start() {
	a.Update()
	// Keep the program alive
	select {}
}
