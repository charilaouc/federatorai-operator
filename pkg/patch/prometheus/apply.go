package prometheus

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func PromApply(k8sCli client.Client) error {
	ok, missingRulesMap, err := RulesCheck(k8sCli)
	if err != nil {
		return err
	} else if !ok {
		if err = PatchMissingRules(k8sCli, missingRulesMap); err != nil {
			return err
		}
	}

	/*
		ok, err = RelabelingCheck(k8sCli)
		if err != nil {
			return err
		} else if !ok {
			if err = PatchRelabelings(k8sCli); err != nil {
				return err
			}
		}
	*/
	return nil
}
