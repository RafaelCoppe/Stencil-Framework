package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	switch command {
	case "create-route":
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run cmd/cli.go create-route <route-path>")
			return
		}
		createRoute(os.Args[2])
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Stencil Framework CLI")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/cli.go create-route <route-path>")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("  go run cmd/cli.go create-route users")
	fmt.Println("  go run cmd/cli.go create-route admin/dashboard")
}

func createRoute(routePath string) {
	// Clean the route path
	routePath = strings.Trim(routePath, "/")
	if routePath == "" {
		fmt.Println("Error: Route path cannot be empty")
		return
	}

	// Create the directory
	dirPath := filepath.Join("app", routePath)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	// Generate package name
	packageName := filepath.Base(routePath)
	packageName = strings.ReplaceAll(packageName, "-", "")
	packageName = strings.ReplaceAll(packageName, "_", "")

	// Create page.go
	pageContent := generatePageContent(packageName, routePath)
	pageFile := filepath.Join(dirPath, "page.go")
	err = os.WriteFile(pageFile, []byte(pageContent), 0644)
	if err != nil {
		fmt.Printf("Error creating page.go: %v\n", err)
		return
	}

	// Create create.go
	createContent := generateCreateContent(packageName, routePath)
	createFile := filepath.Join(dirPath, "create.go")
	err = os.WriteFile(createFile, []byte(createContent), 0644)
	if err != nil {
		fmt.Printf("Error creating create.go: %v\n", err)
		return
	}

	// Create edit.go
	editContent := generateEditContent(packageName, routePath)
	editFile := filepath.Join(dirPath, "edit.go")
	err = os.WriteFile(editFile, []byte(editContent), 0644)
	if err != nil {
		fmt.Printf("Error creating edit.go: %v\n", err)
		return
	}

	fmt.Printf("✅ Route created successfully!\n")
	fmt.Printf("📁 Directory: %s\n", dirPath)
	fmt.Printf("📄 Files created:\n")
	fmt.Printf("  - %s\n", pageFile)
	fmt.Printf("  - %s\n", createFile)
	fmt.Printf("  - %s\n", editFile)
	fmt.Printf("\n💡 Don't forget to register the route in your main RegisterRoutes() function:\n")
	fmt.Printf("framework.RegisterPageRoute(\"/%s\",\n", routePath)
	fmt.Printf("    func() framework.PageInterface { return &%s.%sPage{} },\n", packageName, strings.Title(packageName))
	fmt.Printf("    func() framework.PageInterface { return &%s.%sCreatePage{} },\n", packageName, strings.Title(packageName))
	fmt.Printf("    func() framework.PageInterface { return &%s.%sEditPage{} },\n", packageName, strings.Title(packageName))
	fmt.Printf(")\n")
}

func generatePageContent(packageName, routePath string) string {
	title := strings.Title(strings.ReplaceAll(packageName, "-", " "))
	return fmt.Sprintf(`//go:build js && wasm

package %s

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
)

// %sPage represents the %s page
type %sPage struct {
	framework.BasePage
}

func (p *%sPage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *%sPage) Render() string {
	content := StencilUtils.Join(
		StencilText.Titre1("%s", "text-center", "text-primary", "mb-4"),
		StencilText.Paragraphe("Welcome to the %s page!", "text-center", "lead", "mb-4"),
		
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Page Actions", "mb-3"),
				StencilInteractions.Lien("/%s/create", "Create New", "btn", "btn-success", "me-2"),
				StencilInteractions.Lien("/%s/edit", "Edit", "btn", "btn-warning", "me-2"),
				StencilInteractions.Lien("/", "← Back to Home", "btn", "btn-secondary"),
			),
			"text-center", "bg-light", "p-4", "rounded",
		),
	)

	return StencilPage.Container(content, "container", "my-5")
}
`, packageName, strings.Title(packageName), routePath, strings.Title(packageName), strings.Title(packageName), strings.Title(packageName), title, routePath, routePath, routePath)
}

func generateCreateContent(packageName, routePath string) string {
	title := strings.Title(strings.ReplaceAll(packageName, "-", " "))
	return fmt.Sprintf(`//go:build js && wasm

package %s

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
)

// %sCreatePage represents the create page for %s
type %sCreatePage struct {
	framework.BasePage
}

func (p *%sCreatePage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *%sCreatePage) Render() string {
	content := StencilUtils.Join(
		StencilText.Titre1("Create New %s", "text-center", "text-success", "mb-4"),
		StencilText.Paragraphe("Create a new item in %s", "text-center", "lead", "mb-4"),
		
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Create Form", "mb-3"),
				StencilText.Paragraphe("Add your form components here.", "mb-3"),
			),
			"bg-success", "bg-opacity-10", "p-4", "rounded", "mb-4",
		),
		
		StencilPage.Div(
			StencilUtils.Join(
				StencilInteractions.Lien("/%s", "← Back to %s", "btn", "btn-secondary", "me-2"),
				StencilInteractions.Lien("/", "Home", "btn", "btn-primary"),
			),
			"text-center",
		),
	)

	return StencilPage.Container(content, "container", "my-5")
}
`, packageName, strings.Title(packageName), routePath, strings.Title(packageName), strings.Title(packageName), strings.Title(packageName), title, routePath, routePath, title)
}

func generateEditContent(packageName, routePath string) string {
	title := strings.Title(strings.ReplaceAll(packageName, "-", " "))
	return fmt.Sprintf(`//go:build js && wasm

package %s

import (
	"github.com/RafaelCoppe/Stencil-Framework/framework"
	StencilPage "github.com/RafaelCoppe/Stencil-Go/pkg/page"
	StencilText "github.com/RafaelCoppe/Stencil-Go/pkg/text"
	StencilUtils "github.com/RafaelCoppe/Stencil-Go/pkg/utils"
	StencilInteractions "github.com/RafaelCoppe/Stencil-Go/pkg/interactions"
)

// %sEditPage represents the edit page for %s
type %sEditPage struct {
	framework.BasePage
}

func (p *%sEditPage) GetInitialState() map[string]interface{} {
	return map[string]interface{}{}
}

func (p *%sEditPage) Render() string {
	content := StencilUtils.Join(
		StencilText.Titre1("Edit %s", "text-center", "text-warning", "mb-4"),
		StencilText.Paragraphe("Edit an existing item in %s", "text-center", "lead", "mb-4"),
		
		StencilPage.Div(
			StencilUtils.Join(
				StencilText.Titre2("Edit Form", "mb-3"),
				StencilText.Paragraphe("Add your form components here.", "mb-3"),
			),
			"bg-warning", "bg-opacity-10", "p-4", "rounded", "mb-4",
		),
		
		StencilPage.Div(
			StencilUtils.Join(
				StencilInteractions.Lien("/%s", "← Back to %s", "btn", "btn-secondary", "me-2"),
				StencilInteractions.Lien("/", "Home", "btn", "btn-primary"),
			),
			"text-center",
		),
	)

	return StencilPage.Container(content, "container", "my-5")
}
`, packageName, strings.Title(packageName), routePath, strings.Title(packageName), strings.Title(packageName), strings.Title(packageName), title, routePath, routePath, title)
}
