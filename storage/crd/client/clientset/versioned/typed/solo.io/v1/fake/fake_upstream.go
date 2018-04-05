/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	solo_io_v1 "github.com/solo-io/gloo-storage/crd/solo.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUpstreams implements UpstreamInterface
type FakeUpstreams struct {
	Fake *FakeGlooV1
	ns   string
}

var upstreamsResource = schema.GroupVersionResource{Group: "gloo.solo.io", Version: "v1", Resource: "upstreams"}

var upstreamsKind = schema.GroupVersionKind{Group: "gloo.solo.io", Version: "v1", Kind: "Upstream"}

// Get takes name of the upstream, and returns the corresponding upstream object, and an error if there is any.
func (c *FakeUpstreams) Get(name string, options v1.GetOptions) (result *solo_io_v1.Upstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(upstreamsResource, c.ns, name), &solo_io_v1.Upstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Upstream), err
}

// List takes label and field selectors, and returns the list of Upstreams that match those selectors.
func (c *FakeUpstreams) List(opts v1.ListOptions) (result *solo_io_v1.UpstreamList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(upstreamsResource, upstreamsKind, c.ns, opts), &solo_io_v1.UpstreamList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &solo_io_v1.UpstreamList{}
	for _, item := range obj.(*solo_io_v1.UpstreamList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested upstreams.
func (c *FakeUpstreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(upstreamsResource, c.ns, opts))

}

// Create takes the representation of a upstream and creates it.  Returns the server's representation of the upstream, and an error, if there is any.
func (c *FakeUpstreams) Create(upstream *solo_io_v1.Upstream) (result *solo_io_v1.Upstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(upstreamsResource, c.ns, upstream), &solo_io_v1.Upstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Upstream), err
}

// Update takes the representation of a upstream and updates it. Returns the server's representation of the upstream, and an error, if there is any.
func (c *FakeUpstreams) Update(upstream *solo_io_v1.Upstream) (result *solo_io_v1.Upstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(upstreamsResource, c.ns, upstream), &solo_io_v1.Upstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Upstream), err
}

// Delete takes name of the upstream and deletes it. Returns an error if one occurs.
func (c *FakeUpstreams) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(upstreamsResource, c.ns, name), &solo_io_v1.Upstream{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUpstreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(upstreamsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &solo_io_v1.UpstreamList{})
	return err
}

// Patch applies the patch and returns the patched upstream.
func (c *FakeUpstreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *solo_io_v1.Upstream, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(upstreamsResource, c.ns, name, data, subresources...), &solo_io_v1.Upstream{})

	if obj == nil {
		return nil, err
	}
	return obj.(*solo_io_v1.Upstream), err
}
