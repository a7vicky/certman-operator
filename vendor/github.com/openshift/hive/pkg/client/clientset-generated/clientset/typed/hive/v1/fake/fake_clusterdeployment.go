// Code generated by main. DO NOT EDIT.

package fake

import (
	hivev1 "github.com/openshift/hive/pkg/apis/hive/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterDeployments implements ClusterDeploymentInterface
type FakeClusterDeployments struct {
	Fake *FakeHiveV1
	ns   string
}

var clusterdeploymentsResource = schema.GroupVersionResource{Group: "hive.openshift.io", Version: "v1", Resource: "clusterdeployments"}

var clusterdeploymentsKind = schema.GroupVersionKind{Group: "hive.openshift.io", Version: "v1", Kind: "ClusterDeployment"}

// Get takes name of the clusterDeployment, and returns the corresponding clusterDeployment object, and an error if there is any.
func (c *FakeClusterDeployments) Get(name string, options v1.GetOptions) (result *hivev1.ClusterDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterdeploymentsResource, c.ns, name), &hivev1.ClusterDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterDeployment), err
}

// List takes label and field selectors, and returns the list of ClusterDeployments that match those selectors.
func (c *FakeClusterDeployments) List(opts v1.ListOptions) (result *hivev1.ClusterDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterdeploymentsResource, clusterdeploymentsKind, c.ns, opts), &hivev1.ClusterDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hivev1.ClusterDeploymentList{ListMeta: obj.(*hivev1.ClusterDeploymentList).ListMeta}
	for _, item := range obj.(*hivev1.ClusterDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterDeployments.
func (c *FakeClusterDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterdeploymentsResource, c.ns, opts))

}

// Create takes the representation of a clusterDeployment and creates it.  Returns the server's representation of the clusterDeployment, and an error, if there is any.
func (c *FakeClusterDeployments) Create(clusterDeployment *hivev1.ClusterDeployment) (result *hivev1.ClusterDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterdeploymentsResource, c.ns, clusterDeployment), &hivev1.ClusterDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterDeployment), err
}

// Update takes the representation of a clusterDeployment and updates it. Returns the server's representation of the clusterDeployment, and an error, if there is any.
func (c *FakeClusterDeployments) Update(clusterDeployment *hivev1.ClusterDeployment) (result *hivev1.ClusterDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterdeploymentsResource, c.ns, clusterDeployment), &hivev1.ClusterDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterDeployment), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterDeployments) UpdateStatus(clusterDeployment *hivev1.ClusterDeployment) (*hivev1.ClusterDeployment, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterdeploymentsResource, "status", c.ns, clusterDeployment), &hivev1.ClusterDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterDeployment), err
}

// Delete takes name of the clusterDeployment and deletes it. Returns an error if one occurs.
func (c *FakeClusterDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(clusterdeploymentsResource, c.ns, name), &hivev1.ClusterDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterdeploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &hivev1.ClusterDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched clusterDeployment.
func (c *FakeClusterDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hivev1.ClusterDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterdeploymentsResource, c.ns, name, pt, data, subresources...), &hivev1.ClusterDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hivev1.ClusterDeployment), err
}
