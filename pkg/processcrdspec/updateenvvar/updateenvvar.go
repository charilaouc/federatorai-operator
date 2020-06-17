package updateenvvar

import (
	"strings"

	"github.com/containers-ai/federatorai-operator/pkg/util"
	securityv1 "github.com/openshift/api/security/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func AssignServiceAccountsToSecurityContextConstraints(scc *securityv1.SecurityContextConstraints, ns string) {
	serviceAccount := "serviceaccount:" + ns
	for index, value := range scc.Users {
		if strings.Contains(value, util.NamespaceServiceAccount) {
			newUser := strings.Replace(scc.Users[index], util.NamespaceServiceAccount, serviceAccount, -1)
			scc.Users[index] = newUser
		}
	}
}

func UpdateEnvVarsToDeployment(dep *appsv1.Deployment, envVars []corev1.EnvVar) {

	for containerIndex, container := range dep.Spec.Template.Spec.Containers {
		for _, envVar := range envVars {
			exist := false
			for envIndex, containerEnvVar := range container.Env {
				if envVar.Name == containerEnvVar.Name {
					exist = true
					dep.Spec.Template.Spec.Containers[containerIndex].Env[envIndex] = envVar
					break
				}
			}
			if !exist {
				dep.Spec.Template.Spec.Containers[containerIndex].Env = append(
					dep.Spec.Template.Spec.Containers[containerIndex].Env,
					envVar,
				)
			}
		}
	}
}
