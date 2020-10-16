package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/svang/api/internal/command"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate the sample yaml file",
	Long: `Generate the sample yaml file and provide the values
which will be used inside the application to set the
email user, password, service account and firebase database details.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("init called")

		err := command.New(cmd.Flags()).
			CreateSampleConfigFile().
			TimeTaken().
			Error

		if err != nil {
			return err
		}
		return nil

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
