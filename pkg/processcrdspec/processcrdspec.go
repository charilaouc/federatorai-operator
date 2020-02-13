package processcrdspec

import (
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec/alamedaserviceparamter"
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec/componentsectionset"
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec/globalsectionset"
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec/updateenvvar"
	securityv1 "github.com/openshift/api/security/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func ParamterToDeployment(dep *appsv1.Deployment, asp *alamedaserviceparamter.AlamedaServiceParamter) *appsv1.Deployment {
	globalsectionset.GlobalSectionSetParamterToDeployment(dep, asp) //DeploymentSpec's Global Section Set
	componentsectionset.SectionSetParamterToDeployment(dep, asp)    //DeploymentSpec's Component Section Set
	return dep
}
func ParamterToDaemonSet(ds *appsv1.DaemonSet, asp *alamedaserviceparamter.AlamedaServiceParamter) *appsv1.DaemonSet {
	globalsectionset.GlobalSectionSetParamterToDaemonSet(ds, asp) //DaemonSet's Global Section Set
	componentsectionset.SectionSetParamterToDaemonSet(ds, asp)    //DaemonSet's Component Section Set
	return ds
}
func ParamterToSecurityContextConstraints(scc *securityv1.SecurityContextConstraints, asp *alamedaserviceparamter.AlamedaServiceParamter) *securityv1.SecurityContextConstraints {
	updateenvvar.AssignServiceAccountsToSecurityContextConstraints(scc, asp.NameSpace)
	return scc
}

func ParamterToPersistentVolumeClaim(pvc *corev1.PersistentVolumeClaim, asp *alamedaserviceparamter.AlamedaServiceParamter) *corev1.PersistentVolumeClaim {
	globalsectionset.GlobalSectionSetParamterToPersistentVolumeClaim(pvc, asp)
	componentsectionset.SectionSetParamterToPersistentVolumeClaim(pvc, asp) //PersistentVolumeClaim's Component Section Set
	return pvc
}
func ParamterToStatefulset(ss *appsv1.StatefulSet, asp *alamedaserviceparamter.AlamedaServiceParamter) *appsv1.StatefulSet {
	globalsectionset.GlobalSectionSetParamterToStatefulset(ss, asp)
	componentsectionset.SectionSetParamterToStatefulSet(ss, asp) //StatefulSet's Component Section Set
	return ss
}
