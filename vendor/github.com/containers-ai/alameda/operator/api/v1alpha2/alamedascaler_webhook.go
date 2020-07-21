/*
Copyright 2019 The Alameda Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var alamedascalerlog = logf.Log.WithName("alamedascaler-resource")
var k8sClient client.Client

func (r *AlamedaScaler) SetupWebhookWithManager(mgr ctrl.Manager) error {
	k8sClient = mgr.GetClient()
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-autoscaling-containers-ai-v1alpha2-alamedascaler,mutating=true,failurePolicy=fail,groups=autoscaling.containers.ai,resources=alamedascalers,verbs=create;update,versions=v1alpha2,name=malamedascaler.containers.ai

var _ webhook.Defaulter = &AlamedaScaler{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *AlamedaScaler) Default() {
	alamedascalerlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-autoscaling-containers-ai-v1alpha2-alamedascaler,mutating=false,failurePolicy=fail,groups=autoscaling.containers.ai,resources=alamedascalers,versions=v1alpha2,name=valamedascaler.containers.ai

var _ webhook.Validator = &AlamedaScaler{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AlamedaScaler) ValidateCreate() error {
	alamedascalerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AlamedaScaler) ValidateUpdate(old runtime.Object) error {
	alamedascalerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AlamedaScaler) ValidateDelete() error {
	alamedascalerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
