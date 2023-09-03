package kor

import (
	"github.com/spf13/cobra"
	"github.com/yonahd/kor/pkg/kor"
)

var stsCmd = &cobra.Command{
	Use:     "statefulsets",
	Aliases: []string{"sts"},
	Short:   "Gets unused statefulsets",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if outputFormat == "json" {
			kor.GetUnusedStatefulsetsJSON(includeExcludeLists, kubeconfig)
		} else {
			kor.GetUnusedStatefulsets(includeExcludeLists, kubeconfig)
		}

	},
}

func init() {
	rootCmd.AddCommand(stsCmd)
}
