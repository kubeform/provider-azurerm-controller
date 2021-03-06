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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/data/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FactoryLinkedServiceSynapsesGetter has a method to return a FactoryLinkedServiceSynapseInterface.
// A group's client should implement this interface.
type FactoryLinkedServiceSynapsesGetter interface {
	FactoryLinkedServiceSynapses(namespace string) FactoryLinkedServiceSynapseInterface
}

// FactoryLinkedServiceSynapseInterface has methods to work with FactoryLinkedServiceSynapse resources.
type FactoryLinkedServiceSynapseInterface interface {
	Create(ctx context.Context, factoryLinkedServiceSynapse *v1alpha1.FactoryLinkedServiceSynapse, opts v1.CreateOptions) (*v1alpha1.FactoryLinkedServiceSynapse, error)
	Update(ctx context.Context, factoryLinkedServiceSynapse *v1alpha1.FactoryLinkedServiceSynapse, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceSynapse, error)
	UpdateStatus(ctx context.Context, factoryLinkedServiceSynapse *v1alpha1.FactoryLinkedServiceSynapse, opts v1.UpdateOptions) (*v1alpha1.FactoryLinkedServiceSynapse, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FactoryLinkedServiceSynapse, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FactoryLinkedServiceSynapseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceSynapse, err error)
	FactoryLinkedServiceSynapseExpansion
}

// factoryLinkedServiceSynapses implements FactoryLinkedServiceSynapseInterface
type factoryLinkedServiceSynapses struct {
	client rest.Interface
	ns     string
}

// newFactoryLinkedServiceSynapses returns a FactoryLinkedServiceSynapses
func newFactoryLinkedServiceSynapses(c *DataV1alpha1Client, namespace string) *factoryLinkedServiceSynapses {
	return &factoryLinkedServiceSynapses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the factoryLinkedServiceSynapse, and returns the corresponding factoryLinkedServiceSynapse object, and an error if there is any.
func (c *factoryLinkedServiceSynapses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FactoryLinkedServiceSynapse, err error) {
	result = &v1alpha1.FactoryLinkedServiceSynapse{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FactoryLinkedServiceSynapses that match those selectors.
func (c *factoryLinkedServiceSynapses) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FactoryLinkedServiceSynapseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FactoryLinkedServiceSynapseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested factoryLinkedServiceSynapses.
func (c *factoryLinkedServiceSynapses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a factoryLinkedServiceSynapse and creates it.  Returns the server's representation of the factoryLinkedServiceSynapse, and an error, if there is any.
func (c *factoryLinkedServiceSynapses) Create(ctx context.Context, factoryLinkedServiceSynapse *v1alpha1.FactoryLinkedServiceSynapse, opts v1.CreateOptions) (result *v1alpha1.FactoryLinkedServiceSynapse, err error) {
	result = &v1alpha1.FactoryLinkedServiceSynapse{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(factoryLinkedServiceSynapse).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a factoryLinkedServiceSynapse and updates it. Returns the server's representation of the factoryLinkedServiceSynapse, and an error, if there is any.
func (c *factoryLinkedServiceSynapses) Update(ctx context.Context, factoryLinkedServiceSynapse *v1alpha1.FactoryLinkedServiceSynapse, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceSynapse, err error) {
	result = &v1alpha1.FactoryLinkedServiceSynapse{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		Name(factoryLinkedServiceSynapse.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(factoryLinkedServiceSynapse).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *factoryLinkedServiceSynapses) UpdateStatus(ctx context.Context, factoryLinkedServiceSynapse *v1alpha1.FactoryLinkedServiceSynapse, opts v1.UpdateOptions) (result *v1alpha1.FactoryLinkedServiceSynapse, err error) {
	result = &v1alpha1.FactoryLinkedServiceSynapse{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		Name(factoryLinkedServiceSynapse.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(factoryLinkedServiceSynapse).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the factoryLinkedServiceSynapse and deletes it. Returns an error if one occurs.
func (c *factoryLinkedServiceSynapses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *factoryLinkedServiceSynapses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched factoryLinkedServiceSynapse.
func (c *factoryLinkedServiceSynapses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FactoryLinkedServiceSynapse, err error) {
	result = &v1alpha1.FactoryLinkedServiceSynapse{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("factorylinkedservicesynapses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
