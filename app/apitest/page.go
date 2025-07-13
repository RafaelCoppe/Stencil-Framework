//go:build js && wasm

package apitest

import (
	"fmt"
	"syscall/js"

	"github.com/RafaelCoppe/Stencil-Framework/core/framework"
	"github.com/RafaelCoppe/Stencil-Framework/core/http"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// Todo représente un élément todo de l'API
type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"userId"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// ApitestPage represents the apitest page
type ApitestPage struct {
	framework.BasePage
}

func (p *ApitestPage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{
		"todos":       []Todo{},
		"loading":     false,
		"error":       "",
		"apiResponse": "",
	}
}

func (p *ApitestPage) HandleEvent(eventName string, event js.Value) {
	switch eventName {
	case "loadTodos":
		p.loadTodos()
	case "loadUsers":
		p.loadUsers()
	case "createPost":
		p.createPost()
	case "clearTodos":
		framework.SetState("todos", []Todo{})
		framework.SetState("apiResponse", "")
		framework.SetState("error", "")
	}
}

func (p *ApitestPage) loadTodos() {
	// Mettre l'état en chargement
	framework.SetState("loading", true)
	framework.SetState("error", "")

	// Utiliser le module HTTP global (configuré dans main.go)
	go func() {
		// Faire l'appel GET avec le module HTTP global - ULTRA SIMPLE !
		response := http.GET("/todos", map[string]string{
			"_limit": "10",
		})

		// Arrêter le chargement
		framework.SetState("loading", false)

		// Vérifier les erreurs
		if response.Error != nil {
			framework.SetState("error", fmt.Sprintf("Erreur réseau: %v", response.Error))
			return
		}

		if !response.IsSuccess() {
			framework.SetState("error", fmt.Sprintf("Erreur API: Status %d", response.StatusCode))
			return
		}

		// Parser les todos
		var todos []Todo
		if err := response.JSON(&todos); err != nil {
			framework.SetState("error", fmt.Sprintf("Erreur de parsing JSON: %v", err))
			return
		}

		// Mettre à jour l'état avec les todos
		framework.SetState("todos", todos)
	}()
}

func (p *ApitestPage) loadUsers() {
	framework.SetState("loading", true)
	framework.SetState("error", "")

	go func() {
		// Utiliser le module HTTP global pour récupérer les utilisateurs
		response := http.GET("/users", map[string]string{
			"_limit": "5",
		})

		framework.SetState("loading", false)

		if response.Error != nil {
			framework.SetState("error", fmt.Sprintf("Erreur: %v", response.Error))
			return
		}

		if response.IsSuccess() {
			// Pour simplifier, on affiche juste la réponse brute
			framework.SetState("todos", []Todo{}) // Vider les todos
			framework.SetState("apiResponse", response.String())
		}
	}()
}

func (p *ApitestPage) createPost() {
	framework.SetState("loading", true)
	framework.SetState("error", "")

	go func() {
		// Créer un post avec le module HTTP global
		postData := map[string]interface{}{
			"title":  "Post créé avec Stencil Framework",
			"body":   "Ceci est un exemple de création de post via le module HTTP",
			"userId": 1,
		}

		response := http.POST("/posts", postData)

		framework.SetState("loading", false)

		if response.Error != nil {
			framework.SetState("error", fmt.Sprintf("Erreur: %v", response.Error))
			return
		}

		if response.IsSuccess() {
			framework.SetState("todos", []Todo{}) // Vider les todos
			framework.SetState("apiResponse", response.String())
		}
	}()
}

func (p *ApitestPage) Render() string {
	loading := framework.GetStateBool("loading")
	errorMsg := framework.GetStateString("error")
	todos := framework.GetState("todos")
	apiResponse := framework.GetStateString("apiResponse")

	// Header de la page
	header := StencilUtils.Join(
		StencilText.Titre1("API Test - Client HTTP Global", "text-center", "text-primary", "mb-4"),
		StencilText.Paragraphe("Démonstration du client HTTP global configuré dans main.go", "text-center", "lead", "mb-4"),
	)

	// Boutons de contrôle
	buttons := StencilPage.Div(
		StencilUtils.Join(
			StencilInteractions.Bouton("🔄 Charger Todos", "loadTodos", "btn", "btn-primary", "me-2"),
			StencilInteractions.Bouton("👥 Charger Users", "loadUsers", "btn", "btn-info", "me-2"),
			StencilInteractions.Bouton("📝 Créer Post", "createPost", "btn", "btn-success", "me-2"),
			StencilInteractions.Bouton("🗑️ Vider", "clearTodos", "btn", "btn-secondary", "me-2"),
			StencilInteractions.Lien("/", "← Retour", "btn", "btn-outline-secondary"),
		),
		"text-center", "mb-4",
	)

	// Zone de contenu principal
	var contentArea string

	if loading {
		// Affichage du chargement
		contentArea = StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("⏳ Chargement en cours...", "text-center", "text-info"),
				StencilText.Paragraphe("Récupération des données depuis l'API", "text-center", "text-muted"),
			),
			"text-center", "p-4", "bg-light", "rounded",
		)
	} else if errorMsg != "" {
		// Affichage des erreurs
		contentArea = StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("❌ Erreur", "text-center", "text-danger"),
				StencilText.Paragraphe(errorMsg, "text-center", "text-danger"),
			),
			"text-center", "p-4", "bg-danger", "bg-opacity-10", "border", "border-danger", "rounded",
		)
	} else if todosSlice, ok := todos.([]Todo); ok && len(todosSlice) > 0 {
		// Affichage des todos
		contentArea = p.renderTodos(todosSlice)
	} else if apiResponse != "" {
		// Affichage de la réponse API générique
		contentArea = StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("📡 Réponse API", "text-center", "mb-4"),
				StencilPage.Div(
					StencilText.Paragraphe(apiResponse, "font-monospace", "small"),
					"p-3", "bg-light", "border", "rounded", "overflow-auto", "text-start",
				),
			),
			"",
		)
	} else {
		// État initial
		contentArea = StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("📋 Todos", "text-center", "mb-3"),
				StencilText.Paragraphe("Cliquez sur 'Charger les Todos' pour récupérer les données depuis l'API", "text-center", "text-muted"),
			),
			"text-center", "p-4", "bg-light", "rounded",
		)
	}

	// API Info
	apiInfo := StencilPage.Div(
		StencilUtils.Join(
			StencilText.Titre3("🔗 Client HTTP Global", "mb-3"),
			StencilText.Paragraphe("✅ Client initialisé dans main.go", "mb-2", "text-success"),
			StencilText.Paragraphe("🌐 Base URL: https://jsonplaceholder.typicode.com", "mb-2", "font-monospace"),
			StencilText.Paragraphe("📋 Endpoints: /todos, /users, /posts", "mb-2"),
			StencilText.Paragraphe("🚀 Réutilisable partout dans l'application", "text-primary"),
			StencilText.Paragraphe("💡 Usage: framework.HTTPGet(), HTTPPost(), etc.", "text-muted", "small"),
		),
		"mt-5", "p-4", "bg-info", "bg-opacity-10", "border", "border-info", "rounded",
	)

	content := StencilUtils.Join(header, buttons, contentArea, apiInfo)

	return StencilPage.Container(content, "container", "my-5")
}

