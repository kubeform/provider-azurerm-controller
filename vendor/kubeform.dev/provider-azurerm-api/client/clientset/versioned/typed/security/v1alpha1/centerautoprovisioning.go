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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/security/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CenterAutoProvisioningsGetter has a method to return a CenterAutoProvisioningInterface.
// A group's client should implement this interface.
type CenterAutoProvisioningsGetter interface {
	CenterAutoProvisionings(namespace string) CenterAutoProvisioningInterface
}

// CenterAutoProvisioningInterface has methods to work with CenterAutoProvisioning resources.
type CenterAutoProvisioningInterface interface {
	Create(ctx context.Context, centerAutoProvisioning *v1alpha1.CenterAutoProvisioning, opts v1.CreateOptions) (*v1alpha1.CenterAutoProvisioning, error)
	Update(ctx context.Context, centerAutoProvisioning *v1alpha1.CenterAutoProvisioning, opts v1.UpdateOptions) (*v1alpha1.CenterAutoProvisioning, error)
	UpdateStatus(ctx context.Context, centerAutoProvisioning *v1alpha1.CenterAutoProvisioning, opts v1.UpdateOptions) (*v1alpha1.CenterAutoProvisioning, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CenterAutoProvisioning, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CenterAutoProvisioningList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CenterAutoProvisioning, err error)
	CenterAutoProvisioningExpansion
}

// centerAutoProvisionings implements CenterAutoProvisioningInterface
type centerAutoProvisionings struct {
	client rest.Interface
	ns     string
}

// newCenterAutoProvisionings returns a CenterAutoProvisionings
func newCenterAutoProvisionings(c *SecurityV1alpha1Client, namespace string) *centerAutoProvisionings {
	return &centerAutoProvisionings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the centerAutoProvisioning, and returns the corresponding centerAutoProvisioning object, and an error if there is any.
func (c *centerAutoProvisionings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CenterAutoProvisioning, err error) {
	result = &v1alpha1.CenterAutoProvisioning{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CenterAutoProvisionings that match those selectors.
func (c *centerAutoProvisionings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CenterAutoProvisioningList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CenterAutoProvisioningList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested centerAutoProvisionings.
func (c *centerAutoProvisionings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a centerAutoProvisioning and creates it.  Returns the server's representation of the centerAutoProvisioning, and an error, if there is any.
func (c *centerAutoProvisionings) Create(ctx context.Context, centerAutoProvisioning *v1alpha1.CenterAutoProvisioning, opts v1.CreateOptions) (result *v1alpha1.CenterAutoProvisioning, err error) {
	result = &v1alpha1.CenterAutoProvisioning{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(centerAutoProvisioning).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a centerAutoProvisioning and updates it. Returns the server's representation of the centerAutoProvisioning, and an error, if there is any.
func (c *centerAutoProvisionings) Update(ctx context.Context, centerAutoProvisioning *v1alpha1.CenterAutoProvisioning, opts v1.UpdateOptions) (result *v1alpha1.CenterAutoProvisioning, err error) {
	result = &v1alpha1.CenterAutoProvisioning{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		Name(centerAutoProvisioning.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(centerAutoProvisioning).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *centerAutoProvisionings) UpdateStatus(ctx context.Context, centerAutoProvisioning *v1alpha1.CenterAutoProvisioning, opts v1.UpdateOptions) (result *v1alpha1.CenterAutoProvisioning, err error) {
	result = &v1alpha1.CenterAutoProvisioning{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		Name(centerAutoProvisioning.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(centerAutoProvisioning).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the centerAutoProvisioning and deletes it. Returns an error if one occurs.
func (c *centerAutoProvisionings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *centerAutoProvisionings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched centerAutoProvisioning.
func (c *centerAutoProvisionings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CenterAutoProvisioning, err error) {
	result = &v1alpha1.CenterAutoProvisioning{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("centerautoprovisionings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
