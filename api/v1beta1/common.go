/*
Copyright 2022.

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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	"github.com/openstack-k8s-operators/lib-common/modules/common/endpoint"
)

// GlanceAPITemplate defines the desired state of GlanceAPI
type GlanceAPITemplate struct {

	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=internal;external
	// +kubebuilder:default=external
	APIType string `json:"apiType"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=1
	// +kubebuilder:validation:Maximum=32
	// +kubebuilder:validation:Minimum=0
	// Replicas of glance API to run
	Replicas int32 `json:"replicas"`

	// +kubebuilder:validation:Required
	// Glance Container Image URL (will be set to environmental default if empty)
	ContainerImage string `json:"containerImage"`

	// +kubebuilder:validation:Optional
	// NodeSelector to target subset of worker nodes running this service
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// +kubebuilder:validation:Optional
	// Debug - enable debug for different deploy stages. If an init container is used, it runs and the
	// actual action pod gets started with sleep infinity
	Debug GlanceAPIDebug `json:"debug,omitempty"`

	// +kubebuilder:validation:Optional
	// Pvc - Storage claim for file-backed Glance
	Pvc string `json:"pvc,omitempty"`

	// +kubebuilder:validation:Optional
	// CustomServiceConfig - customize the service config using this parameter to change service defaults,
	// or overwrite rendered information using raw OpenStack config format. The content gets added to
	// to /etc/<service>/<service>.conf.d directory as custom.conf file.
	CustomServiceConfig string `json:"customServiceConfig,omitempty"`

	// +kubebuilder:validation:Optional
	// ConfigOverwrite - interface to overwrite default config files like e.g. logging.conf or policy.json.
	// But can also be used to add additional files. Those get added to the service config dir in /etc/<service> .
	// TODO: -> implement
	DefaultConfigOverwrite map[string]string `json:"defaultConfigOverwrite,omitempty"`

	// +kubebuilder:validation:Optional
	// Resources - Compute Resources required by this service (Limits/Requests).
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// +kubebuilder:validation:Optional
	// NetworkAttachments is a list of NetworkAttachment resource names to expose the services to the given network
	NetworkAttachments []string `json:"networkAttachments,omitempty"`

	// +kubebuilder:validation:Optional
	// ExternalEndpoints, expose a VIP via MetalLB on the pre-created address pool
	ExternalEndpoints []MetalLBConfig `json:"externalEndpoints,omitempty"`
}

// MetalLBConfig to configure the MetalLB loadbalancer service
type MetalLBConfig struct {
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Enum=internal;public
	// Endpoint, OpenStack endpoint this service maps to
	Endpoint endpoint.Endpoint `json:"endpoint"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinLength=1
	// IPAddressPool expose VIP via MetalLB on the IPAddressPool
	IPAddressPool string `json:"ipAddressPool"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=true
	// SharedIP if true, VIP/VIPs get shared with multiple services
	SharedIP bool `json:"sharedIP"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default=""
	// SharedIPKey specifies the sharing key which gets set as the annotation on the LoadBalancer service.
	// Services which share the same VIP must have the same SharedIPKey. Defaults to the IPAddressPool if
	// SharedIP is true, but no SharedIPKey specified.
	SharedIPKey string `json:"sharedIPKey"`

	// +kubebuilder:validation:Optional
	// LoadBalancerIPs, request given IPs from the pool if available. Using a list to allow dual stack (IPv4/IPv6) support
	LoadBalancerIPs []string `json:"loadBalancerIPs,omitempty"`
}