package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ytctl",
	Short: "ytctl downloads files",
}
