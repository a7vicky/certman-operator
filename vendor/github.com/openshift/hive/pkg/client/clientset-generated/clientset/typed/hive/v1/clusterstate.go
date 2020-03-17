// Code generated by main. DO NOT EDIT.

package v1

import (
	"time"

	v1 "github.com/openshift/hive/pkg/apis/hive/v1"
	scheme "github.com/openshift/hive/pkg/client/clientset-generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterStatesGetter has a method to return a ClusterStateInterface.
// A group's client should implement this interface.
type ClusterStatesGetter interface {
	ClusterStates(namespace string) ClusterStateInterface
}

// ClusterStateInterface has methods to work with ClusterState resources.
type ClusterStateInterface interface {
	Create(*v1.ClusterState) (*v1.ClusterState, error)
	Update(*v1.ClusterState) (*v1.ClusterState, error)
	UpdateStatus(*v1.ClusterState) (*v1.ClusterState, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterState, error)
	List(opts metav1.ListOptions) (*v1.ClusterStateList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterState, err error)
	ClusterStateExpansion
}

// clusterStates implements ClusterStateInterface
type clusterStates struct {
	client rest.Interface
	ns     string
}

// newClusterStates returns a ClusterStates
func newClusterStates(c *HiveV1Client, namespace string) *clusterStates {
	return &clusterStates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterState, and returns the corresponding clusterState object, and an error if there is any.
func (c *clusterStates) Get(name string, options metav1.GetOptions) (result *v1.ClusterState, err error) {
	result = &v1.ClusterState{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterstates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterStates that match those selectors.
func (c *clusterStates) List(opts metav1.ListOptions) (result *v1.ClusterStateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ClusterStateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterstates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterStates.
func (c *clusterStates) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterstates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterState and creates it.  Returns the server's representation of the clusterState, and an error, if there is any.
func (c *clusterStates) Create(clusterState *v1.ClusterState) (result *v1.ClusterState, err error) {
	result = &v1.ClusterState{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterstates").
		Body(clusterState).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterState and updates it. Returns the server's representation of the clusterState, and an error, if there is any.
func (c *clusterStates) Update(clusterState *v1.ClusterState) (result *v1.ClusterState, err error) {
	result = &v1.ClusterState{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterstates").
		Name(clusterState.Name).
		Body(clusterState).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterStates) UpdateStatus(clusterState *v1.ClusterState) (result *v1.ClusterState, err error) {
	result = &v1.ClusterState{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterstates").
		Name(clusterState.Name).
		SubResource("status").
		Body(clusterState).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterState and deletes it. Returns an error if one occurs.
func (c *clusterStates) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterstates").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterStates) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterstates").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterState.
func (c *clusterStates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterState, err error) {
	result = &v1.ClusterState{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterstates").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
