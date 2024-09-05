package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sevenreup/panga/src/cmd/ui/list"
	"github.com/sevenreup/panga/src/pkg/templates"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all available templates",
	Run: func(cmd *cobra.Command, args []string) {
		var tprogram = tea.NewProgram(list.CreateList(templates.FetchTemplates()))
		if _, err := tprogram.Run(); err != nil {
			fmt.Println("Error running program:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
