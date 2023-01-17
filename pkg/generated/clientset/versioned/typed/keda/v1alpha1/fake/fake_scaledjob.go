/*
Copyright 2023 The KEDA Authors

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

package fake

import (
	"context"

	v1alpha1 "github.com/kedacore/keda/v2/apis/keda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScaledJobs implements ScaledJobInterface
type FakeScaledJobs struct {
	Fake *FakeKedaV1alpha1
	ns   string
}

var scaledjobsResource = schema.GroupVersionResource{Group: "keda", Version: "v1alpha1", Resource: "scaledjobs"}

var scaledjobsKind = schema.GroupVersionKind{Group: "keda", Version: "v1alpha1", Kind: "ScaledJob"}

// Get takes name of the scaledJob, and returns the corresponding scaledJob object, and an error if there is any.
func (c *FakeScaledJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScaledJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(scaledjobsResource, c.ns, name), &v1alpha1.ScaledJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaledJob), err
}

// List takes label and field selectors, and returns the list of ScaledJobs that match those selectors.
func (c *FakeScaledJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScaledJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(scaledjobsResource, scaledjobsKind, c.ns, opts), &v1alpha1.ScaledJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScaledJobList{ListMeta: obj.(*v1alpha1.ScaledJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScaledJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scaledJobs.
func (c *FakeScaledJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(scaledjobsResource, c.ns, opts))

}

// Create takes the representation of a scaledJob and creates it.  Returns the server's representation of the scaledJob, and an error, if there is any.
func (c *FakeScaledJobs) Create(ctx context.Context, scaledJob *v1alpha1.ScaledJob, opts v1.CreateOptions) (result *v1alpha1.ScaledJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(scaledjobsResource, c.ns, scaledJob), &v1alpha1.ScaledJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaledJob), err
}

// Update takes the representation of a scaledJob and updates it. Returns the server's representation of the scaledJob, and an error, if there is any.
func (c *FakeScaledJobs) Update(ctx context.Context, scaledJob *v1alpha1.ScaledJob, opts v1.UpdateOptions) (result *v1alpha1.ScaledJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(scaledjobsResource, c.ns, scaledJob), &v1alpha1.ScaledJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaledJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScaledJobs) UpdateStatus(ctx context.Context, scaledJob *v1alpha1.ScaledJob, opts v1.UpdateOptions) (*v1alpha1.ScaledJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(scaledjobsResource, "status", c.ns, scaledJob), &v1alpha1.ScaledJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaledJob), err
}

// Delete takes name of the scaledJob and deletes it. Returns an error if one occurs.
func (c *FakeScaledJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(scaledjobsResource, c.ns, name, opts), &v1alpha1.ScaledJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScaledJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(scaledjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScaledJobList{})
	return err
}

// Patch applies the patch and returns the patched scaledJob.
func (c *FakeScaledJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScaledJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(scaledjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ScaledJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScaledJob), err
}
