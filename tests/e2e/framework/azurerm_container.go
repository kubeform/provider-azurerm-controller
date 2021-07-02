/*
Copyright AppsCode Inc. and Contributors

Licensed under the AppsCode Free Trial License 1.0.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://github.com/appscode/licenses/raw/1.0.0/AppsCode-Free-Trial-1.0.0.md

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	"context"
	"time"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
	"kubeform.dev/provider-azurerm-api/apis/storage/v1alpha1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
)

func (i *Invocation) Container(name string, secretName string) *v1alpha1.Container {
	storageACC := "kubedbci"

	return &v1alpha1.Container{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.ContainerSpec{
			ProviderRef: corev1.LocalObjectReference{
				Name: secretName,
			},
			Resource: v1alpha1.ContainerSpecResource{
				Name:               &name,
				StorageAccountName: &storageACC,
			},
		},
	}
}

func (f *Framework) CreateContainer(obj *v1alpha1.Container) error {
	_, err := f.kfClient.StorageV1alpha1().Containers(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) UpdateContainer(obj *v1alpha1.Container) error {
	obj, err := f.kfClient.StorageV1alpha1().Containers(obj.Namespace).Get(context.TODO(), obj.Name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	updataName := *obj.Spec.Resource.Name + "-update"
	obj.Spec.Resource.Name = &updataName

	_, err = f.kfClient.StorageV1alpha1().Containers(obj.Namespace).Update(context.TODO(), obj, metav1.UpdateOptions{})
	return err
}

func (f *Framework) DeleteContainer(meta metav1.ObjectMeta) error {
	return f.kfClient.StorageV1alpha1().Containers(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInBackground())
}

func (f *Framework) EventuallyContainerRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			container, err := f.kfClient.StorageV1alpha1().Containers(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return container.Status.Phase == status.CurrentStatus
		},
		time.Minute*30,
		time.Second*10,
	)
}

func (f *Framework) EventuallyContainerDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.kfClient.StorageV1alpha1().Containers(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*30,
		time.Second*10,
	)
}
