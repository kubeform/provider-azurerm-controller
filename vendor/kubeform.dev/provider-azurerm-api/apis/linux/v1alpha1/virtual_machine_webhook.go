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

func (r *VirtualMachine) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-linux-azurerm-kubeform-com-v1alpha1-virtualmachine,mutating=false,failurePolicy=fail,groups=linux.azurerm.kubeform.com,resources=virtualmachines,versions=v1alpha1,name=virtualmachine.linux.azurerm.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &VirtualMachine{}

var virtualmachineForceNewList = map[string]bool{
	"/admin_ssh_key/*/public_key":            true,
	"/admin_ssh_key/*/username":              true,
	"/admin_username":                        true,
	"/availability_set_id":                   true,
	"/computer_name":                         true,
	"/disable_password_authentication":       true,
	"/eviction_policy":                       true,
	"/location":                              true,
	"/name":                                  true,
	"/os_disk/*/diff_disk_settings/*/option": true,
	"/os_disk/*/name":                        true,
	"/os_disk/*/storage_account_type":        true,
	"/plan/*/name":                           true,
	"/plan/*/product":                        true,
	"/plan/*/publisher":                      true,
	"/platform_fault_domain":                 true,
	"/priority":                              true,
	"/provision_vm_agent":                    true,
	"/resource_group_name":                   true,
	"/source_image_id":                       true,
	"/source_image_reference/*/offer":        true,
	"/source_image_reference/*/publisher":    true,
	"/source_image_reference/*/sku":          true,
	"/source_image_reference/*/version":      true,
	"/virtual_machine_scale_set_id":          true,
	"/zone":                                  true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachine) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachine) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*VirtualMachine)
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

	for key := range virtualmachineForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`virtualmachine "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachine) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`virtualmachine "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
