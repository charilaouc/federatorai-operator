package prometheus

import (
	"github.com/spf13/cobra"
)

var configurationFilePath string

func init() {
	PromApplyCmd.Flags().StringVar(&configurationFilePath, "config", "/etc/federatorai/operator/operator.toml", "File path to federatorai-operator coniguration")
}

var (
	PromApplyCmd = &cobra.Command{
		Use:   "prom_apply",
		Short: "prom_apply ",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			InitConfig()
			InitK8SClient()

			ok, missingRulesMap, err := RulesCheck(k8sCli)
			if err != nil {
				return err
			} else if ok {
				return nil
			}
			err = PatchMissingRules(k8sCli, missingRulesMap)
			return err
		},
	}
)
