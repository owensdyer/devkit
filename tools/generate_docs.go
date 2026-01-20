/*
	Automatically generate documentation for the commands
*/
package main

import (
	"log"
	"github.com/owensdyer/devkit/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	// Output directory
	docsDir := "./docs"

	// Generate Markdown for all commands
	err := doc.GenMarkdownTree(cmd.RootCmd, docsDir)
	if err != nil {
		log.Fatalf("Error generating Markdown docs: %v", err)
	}

	log.Printf("Docs generated successfully in %s", docsDir)
}