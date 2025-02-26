/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

	machine "github.com/gardener/machine-controller-manager/pkg/apis/machine"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMachines implements MachineInterface
type FakeMachines struct {
	Fake *FakeMachine
	ns   string
}

var machinesResource = schema.GroupVersionResource{Group: "machine.sapcloud.io", Version: "", Resource: "machines"}

var machinesKind = schema.GroupVersionKind{Group: "machine.sapcloud.io", Version: "", Kind: "Machine"}

// Get takes name of the machine, and returns the corresponding machine object, and an error if there is any.
func (c *FakeMachines) Get(ctx context.Context, name string, options v1.GetOptions) (result *machine.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(machinesResource, c.ns, name), &machine.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.Machine), err
}

// List takes label and field selectors, and returns the list of Machines that match those selectors.
func (c *FakeMachines) List(ctx context.Context, opts v1.ListOptions) (result *machine.MachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(machinesResource, machinesKind, c.ns, opts), &machine.MachineList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machine.MachineList{ListMeta: obj.(*machine.MachineList).ListMeta}
	for _, item := range obj.(*machine.MachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested machines.
func (c *FakeMachines) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(machinesResource, c.ns, opts))

}

// Create takes the representation of a machine and creates it.  Returns the server's representation of the machine, and an error, if there is any.
func (c *FakeMachines) Create(ctx context.Context, machine *machine.Machine, opts v1.CreateOptions) (result *machine.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(machinesResource, c.ns, machine), &machine.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.Machine), err
}

// Update takes the representation of a machine and updates it. Returns the server's representation of the machine, and an error, if there is any.
func (c *FakeMachines) Update(ctx context.Context, machine *machine.Machine, opts v1.UpdateOptions) (result *machine.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(machinesResource, c.ns, machine), &machine.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.Machine), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMachines) UpdateStatus(ctx context.Context, machine *machine.Machine, opts v1.UpdateOptions) (*machine.Machine, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(machinesResource, "status", c.ns, machine), &machine.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.Machine), err
}

// Delete takes name of the machine and deletes it. Returns an error if one occurs.
func (c *FakeMachines) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(machinesResource, c.ns, name), &machine.Machine{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMachines) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(machinesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &machine.MachineList{})
	return err
}

// Patch applies the patch and returns the patched machine.
func (c *FakeMachines) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machine.Machine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(machinesResource, c.ns, name, pt, data, subresources...), &machine.Machine{})

	if obj == nil {
		return nil, err
	}
	return obj.(*machine.Machine), err
}
