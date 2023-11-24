package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"os"
	"path"
	"strings"
)

func NewGenDocsCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gen-docs",
		Short: "Generate the documentation [Development purpose]",
		Long: "Generate the documentation [Development purpose]\n\n" +
			"The generated documentations files are placed in the `docs/ecosystem/cli` directory.\n" +
			"To execute the command and get the right format, you can use the NX target `gen-docs` with this command:\n" +
			"```bash\n" +
			"nx run selflow-cli:gen-docs\n" +
			"```",
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

const fileHeadingFormat = "---\nslug: %s\ntitle: '⌨ %s'\n---\n\n# ⌨️ `%s`\n\n"

func generateDocs() error {
	docsSourcePath := path.Join(".", "docs", "ecosystem", "cli")

	return doc.GenMarkdownTreeCustom(rootCmd, docsSourcePath, func(filepath string) string {
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
