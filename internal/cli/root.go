package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "stencil",
	Short: "Stencil Framework CLI tool",
	Long:  `A CLI tool for the Stencil Framework to help create and manage WebAssembly applications.`,
}

var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new Stencil Framework project",
	Long:  `Create a new Stencil Framework project with all necessary files including index.html, Makefile, wasm_exec.js, and app folder structure.`,
	Args:  cobra.MaximumNArgs(1),
	RunE:  runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func Execute() error {
	return rootCmd.Execute()
}

func runInit(cmd *cobra.Command, args []string) error {
	var projectName string
	var projectDir string

	if len(args) > 0 {
		projectName = args[0]
		projectDir = projectName
	} else {
		// Use current directory
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %v", err)
		}
		projectName = filepath.Base(cwd)
		projectDir = "."
	}

	fmt.Printf("Initializing Stencil Framework project '%s'...\n", projectName)

	// Create project directory if it doesn't exist
	if projectDir != "." {
		if err := os.MkdirAll(projectDir, 0755); err != nil {
			return fmt.Errorf("failed to create project directory: %v", err)
		}
	}

	// Create app directory
	appDir := filepath.Join(projectDir, "app")
	if err := os.MkdirAll(appDir, 0755); err != nil {
		return fmt.Errorf("failed to create app directory: %v", err)
	}

	// Create files
	if err := createIndexHTML(projectDir, projectName); err != nil {
		return err
	}

	if err := createMakefile(projectDir); err != nil {
		return err
	}

	if err := createAppGoFile(appDir, projectName); err != nil {
		return err
	}

	if err := createGoMod(projectDir, projectName); err != nil {
		return err
	}

	// Copy wasm_exec.js (will be handled by make setup)
	fmt.Println("✅ Project initialized successfully!")
	fmt.Println("\nNext steps:")
	if projectDir != "." {
		fmt.Printf("  cd %s\n", projectName)
	}
	fmt.Println("  make setup    # Setup the project (downloads wasm_exec.js)")
	fmt.Println("  make build    # Build the WebAssembly application")
	fmt.Println("  make serve    # Start the development server")

	return nil
}
