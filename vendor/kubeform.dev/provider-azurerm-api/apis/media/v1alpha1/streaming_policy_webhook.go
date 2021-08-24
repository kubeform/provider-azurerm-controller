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
	"kubeform.dev/provider-azurerm-api/util"

	jsoniter "github.com/json-iterator/go"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func (r *StreamingPolicy) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:verbs=create;update;delete,path=/validate-media-azurerm-kubeform-com-v1alpha1-streamingpolicy,mutating=false,failurePolicy=fail,groups=media.azurerm.kubeform.com,resources=streamingpolicies,versions=v1alpha1,name=streamingpolicy.media.azurerm.kubeform.io,sideEffects=None,admissionReviewVersions=v1

var _ webhook.Validator = &StreamingPolicy{}

var streamingpolicyForceNewList = map[string]bool{
	"/common_encryption_cbcs/*/default_content_key/*/label":                             true,
	"/common_encryption_cbcs/*/default_content_key/*/policy_name":                       true,
	"/common_encryption_cbcs/*/drm_fairplay/*/allow_persistent_license":                 true,
	"/common_encryption_cbcs/*/drm_fairplay/*/custom_license_acquisition_url_template":  true,
	"/common_encryption_cbcs/*/enabled_protocols/*/dash":                                true,
	"/common_encryption_cbcs/*/enabled_protocols/*/download":                            true,
	"/common_encryption_cbcs/*/enabled_protocols/*/hls":                                 true,
	"/common_encryption_cbcs/*/enabled_protocols/*/smooth_streaming":                    true,
	"/common_encryption_cenc/*/default_content_key/*/label":                             true,
	"/common_encryption_cenc/*/default_content_key/*/policy_name":                       true,
	"/common_encryption_cenc/*/drm_playready/*/custom_attributes":                       true,
	"/common_encryption_cenc/*/drm_playready/*/custom_license_acquisition_url_template": true,
	"/common_encryption_cenc/*/drm_widevine_custom_license_acquisition_url_template":    true,
	"/common_encryption_cenc/*/enabled_protocols/*/dash":                                true,
	"/common_encryption_cenc/*/enabled_protocols/*/download":                            true,
	"/common_encryption_cenc/*/enabled_protocols/*/hls":                                 true,
	"/common_encryption_cenc/*/enabled_protocols/*/smooth_streaming":                    true,
	"/default_content_key_policy_name":                                                  true,
	"/media_services_account_name":                                                      true,
	"/name":                                                                             true,
	"/no_encryption_enabled_protocols/*/dash":                                           true,
	"/no_encryption_enabled_protocols/*/download":                                       true,
	"/no_encryption_enabled_protocols/*/hls":                                            true,
	"/no_encryption_enabled_protocols/*/smooth_streaming":                               true,
	"/resource_group_name":                                                              true,
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *StreamingPolicy) ValidateCreate() error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *StreamingPolicy) ValidateUpdate(old runtime.Object) error {
	if r.Spec.Resource.ID == "" {
		return nil
	}
	newObj := r.Spec.Resource
	res := old.(*StreamingPolicy)
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

	for key := range streamingpolicyForceNewList {
		keySplit := strings.Split(key, "/*")
		length := len(keySplit)
		checkIfAnyDif := false
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempOld, tempOld, tempNew)
		util.CheckIfAnyDifference("", keySplit, 0, length, &checkIfAnyDif, tempNew, tempOld, tempNew)

		if checkIfAnyDif && r.Spec.UpdatePolicy == base.UpdatePolicyDoNotDestroy {
			return fmt.Errorf(`streamingpolicy "%v/%v" immutable field can't be updated. To update, change spec.updatePolicy to Destroy`, r.Namespace, r.Name)
		}
	}
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *StreamingPolicy) ValidateDelete() error {
	if r.Spec.TerminationPolicy == base.TerminationPolicyDoNotTerminate {
		return fmt.Errorf(`streamingpolicy "%v/%v" can't be terminated. To delete, change spec.terminationPolicy to Delete`, r.Namespace, r.Name)
	}
	return nil
}
