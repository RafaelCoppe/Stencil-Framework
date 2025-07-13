//go:build js && wasm

package components

import (
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
)

// ComponentProps représente les propriétés d'un composant
type ComponentProps map[string]interface{}

// PropString récupère une propriété string avec une valeur par défaut
func PropString(props ComponentProps, key string, defaultValue string) string {
	if val, exists := props[key]; exists {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return defaultValue
}

// FeatureCard crée une carte de fonctionnalité avec icône, titre et description
func FeatureCard(props ComponentProps) string {
	icon := PropString(props, "icon", "⭐")
	title := PropString(props, "title", "Fonctionnalité")
	description := PropString(props, "description", "Description de la fonctionnalité")

	return StencilPage.Div(
		StencilUtils.Join(
			StencilPage.Div(icon, "text-4xl", "mb-4"),
			StencilText.Titre3(title, "text-xl", "font-semibold", "text-gray-800", "mb-2"),
			StencilText.Paragraphe(description, "text-gray-600"),
		),
		"text-center", "p-6", "bg-gradient-to-br", "from-blue-50", "to-blue-100", "rounded-xl",
	)
}

// StepItem crée un élément d'étape avec numéro, titre et commande
func StepItem(props ComponentProps) string {
	number := PropString(props, "number", "1")
	title := PropString(props, "title", "Étape")
	command := PropString(props, "command", "commande")

	return StencilPage.Div(
		StencilUtils.Join(
			StencilPage.Div(number, "bg-blue-100", "text-blue-600", "rounded-full", "w-8", "h-8", "flex", "items-center", "justify-center", "font-bold", "text-sm"),
			StencilPage.Div(
				StencilUtils.Join(
					StencilText.Paragraphe(title, "font-semibold", "text-gray-800", "mb-2"),
					`<code class="bg-gray-100 px-3 py-2 rounded text-sm">`+command+`</code>`,
				),
			),
		),
		"flex", "items-start", "space-x-4",
	)
}

// HeroSection crée la section héro avec titre, sous-titre et boutons d'action
func HeroSection(props ComponentProps) string {
	title := PropString(props, "title", "Mon Application")
	subtitle := PropString(props, "subtitle", "Description de l'application")
	primaryButtonText := PropString(props, "primaryButtonText", "Commencer")
	primaryButtonHref := PropString(props, "primaryButtonHref", "#")
	secondaryButtonText := PropString(props, "secondaryButtonText", "En savoir plus")
	secondaryButtonHref := PropString(props, "secondaryButtonHref", "#")

	return StencilPage.Div(
		StencilUtils.Join(
			StencilText.Titre1(title, "text-5xl", "font-bold", "text-white", "mb-4"),
			StencilText.Paragraphe(subtitle, "text-xl", "text-blue-100", "mb-8"),

			// Boutons d'action
			StencilPage.Div(
				StencilUtils.Join(
					StencilInteractions.Lien(primaryButtonHref, primaryButtonText, "bg-white", "text-blue-600", "px-8", "py-3", "rounded-lg", "font-semibold", "hover:bg-blue-50", "transition-all", "transform", "hover:-translate-y-1", "shadow-lg", "inline-block", "text-center", "no-underline"),
					StencilInteractions.Lien(secondaryButtonHref, secondaryButtonText, "bg-blue-700", "text-white", "px-8", "py-3", "rounded-lg", "font-semibold", "hover:bg-blue-800", "transition-all", "transform", "hover:-translate-y-1", "shadow-lg", "inline-block", "text-center", "no-underline"),
				),
				"flex", "flex-wrap", "justify-center", "gap-4", "mb-8",
			),
		),
		"text-center", "mb-12",
	)
}
