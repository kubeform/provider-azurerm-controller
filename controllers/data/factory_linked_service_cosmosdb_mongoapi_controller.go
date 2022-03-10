/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Community License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Community-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package data

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	tfschema "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	auditlib "go.bytebuilders.dev/audit/lib"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/klog/v2"
	meta_util "kmodules.xyz/client-go/meta"
	datav1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"
	"kubeform.dev/provider-azurerm-controller/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

// FactoryLinkedServiceCosmosdbMongoapiReconciler reconciles a FactoryLinkedServiceCosmosdbMongoapi object
type FactoryLinkedServiceCosmosdbMongoapiReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme

	Gvk      schema.GroupVersionKind // GVK of the Resource
	Provider *tfschema.Provider      // returns a *schema.Provider from the provider package
	Resource *tfschema.Resource      // returns *schema.Resource
	TypeName string                  // resource type
}

// +kubebuilder:rbac:groups=data.azurerm.kubeform.com,resources=factorylinkedservicecosmosdbmongoapis,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=data.azurerm.kubeform.com,resources=factorylinkedservicecosmosdbmongoapis/status,verbs=get;update;patch

func (r *FactoryLinkedServiceCosmosdbMongoapiReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("factorylinkedservicecosmosdbmongoapi", req.NamespacedName)

	var unstructuredObj unstructured.Unstructured
	unstructuredObj.SetGroupVersionKind(r.Gvk)

	if err := r.Get(ctx, req.NamespacedName, &unstructuredObj); err != nil {
		log.Error(err, "unable to fetch FactoryLinkedServiceCosmosdbMongoapi")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	rClient := r.Client
	provider := r.Provider
	res := r.Resource
	gv := r.Gvk.GroupVersion()
	tName := r.TypeName
	jsonit := controllers.GetJSONItr(datav1alpha1.GetEncoder(), datav1alpha1.GetDecoder())
	err := controllers.StartProcess(rClient, provider, ctx, res, gv, &unstructuredObj, tName, jsonit)
	return ctrl.Result{}, err
}

func (r *FactoryLinkedServiceCosmosdbMongoapiReconciler) SetupWithManager(ctx context.Context, mgr ctrl.Manager, auditor *auditlib.EventPublisher, restrictToNamespace string) error {
	if auditor != nil {
		if err := auditor.SetupWithManager(ctx, mgr, &datav1alpha1.FactoryLinkedServiceCosmosdbMongoapi{}); err != nil {
			klog.Error(err, "unable to set up auditor", datav1alpha1.FactoryLinkedServiceCosmosdbMongoapi{}.APIVersion, datav1alpha1.FactoryLinkedServiceCosmosdbMongoapi{}.Kind)
			return err
		}
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&datav1alpha1.FactoryLinkedServiceCosmosdbMongoapi{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(e event.CreateEvent) bool {
				return !meta_util.MustAlreadyReconciled(e.Object)
			},
			UpdateFunc: func(e event.UpdateEvent) bool {
				return (e.ObjectNew.(metav1.Object)).GetDeletionTimestamp() != nil || !meta_util.MustAlreadyReconciled(e.ObjectNew)
			},
		}).
		WithEventFilter(predicate.NewPredicateFuncs(func(e client.Object) bool {
			if restrictToNamespace != "" && e.GetNamespace() != restrictToNamespace {
				klog.Infof("Only %s namespace is supported for Kubeform Community. Please upgrade to Kubeform Enterprise to use any namespace.", restrictToNamespace)
				return false
			}
			return true
		})).
		Watches(
			&source.Kind{Type: &v1.Secret{}},
			handler.EnqueueRequestsFromMapFunc(r.SensitiveSecretWatch(ctx)),
		).
		Complete(r)
}

func (r *FactoryLinkedServiceCosmosdbMongoapiReconciler) SensitiveSecretWatch(ctx context.Context) handler.MapFunc {
	log := ctrl.LoggerFrom(ctx)
	return func(o client.Object) []ctrl.Request {
		result := []ctrl.Request{}

		sensSec, ok := o.(*v1.Secret)
		if !ok {
			log.Error(fmt.Errorf("expected a Secret but go a %T", o), "failed to get secret for FactoryLinkedServiceCosmosdbMongoapi")
			return nil
		}

		secName := sensSec.Name
		secNamespace := sensSec.Namespace

		resourceList := &datav1alpha1.FactoryLinkedServiceCosmosdbMongoapiList{}
		if err := r.List(ctx, resourceList, client.InNamespace(secNamespace)); err != nil {
			log.Error(err, "failed to list FactoryLinkedServiceCosmosdbMongoapi")
			return nil
		}

		for _, res := range resourceList.Items {
			if res.Spec.SecretRef.Name == secName {
				name := client.ObjectKey{Namespace: res.Namespace, Name: res.Name}
				result = append(result, ctrl.Request{NamespacedName: name})
			}
		}
		return result
	}
}