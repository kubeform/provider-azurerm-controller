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

	v1alpha1 "kubeform.dev/provider-azurerm-api/apis/notificationhub/v1alpha1"
	scheme "kubeform.dev/provider-azurerm-api/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NotificationHubsGetter has a method to return a NotificationHubInterface.
// A group's client should implement this interface.
type NotificationHubsGetter interface {
	NotificationHubs(namespace string) NotificationHubInterface
}

// NotificationHubInterface has methods to work with NotificationHub resources.
type NotificationHubInterface interface {
	Create(ctx context.Context, notificationHub *v1alpha1.NotificationHub, opts v1.CreateOptions) (*v1alpha1.NotificationHub, error)
	Update(ctx context.Context, notificationHub *v1alpha1.NotificationHub, opts v1.UpdateOptions) (*v1alpha1.NotificationHub, error)
	UpdateStatus(ctx context.Context, notificationHub *v1alpha1.NotificationHub, opts v1.UpdateOptions) (*v1alpha1.NotificationHub, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.NotificationHub, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NotificationHubList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NotificationHub, err error)
	NotificationHubExpansion
}

// notificationHubs implements NotificationHubInterface
type notificationHubs struct {
	client rest.Interface
	ns     string
}

// newNotificationHubs returns a NotificationHubs
func newNotificationHubs(c *NotificationhubV1alpha1Client, namespace string) *notificationHubs {
	return &notificationHubs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the notificationHub, and returns the corresponding notificationHub object, and an error if there is any.
func (c *notificationHubs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NotificationHubs that match those selectors.
func (c *notificationHubs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NotificationHubList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NotificationHubList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested notificationHubs.
func (c *notificationHubs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("notificationhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a notificationHub and creates it.  Returns the server's representation of the notificationHub, and an error, if there is any.
func (c *notificationHubs) Create(ctx context.Context, notificationHub *v1alpha1.NotificationHub, opts v1.CreateOptions) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("notificationhubs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(notificationHub).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a notificationHub and updates it. Returns the server's representation of the notificationHub, and an error, if there is any.
func (c *notificationHubs) Update(ctx context.Context, notificationHub *v1alpha1.NotificationHub, opts v1.UpdateOptions) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(notificationHub.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(notificationHub).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *notificationHubs) UpdateStatus(ctx context.Context, notificationHub *v1alpha1.NotificationHub, opts v1.UpdateOptions) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(notificationHub.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(notificationHub).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the notificationHub and deletes it. Returns an error if one occurs.
func (c *notificationHubs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *notificationHubs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("notificationhubs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched notificationHub.
func (c *notificationHubs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NotificationHub, err error) {
	result = &v1alpha1.NotificationHub{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("notificationhubs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