func (p *ApitestPage) renderTodos(todos []Todo) string {
	var todoItems []string

	for _, todo := range todos {
		status := "❌"
		statusClass := "text-danger"
		if todo.Completed {
			status = "✅"
			statusClass = "text-success"
		}

		todoItem := StencilPage.Div(
			StencilUtils.Join(
				StencilPage.Div(
					StencilUtils.Join(
						StencilText.Span(fmt.Sprintf("#%d", todo.ID), "badge", "bg-secondary", "me-2"),
						StencilText.Span(fmt.Sprintf("User %d", todo.UserID), "badge", "bg-info", "me-2"),
					),
					"mb-2",
				),
				StencilText.Titre4(todo.Title, "mb-2"),
				StencilText.Paragraphe(
					StencilUtils.Join(
						StencilText.Span("Status: ", "text-muted"),
						StencilText.Span(status, statusClass, "me-2"),
						StencilText.Span(
							map[bool]string{true: "Complété", false: "En cours"}[todo.Completed],
							statusClass,
						),
					),
					"mb-0",
				),
			),
			"col-md-6", "col-lg-4", "mb-4",
		)

		todoItems = append(todoItems, todoItem)
	}

	todoGrid := StencilPage.Div(
		StencilUtils.Join(todoItems...),
		"row",
	)

	return StencilPage.Div(
		StencilUtils.Join(
			StencilText.Titre2("📋 Todos récupérés", "text-center", "mb-4"),
			StencilText.Paragraphe(fmt.Sprintf("Affichage de %d todos", len(todos)), "text-center", "text-muted", "mb-4"),
			todoGrid,
		),
		"",
	)
}
