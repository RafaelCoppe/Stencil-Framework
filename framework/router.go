//go:build js && wasm

package framework

import (
	"strings"
	"syscall/js"
)

// RouteHandler represents a function that returns a PageInterface
type RouteHandler func() PageInterface

// Router manages application routing
type Router struct {
	routes      map[string]RouteHandler
	currentPath string
	basePath    string
}

// NewRouter creates a new router instance
func NewRouter() *Router {
	return &Router{
		routes:   make(map[string]RouteHandler),
		basePath: "/app",
	}
}

// Global router instance
var globalRouter *Router

// InitRouter initializes the global router
func InitRouter() *Router {
	if globalRouter == nil {
		globalRouter = NewRouter()
		setupBrowserRouting()
	}
	return globalRouter
}

// RegisterRoute registers a route with its handler
func (r *Router) RegisterRoute(path string, handler RouteHandler) {
	// Normalize path
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	r.routes[path] = handler
}

// RegisterPageRoute registers routes for page, create, and edit actions
func (r *Router) RegisterPageRoute(basePath string, pageHandler, createHandler, editHandler RouteHandler) {
	if !strings.HasPrefix(basePath, "/") {
		basePath = "/" + basePath
	}

	if pageHandler != nil {
		r.RegisterRoute(basePath, pageHandler)
		r.RegisterRoute(basePath+"/", pageHandler) // Also handle trailing slash
	}

	if createHandler != nil {
		r.RegisterRoute(basePath+"/create", createHandler)
	}

	if editHandler != nil {
		r.RegisterRoute(basePath+"/edit", editHandler)
	}
}

// Navigate to a specific path
func (r *Router) Navigate(path string) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	r.currentPath = path

	// Update browser URL without reloading
	js.Global().Get("history").Call("pushState", js.Null(), "", path)

	// Render the new page
	r.render()
}

// GetCurrentPath returns the current path
func (r *Router) GetCurrentPath() string {
	if r.currentPath == "" {
		// Get current path from browser
		location := js.Global().Get("location")
		pathname := location.Get("pathname").String()
		r.currentPath = pathname
	}
	return r.currentPath
}

// render renders the current route
func (r *Router) render() {
	path := r.GetCurrentPath()

	// Find matching route
	handler := r.findRoute(path)
	if handler == nil {
		// Try default route
		handler = r.routes["/"]
		if handler == nil {
			// Render 404 page
			r.render404()
			return
		}
	}

	// Get page instance and render
	page := handler()
	if appInstance != nil {
		appInstance.setPage(page)
		appInstance.update()
	}
}

// findRoute finds the best matching route for a path
func (r *Router) findRoute(path string) RouteHandler {
	// First try exact match
	if handler, exists := r.routes[path]; exists {
		return handler
	}

	// Remove trailing slash and try again
	if strings.HasSuffix(path, "/") && len(path) > 1 {
		trimmedPath := strings.TrimSuffix(path, "/")
		if handler, exists := r.routes[trimmedPath]; exists {
			return handler
		}
	}

	// Try with trailing slash
	if !strings.HasSuffix(path, "/") {
		pathWithSlash := path + "/"
		if handler, exists := r.routes[pathWithSlash]; exists {
			return handler
		}
	}

	return nil
}

// render404 renders a 404 page
func (r *Router) render404() {
	notFoundPage := &notFoundPage{}
	if appInstance != nil {
		appInstance.setPage(notFoundPage)
		appInstance.update()
	}
}

// setupBrowserRouting sets up browser navigation event listeners
func setupBrowserRouting() {
	// Handle back/forward navigation
	js.Global().Call("addEventListener", "popstate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if globalRouter != nil {
			globalRouter.currentPath = ""
			globalRouter.render()
		}
		return nil
	}))

	// Handle initial load - execute immediately if DOM is already loaded
	if js.Global().Get("document").Get("readyState").String() == "complete" {
		if globalRouter != nil {
			globalRouter.render()
		}
	} else {
		// Handle initial load when DOM becomes ready
		js.Global().Call("addEventListener", "DOMContentLoaded", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			if globalRouter != nil {
				globalRouter.render()
			}
			return nil
		}))
	}
}

// NavigateTo navigates to a path (global function)
func NavigateTo(path string) {
	if globalRouter != nil {
		globalRouter.Navigate(path)
	}
}

// RegisterRoute registers a route globally
func RegisterRoute(path string, handler RouteHandler) {
	router := InitRouter()
	router.RegisterRoute(path, handler)
}

// RegisterPageRoute registers page routes globally
func RegisterPageRoute(basePath string, pageHandler, createHandler, editHandler RouteHandler) {
	router := InitRouter()
	router.RegisterPageRoute(basePath, pageHandler, createHandler, editHandler)
}

// GetRouter returns the global router instance
func GetRouter() *Router {
	return InitRouter()
}

// notFoundPage represents a 404 page
type notFoundPage struct {
	BasePage
}

func (p *notFoundPage) Render() string {
	return `
	<div class="container text-center mt-5">
		<h1 class="display-1">404</h1>
		<h2>Page Not Found</h2>
		<p class="lead">The page you are looking for does not exist.</p>
		<a href="/" class="btn btn-primary">Go Home</a>
	</div>
	`
}
