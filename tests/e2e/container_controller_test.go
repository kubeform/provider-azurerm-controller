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

package e2e_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	azureresource "kubeform.dev/provider-azurerm-api/apis/storage/v1alpha1"
	"kubeform.dev/provider-azurerm-controller/tests/e2e/framework"
)

var _ = Describe("Test", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
	})
	Describe("Azurerm", func() {
		Context("ContainerController", func() {
			var (
				providerRef   *core.Secret
				containerName string
				secretName    string
				container     *azureresource.Container
			)

			BeforeEach(func() {
				containerName = f.GetRandomName("")
				secretName = f.GetRandomName("secret")
				providerRef = f.AzureProviderRef(secretName)
				container = f.Container(containerName, secretName)
			})

			AfterEach(func() {
				By("Deleting Container")
				err = f.DeleteContainer(container.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting Container")
				f.EventuallyContainerDeleted(container.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})

			It("should create and delete Container successfully", func() {
				By("Creating AzureProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Container")
				err = f.CreateContainer(container)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Container")
				f.EventuallyContainerRunning(container.ObjectMeta).Should(BeTrue())
			})

			It("should create, update and delete Container successfully", func() {
				By("Creating AzureProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating Container")
				err = f.CreateContainer(container)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Container")
				f.EventuallyContainerRunning(container.ObjectMeta).Should(BeTrue())

				By("Updating Container")
				err = f.UpdateContainer(container)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running Container")
				f.EventuallyContainerRunning(container.ObjectMeta).Should(BeTrue())
			})
		})
	})
})
