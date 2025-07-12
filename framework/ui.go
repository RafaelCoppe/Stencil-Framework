package framework

import (
	"fmt"

	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// UI provides convenient methods for creating UI elements
type UI struct{}

// NewUI creates a new UI helper instance
func NewUI() *UI {
	return &UI{}
}

// Container creates a container element
func (ui *UI) Container(content string, classes ...string) string {
	return StencilPage.Container(content, classes...)
}

// Div creates a div element
func (ui *UI) Div(content string, classes ...string) string {
	return StencilPage.Div(content, classes...)
}

// H1 creates an H1 title
func (ui *UI) H1(text string, classes ...string) string {
	return StencilText.Titre1(text, classes...)
}

// H2 creates an H2 title
func (ui *UI) H2(text string, classes ...string) string {
	return StencilText.Titre2(text, classes...)
}

// P creates a paragraph
func (ui *UI) P(text string, classes ...string) string {
	return StencilText.Paragraphe(text, classes...)
}

// Span creates a span element
func (ui *UI) Span(text string, classes ...string) string {
	return StencilText.Span(text, classes...)
}

// Link creates a link element
func (ui *UI) Link(href, text string, classes ...string) string {
	return StencilInteractions.Lien(href, text, classes...)
}

// Button creates a button with an event
func (ui *UI) Button(label, event string, classes ...string) string {
	classStr := StencilUtils.Join(classes...)
	return fmt.Sprintf(`<button data-onclick="%s" class="%s">%s</button>`, event, classStr, label)
}

// Input creates an input field with an event
func (ui *UI) Input(inputType, name, placeholder, value, event string, classes ...string) string {
	classStr := StencilUtils.Join(classes...)
	if event != "" {
		return fmt.Sprintf(`<input type="%s" name="%s" placeholder="%s" value="%s" data-onchange="%s" class="%s">`,
			inputType, name, placeholder, value, event, classStr)
	}
	return fmt.Sprintf(`<input type="%s" name="%s" placeholder="%s" value="%s" class="%s">`,
		inputType, name, placeholder, value, classStr)
}

// TextInput creates a text input field
func (ui *UI) TextInput(name, placeholder, value, event string, classes ...string) string {
	return ui.Input("text", name, placeholder, value, event, classes...)
}

// Hr creates a horizontal rule
func (ui *UI) Hr(classes ...string) string {
	return StencilUtils.Hr(classes...)
}

// Join combines multiple HTML strings
func (ui *UI) Join(parts ...string) string {
	return StencilUtils.Join(parts...)
}

// Global UI instance for convenience
var UI_INSTANCE = NewUI()

// Convenience functions for global access
func Container(content string, classes ...string) string {
	return UI_INSTANCE.Container(content, classes...)
}

func Div(content string, classes ...string) string {
	return UI_INSTANCE.Div(content, classes...)
}

func H1(text string, classes ...string) string {
	return UI_INSTANCE.H1(text, classes...)
}

func H2(text string, classes ...string) string {
	return UI_INSTANCE.H2(text, classes...)
}

func P(text string, classes ...string) string {
	return UI_INSTANCE.P(text, classes...)
}

func Span(text string, classes ...string) string {
	return UI_INSTANCE.Span(text, classes...)
}

func Link(href, text string, classes ...string) string {
	return UI_INSTANCE.Link(href, text, classes...)
}

func Button(label, event string, classes ...string) string {
	return UI_INSTANCE.Button(label, event, classes...)
}

func TextInput(name, placeholder, value, event string, classes ...string) string {
	return UI_INSTANCE.TextInput(name, placeholder, value, event, classes...)
}

func Hr(classes ...string) string {
	return UI_INSTANCE.Hr(classes...)
}

func Join(parts ...string) string {
	return UI_INSTANCE.Join(parts...)
}
