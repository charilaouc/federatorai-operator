package upgrader

import (
	"github.com/spf13/cobra"

	"github.com/containers-ai/federatorai-operator/cmd/upgrader/influxdb"
)

var (
	UpgradeRootCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade ",
		Long:  "",
	}

	logOutputPath = "/var/log/alameda/federatorai-operator-upgrade.log"
)

func init() {
	UpgradeRootCmd.AddCommand(influxdb.UpgradeInfluxDBSchemaCMD)
	UpgradeRootCmd.PersistentFlags().StringVar(&logOutputPath, "log-output", logOutputPath, "File path to federatorai-operator upgrade log output")
}
