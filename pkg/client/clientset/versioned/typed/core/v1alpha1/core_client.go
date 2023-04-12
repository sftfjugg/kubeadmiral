// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kubewharf/kubeadmiral/pkg/apis/core/v1alpha1"
	"github.com/kubewharf/kubeadmiral/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type CoreV1alpha1Interface interface {
	RESTClient() rest.Interface
	ClusterOverridePoliciesGetter
	ClusterPropagatedVersionsGetter
	ClusterPropagationPoliciesGetter
	FederatedClustersGetter
	FederatedTypeConfigsGetter
	OverridePoliciesGetter
	PropagatedVersionsGetter
	PropagationPoliciesGetter
	SchedulerPluginWebhookConfigurationsGetter
	SchedulingProfilesGetter
}

// CoreV1alpha1Client is used to interact with features provided by the core.kubeadmiral.io group.
type CoreV1alpha1Client struct {
	restClient rest.Interface
}

func (c *CoreV1alpha1Client) ClusterOverridePolicies() ClusterOverridePolicyInterface {
	return newClusterOverridePolicies(c)
}

func (c *CoreV1alpha1Client) ClusterPropagatedVersions() ClusterPropagatedVersionInterface {
	return newClusterPropagatedVersions(c)
}

func (c *CoreV1alpha1Client) ClusterPropagationPolicies() ClusterPropagationPolicyInterface {
	return newClusterPropagationPolicies(c)
}

func (c *CoreV1alpha1Client) FederatedClusters() FederatedClusterInterface {
	return newFederatedClusters(c)
}

func (c *CoreV1alpha1Client) FederatedTypeConfigs() FederatedTypeConfigInterface {
	return newFederatedTypeConfigs(c)
}

func (c *CoreV1alpha1Client) OverridePolicies(namespace string) OverridePolicyInterface {
	return newOverridePolicies(c, namespace)
}

func (c *CoreV1alpha1Client) PropagatedVersions(namespace string) PropagatedVersionInterface {
	return newPropagatedVersions(c, namespace)
}

func (c *CoreV1alpha1Client) PropagationPolicies(namespace string) PropagationPolicyInterface {
	return newPropagationPolicies(c, namespace)
}

func (c *CoreV1alpha1Client) SchedulerPluginWebhookConfigurations() SchedulerPluginWebhookConfigurationInterface {
	return newSchedulerPluginWebhookConfigurations(c)
}

func (c *CoreV1alpha1Client) SchedulingProfiles() SchedulingProfileInterface {
	return newSchedulingProfiles(c)
}

// NewForConfig creates a new CoreV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*CoreV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &CoreV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new CoreV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *CoreV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new CoreV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *CoreV1alpha1Client {
	return &CoreV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *CoreV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
