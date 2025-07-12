# Stencil WebAssembly Framework

A React/Vue-like framework for building WebAssembly applications in Go using the Stencil-Go library.

## Framework Architecture

The framework now provides a component-based architecture similar to React or Vue.js:

### Key Features

1. **Component-Based**: Create reusable UI components
2. **State Management**: Reactive state with automatic re-rendering  
3. **Event Handling**: Simple event registration and handling
4. **Declarative UI**: Build UI with simple function calls
5. **Page-Based Routing**: Organize your app into pages

### Framework Structure

```
framework/
├── app.go          # Main application framework
├── component.go    # Component base classes
├── state.go        # State management
├── ui.go          # UI helper functions
└── go.mod         # Framework dependencies
```

## How to Use the Framework

### 1. Simple Usage (Like React/Vue)

Create a single Go file for your application:

```go
package main

import (
    "fmt"
    "syscall/js"
    
    "github.com/RafaelCoppe/Stencil-Framework/framework"
)

// Define your page component
type HomePage struct {
    *framework.Page
}

func NewHomePage(app *framework.App) *HomePage {
    page := &HomePage{
        Page: framework.NewPage(app, "Home"),
    }
    
    // Initialize state
    page.State.Set("count", 0)
    
    // Register event handlers
    framework.RegisterEventHandler("increment", page.handleIncrement)
    
    return page
}

func (h *HomePage) handleIncrement(eventName string, event js.Value, app *framework.App) {
    count := h.State.GetInt("count")
    h.State.Set("count", count+1)
}

func (h *HomePage) Render() string {
    count := h.State.GetInt("count")
    
    return framework.Container(framework.Join(
        framework.H1("My App", "text-2xl", "font-bold", "mb-4"),
        framework.P(fmt.Sprintf("Count: %d", count), "text-lg", "mb-4"),
        framework.Button("Increment", "increment", "bg-blue-500", "text-white", "px-4", "py-2", "rounded"),
    ), "max-w-md", "mx-auto", "p-8")
}

func main() {
    app := framework.NewApp("app")
    homePage := NewHomePage(app)
    app.SetRootComponent(homePage)
    app.Start()
}
```

### 2. Available UI Components

```go
// Layout
framework.Container(content, classes...)
framework.Div(content, classes...)

// Text  
framework.H1(text, classes...)
framework.H2(text, classes...)
framework.P(text, classes...)
framework.Span(text, classes...)

// Interactive
framework.Button(label, event, classes...)
framework.TextInput(name, placeholder, value, event, classes...)
framework.Link(href, text, classes...)

// Utilities
framework.Hr(classes...)
framework.Join(parts...)  // Combine multiple elements
```

### 3. State Management

```go
// Set state (triggers re-render)
page.State.Set("key", value)

// Get state
value := page.State.Get("key")
text := page.State.GetString("key")
number := page.State.GetInt("key")
flag := page.State.GetBool("key")

// Check if state exists
exists := page.State.Has("key")
```

### 4. Event Handling

```go
// Register an event handler
framework.RegisterEventHandler("myEvent", func(eventName string, event js.Value, app *framework.App) {
    // Handle the event
    // Update state
    // State changes automatically trigger re-render
})

// Create a button that triggers the event
framework.Button("Click me", "myEvent", "btn", "btn-primary")
```

## Comparison with React/Vue

### React Style
```jsx
const [count, setCount] = useState(0);

return (
  <div className="container">
    <h1>My App</h1>
    <p>Count: {count}</p>
    <button onClick={() => setCount(count + 1)}>Increment</button>
  </div>
);
```

### Our Framework Style
```go
count := h.State.GetInt("count")

return framework.Container(framework.Join(
    framework.H1("My App"),
    framework.P(fmt.Sprintf("Count: %d", count)),
    framework.Button("Increment", "increment"),
), "container")
```

## Building and Running

### For Development
```bash
# Setup
make setup

# Build WASM
GOOS=js GOARCH=wasm go build -o app.wasm your-app.go

# Serve  
make serve
```

### Project Structure
```
your-project/
├── main.go           # Your application code
├── go.mod           # Dependencies
├── index.html       # HTML page
├── app.wasm        # Generated WebAssembly
└── wasm_exec.js    # WebAssembly runtime
```

## Advantages Over the Original

1. **Separation of Concerns**: Framework code is separate from application code
2. **Reusable**: Import the framework into any project
3. **Component-Based**: Build reusable UI components  
4. **Reactive**: Automatic re-rendering when state changes
5. **Simple API**: Clean, intuitive API similar to React/Vue
6. **Type Safe**: Full Go type safety
7. **Event System**: Easy event handling and registration

## Next Steps

To make this production-ready, we could add:

1. **Router**: Client-side routing between pages
2. **Lifecycle Methods**: componentDidMount, componentWillUnmount equivalent
3. **Props**: Pass data between components
4. **Context**: Global state management
5. **Hooks**: Reusable stateful logic
6. **Hot Reload**: Development experience improvements
7. **Testing**: Unit testing framework for components

This framework now provides a React/Vue-like development experience for WebAssembly applications in Go!
