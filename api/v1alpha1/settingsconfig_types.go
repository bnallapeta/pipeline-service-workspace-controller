package v1alpha1

import (
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cfg "sigs.k8s.io/controller-runtime/pkg/config/v1alpha1"
)

// NOTE: json tags are required. Any new fields you add must have json tags for the fields to be serialized.

type SettingsNetPolConfig struct {
	// List of egress rules to be applied to the selected pods. Outgoing traffic is
	// allowed if there are no NetworkPolicies selecting the pod (and cluster policy
	// otherwise allows the traffic), OR if the traffic matches at least one egress rule
	// across all of the NetworkPolicy objects whose podSelector matches the pod. If
	// this field is empty then this NetworkPolicy limits all outgoing traffic (and serves
	// solely to ensure that the pods it selects are isolated by default).
	// This field is beta-level in 1.8
	// +optional
	Egress []netv1.NetworkPolicyEgressRule `json:"egress,omitempty" protobuf:"bytes,3,rep,name=egress"`
}

//+kubebuilder:object:root=true

// SettingsConfig is the Schema for the settingsconfigs API
type SettingsConfig struct {
	metav1.TypeMeta `json:",inline"`

	// ControllerManagerConfigurationSpec returns the generic configuration for controllers
	cfg.ControllerManagerConfigurationSpec `json:",inline"`

	// Namespace defines the space within which each name must be unique. An empty namespace is
	// equivalent to the "default" namespace, but "default" is the canonical representation.
	// Not all objects are required to be scoped to a namespace - the value of this field for
	// those objects will be empty.
	//
	// Must be a DNS_LABEL.
	// Cannot be updated.
	// More info: http://kubernetes.io/docs/user-guide/namespaces
	// +optional
	Namespace string `json:"namespace,omitempty"`

	NetPolConfig SettingsNetPolConfig `json:"networkPolicyConfig,omitempty"`
}

//+kubebuilder:object:root=true

func init() {
	SchemeBuilder.Register(&SettingsConfig{})
}