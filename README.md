# Framework WebAssembly avec Stencil-Go

Un framework moderne pour créer des applications web interactives en Go en utilisant WebAssembly et la bibliothèque [Stencil-Go](https://github.com/RafaelCoppe/Stencil-Go).

## 🚀 Fonctionnalités

- **WebAssembly** : Exécution native du code Go dans le navigateur
- **Stencil-Go** : Génération déclarative de HTML avec une syntaxe fluide
- **Réactivité** : Gestion d'état et re-rendu automatique
- **Événements** : Système d'événements personnalisés
- **Routing** : Système de routage inspiré de Next.js avec support de routes basées sur fichiers
- **Styling** : Intégration Tailwind CSS
- **Performance** : Rendu optimisé avec DOM virtuel

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
├── main.go         # Point d'entrée principal avec routage
├── app/
│   └── page.go     # Page d'accueil (/)
├── framework/      # Code du framework
├── go.mod          # Dépendances Go
├── index.html      # Page HTML d'entrée
├── Makefile        # Automatisation des tâches
├── README.md       # Documentation
├── ROUTING.md      # Documentation du routage
└── wasm_exec.js    # Runtime WebAssembly Go
```

## 🛠 Installation et configuration

### Prérequis

- Go 1.21 ou supérieur
- Git
- Un serveur HTTP (Python, PHP, ou Node.js)

### Installation

1. **Clonez ou créez le projet** :
```bash
mkdir stencil-wasm
cd stencil-wasm
```

2. **Créez les fichiers** en copiant le contenu des artefacts fournis

3. **Configuration automatique** :
```bash
make setup
```

Cette commande :
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

### Structure de base

```go
package main

import (
    "stencil"
    "syscall/js"
)

type App struct {
    container js.Value
    state     map[string]interface{}
}

func (a *App) buildHTML() string {
    return stencil.Container([]string{
        stencil.Titre1("Mon App", []string{"text-2xl", "font-bold"}),
        stencil.Paragraphe("Contenu de l'application", nil),
    }, []string{"max-w-4xl", "mx-auto", "p-8"})
}
```

### Gestion d'état

```go
// Définir l'état
app.SetState("count", 0)
app.SetState("text", "Hello World")

// Récupérer l'état
count := app.GetState("count").(int)
text := app.GetState("text").(string)
```

### Événements personnalisés

```go
// Bouton avec événement
BoutonAvecEvent("Cliquer", "monEvenement", []string{"btn", "btn-primary"})

// Gestionnaire d'événement
func (a *App) handleEvent(eventName string, event js.Value) {
    switch eventName {
    case "monEvenement":
        // Logique de l'événement
        a.SetState("clicked", true)
        a.update()
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

## 📚 Exemples

### Compteur simple
```go
func (a *App) buildCounter() string {
    count := a.GetState("count").(int)
    
    return stencil.Section([]string{
        stencil.Titre2("Compteur", []string{"text-xl", "mb-4"}),
        stencil.Paragraphe(fmt.Sprintf("Valeur : %d", count), []string{"text-lg", "mb-4"}),
        stencil.HorizontalBar([]string{
            BoutonAvecEvent("Incrémenter", "increment", []string{"bg-blue-500", "text-white", "px-4", "py-2", "rounded"}),
            BoutonAvecEvent("Décrémenter", "decrement", []string{"bg-red-500", "text-white", "px-4", "py-2", "rounded"}),
        }, []string{"gap-4"}),
    }, []string{"p-6", "bg-gray-100", "rounded-lg"})
}
```

### Formulaire
```go
func (a *App) buildForm() string {
    return stencil.Form([]string{
        stencil.InputText("name", "Votre nom", "", []string{"border", "p-2", "rounded", "mb-4"}),
        stencil.TextArea("message", "Votre message", "", []string{"border", "p-2", "rounded", "mb-4"}),
        stencil.Bouton("Envoyer", "submitForm()", []string{"bg-green-500", "text-white", "px-4", "py-2", "rounded"}),
    }, "/submit", "POST", []string{"max-w-md"})
}
```

## 🔧 Commandes Make

- `make setup` - Configuration initiale
- `make build` - Compilation WebAssembly
- `make serve` - Serveur de développement
- `make dev` - Configuration + compilation + serveur
- `make clean` - Nettoyage des fichiers générés
- `make test` - Test de compilation
- `make info` - Informations sur le projet
- `make help` - Aide

## 🐛 Dépannage

### Erreur "Module not found"
```bash
make setup
go mod tidy
```

### Erreur de chargement WebAssembly
Vérifiez que les fichiers `app.wasm` et `wasm_exec.js` sont présents dans le même dossier que `index.html`.

### Problème de CORS
Utilisez toujours un serveur HTTP local, ne pas ouvrir directement le fichier HTML.

## 🤝 Contribution

Les contributions sont les bienvenues ! N'hésitez pas à :
- Ouvrir une issue pour rapporter un bug
- Proposer de nouvelles fonctionnalités
- Soumettre des pull requests

## 📄 Licence

Ce projet est sous licence MIT. Voir le fichier LICENSE pour plus de détails.

## 🔗 Liens utiles

- [Stencil-Go](https://github.com/RafaelCoppe/Stencil-Go) - Bibliothèque de templating
- [WebAssembly](https://webassembly.org/) - Documentation WebAssembly
- [Tailwind CSS](https://tailwindcss.com/) - Framework CSS
- [Go WebAssembly](https://pkg.go.dev/syscall/js) - Documentation Go WebAssembly