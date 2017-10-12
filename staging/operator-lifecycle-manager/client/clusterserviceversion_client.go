package client

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"

	"github.com/coreos-inc/alm/apis/clusterserviceversion/v1alpha1"
)

type ClusterServiceVersionInterface interface {
	UpdateCSV(csv *v1alpha1.ClusterServiceVersion) (result *v1alpha1.ClusterServiceVersion, err error)
}

type ClusterServiceVersionClient struct {
	*rest.RESTClient
}

var _ ClusterServiceVersionInterface = &ClusterServiceVersionClient{}

// NewClusterServiceVersionClient creates a client that can interact with the ClusterServiceVersion resource in k8s api
func NewClusterServiceVersionClient(kubeconfig string) (client *ClusterServiceVersionClient, err error) {
	var config *rest.Config
	config, err = getConfig(kubeconfig)
	if err != nil {
		return
	}

	scheme := runtime.NewScheme()
	if err = v1alpha1.AddToScheme(scheme); err != nil {
		return
	}

	config.GroupVersion = &v1alpha1.SchemeGroupVersion
	config.APIPath = "/apis"
	config.ContentType = runtime.ContentTypeJSON
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(scheme)}

	var restClient *rest.RESTClient
	restClient, err = rest.RESTClientFor(config)
	if err != nil {
		return
	}

	return &ClusterServiceVersionClient{restClient}, nil
}

func (c *ClusterServiceVersionClient) UpdateCSV(in *v1alpha1.ClusterServiceVersion) (out *v1alpha1.ClusterServiceVersion, err error) {
	out = &v1alpha1.ClusterServiceVersion{}
	if err = c.RESTClient.
		Put().
		Context(context.TODO()).
		Namespace(in.Namespace).
		Resource("clusterserviceversion-v1s").
		Name(in.Name).
		Body(in).
		Do().
		Into(out); err != nil {
		err = errors.New("failed to update CR status: " + err.Error())
	}

	return
}
