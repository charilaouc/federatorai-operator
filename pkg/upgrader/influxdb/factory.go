package influxdb

import (
	federatoraiv1alpha1 "github.com/containers-ai/federatorai-operator/api/v1alpha1"
	"github.com/containers-ai/federatorai-operator/pkg/component"
	"github.com/containers-ai/federatorai-operator/pkg/influxdb"
	"github.com/containers-ai/federatorai-operator/pkg/processcrdspec/alamedaserviceparamter"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// resourceFacetory encapsulates componentes that are need while creating resource
type resourceFactory struct {
	asp alamedaserviceparamter.AlamedaServiceParamter
	c   *component.ComponentConfig
}

func newResourceFactoryByAlamedaService(
	alamedaService federatoraiv1alpha1.AlamedaService) (resourceFactory, error) {

	c := component.NewComponentConfig(
		component.PodTemplateConfig{},
		alamedaService, component.WithNamespace(alamedaService.GetNamespace()))

	asp := alamedaserviceparamter.NewAlamedaServiceParamter(&alamedaService)

	return resourceFactory{
		asp: *asp,
		c:   c,
	}, nil
}

func (f resourceFactory) getAlamedaInfluxdbSts() *appsv1.StatefulSet {
	influxdbStsAssets := alamedaserviceparamter.GetAlamedaInfluxdbStatefulSet()
	influxdbSts := f.c.NewStatefulSet(influxdbStsAssets)
	return influxdbSts
}

func (f resourceFactory) getAlamedaInfluxdbService() corev1.Service {
	influxdbServiceAssets := alamedaserviceparamter.GetAlamedaInfluxdbService()
	influxdbService := f.c.NewService(influxdbServiceAssets)
	return *influxdbService
}

func (f resourceFactory) getAlamedaInfluxdbConfig() influxdb.Config {
	influxdbStsAssets := alamedaserviceparamter.GetAlamedaInfluxdbStatefulSet()
	influxdbSts := f.c.NewStatefulSet(influxdbStsAssets)
	influxdbServiceAssets := alamedaserviceparamter.GetAlamedaInfluxdbService()
	influxdbService := f.c.NewService(influxdbServiceAssets)
	influxdbConfig := getInfluxdbConfigFromStatefulSetAndService(*influxdbSts, *influxdbService)
	return influxdbConfig
}

func (f resourceFactory) listWorkloadControllersWithoutAlamedaInfluxdb() []workloadController {

	resource := f.asp.GetInstallResource()
	resource.Delete(alamedaserviceparamter.GetAlamedaInfluxdbResource())

	workloadControllers := make([]workloadController, 0,
		len(resource.DeploymentList)+len(resource.StatefulSetList))
	for _, deployment := range resource.DeploymentList {
		d := f.c.NewDeployment(deployment)
		workloadControllers = append(
			workloadControllers, workloadController{
				Kind:      workloadControllerKind(workloadControllerDeployment),
				Namespace: d.GetNamespace(),
				Name:      d.GetName(),
			})
	}
	for _, statefulSet := range resource.StatefulSetList {
		s := f.c.NewStatefulSet(statefulSet)
		workloadControllers = append(workloadControllers,
			workloadController{
				Kind:      workloadControllerKind(workloadControllerStatefulSet),
				Namespace: s.GetNamespace(),
				Name:      s.GetName(),
			})
	}

	return workloadControllers
}
