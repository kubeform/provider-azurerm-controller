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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/kusto/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AttachedDatabaseConfigurationsGetter has a method to return a AttachedDatabaseConfigurationInterface.
// A group's client should implement this interface.
type AttachedDatabaseConfigurationsGetter interface {
	AttachedDatabaseConfigurations(namespace string) AttachedDatabaseConfigurationInterface
}

// AttachedDatabaseConfigurationInterface has methods to work with AttachedDatabaseConfiguration resources.
type AttachedDatabaseConfigurationInterface interface {
	Create(ctx context.Context, attachedDatabaseConfiguration *v1alpha1.AttachedDatabaseConfiguration, opts v1.CreateOptions) (*v1alpha1.AttachedDatabaseConfiguration, error)
	Update(ctx context.Context, attachedDatabaseConfiguration *v1alpha1.AttachedDatabaseConfiguration, opts v1.UpdateOptions) (*v1alpha1.AttachedDatabaseConfiguration, error)
	UpdateStatus(ctx context.Context, attachedDatabaseConfiguration *v1alpha1.AttachedDatabaseConfiguration, opts v1.UpdateOptions) (*v1alpha1.AttachedDatabaseConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AttachedDatabaseConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AttachedDatabaseConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AttachedDatabaseConfiguration, err error)
	AttachedDatabaseConfigurationExpansion
}

// attachedDatabaseConfigurations implements AttachedDatabaseConfigurationInterface
type attachedDatabaseConfigurations struct {
	client rest.Interface
	ns     string
}

// newAttachedDatabaseConfigurations returns a AttachedDatabaseConfigurations
func newAttachedDatabaseConfigurations(c *KustoV1alpha1Client, namespace string) *attachedDatabaseConfigurations {
	return &attachedDatabaseConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the attachedDatabaseConfiguration, and returns the corresponding attachedDatabaseConfiguration object, and an error if there is any.
func (c *attachedDatabaseConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AttachedDatabaseConfiguration, err error) {
	result = &v1alpha1.AttachedDatabaseConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AttachedDatabaseConfigurations that match those selectors.
func (c *attachedDatabaseConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AttachedDatabaseConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AttachedDatabaseConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested attachedDatabaseConfigurations.
func (c *attachedDatabaseConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a attachedDatabaseConfiguration and creates it.  Returns the server's representation of the attachedDatabaseConfiguration, and an error, if there is any.
func (c *attachedDatabaseConfigurations) Create(ctx context.Context, attachedDatabaseConfiguration *v1alpha1.AttachedDatabaseConfiguration, opts v1.CreateOptions) (result *v1alpha1.AttachedDatabaseConfiguration, err error) {
	result = &v1alpha1.AttachedDatabaseConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(attachedDatabaseConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a attachedDatabaseConfiguration and updates it. Returns the server's representation of the attachedDatabaseConfiguration, and an error, if there is any.
func (c *attachedDatabaseConfigurations) Update(ctx context.Context, attachedDatabaseConfiguration *v1alpha1.AttachedDatabaseConfiguration, opts v1.UpdateOptions) (result *v1alpha1.AttachedDatabaseConfiguration, err error) {
	result = &v1alpha1.AttachedDatabaseConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		Name(attachedDatabaseConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(attachedDatabaseConfiguration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *attachedDatabaseConfigurations) UpdateStatus(ctx context.Context, attachedDatabaseConfiguration *v1alpha1.AttachedDatabaseConfiguration, opts v1.UpdateOptions) (result *v1alpha1.AttachedDatabaseConfiguration, err error) {
	result = &v1alpha1.AttachedDatabaseConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		Name(attachedDatabaseConfiguration.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(attachedDatabaseConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the attachedDatabaseConfiguration and deletes it. Returns an error if one occurs.
func (c *attachedDatabaseConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *attachedDatabaseConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched attachedDatabaseConfiguration.
func (c *attachedDatabaseConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AttachedDatabaseConfiguration, err error) {
	result = &v1alpha1.AttachedDatabaseConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("attacheddatabaseconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
