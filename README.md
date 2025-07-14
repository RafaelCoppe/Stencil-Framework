# Framework WebAssembly avec Stencil-Go

Un framework moderne pour créer des applications web interactives en Go en utilisant WebAssembly et la bibliothèque [Stencil-Go](https://github.com/RafaelCoppe/Stencil-Go).

## 🚀 Fonctionnalités

- **WebAssembly** : Exécution native du code Go dans le navigateur
- **Stencil-Go** : Génération déclarative de HTML avec une syntaxe fluide et modulaire
- **Réactivité** : Gestion d'état et re-rendu automatique
- **Routage** : Système de routage inspiré de Next.js avec support de routes basées sur fichiers
- **Composants** : Système de composants réutilisables avec props
- **HTTP Client** : Client HTTP intégré avec support API (JSONPlaceholder, GitHub, etc.)
- **CLI Tools** : Outils en ligne de commande pour créer routes et composants
- **Styling** : Intégration Tailwind CSS avec classes utilitaires
- **Performance** : Rendu optimisé pour WebAssembly

## 🧭 Système de Routage Next.js

Le framework inclut un système de routage inspiré de Next.js :

- **Routes basées sur fichiers** : Organisez vos pages dans le dossier `app/`
- **Support page/create/edit** : Chaque route peut avoir `page.go`, `create.go`, et `edit.go`
- **Navigation client** : Navigation sans rechargement de page
- **Historique navigateur** : Support complet des boutons précédent/suivant

Voir [ROUTING.md](ROUTING.md) pour la documentation complète.

## 📁 Structure du projet

```text
stencil-framework/
├── main.go                    # Point d'entrée avec config HTTP et routage
├── app/
│   ├── page.go               # Page d'accueil avec démo interactive
│   ├── about/                # Pages about avec CRUD
│   └── apitest/              # Page de démonstration HTTP client
├── components/               # Composants réutilisables
│   ├── examples.go          # Composants de démonstration
│   └── go.mod               # Module composants
├── core/                    # Framework (ne pas modifier)
│   ├── framework/           # Code du framework
│   ├── http/                # Client HTTP global
│   ├── cmd/                 # Outils CLI
│   ├── index.html           # Page HTML d'entrée
│   └── wasm_exec.js         # Runtime WebAssembly Go
├── go.mod                   # Dépendances Go
├── Makefile                 # Commandes automatisées
└── README.md                # Documentation
```

## 🛠 Installation et configuration

### Prérequis

- Go 1.21 ou supérieur
- Git
- Un serveur HTTP (Python, PHP, ou Node.js)

### Installation rapide

1. **Clonez le projet** :

```bash
git clone https://github.com/RafaelCoppe/Stencil-Framework.git
cd Stencil-Framework
```

2. **Configuration automatique** :

```bash
make setup
```

3. **Compilation du projet** :

```bash
make build
```

4. **Lancement du serveur** :

```bash
make serve
```

L'application sera accessible sur `http://localhost:8080`

### Configuration détaillée

La commande `make setup` :

- Copie le fichier `wasm_exec.js` nécessaire depuis l'installation Go
- Met à jour les dépendances du module
- Prépare l'environnement de développement

## 🏗 Compilation

### Compilation simple
```bash
make build
```

### Mode développement (compilation + serveur)
```bash
make dev
```

### Test de compilation
```bash
make test
```

## 🌐 Exécution

### Serveur de développement
```bash
make serve
```

L'application sera disponible sur `http://localhost:8080`

### Serveurs alternatifs

**Python 3 :**
```bash
python3 -m http.server 8080
```

**Python 2 :**
```bash
python -m SimpleHTTPServer 8080
```

**PHP :**
```bash
php -S localhost:8080
```

**Node.js (avec http-server) :**
```bash
npx http-server -p 8080
```

## 🎯 Utilisation

### Structure d'une page

Chaque page du framework étend `framework.BasePage` et implémente trois méthodes :

```go
type MyPage struct {
    framework.BasePage
}

func (p *MyPage) GetInitialState() map[string]interface{} {
    return map[string]interface{}{
        "count": 0,
        "message": "Hello World",
    }
}

func (p *MyPage) HandleEvent(eventName string, event js.Value) {
    switch eventName {
    case "increment":
        count := framework.GetState("count").(int)
        framework.SetState("count", count + 1)
    case "decrement":
        count := framework.GetState("count").(int)
        framework.SetState("count", count - 1)
    }
}

func (p *MyPage) Render() string {
    count := framework.GetState("count").(int)
    
    return StencilPage.Container(
        StencilUtils.Join(
            StencilText.Titre1("Compteur", "text-2xl", "font-bold", "mb-4"),
            StencilText.Paragraphe(fmt.Sprintf("Valeur: %d", count), "text-lg", "mb-4"),
            StencilInteractions.Bouton("Incrémenter", "increment", "btn", "btn-primary"),
        ),
        "max-w-md", "mx-auto", "p-6",
    )
}
```

### Client HTTP global

Le framework inclut un client HTTP configurable dans `main.go` :

```go
func main() {
    // Option 1: Client JSONPlaceholder (recommandé pour débuter)
    http.InitJSONPlaceholder()
    
    // Option 2: Client GitHub avec token
    // http.InitGitHub("your-github-token")
    
    // Option 3: Client personnalisé
    // http.InitCustom("https://api.example.com", func(client *http.Client) {
    //     client.SetAuth("your-token")
    //     client.SetHeader("X-API-Version", "v1")
    // })
    
    // Démarrage de l'application
    app.RegisterRoutes()
    framework.RunWithRouter("app")
}
```

### Utilisation du client HTTP

```go
// Dans une page
func (p *MyPage) HandleEvent(eventName string, event js.Value) {
    switch eventName {
    case "loadData":
        // GET request
        data, err := framework.HTTPGet("/todos/1")
        if err == nil {
            framework.SetState("data", data)
        }
        
        // POST request
        newData := map[string]interface{}{
            "title": "New Todo",
            "completed": false,
        }
        response, err := framework.HTTPPost("/todos", newData)
    }
}
```

### Système de composants

Créez des composants réutilisables avec le système de props :

```go
// Définir un composant
func MyCard(props components.ComponentProps) string {
    title := components.PropString(props, "title", "Titre par défaut")
    content := components.PropString(props, "content", "Contenu")
    
    return StencilPage.Div(
        StencilUtils.Join(
            StencilText.Titre3(title, "text-lg", "font-bold", "mb-2"),
            StencilText.Paragraphe(content, "text-gray-600"),
        ),
        "bg-white", "p-4", "rounded", "shadow",
    )
}

// Utiliser le composant
func (p *MyPage) Render() string {
    return MyCard(components.ComponentProps{
        "title": "Ma carte",
        "content": "Contenu de ma carte",
    })
}
```

### Navigation programmatique

```go
func (p *MyPage) HandleEvent(eventName string, event js.Value) {
    switch eventName {
    case "goToAbout":
        framework.NavigateTo("/about")
    case "goToEdit":
        framework.NavigateTo("/about/edit")
    }
}
```

### Composants Stencil disponibles

#### Layout
- `Container()` - Conteneur centré
- `HorizontalBar()` - Flex row
- `VerticalBar()` - Flex column
- `Section()` - Section
- `Div()` - Division générique

#### Texte
- `Titre1()`, `Titre2()` - Titres H1, H2
- `Paragraphe()` - Paragraphe
- `Span()` - Span inline

#### Interaction
- `Bouton()` - Bouton
- `Lien()` - Lien
- `Form()` - Formulaire
- `InputText()` - Champ texte
- `TextArea()` - Zone de texte
- `CheckBox()` - Case à cocher

#### Média
- `Image()` - Image
- `Video()` - Vidéo

#### Utilitaires
- `HR()` - Ligne horizontale
- `Br()` - Saut de ligne
- `Join()` - Concaténation

## 🎨 Styling

Le framework utilise Tailwind CSS pour le styling. Vous pouvez utiliser toutes les classes Tailwind :

```go
stencil.Bouton("Mon bouton", "alert('click')", []string{
    "bg-blue-500",
    "text-white",
    "px-4",
    "py-2",
    "rounded",
    "hover:bg-blue-600",
    "transition-colors"
})
```

## 📚 Pages de démonstration

Le framework inclut plusieurs pages de démonstration :

### Page d'accueil (/)

La page d'accueil montre toutes les fonctionnalités du framework :

- **Section Hero** : Titre animé avec boutons d'action
- **Navigation Demo** : Liens vers toutes les pages avec styles Tailwind
- **Getting Started** : Guide étape par étape avec composants `StepItem`
- **Features** : Cartes de fonctionnalités avec composants `FeatureCard`
- **Toggle Details** : Bouton interactif pour afficher/masquer les détails
- **Components Overview** : Liste des composants disponibles

### Page About (/about)

Démontre le système de routage Next.js avec :

- `/about` - Page principale
- `/about/create` - Page de création
- `/about/edit` - Page d'édition

### Page API Test (/apitest)

Démontre le client HTTP global avec :

- Chargement de données depuis JSONPlaceholder
- Gestion des états de chargement
- Affichage des réponses API
- Gestion des erreurs

### Création de nouvelles routes

```bash
# Créer une nouvelle route
make create-route ROUTE=blog

# Créer une route imbriquée
make create-route ROUTE=admin/users
```

---

## 📚 Exemples

### Page complète avec état et événements

```go
type CounterPage struct {
    framework.BasePage
}

func (p *CounterPage) GetInitialState() map[string]interface{} {
    return map[string]interface{}{
        "count": 0,
        "message": "Compteur initialisé",
    }
}

func (p *CounterPage) HandleEvent(eventName string, event js.Value) {
    count := framework.GetState("count").(int)
    
    switch eventName {
    case "increment":
        framework.SetState("count", count + 1)
        framework.SetState("message", "Incrémenté!")
    case "decrement":
        framework.SetState("count", count - 1)
        framework.SetState("message", "Décrémenté!")
    case "reset":
        framework.SetState("count", 0)
        framework.SetState("message", "Remis à zéro!")
    }
}

func (p *CounterPage) Render() string {
    count := framework.GetState("count").(int)
    message := framework.GetState("message").(string)
    
    return StencilPage.Container(
        StencilUtils.Join(
            StencilText.Titre1("Compteur Interactif", "text-3xl", "font-bold", "text-center", "mb-6"),
            StencilPage.Div(
                StencilUtils.Join(
                    StencilText.Titre2(fmt.Sprintf("%d", count), "text-6xl", "font-bold", "text-blue-600", "text-center"),
                    StencilText.Paragraphe(message, "text-center", "text-gray-600", "mt-2"),
                ),
                "bg-gray-100", "p-8", "rounded-lg", "mb-6",
            ),
            StencilPage.Div(
                StencilUtils.Join(
                    StencilInteractions.Bouton("-", "decrement", "bg-red-500", "text-white", "px-6", "py-2", "rounded", "mr-2"),
                    StencilInteractions.Bouton("Reset", "reset", "bg-gray-500", "text-white", "px-6", "py-2", "rounded", "mr-2"),
                    StencilInteractions.Bouton("+", "increment", "bg-green-500", "text-white", "px-6", "py-2", "rounded"),
                ),
                "text-center",
            ),
        ),
        "max-w-md", "mx-auto", "mt-10",
    )
}
```

## 🔧 Commandes Make

Le framework inclut un Makefile complet pour automatiser le développement :

| Commande | Description |
|----------|-------------|
| `make setup` | Configuration initiale (copie wasm_exec.js, go mod tidy) |
| `make build` | Compilation WebAssembly vers `core/app.wasm` |
| `make serve` | Serveur de développement avec routage SPA |
| `make dev` | Configuration + compilation + serveur (mode complet) |
| `make clean` | Nettoyage des fichiers générés |
| `make test` | Test de compilation |
| `make info` | Informations sur le projet |
| `make help` | Aide complète |
| `make create-route ROUTE=nom` | Création d'une nouvelle route |

### Outils CLI

```bash
# Créer une route simple
make create-route blog

# Créer une route avec sous-pages
make create-route admin/dashboard

# Créer manuellement avec Go
go run core/cmd/cli.go create-route users
```

### Serveurs supportés

Le framework détecte automatiquement et utilise le meilleur serveur disponible :

1. **Python 3** avec serveur SPA personnalisé (recommandé)
2. **Python 2** avec serveur SPA personnalisé
3. **PHP** avec routeur SPA intégré
4. Fallback vers serveur basique si aucun n'est disponible

## 🐛 Dépannage

### Erreur "Module not found"

```bash
make setup
go mod tidy
```

### Erreur de chargement WebAssembly

Vérifiez que les fichiers `app.wasm` et `wasm_exec.js` sont présents dans le dossier `core/`.

### Problème de CORS

Utilisez toujours un serveur HTTP local avec `make serve`, ne pas ouvrir directement le fichier HTML.

---

## 🎯 Bonnes pratiques

### Structure des pages

- Utilisez `framework.BasePage` comme base pour toutes vos pages
- Initialisez l'état dans `GetInitialState()`
- Gérez les événements dans `HandleEvent()`
- Construisez le HTML dans `Render()`

### Gestion d'état

- Utilisez `framework.SetState()` pour modifier l'état
- Utilisez `framework.GetState()` pour lire l'état
- L'état est global et persistant entre les événements

### Composants

- Créez des composants réutilisables avec le système de props
- Utilisez les helpers `PropString()`, `PropInt()`, `PropBool()`
- Organisez vos composants dans le dossier `components/`

### Performance

- Évitez les re-rendus excessifs
- Utilisez `StencilUtils.Join()` pour concatener efficacement
- Préférez les classes Tailwind aux styles inline

---

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :

- Ouvrir une issue pour rapporter un bug
- Proposer de nouvelles fonctionnalités
- Soumettre des pull requests

### 🚀 Pour contribuer

1. Fork le projet
2. Créez une branche pour votre fonctionnalité
3. Committez vos changements
4. Testez avec `make test`
5. Ouvrez une Pull Request

---

## 📄 Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.

---

## 🔗 Liens utiles

- [Stencil-Go](https://github.com/RafaelCoppe/Stencil-Go) - Bibliothèque de templating
- [WebAssembly](https://webassembly.org/) - Documentation WebAssembly
- [Tailwind CSS](https://tailwindcss.com/) - Framework CSS
- [Go WebAssembly](https://pkg.go.dev/syscall/js) - Documentation Go WebAssembly