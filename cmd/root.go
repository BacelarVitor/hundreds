/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "hundreds",
	Short: "List the counting commited days of the challenge between me and Matheus",
	Long: `List how many days me and Matheus has committed and how much days are left for the 100th since starting the #100DaysOfCode challenge .`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Display more verbose output in console output. (default: false)")
	// viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// rootCmd.PersistentFlags().BoolVarP(&Debug, "debug", "d", false, "Display debugging output in the console. (default: false)")
	// viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


