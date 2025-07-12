# Next.js-Style Routing for Stencil Framework

The Stencil Framework now supports Next.js-style file-based routing, making it easy to organize your application with a familiar pattern.

## Directory Structure

```text
app/
├── page.go          # Root page (/)
├── about/
│   ├── page.go      # About page (/about)
│   ├── create.go    # Create page (/about/create)
│   └── edit.go      # Edit page (/about/edit)
└── other-route/
    ├── page.go      # Other route page (/other-route)
    ├── create.go    # Create page (/other-route/create)
    └── edit.go      # Edit page (/other-route/edit)
```

## File Types

### page.go

The main page for a route. This is the default page that loads when someone visits the route.

### create.go

A specialized page for creating new items. Accessible at `{route}/create`.

### edit.go

A specialized page for editing existing items. Accessible at `{route}/edit`.

## How to Use

### 1. Create Your Page Structure

Create your pages in the `app/` directory following the structure above. Each page should be in its own package:

```go
//go:build js && wasm

package about

import (
    "github.com/RafaelCoppe/Stencil-Framework/framework"
    // ... other imports
)

type AboutPage struct {
    framework.BasePage
}

func (p *AboutPage) Render() string {
    // Your page content here
    return "Hello from About page!"
}
```

### 2. Register Your Routes

In your main `app/page.go` file, register all your routes in the `RegisterRoutes()` function:

```go
func RegisterRoutes() {
    // Register root page
    framework.RegisterPageRoute("/", 
        func() framework.PageInterface { return &WelcomePage{} },
        nil, // no create.go
        nil, // no edit.go
    )
    
    // Register about page with all three variants
    framework.RegisterPageRoute("/about", 
        func() framework.PageInterface { return &about.AboutPage{} },
        func() framework.PageInterface { return &about.CreatePage{} },
        func() framework.PageInterface { return &about.EditPage{} },
    )
}
```

### 3. Start the Application

In your `main.go`:

```go
func main() {
    // Register all routes
    app.RegisterRoutes()
    
    // Start the application with router
    framework.RunWithRouter("app")
}
```

## Navigation

### Creating Links

Use the Stencil-Go link component to create navigation links:

```go
StencilInteractions.Lien("/about", "About Page", "btn", "btn-primary")
StencilInteractions.Lien("/about/create", "Create Something", "btn", "btn-success")
```

### Programmatic Navigation

You can also navigate programmatically:

```go
framework.NavigateTo("/about/edit")
```

## Router Features

- **Client-side navigation**: Navigation happens without page reloads
- **Browser history support**: Back/forward buttons work correctly
- **URL updates**: The browser URL updates to reflect the current page
- **404 handling**: Automatic 404 page for non-existent routes
- **Nested routes**: Support for deeply nested route structures

## Route Registration

The `RegisterPageRoute` function accepts:

1. **Base path**: The main route path (e.g., "/about")
2. **Page handler**: Function returning the main page component
3. **Create handler**: Function returning the create page component (optional)
4. **Edit handler**: Function returning the edit page component (optional)

Pass `nil` for handlers you don't need.

## Examples

See the example implementation in:

- `app/page.go` - Root page with navigation
- `app/about/page.go` - About page
- `app/about/create.go` - Create page
- `app/about/edit.go` - Edit page

## Building and Running

1. Build: `make build`
2. Serve: `make serve`
3. Visit: `http://localhost:8080`

The routing system will automatically handle navigation between pages based on the URL structure you've defined.

## CLI Tool for Route Generation

You can use the included CLI tool to quickly generate new routes:

```bash
# Create a new route with all three files (page, create, edit)
go run cmd/cli.go create-route users

# Create a nested route
go run cmd/cli.go create-route admin/dashboard
```

This will automatically generate:

- The directory structure
- `page.go`, `create.go`, and `edit.go` files with boilerplate code
- Instructions on how to register the route

After running the CLI tool, you'll need to manually add the route registration to your `RegisterRoutes()` function as shown in the CLI output.
