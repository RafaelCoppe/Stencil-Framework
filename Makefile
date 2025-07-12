# Makefile pour le framework WebAssembly Stencil-Go

# Variables
GOOS = js
GOARCH = wasm
BINARY_NAME = app.wasm
MAIN_FILE = main.go
PORT = 8080

# Cibles principales
.PHONY: all build serve clean setup dev create-route help

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
	@echo "� Mise à jour des dépendances..."
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
		echo "❌ Aucun serveur HTTP disponible. Installez Python ou PHP."; \
		exit 1; \
	fi

# Mode développement (compilation + serveur)
dev: setup build serve

# Nettoyage
clean:
	@echo "🧹 Nettoyage des fichiers générés..."
	@rm -f $(BINARY_NAME)
	@echo "✅ Nettoyage terminé"

# CLI pour créer des routes
create-route:
	@if [ -z "$(ROUTE)" ]; then \
		echo "❌ Veuillez spécifier le nom de la route:"; \
		echo "   make create-route ROUTE=nom-de-la-route"; \
		echo "   make create-route ROUTE=admin/dashboard"; \
		exit 1; \
	fi
	@echo "🚀 Création de la route: $(ROUTE)"
	@go run cmd/cli.go create-route $(ROUTE)

# Test de la compilation
test:
	@echo "🧪 Test de la compilation..."
	@GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o /tmp/test_$(BINARY_NAME) $(MAIN_FILE)
	@echo "✅ Test de compilation réussi"
	@rm -f /tmp/test_$(BINARY_NAME)

# Information sur les dépendances
info:
	@echo "📋 Informations sur le projet:"
	@echo "  - GOOS: $(GOOS)"
	@echo "  - GOARCH: $(GOARCH)"
	@echo "  - Binary: $(BINARY_NAME)"
	@echo "  - Main file: $(MAIN_FILE)"
	@echo "  - Port: $(PORT)"
	@echo "  - Go version: $$(go version)"

# Affichage de l'aide
help:
	@echo "📚 Stencil Framework - Commandes disponibles:"
	@echo ""
	@echo "🔨 Compilation:"
	@echo "  make build         - Compiler le projet WebAssembly"
	@echo "  make clean         - Nettoyer les fichiers générés"
	@echo ""
	@echo "🚀 Développement:"
	@echo "  make setup         - Configuration initiale du projet"
	@echo "  make serve         - Démarrer le serveur de développement"
	@echo "  make dev           - Compilation + serveur (mode développement)"
	@echo ""
	@echo "🧭 Routage:"
	@echo "  make create-route ROUTE=nom     - Créer une nouvelle route"
	@echo "  make create-route ROUTE=admin/users - Créer une route imbriquée"
	@echo ""
	@echo "🔧 Autres:"
	@echo "  make help          - Afficher cette aide"