package cmd

import (
	"fmt"

	"github.com/sevenreup/panga/src/pkg/engine"
	"github.com/sevenreup/panga/src/pkg/templates"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new scaffold",
	Long:  `Create a new scaffold`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a template name")
			return
		}
		templateName := args[0]
		tem := templates.GetTemplate(templateName)
		if tem == nil {
			fmt.Println("Template not found")
			return
		}
		eng := engine.NewEngine(*tem)
		eng.Run()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
