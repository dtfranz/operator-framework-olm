package client

import (
	"context"

	log "github.com/sirupsen/logrus"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"

	"errors"

	"github.com/coreos-inc/alm/apis/alphacatalogentry/v1alpha1"
)

type AlphaCatalogEntryInterface interface {
	UpdateEntry(csv *v1alpha1.AlphaCatalogEntry) (*v1alpha1.AlphaCatalogEntry, error)
}

type AlphaCatalogEntryClient struct {
	*rest.RESTClient
}

var _ AlphaCatalogEntryInterface = &AlphaCatalogEntryClient{}

// NewAlphaCatalogEntryClient creates a client that can interact with the AlphaCatalogEntry resource in k8s api
func NewAlphaCatalogEntryClient(kubeconfig string) (client *AlphaCatalogEntryClient, err error) {
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
	return &AlphaCatalogEntryClient{restClient}, nil
}

func (c *AlphaCatalogEntryClient) UpdateEntry(in *v1alpha1.AlphaCatalogEntry) (result *v1alpha1.AlphaCatalogEntry, err error) {
	result = &v1alpha1.AlphaCatalogEntry{}
	log.Infof("%#v", in)
	err = c.RESTClient.Post().Context(context.TODO()).
		Namespace(in.Namespace).
		Resource(v1alpha1.AlphaCatalogEntryCRDName).
		Name(in.Name).
		Body(in).
		Do().
		Into(result)
	if err == nil {
		return
	}
	if k8serrors.IsAlreadyExists(err) {
		err = c.RESTClient.Delete().Context(context.TODO()).
			Namespace(in.Namespace).
			Resource(v1alpha1.AlphaCatalogEntryCRDName).
			Name(in.Name).
			Body(in).
			Do().
			Into(result)
		if err != nil {
			return
		}

		err = c.RESTClient.Post().Context(context.TODO()).
			Namespace(in.Namespace).
			Resource(v1alpha1.AlphaCatalogEntryCRDName).
			Name(in.Name).
			Body(in).
			Do().
			Into(result)

		if err != nil {
			err = errors.New("failed to update CR status: " + err.Error())
		}
	} else {
		log.Infof("error creating alphacatalogentries %s", err.Error())
	}
	return
}
