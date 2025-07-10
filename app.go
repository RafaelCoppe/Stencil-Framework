package main

import (
	"fmt"
	"syscall/js"

	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// App représente l'application WebAssembly
type App struct {
	container js.Value
	state     map[string]interface{}
}

// NewApp crée une nouvelle instance d'application
func NewApp(containerId string) *App {
	container := js.Global().Get("document").Call("getElementById", containerId)
	if container.IsNull() {
		panic(fmt.Sprintf("Element avec l'ID '%s' non trouvé", containerId))
	}

	return &App{
		container: container,
		state:     make(map[string]interface{}),
	}
}

// SetState met à jour l'état de l'application
func (a *App) SetState(key string, value interface{}) {
	a.state[key] = value
}

// GetState récupère une valeur de l'état
func (a *App) GetState(key string) interface{} {
	return a.state[key]
}

// Render met à jour le DOM avec le contenu HTML généré
func (a *App) Render(html string) {
	a.container.Set("innerHTML", html)
	a.attachEventListeners()
}

// attachEventListeners attache les écouteurs d'événements après le rendu
func (a *App) attachEventListeners() {
	// Attacher les événements click pour les boutons avec data-onclick
	buttons := a.container.Call("querySelectorAll", "[data-onclick]")
	for i := 0; i < buttons.Length(); i++ {
		button := buttons.Index(i)
		onclick := button.Get("dataset").Get("onclick").String()

		// Créer un wrapper pour l'événement
		eventFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
			a.handleEvent(onclick, args[0])
			return nil
		})

		button.Call("addEventListener", "click", eventFunc)
	}

	// Attacher les événements change pour les inputs
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

// handleEvent gère les événements personnalisés
func (a *App) handleEvent(eventName string, event js.Value) {
	switch eventName {
	case "increment":
		count := a.GetState("count")
		if count == nil {
			count = 0
		}
		a.SetState("count", count.(int)+1)
		a.update()
	case "decrement":
		count := a.GetState("count")
		if count == nil {
			count = 0
		}
		a.SetState("count", count.(int)-1)
		a.update()
	case "reset":
		a.SetState("count", 0)
		a.update()
	case "updateText":
		target := event.Get("target")
		value := target.Get("value").String()
		a.SetState("text", value)
		a.update()
	}
}

// update re-rend l'application avec le nouvel état
func (a *App) update() {
	html := a.buildHTML()
	a.Render(html)
}

// buildHTML construit le HTML de l'application en utilisant Stencil
func (a *App) buildHTML() string {
	count := a.GetState("count")
	if count == nil {
		count = 0
	}

	text := a.GetState("text")
	if text == nil {
		text = "Tapez quelque chose..."
	}

	return StencilUtils.Join(
		StencilPage.Container(StencilUtils.Join(
			StencilText.Titre1("Framework WebAssembly avec Stencil-Go", "text-3xl", "font-bold", "mb-8", "text-center"),

			// Section compteur
			StencilPage.Div(StencilUtils.Join(
				StencilText.Titre2("Compteur", "text-xl", "font-semibold", "mb-4"),
				StencilText.Paragraphe(fmt.Sprintf("Valeur actuelle : %d", count), "text-lg", "mb-4"),
				StencilPage.Div(StencilUtils.Join(
					BoutonAvecEvent("Incrémenter", "increment", []string{"bg-blue-500", "text-white", "px-4", "py-2", "rounded", "mr-2"}),
					BoutonAvecEvent("Décrémenter", "decrement", []string{"bg-red-500", "text-white", "px-4", "py-2", "rounded", "mr-2"}),
					BoutonAvecEvent("Reset", "reset", []string{"bg-gray-500", "text-white", "px-4", "py-2", "rounded"}),
				), "justify-center", "gap-2"),
			), "mb-8", "p-6", "bg-gray-100", "rounded-lg"),

			// Section texte
			StencilPage.Div(StencilUtils.Join(
				StencilText.Titre2("Texte dynamique", "text-xl", "font-semibold", "mb-4"),
				InputTextAvecEvent("textInput", "Tapez votre texte...", "", "updateText", []string{"border", "p-2", "rounded", "w-full", "mb-4"}),
				StencilText.Paragraphe(fmt.Sprintf("Texte saisi : %s", text), "text-lg", "p-4", "bg-yellow-100", "rounded"),
			), "mb-8", "p-6", "bg-gray-100", "rounded-lg"),

			// Section demo
			StencilPage.Div(StencilUtils.Join(
				StencilText.Titre2("Autres composants", "text-xl", "font-semibold", "mb-4"),
				StencilPage.Div(StencilUtils.Join(
					StencilText.Paragraphe("Voici une démonstration des autres composants Stencil :", "mb-4"),
					StencilInteractions.Lien("https://github.com/RafaelCoppe/Stencil-Go", "Documentation Stencil", "text-blue-500", "underline", "mb-2", "block"),
					StencilUtils.Hr("my-4"),
					StencilText.Paragraphe("Liste d'exemple :", "font-semibold", "mb-2"),
					StencilPage.Div(StencilUtils.Join(
						StencilText.Span("• Item 1", "block", "mb-1"),
						StencilText.Span("• Item 2", "block", "mb-1"),
						StencilText.Span("• Item 3", "block", "mb-1"),
					), "ml-4"),
				)),
			), "p-6", "bg-gray-100", "rounded-lg"),
		), "max-w-4xl", "mx-auto", "p-8"),
	)
}

// BoutonAvecEvent crée un bouton avec un événement personnalisé
func BoutonAvecEvent(label, event string, classes []string) string {
	return fmt.Sprintf(`<button data-onclick="%s" class="%s">%s</button>`,
		event,
		StencilUtils.Join(classes...),
		label)
}

// InputTextAvecEvent crée un input avec un événement personnalisé
func InputTextAvecEvent(name, placeholder, value, event string, classes []string) string {
	return fmt.Sprintf(`<input type="text" name="%s" placeholder="%s" value="%s" data-onchange="%s" class="%s">`,
		name,
		placeholder,
		value,
		event,
		StencilUtils.Join(classes...))
}

func main() {
	// Créer l'application
	app := NewApp("app")

	// Initialiser l'état
	app.SetState("count", 0)
	app.SetState("text", "")

	// Premier rendu
	app.update()

	// Garder le programme en vie
	select {}
}
