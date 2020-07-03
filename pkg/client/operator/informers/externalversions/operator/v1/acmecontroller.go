// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	operatorv1 "github.com/tnozicka/openshift-acme/pkg/api/operator/v1"
	internalinterfaces "github.com/tnozicka/openshift-acme/pkg/client/operator/informers/externalversions/internalinterfaces"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
	v1 "k8s.io/kubernetes/pkg/client/listers/operator/v1"
)

// ACMEControllerInformer provides access to a shared informer and lister for
// ACMEControllers.
type ACMEControllerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ACMEControllerLister
}

type aCMEControllerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewACMEControllerInformer constructs a new informer for ACMEController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewACMEControllerInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredACMEControllerInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredACMEControllerInformer constructs a new informer for ACMEController type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredACMEControllerInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().ACMEControllers().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorV1().ACMEControllers().Watch(context.TODO(), options)
			},
		},
		&operatorv1.ACMEController{},
		resyncPeriod,
		indexers,
	)
}

func (f *aCMEControllerInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredACMEControllerInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *aCMEControllerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorv1.ACMEController{}, f.defaultInformer)
}

func (f *aCMEControllerInformer) Lister() v1.ACMEControllerLister {
	return v1.NewACMEControllerLister(f.Informer().GetIndexer())
}