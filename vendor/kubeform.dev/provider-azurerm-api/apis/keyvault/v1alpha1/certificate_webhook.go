/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	"fmt"
	"strings"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/apimachinery/pkg/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *Certificate) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-keyvault-azurerm-kubeform-com-v1alpha1-certificate,mutating=false,failurePolicy=fail,groups=keyvault.azurerm.kubeform.com,resources=certificates,versions=v1alpha1,name=certificate.keyvault.azurerm.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &Certificate{}

var certificateForceNewList = map[string]bool{
	"/certificate_policy/*/issuer_parameters/*/name":                                            true,
	"/certificate_policy/*/key_properties/*/curve":                                              true,
	"/certificate_policy/*/key_properties/*/exportable":                                         true,
	"/certificate_policy/*/key_properties/*/key_size":                                           true,
	"/certificate_policy/*/key_properties/*/key_type":                                           true,
	"/certificate_policy/*/key_properties/*/reuse_key":                                          true,
	"/certificate_policy/*/lifetime_action/*/action/*/action_type":                              true,
	"/certificate_policy/*/lifetime_action/*/trigger/*/days_before_expiry":                      true,
	"/certificate_policy/*/lifetime_action/*/trigger/*/lifetime_percentage":                     true,
	"/certificate_policy/*/secret_properties/*/content_type":                                    true,
	"/certificate_policy/*/x509_certificate_properties/*/extended_key_usage":                    true,
	"/certificate_policy/*/x509_certificate_properties/*/key_usage":                             true,
	"/certificate_policy/*/x509_certificate_properties/*/subject":                               true,
	"/certificate_policy/*/x509_certificate_properties/*/subject_alternative_names/*/dns_names": true,
	"/certificate_policy/*/x509_certificate_properties/*/subject_alternative_names/*/emails":    true,
	"/certificate_policy/*/x509_certificate_properties/*/subject_alternative_names/*/upns":      true,
	"/certificate_policy/*/x509_certificate_properties/*/validity_in_months":                    true,
	"/key_vault_id": true,
	"/name":         true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Certificate) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Certificate) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*Certificate)
	oldObj := res.Spec.Resource

	jsnitr := jsoniter.Config{
		EscapeHTML:             true,
		SortMapKeys:            true,
		TagKey:                 "tf",
		ValidateJsonRawMessage: true,
		TypeEncoders:           GetEncoder(),
		TypeDecoders:           GetDecoder(),
	}.Froze()

	byteNew, err := jsnitr.Marshal(newObj)
	if err != nil {
		return err
	}
	tempNew := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteNew, &tempNew)
	if err != nil {
		return err
	}

	byteOld, err := jsnitr.Marshal(oldObj)
	if err != nil {
		return err
	}
	tempOld := make(map[string]interface{})
	err = jsnitr.Unmarshal(byteOld, &tempOld)
	if err != nil {
		return err
	}

	for key, _ := range certificateForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`certificate "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Certificate) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`certificate "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
