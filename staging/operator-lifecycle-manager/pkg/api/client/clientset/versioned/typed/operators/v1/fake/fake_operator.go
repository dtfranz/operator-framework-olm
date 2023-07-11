/*
Copyright Red Hat, Inc.

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

	v1 "github.com/operator-framework/api/pkg/operators/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeOperators implements OperatorInterface
type FakeOperators struct {
	Fake *FakeOperatorsV1
}

var operatorsResource = v1.SchemeGroupVersion.WithResource("operators")

var operatorsKind = v1.SchemeGroupVersion.WithKind("Operator")

// Get takes name of the operator, and returns the corresponding operator object, and an error if there is any.
func (c *FakeOperators) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(operatorsResource, name), &v1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Operator), err
}

// List takes label and field selectors, and returns the list of Operators that match those selectors.
func (c *FakeOperators) List(ctx context.Context, opts metav1.ListOptions) (result *v1.OperatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(operatorsResource, operatorsKind, opts), &v1.OperatorList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.OperatorList{ListMeta: obj.(*v1.OperatorList).ListMeta}
	for _, item := range obj.(*v1.OperatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested operators.
func (c *FakeOperators) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(operatorsResource, opts))
}

// Create takes the representation of a operator and creates it.  Returns the server's representation of the operator, and an error, if there is any.
func (c *FakeOperators) Create(ctx context.Context, operator *v1.Operator, opts metav1.CreateOptions) (result *v1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(operatorsResource, operator), &v1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Operator), err
}

// Update takes the representation of a operator and updates it. Returns the server's representation of the operator, and an error, if there is any.
func (c *FakeOperators) Update(ctx context.Context, operator *v1.Operator, opts metav1.UpdateOptions) (result *v1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(operatorsResource, operator), &v1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Operator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeOperators) UpdateStatus(ctx context.Context, operator *v1.Operator, opts metav1.UpdateOptions) (*v1.Operator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(operatorsResource, "status", operator), &v1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Operator), err
}

// Delete takes name of the operator and deletes it. Returns an error if one occurs.
func (c *FakeOperators) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(operatorsResource, name, opts), &v1.Operator{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeOperators) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(operatorsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.OperatorList{})
	return err
}

// Patch applies the patch and returns the patched operator.
func (c *FakeOperators) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Operator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(operatorsResource, name, pt, data, subresources...), &v1.Operator{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Operator), err
}
