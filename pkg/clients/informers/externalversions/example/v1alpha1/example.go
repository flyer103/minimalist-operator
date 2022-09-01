/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	examplev1alpha1 "github.com/flyer103/minimalist-operator/pkg/apis/example/v1alpha1"
	versioned "github.com/flyer103/minimalist-operator/pkg/clients/clientset/versioned"
	internalinterfaces "github.com/flyer103/minimalist-operator/pkg/clients/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/flyer103/minimalist-operator/pkg/clients/listers/example/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ExampleInformer provides access to a shared informer and lister for
// Examples.
type ExampleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ExampleLister
}

type exampleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewExampleInformer constructs a new informer for Example type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExampleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExampleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredExampleInformer constructs a new informer for Example type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExampleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProductV1alpha1().Examples(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProductV1alpha1().Examples(namespace).Watch(context.TODO(), options)
			},
		},
		&examplev1alpha1.Example{},
		resyncPeriod,
		indexers,
	)
}

func (f *exampleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExampleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *exampleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&examplev1alpha1.Example{}, f.defaultInformer)
}

func (f *exampleInformer) Lister() v1alpha1.ExampleLister {
	return v1alpha1.NewExampleLister(f.Informer().GetIndexer())
}