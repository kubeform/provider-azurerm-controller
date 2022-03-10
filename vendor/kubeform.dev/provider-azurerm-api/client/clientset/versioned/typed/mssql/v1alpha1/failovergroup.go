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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/mssql/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FailoverGroupsGetter has a method to return a FailoverGroupInterface.
// A group's client should implement this interface.
type FailoverGroupsGetter interface {
	FailoverGroups(namespace string) FailoverGroupInterface
}

// FailoverGroupInterface has methods to work with FailoverGroup resources.
type FailoverGroupInterface interface {
	Create(ctx context.Context, failoverGroup *v1alpha1.FailoverGroup, opts v1.CreateOptions) (*v1alpha1.FailoverGroup, error)
	Update(ctx context.Context, failoverGroup *v1alpha1.FailoverGroup, opts v1.UpdateOptions) (*v1alpha1.FailoverGroup, error)
	UpdateStatus(ctx context.Context, failoverGroup *v1alpha1.FailoverGroup, opts v1.UpdateOptions) (*v1alpha1.FailoverGroup, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.FailoverGroup, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.FailoverGroupList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FailoverGroup, err error)
	FailoverGroupExpansion
}

// failoverGroups implements FailoverGroupInterface
type failoverGroups struct {
	client rest.Interface
	ns     string
}

// newFailoverGroups returns a FailoverGroups
func newFailoverGroups(c *MssqlV1alpha1Client, namespace string) *failoverGroups {
	return &failoverGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the failoverGroup, and returns the corresponding failoverGroup object, and an error if there is any.
func (c *failoverGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.FailoverGroup, err error) {
	result = &v1alpha1.FailoverGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("failovergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FailoverGroups that match those selectors.
func (c *failoverGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.FailoverGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FailoverGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("failovergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested failoverGroups.
func (c *failoverGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("failovergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a failoverGroup and creates it.  Returns the server's representation of the failoverGroup, and an error, if there is any.
func (c *failoverGroups) Create(ctx context.Context, failoverGroup *v1alpha1.FailoverGroup, opts v1.CreateOptions) (result *v1alpha1.FailoverGroup, err error) {
	result = &v1alpha1.FailoverGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("failovergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(failoverGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a failoverGroup and updates it. Returns the server's representation of the failoverGroup, and an error, if there is any.
func (c *failoverGroups) Update(ctx context.Context, failoverGroup *v1alpha1.FailoverGroup, opts v1.UpdateOptions) (result *v1alpha1.FailoverGroup, err error) {
	result = &v1alpha1.FailoverGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("failovergroups").
		Name(failoverGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(failoverGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *failoverGroups) UpdateStatus(ctx context.Context, failoverGroup *v1alpha1.FailoverGroup, opts v1.UpdateOptions) (result *v1alpha1.FailoverGroup, err error) {
	result = &v1alpha1.FailoverGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("failovergroups").
		Name(failoverGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(failoverGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the failoverGroup and deletes it. Returns an error if one occurs.
func (c *failoverGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("failovergroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *failoverGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("failovergroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched failoverGroup.
func (c *failoverGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.FailoverGroup, err error) {
	result = &v1alpha1.FailoverGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("failovergroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
