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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	settings_v1alpha1 "k8s.io/api/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	kubernetes "k8s.io/client-go/kubernetes"
	v1alpha1 "k8s.io/client-go/listers/settings/v1alpha1"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// PodPresetInformer provides access to a shared informer and lister for
// PodPresets.
type PodPresetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodPresetLister
}

type podPresetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodPresetInformer constructs a new informer for PodPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodPresetInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodPresetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodPresetInformer constructs a new informer for PodPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodPresetInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().PodPresets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().PodPresets(namespace).Watch(options)
			},
		},
		&settings_v1alpha1.PodPreset{},
		resyncPeriod,
		indexers,
	)
}

func (f *podPresetInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodPresetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podPresetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&settings_v1alpha1.PodPreset{}, f.defaultInformer)
}

func (f *podPresetInformer) Lister() v1alpha1.PodPresetLister {
	return v1alpha1.NewPodPresetLister(f.Informer().GetIndexer())
}
