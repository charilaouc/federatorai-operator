package updateenvvar

import (
	"strings"

	"github.com/containers-ai/federatorai-operator/pkg/util"
	securityv1 "github.com/openshift/api/security/v1"
	appsv1 "k8s.io/api/apps/v1"
)

func AssignServiceToDeployment(dep *appsv1.Deployment, ns string) {
	if len(dep.Spec.Template.Spec.Containers[0].Env) > 0 {
		for index, value := range dep.Spec.Template.Spec.Containers[0].Env {
			if strings.Contains(value.String(), util.NamespaceService) {
				dep.Spec.Template.Spec.Containers[0].Env[index].Value = strings.Replace(dep.Spec.Template.Spec.Containers[0].Env[index].Value, util.NamespaceService, ns+".svc", -1)
			}
		}
	}

	for containerIdx, _ := range dep.Spec.Template.Spec.Containers {
		for index, value := range dep.Spec.Template.Spec.Containers[containerIdx].Args {
			if strings.Contains(value, util.NamespaceService) {
				newArg := strings.Replace(dep.Spec.Template.Spec.Containers[containerIdx].Args[index], util.NamespaceService, ns+".svc", -1)
				dep.Spec.Template.Spec.Containers[containerIdx].Args[index] = newArg
			}
		}
	}
}
func AssignServiceToDaemonSet(ds *appsv1.DaemonSet, ns string) {
	if len(ds.Spec.Template.Spec.Containers[0].Env) > 0 {
		for index, value := range ds.Spec.Template.Spec.Containers[0].Env {
			if strings.Contains(value.String(), util.NamespaceService) {
				ds.Spec.Template.Spec.Containers[0].Env[index].Value = strings.Replace(ds.Spec.Template.Spec.Containers[0].Env[index].Value, util.NamespaceService, ns+".svc", -1)
			}
		}
	}

	for containerIdx, _ := range ds.Spec.Template.Spec.Containers {
		for index, value := range ds.Spec.Template.Spec.Containers[containerIdx].Args {
			if strings.Contains(value, util.NamespaceService) {
				newArg := strings.Replace(ds.Spec.Template.Spec.Containers[containerIdx].Args[index], util.NamespaceService, ns+".svc", -1)
				ds.Spec.Template.Spec.Containers[containerIdx].Args[index] = newArg
			}
		}
	}
}
func AssignServiceAccountsToSecurityContextConstraints(scc *securityv1.SecurityContextConstraints, ns string) {
	serviceAccount := "serviceaccount:" + ns
	for index, value := range scc.Users {
		if strings.Contains(value, util.NamespaceServiceAccount) {
			newUser := strings.Replace(scc.Users[index], util.NamespaceServiceAccount, serviceAccount, -1)
			scc.Users[index] = newUser
		}
	}
}
