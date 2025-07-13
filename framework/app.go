//go:build js && wasm

package framework

import (
	"fmt"
	"syscall/js"
)

// PageInterface represents the interface that user pages must implement
type PageInterface interface {
	Render() string
	HandleEvent(eventName string, event js.Value)
	GetInitialState() map[string]interface{}
}

// App represents the main application framework (internal use only)
type app struct {
	container js.Value
	state     map[string]interface{}
	page      PageInterface
}

// Global app instance
var appInstance *app

// newApp creates a new application instance (internal)
func newApp(containerId string) *app {
	container := js.Global().Get("document").Call("getElementById", containerId)
	if container.IsNull() {
		panic(fmt.Sprintf("Element with ID '%s' not found", containerId))
	}

	return &app{
		container: container,
		state:     make(map[string]interface{}),
	}
}

// SetRootComponent sets the root component to render
func (a *app) setPage(page PageInterface) {
	a.page = page
}

// setState updates the application state
func (a *app) setState(key string, value interface{}) {
	a.state[key] = value
}

// getState retrieves a value from the state
func (a *app) getState(key string) interface{} {
	return a.state[key]
}

// update re-renders the application
func (a *app) update() {
	if a.page != nil {
		html := a.page.Render()
		a.render(html)
	}
}

// render updates the DOM with the generated HTML
func (a *app) render(html string) {
	a.container.Set("innerHTML", html)
	a.attachEventListeners()
}

// attachEventListeners attaches event listeners after rendering
func (a *app) attachEventListeners() {
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

	// Attach click events for router navigation links
	links := a.container.Call("querySelectorAll", "a[href^='/']")
	for i := 0; i < links.Length(); i++ {
		link := links.Index(i)

		eventFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			event := args[0]
			event.Call("preventDefault")

			href := link.Get("href").String()
			// Extract path from full URL
			url := js.Global().Get("URL").New(href)
			path := url.Get("pathname").String()

			// Navigate using router
			if globalRouter != nil {
				globalRouter.Navigate(path)
			}
			return nil
		})

		link.Call("addEventListener", "click", eventFunc)
	}
}

// handleEvent handles custom events by delegating to the user's page
func (a *app) handleEvent(eventName string, event js.Value) {
	if a.page != nil {
		a.page.HandleEvent(eventName, event)
		a.update() // Auto re-render after event handling
	}
}

// start initializes and starts the application
func (a *app) start() {
	// Initialize state with user's initial state
	if a.page != nil {
		initialState := a.page.GetInitialState()
		for key, value := range initialState {
			a.state[key] = value
		}
	}

	a.update()
	// Keep the program alive
	select {}
}

// startWithRouter starts the application in router mode
func (a *app) startWithRouter() {
	// Let the router handle the initial render
	if globalRouter != nil {
		// Force immediate render for the current path
		globalRouter.currentPath = js.Global().Get("location").Get("pathname").String()
		globalRouter.render()
	}

	// Keep the program alive
	select {}
}

// Public API for users

// Run starts the Stencil application with the provided page
// This is the main entry point for users
func Run(page PageInterface, containerId ...string) {
	containerID := "app" // default
	if len(containerId) > 0 {
		containerID = containerId[0]
	}

	appInstance = newApp(containerID)

	// Initialize router
	router := InitRouter()

	// If a page is provided, register it as the default route
	if page != nil {
		router.RegisterRoute("/", func() PageInterface { return page })
	}

	// Start the application
	appInstance.page = page
	appInstance.start()
}

// RunWithRouter starts the framework with routing capabilities
func RunWithRouter(containerId ...string) {
	containerID := "app" // default
	if len(containerId) > 0 {
		containerID = containerId[0]
	}

	appInstance = newApp(containerID)

	// Initialize router
	InitRouter()

	// Start the application in router mode
	appInstance.startWithRouter()
}

// SetState updates the application state (available to user pages)
func SetState(key string, value interface{}) {
	if appInstance != nil {
		appInstance.setState(key, value)
		appInstance.update()
	}
}

// GetState retrieves a value from the application state (available to user pages)
func GetState(key string) interface{} {
	if appInstance != nil {
		return appInstance.getState(key)
	}
	return nil
}

// GetStateString retrieves a state value as string
func GetStateString(key string) string {
	if val := GetState(key); val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

// GetStateInt retrieves a state value as int
func GetStateInt(key string) int {
	if val := GetState(key); val != nil {
		if i, ok := val.(int); ok {
			return i
		}
	}
	return 0
}

// GetStateBool retrieves a state value as bool
func GetStateBool(key string) bool {
	if val := GetState(key); val != nil {
		if b, ok := val.(bool); ok {
			return b
		}
	}
	return false
}
