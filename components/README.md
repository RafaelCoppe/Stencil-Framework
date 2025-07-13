# Stencil Components

Un système de composants basé sur des fonctions pour le framework Stencil.

## Installation

Le module components fait partie du framework Stencil. Pour l'utiliser dans vos pages :

```go
import (
    "github.com/RafaelCoppe/Stencil-Framework/components"
)
```

## Concept

Les composants Stencil sont des fonctions qui acceptent des propriétés (`ComponentProps`) et retournent du HTML sous forme de string. Cela permet de créer des éléments d'interface réutilisables et configurables.

## Utilisation de base

### Créer un composant simple

```go
func MyButton(props components.ComponentProps) string {
    text := components.PropString(props, "text", "Cliquer")
    class := components.PropString(props, "class", "btn")
    
    return fmt.Sprintf(`<button class="%s">%s</button>`, class, text)
}
```

### Utiliser un composant

```go
button := components.Button(components.ComponentProps{
    "text": "Mon bouton",
    "href": "/page",
    "class": "btn btn-primary",
})
```

## Composants disponibles

### Composants de base (common.go)

- **Button** : Bouton/lien cliquable
- **Card** : Carte avec titre et contenu
- **Header** : En-tête avec navigation
- **InputField** : Champ de saisie avec label
- **Alert** : Message d'alerte
- **Modal** : Fenêtre modale
- **List** : Liste ordonnée ou non ordonnée

### Composants de formulaire (forms.go)

- **ContactForm** : Formulaire de contact complet
- **LoginForm** : Formulaire de connexion
- **SearchForm** : Formulaire de recherche

### Exemples (examples.go)

- **ExamplePage** : Page de démonstration
- **SimpleForm** : Exemple de formulaire
- **LoginPage** : Page de connexion

## Propriétés (Props)

Les composants utilisent un système de propriétés flexible avec des fonctions d'aide :

- `PropString(props, "key", "defaultValue")` : Récupère une string
- `PropInt(props, "key", defaultValue)` : Récupère un int
- `PropBool(props, "key", defaultValue)` : Récupère un bool
- `PropStringSlice(props, "key", defaultSlice)` : Récupère un []string

## Utilitaires (utils.go)

- `MergeProps(props...)` : Fusionne plusieurs ComponentProps
- `ConditionalClass(condition, trueClass, falseClass)` : Classe conditionnelle
- `BuildClassString(classes...)` : Construit une chaîne de classes CSS
- `WithAttributes(element, attributes)` : Ajoute des attributs HTML
- `RenderIf(condition, component)` : Rendu conditionnel
- `RenderList(items)` : Rend une liste de composants
- `WithID(component, id)` : Ajoute un ID
- `WithDataAttributes(component, dataAttrs)` : Ajoute des attributs data-*

## Exemple complet

```go
package main

import (
    "github.com/RafaelCoppe/Stencil-Framework/components"
    StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
)

func MyPage() string {
    // En-tête
    header := components.Header(components.ComponentProps{
        "title": "Mon Site",
        "links": []string{"Accueil", "À propos", "Contact"},
    })
    
    // Contenu principal
    content := components.Card(components.ComponentProps{
        "title": "Bienvenue",
        "content": "Ceci est ma page d'accueil.",
    })
    
    // Formulaire
    form := components.ContactForm(components.ComponentProps{
        "title": "Nous contacter",
        "action": "/contact",
    })
    
    // Assemblage
    page := components.RenderList([]string{
        header,
        content,
        form,
    })
    
    return StencilPage.Container(page)
}
```

## Créer ses propres composants

1. Créez un fichier dans le dossier `components/`
2. Définissez vos fonctions de composants
3. Utilisez `ComponentProps` pour les propriétés
4. Retournez du HTML sous forme de string
5. Importez et utilisez vos composants dans vos pages

Exemple de composant personnalisé :

```go
//go:build js && wasm

package components

import (
    StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
    StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
)

// ProductCard crée une carte produit
func ProductCard(props ComponentProps) string {
    name := PropString(props, "name", "Produit")
    price := PropString(props, "price", "0€")
    image := PropString(props, "image", "")
    description := PropString(props, "description", "")
    
    content := ""
    if image != "" {
        content += fmt.Sprintf(`<img src="%s" alt="%s" class="product-image">`, image, name)
    }
    content += StencilText.Titre3(name)
    content += StencilPage.Div(description, "product-description")
    content += StencilPage.Div(price, "product-price")
    
    return StencilPage.Div(content, "product-card")
}
```

Ce système vous permet de créer des composants réutilisables et maintenables pour votre application Stencil !
