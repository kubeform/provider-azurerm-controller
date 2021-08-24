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

func (r *VirtualMachineScaleSet) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-linux-azurerm-kubeform-com-v1alpha1-virtualmachinescaleset,mutating=false,failurePolicy=fail,groups=linux.azurerm.kubeform.com,resources=virtualmachinescalesets,versions=v1alpha1,name=virtualmachinescaleset.linux.azurerm.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &VirtualMachineScaleSet{}

var virtualmachinescalesetForceNewList = map[string]bool{
	"/additional_capabilities/*/ultra_ssd_enabled": true,
	"/admin_username":                     true,
	"/computer_name_prefix":               true,
	"/data_disk/*/disk_encryption_set_id": true,
	"/eviction_policy":                    true,
	"/location":                           true,
	"/name":                               true,
	"/network_interface/*/ip_configuration/*/public_ip_address/*/ip_tag/*/tag":        true,
	"/network_interface/*/ip_configuration/*/public_ip_address/*/ip_tag/*/type":       true,
	"/network_interface/*/ip_configuration/*/public_ip_address/*/public_ip_prefix_id": true,
	"/network_interface/*/name":              true,
	"/os_disk/*/diff_disk_settings/*/option": true,
	"/os_disk/*/disk_encryption_set_id":      true,
	"/os_disk/*/storage_account_type":        true,
	"/plan/*/name":                           true,
	"/plan/*/product":                        true,
	"/plan/*/publisher":                      true,
	"/platform_fault_domain_count":           true,
	"/priority":                              true,
	"/provision_vm_agent":                    true,
	"/proximity_placement_group_id":          true,
	"/resource_group_name":                   true,
	"/upgrade_mode":                          true,
	"/zone_balance":                          true,
	"/zones":                                 true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachineScaleSet) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachineScaleSet) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*VirtualMachineScaleSet)
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

	for key := range virtualmachinescalesetForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`virtualmachinescaleset "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *VirtualMachineScaleSet) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`virtualmachinescaleset "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
