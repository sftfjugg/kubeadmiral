// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=core.kubeadmiral.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("clustercollectedstatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().ClusterCollectedStatuses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterfederatedobjects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().ClusterFederatedObjects().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusteroverridepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().ClusterOverridePolicies().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterpropagatedversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().ClusterPropagatedVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusterpropagationpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().ClusterPropagationPolicies().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("collectedstatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().CollectedStatuses().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedClusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedobjects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedObjects().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedtypeconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedTypeConfigs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("overridepolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().OverridePolicies().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("propagatedversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().PropagatedVersions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("propagationpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().PropagationPolicies().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("schedulerpluginwebhookconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().SchedulerPluginWebhookConfigurations().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("schedulingprofiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().SchedulingProfiles().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
