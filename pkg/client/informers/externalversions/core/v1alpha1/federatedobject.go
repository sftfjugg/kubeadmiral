// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	corev1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	versioned "github.com/kubewharf/kubeadmiral/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubewharf/kubeadmiral/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/client/listers/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FederatedObjectInformer provides access to a shared informer and lister for
// FederatedObjects.
type FederatedObjectInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FederatedObjectLister
}

type federatedObjectInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedObjectInformer constructs a new informer for FederatedObject type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedObjectInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedObjectInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedObjectInformer constructs a new informer for FederatedObject type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedObjectInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().FederatedObjects(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().FederatedObjects(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1alpha1.FederatedObject{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedObjectInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedObjectInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedObjectInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1alpha1.FederatedObject{}, f.defaultInformer)
}

func (f *federatedObjectInformer) Lister() v1alpha1.FederatedObjectLister {
	return v1alpha1.NewFederatedObjectLister(f.Informer().GetIndexer())
}
