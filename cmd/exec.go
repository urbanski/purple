package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/urbanski/purple/purple"

	"github.com/spf13/cobra"
)

var inputFile string
var debugFlag bool

func init() {
	execCmd.Flags().StringVarP(&inputFile, "file", "f", "", "input file")
	execCmd.Flags().BoolVar(&debugFlag, "v", false, "verbose")
	rootCmd.AddCommand(execCmd)
}

var execCmd = &cobra.Command{
	Use:   "test",
	Short: "Execute a test",
	Long:  `Execute a security test`,
	Run: func(cmd *cobra.Command, args []string) {
		if debugFlag == true {
			log.SetLevel(log.DebugLevel)
		}
		if inputFile == "" {
			log.Fatal("You must specify an input file with `-f` eg: `test -f mytest.yaml`")
		} else {
			purple.ExecTest(inputFile)
		}
	},
}

