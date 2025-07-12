package framework

// BaseComponent provides a base implementation for components
type BaseComponent struct {
	State *State
	App   *App
}

// NewBaseComponent creates a new base component
func NewBaseComponent(app *App) *BaseComponent {
	state := NewState(app)
	return &BaseComponent{
		State: state,
		App:   app,
	}
}

// Render is the default render method (should be overridden)
func (c *BaseComponent) Render() string {
	return Div("Base Component - Override the Render method")
}

// Update triggers a re-render
func (c *BaseComponent) Update() {
	if c.App != nil {
		c.App.Update()
	}
}

// Page represents a full page component
type Page struct {
	*BaseComponent
	Title string
}

// NewPage creates a new page
func NewPage(app *App, title string) *Page {
	return &Page{
		BaseComponent: NewBaseComponent(app),
		Title:         title,
	}
}
