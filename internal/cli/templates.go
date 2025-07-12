package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// cleanIdentifier converts a project name to a valid Go identifier
func cleanIdentifier(name string) string {
	// Remove or replace invalid characters
	reg := regexp.MustCompile(`[^a-zA-Z0-9]`)
	cleaned := reg.ReplaceAllString(name, "")

	// Ensure it starts with a letter
	if len(cleaned) > 0 && !regexp.MustCompile(`^[a-zA-Z]`).MatchString(cleaned) {
		cleaned = "App" + cleaned
	}

	// Capitalize first letter
	if len(cleaned) > 0 {
		cleaned = strings.ToUpper(string(cleaned[0])) + cleaned[1:]
	}

	// If empty, use default
	if cleaned == "" {
		cleaned = "App"
	}

	return cleaned
}

func createIndexHTML(projectDir, projectName string) error {
	content := `<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>` + projectName + ` - Stencil Framework</title>
    <script src="https://cdn.tailwindcss.com"></script>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            margin: 0;
            padding: 20px;
        }
        
        #app {
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        
        .loading {
            text-align: center;
            color: white;
            font-size: 1.2em;
        }
        
        .container {
            background: white;
            border-radius: 16px;
            box-shadow: 0 25px 50px rgba(0,0,0,0.1);
            overflow: hidden;
        }
        
        .transition-all {
            transition: all 0.3s ease;
        }
        
        button:hover {
            transform: translateY(-2px);
            box-shadow: 0 10px 20px rgba(0,0,0,0.2);
        }
        
        input:focus {
            outline: none;
            box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.1);
            border-color: #3b82f6;
        }
    </style>
</head>
<body>
    <div id="app">
        <div class="loading">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-white mb-4"></div>
            <p>Chargement de l'application WebAssembly...</p>
        </div>
    </div>

    <script src="wasm_exec.js"></script>
    <script>
        const go = new Go();
        
        // Charger et exécuter le module WebAssembly
        WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject).then((result) => {
            go.run(result.instance);
        }).catch((error) => {
            console.error('Erreur lors du chargement du WebAssembly:', error);
            document.getElementById('app').innerHTML = ` + "`" + `
                <div class="text-center text-white">
                    <h2 class="text-2xl font-bold mb-4">Erreur de chargement</h2>
                    <p class="mb-4">Impossible de charger l'application WebAssembly.</p>
                    <p class="text-sm opacity-75">Vérifiez que les fichiers app.wasm et wasm_exec.js sont présents.</p>
                </div>
            ` + "`" + `;
        });
    </script>
</body>
</html>`

	filePath := filepath.Join(projectDir, "index.html")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func createMakefile(projectDir string) error {
	content := `# Makefile pour le framework WebAssembly Stencil-Go

# Variables
GOOS = js
GOARCH = wasm
BINARY_NAME = app.wasm
MAIN_FILE = app/app.go
PORT = 8080

# Cibles principales
.PHONY: all build serve clean setup dev

all: build

# Compilation du WebAssembly
build:
	@echo "🔨 Compilation du WebAssembly..."
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Compilation terminée : $(BINARY_NAME)"

# Configuration initiale
setup:
	@echo "🚀 Configuration du projet..."
	@if [ ! -f "wasm_exec.js" ]; then \
		echo "📋 Copie de wasm_exec.js..."; \
		if [ -f "$$(go env GOROOT)/lib/wasm/wasm_exec.js" ]; then \
			cp "$$(go env GOROOT)/lib/wasm/wasm_exec.js" .; \
		elif [ -f "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ]; then \
			cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" .; \
		elif [ -f "/usr/local/go/lib/wasm/wasm_exec.js" ]; then \
			cp "/usr/local/go/lib/wasm/wasm_exec.js" .; \
		else \
			echo "❌ wasm_exec.js introuvable. Téléchargement depuis GitHub..."; \
			curl -s -o wasm_exec.js https://raw.githubusercontent.com/golang/go/master/misc/wasm/wasm_exec.js; \
		fi; \
	fi
	@echo "📦 Mise à jour des dépendances..."
	@go mod tidy
	@echo "✅ Configuration terminée"

# Serveur de développement
serve:
	@echo "🌐 Démarrage du serveur de développement sur le port $(PORT)..."
	@if command -v python3 > /dev/null; then \
		python3 -m http.server $(PORT); \
	elif command -v python > /dev/null; then \
		python -m SimpleHTTPServer $(PORT); \
	elif command -v php > /dev/null; then \
		php -S localhost:$(PORT); \
	else \
		echo "❌ Aucun serveur HTTP trouvé. Installez Python ou PHP pour utiliser cette fonctionnalité."; \
		exit 1; \
	fi

# Développement avec rechargement automatique
dev: build
	@echo "🔄 Mode développement activé"
	@echo "📂 Surveillance des changements dans app/"
	@while true; do \
		if [ "$$(find app/ -name '*.go' -newer $(BINARY_NAME) 2>/dev/null | wc -l)" -gt 0 ]; then \
			echo "🔨 Changements détectés, recompilation..."; \
			make build; \
		fi; \
		sleep 2; \
	done

# Nettoyage
clean:
	@echo "🧹 Nettoyage des fichiers générés..."
	@rm -f $(BINARY_NAME)
	@echo "✅ Nettoyage terminé"

# Tests
test:
	@echo "🧪 Exécution des tests..."
	@go test ./...

# Aide
help:
	@echo "📖 Commandes disponibles :"
	@echo "  make setup   - Configuration initiale du projet"
	@echo "  make build   - Compilation du WebAssembly"
	@echo "  make serve   - Démarrer le serveur de développement"
	@echo "  make dev     - Mode développement avec rechargement automatique"
	@echo "  make test    - Exécuter les tests"
	@echo "  make clean   - Nettoyer les fichiers générés"
	@echo "  make help    - Afficher cette aide"
`

	filePath := filepath.Join(projectDir, "Makefile")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func createAppGoFile(appDir, projectName string) error {
	structName := cleanIdentifier(projectName)

	content := `//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// ` + structName + `Page - Welcome page for the ` + projectName + ` application
type ` + structName + `Page struct {
	framework.BasePage
}

func (p *` + structName + `Page) GetInitialState() map[string]interface{} {
	return map[string]interface{}{
		"showDetails": false,
	}
}

func (p *` + structName + `Page) HandleEvent(eventName string, event js.Value) {
	switch eventName {
	case "toggleDetails":
		showDetails := framework.GetStateBool("showDetails")
		framework.SetState("showDetails", !showDetails)
	}
}

func (p *` + structName + `Page) Render() string {
	showDetails := framework.GetStateBool("showDetails")

	// Main content
	content := StencilUtils.Join(
		// Header section
		StencilText.Titre1("Welcome to ` + projectName + `! 🚀", "text-center", "text-primary", "mb-4"),
		StencilText.Paragraphe("Build modern WebAssembly applications with Go", "text-center", "lead", "mb-5"),
		
		StencilUtils.Hr("my-5"),

		// Getting started section
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("🏁 Getting Started", "mb-3"),
				StencilText.Paragraphe("Follow these simple steps to develop your ` + projectName + ` application:", "mb-3"),
				
				// Step list
				StencilPage.Div(
					StencilUtils.Join(
						StencilText.Paragraphe("1. Setup dependencies:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>make setup</code>", "bg-light", "p-2", "rounded", "mb-3"),
						
						StencilText.Paragraphe("2. Build your application:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>make build</code>", "bg-light", "p-2", "rounded", "mb-3"),
						
						StencilText.Paragraphe("3. Start the development server:", "mb-2", "fw-bold"),
						StencilPage.Div("<code>make serve</code>", "bg-light", "p-2", "rounded", "mb-3"),
						
						StencilText.Paragraphe("4. Edit app/app.go to customize your application", "mb-2", "fw-bold"),
						StencilText.Paragraphe("5. Visit http://localhost:8080 to see your changes", "mb-4", "fw-bold"),
					),
					"",
				),
			),
			"bg-white", "p-4", "rounded", "shadow-sm", "mb-4",
		),

		// Features section
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("✨ Framework Features", "mb-3"),
				StencilPage.Div(
					StencilUtils.Join(
						StencilPage.Div(
							StencilUtils.Join(
								StencilText.Titre3("🎯 Simple State Management", "mb-2"),
								StencilText.Paragraphe("Manage your application state with easy-to-use functions", ""),
							),
							"col-md-4", "mb-3",
						),
						StencilPage.Div(
							StencilUtils.Join(
								StencilText.Titre3("🔧 Rich UI Components", "mb-2"),
								StencilText.Paragraphe("Use pre-built components for forms, layout, text, and interactions", ""),
							),
							"col-md-4", "mb-3",
						),
						StencilPage.Div(
							StencilUtils.Join(
								StencilText.Titre3("⚡ WebAssembly Performance", "mb-2"),
								StencilText.Paragraphe("Enjoy near-native performance with Go compiled to WebAssembly", ""),
							),
							"col-md-4", "mb-3",
						),
					),
					"row",
				),
			),
			"bg-light", "p-4", "rounded", "shadow-sm", "mb-4",
		),

		// Toggle button for more details
		StencilPage.Div(
			StencilInteractions.Bouton(
				func() string {
					if showDetails {
						return "Hide Framework Details"
					}
					return "Show Framework Details"
				}(),
				"toggleDetails",
				[]string{"btn", "btn-outline-primary", "btn-lg"},
			),
			"text-center", "mb-4",
		),
	)

	// Conditional details section
	if showDetails {
		details := StencilUtils.Join(
			StencilUtils.Hr("my-4"),
			StencilPage.Div(
				StencilUtils.Join(
					StencilText.Titre2("📚 Available Components", "mb-3"),
					
					// Component categories
					StencilPage.Div(
						StencilUtils.Join(
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("📝 Text Components", "mb-2"),
									StencilText.Paragraphe("• Titre1, Titre2, Titre3", "mb-1"),
									StencilText.Paragraphe("• Paragraphe, Span", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("🎛️ Form Components", "mb-2"),
									StencilText.Paragraphe("• InputText, Checkbox", "mb-1"),
									StencilText.Paragraphe("• Label, Form", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("🔗 Interactive Components", "mb-2"),
									StencilText.Paragraphe("• Bouton (Button)", "mb-1"),
									StencilText.Paragraphe("• Lien (Link)", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
							StencilPage.Div(
								StencilUtils.Join(
									StencilText.Titre3("📐 Layout Components", "mb-2"),
									StencilText.Paragraphe("• Container, Div", "mb-1"),
									StencilText.Paragraphe("• Page components", "mb-1"),
								),
								"col-md-6", "mb-3",
							),
						),
						"row",
					),
				),
				"bg-info", "bg-opacity-10", "p-4", "rounded", "shadow-sm", "mb-4",
			),
			
			StencilPage.Div(
				StencilUtils.Join(
					StencilText.Titre2("🔗 Useful Links", "mb-3"),
					StencilText.Paragraphe("• Framework Repository: github.com/RafaelCoppe/Stencil-Framework", "mb-2"),
					StencilText.Paragraphe("• UI Components: github.com/RafaelCoppe/Stencil-Go", "mb-2"),
					StencilText.Paragraphe("• Documentation: Check the CLI-README.md for detailed usage", "mb-2"),
				),
				"bg-success", "bg-opacity-10", "p-4", "rounded", "shadow-sm",
			),
		)
		content = StencilUtils.Join(content, details)
	}

	return StencilPage.Container(content, "container", "my-5")
}

func main() {
	// Initialize and run the ` + projectName + ` application
	framework.Run(&` + structName + `Page{}, "app")
}
`

	filePath := filepath.Join(appDir, "app.go")
	return os.WriteFile(filePath, []byte(content), 0644)
}

func createGoMod(projectDir, projectName string) error {
	content := fmt.Sprintf(`module %s

go 1.24.1

require (
	github.com/RafaelCoppe/Stencil-Framework/framework v0.0.0-20250712123652-085be36df766
	github.com/RafaelCoppe/Stencil-Go v1.0.3
)
`, projectName)

	filePath := filepath.Join(projectDir, "go.mod")
	return os.WriteFile(filePath, []byte(content), 0644)
}
