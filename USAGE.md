# Stencil Framework - Simplified Usage

The Stencil Framework has been redesigned to work like traditional web frameworks. Users simply import the framework package, create their page logic, and export HTML without any framework-specific boilerplate.

## How to Use

### 1. Import the Framework

```go
import "github.com/RafaelCoppe/Stencil-Framework/framework"
```

### 2. Create Your Page

Create a struct that implements the `PageInterface`:

```go
type MyPage struct {
    framework.BasePage // Embed for default implementations
}
```

### 3. Implement Required Methods

#### GetInitialState() - Set up initial state
```go
func (p *MyPage) GetInitialState() map[string]interface{} {
    return map[string]interface{}{
        "count": 0,
        "text":  "",
    }
}
```

#### HandleEvent() - Handle user interactions
```go
func (p *MyPage) HandleEvent(eventName string, event js.Value) {
    switch eventName {
    case "increment":
        count := framework.GetStateInt("count")
        framework.SetState("count", count+1)
    case "updateText":
        target := event.Get("target")
        value := target.Get("value").String()
        framework.SetState("text", value)
    }
}
```

#### Render() - Generate HTML
```go
func (p *MyPage) Render() string {
    count := framework.GetStateInt("count")
    text := framework.GetStateString("text")
    
    return framework.Join(
        framework.Container(
            framework.H1("My App"),
            framework.P(fmt.Sprintf("Count: %d", count)),
            framework.Button("Increment", "increment", "btn", "btn-primary"),
            framework.TextInput("input", "Type something...", text, "updateText", "form-control"),
        ),
    )
}
```

### 4. Start the Framework

In your main function:

```go
func main() {
    page := &MyPage{}
    framework.Run(page, "app") // "app" is the HTML element ID
}
```

## Available Components

The framework provides convenient UI components:

- `Container()` - Page container
- `Div()` - Generic div
- `H1()`, `H2()` - Headers
- `P()` - Paragraphs
- `Span()` - Inline text
- `Button(label, event, ...classes)` - Interactive buttons
- `TextInput(name, placeholder, value, event, ...classes)` - Text inputs
- `Link(href, text, ...classes)` - Links
- `Hr()` - Horizontal rules
- `Join()` - Combine multiple HTML strings

## State Management

The framework provides global state management:

- `SetState(key, value)` - Update state (triggers re-render)
- `GetState(key)` - Get raw state value
- `GetStateString(key)` - Get state as string
- `GetStateInt(key)` - Get state as integer
- `GetStateBool(key)` - Get state as boolean

## Build Instructions

1. Create your page file (e.g., `app.go`) with the structure above
2. Build for WebAssembly:
   ```bash
   GOOS=js GOARCH=wasm go build -o app.wasm .
   ```
3. Include the WASM file in your HTML page with the appropriate JavaScript glue

## Key Benefits

- **No Framework Boilerplate**: Just business logic and HTML generation
- **Automatic Event Handling**: Framework handles DOM events automatically
- **Reactive State**: State changes trigger automatic re-renders
- **Type Safety**: Full Go type checking
- **Clean Separation**: Framework concerns are completely hidden from user code

## Example

See `app.go` for a complete working example with counters, text inputs, and various UI components.
