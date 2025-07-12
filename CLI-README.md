# Stencil Framework CLI

A command-line interface tool for the Stencil Framework that helps you create and manage WebAssembly applications with Go.

## Installation

1. Clone the Stencil-Framework repository
2. Build the CLI tool:
   ```bash
   cd Stencil-Framework
   go build -o stencil ./cmd/stencil
   ```
3. (Optional) Add the `stencil` binary to your PATH for global access

## Usage

### Initialize a New Project

To create a new Stencil Framework project:

```bash
# Initialize in a new directory
stencil init my-project

# Or initialize in the current directory
stencil init
```

This command creates:
- `index.html` - The main HTML file with WebAssembly loading setup
- `Makefile` - Build and development commands
- `go.mod` - Go module file with necessary dependencies
- `app/` directory containing:
  - `app.go` - Main application file with a sample counter component

### Project Setup and Development

After initializing a project, follow these steps:

1. **Setup the project** (downloads `wasm_exec.js` and installs dependencies):
   ```bash
   make setup
   ```

2. **Build the WebAssembly application**:
   ```bash
   make build
   ```

3. **Start the development server**:
   ```bash
   make serve
   ```

The application will be available at `http://localhost:8080`

### Available Make Commands

- `make setup` - Initial project setup (downloads wasm_exec.js, installs dependencies)
- `make build` - Compile the WebAssembly application
- `make serve` - Start the development server on port 8080
- `make dev` - Development mode with automatic rebuilding on file changes
- `make clean` - Remove generated files
- `make test` - Run tests
- `make help` - Show all available commands

## Project Structure

After initialization, your project will have this structure:

```
my-project/
├── app/
│   └── app.go          # Main application code
├── index.html          # HTML page that loads the WASM app
├── Makefile           # Build and development commands
├── go.mod             # Go module with dependencies
├── wasm_exec.js       # Go WASM runtime (created by make setup)
└── app.wasm           # Compiled WebAssembly binary (created by make build)
```

## Example Application

The generated `app/app.go` includes a simple counter application that demonstrates:
- State management with the Stencil Framework
- Event handling for user interactions
- Component rendering with Stencil-Go UI components

## Dependencies

The generated projects automatically include:
- `github.com/RafaelCoppe/Stencil-Framework/framework` - The main framework
- `github.com/RafaelCoppe/Stencil-Go` - UI components and utilities

## Development Workflow

1. Initialize project: `stencil init my-app`
2. Navigate to project: `cd my-app`
3. Setup: `make setup`
4. Build: `make build`
5. Serve: `make serve`
6. Open browser to `http://localhost:8080`
7. Edit `app/app.go` and rebuild as needed

For continuous development, use `make dev` which will automatically rebuild when Go files change.
