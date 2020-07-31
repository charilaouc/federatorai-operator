package prometheus

import (
	"os"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

func PromCheck(k8sCli client.Client) error {
	ok, _, err := RulesCheck(k8sCli)
	if err == nil && !ok {
		logger.Info("some rules are missing and need to patch")
		os.Exit(1)
	}
	if err != nil {
		logger.Error(err, "prometheus rule check error")
		os.Exit(2)
	}
	/*
		ok, err = RelabelingCheck(k8sCli)
		if err == nil && !ok {
			logger.Info("some reabelings are missing and need to patch")
			os.Exit(1)
		}
	*/
	return err
}
