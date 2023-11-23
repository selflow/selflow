package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
	"strings"
)

func NewGenDocsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gen-docs",
		Short: "Generate the documentation [Development purpose]",
		Run: func(cmd *cobra.Command, args []string) {
			if err := generateDocs(); err != nil {
				if _, err = fmt.Fprintf(os.Stderr, "%v\n", err); err != nil {
					panic(err)
				}
				os.Exit(1)
			}
		},
	}
}

const fileHeadingFormat = "---\nslug: %s\ntitle: \"⌨ %s\"\n---\n\n# ⌨️ `%s`\n\n"

func generateDocs() error {
	return doc.GenMarkdownTreeCustom(rootCmd, "./docs/ecosystem/cli", func(filepath string) string {
		splitFilename := strings.Split(filepath, "/")
		filename := splitFilename[len(splitFilename)-1]
		filename = strings.TrimPrefix(filename, "selflow_")
		filename = strings.TrimSuffix(filename, ".md")

		slug := strings.ReplaceAll(filename, "_", "/")

		return fmt.Sprintf(fileHeadingFormat, slug, filename, filename)
	},
		func(s string) string {
			return s
		})
}
