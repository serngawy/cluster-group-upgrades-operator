/*
Copyright 2021.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SiteConfigSpec defines the desired state of SiteConfig
type SiteConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	PullSecretRef          PullSecretRef          `json:"pullSecretRef,omitempty" yaml:"pullSecretRef,omitempty"`
	ClusterImageSetNameRef string                 `json:"clusterImageSetNameRef,omitempty" yaml:"clusterImageSetNameRef,omitempty"`
	SshPublicKey           string                 `json:"sshPublicKey,omitempty" yaml:"sshPublicKey,omitempty"`
	SshPrivateKeySecretRef SshPrivateKeySecretRef `json:"sshPrivateKeySecretRef,omitempty" yaml:"sshPrivateKeySecretRef,omitempty"`
	Clusters               []Clusters             `json:"clusters,omitempty" yaml:"clusters,omitempty"`
	BaseDomain             string                 `json:"baseDomain,omitempty" yaml:"baseDomain,omitempty"`
}

// PullSecretRef
type PullSecretRef struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// SshPrivateKeySecretRef
type SshPrivateKeySecretRef struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Clusters
type Clusters struct {
	ApiVIP                 string            `json:"apiVIP,omitempty" yaml:"apiVIP,omitempty"`
	IngressVIP             string            `json:"ingressVIP,omitempty" yaml:"ingressVIP,omitempty"`
	ClusterName            string            `json:"clusterName,omitempty" yaml:"clusterName,omitempty"`
	AdditionalNTPSources   []string          `json:"additionalNTPSources,omitempty" yaml:"additionalNTPSources,omitempty"`
	Nodes                  []Node            `json:"nodes,omitempty" yaml:"nodes,omitempty"`
	MachineNetwork         []MachineNetwork  `json:"machineNetwork,omitempty" yaml:"machineNetwork,omitempty"`
	ServiceNetwork         []string          `json:"serviceNetwork,omitempty" yaml:"serviceNetwork,omitempty"`
	ManifestsConfig        ManifestsConfig   `json:"manifestsConfig,omitempty" yaml:"manifestsConfig,omitempty"`
	ClusterType            string            `json:"clusterType,omitempty" yaml:"clusterType,omitempty"`
	NumMasters             uint8             `json:"numMasters,omitempty" yaml:"numMasters,omitempty"`
	NumWorkers             uint8             `json:"numWorkers,omitempty" yaml:"numWorkers,omitempty"`
	ClusterProfile         string            `json:"clusterProfile,omitempty" yaml:"clusterProfile,omitempty"`
	ClusterLabels          map[string]string `json:"clusterLabels,omitempty" yaml:"clusterLabels,omitempty"`
	NetworkType            string            `json:"networkType,omitempty" yaml:"networkType,omitempty"`
	ClusterNetwork         []ClusterNetwork  `json:"clusterNetwork,omitempty" yaml:"clusterNetwork,omitempty"`
	IgnitionConfigOverride string            `json:"ignitionConfigOverride,omitempty" yaml:"ignitionConfigOverride,omitempty"`
	DiskEncryption         DiskEncryption    `json:"diskEncryption,omitempty" yaml:"diskEncryption,omitempty"`
	ProxySettings          ProxySettings     `json:"proxy,omitempty" yaml:"proxy,omitempty"`
}

type DiskEncryption struct {
	Type string       `json:"type,omitempty" yaml:"type,omitempty"`
	Tang []TangConfig `json:"tang,omitempty" yaml:"tang,omitempty"`
}

type ProxySettings struct {
	HttpProxy  string `json:"httpProxy,omitempty" yaml:"httpProxy,omitempty"`
	HttpsProxy string `json:"httpsProxy,omitempty" yaml:"httpsProxy,omitempty"`
	NoProxy    string `json:"noProxy,omitempty" yaml:"noProxy,omitempty"`
}

type TangConfig struct {
	URL        string `yaml:"url,omitempty" json:"url,omitempty"`
	Thumbprint string `yaml:"thumbprint" json:"thumbprint,omitempty"`
}

// Node
type Node struct {
	BmcAddress             string             `json:"bmcAddress,omitempty" yaml:"bmcAddress,omitempty"`
	BootMACAddress         string             `json:"bootMACAddress,omitempty" yaml:"bootMACAddress,omitempty"`
	RootDeviceHints        map[string]string  `json:"rootDeviceHints,omitempty" yaml:"rootDeviceHints,omitempty"`
	Cpuset                 string             `json:"cpuset,omitempty" yaml:"cpuset,omitempty"`
	NodeNetwork            NodeNetwork        `json:"nodeNetwork,omitempty" yaml:"nodeNetwork,omitempty"`
	HostName               string             `json:"hostName,omitempty" yaml:"hostName,omitempty"`
	BmcCredentialsName     BmcCredentialsName `json:"bmcCredentialsName,omitempty" yaml:"bmcCredentialsName,omitempty"`
	BootMode               string             `json:"bootMode,omitempty" yaml:"bootMode,omitempty"`
	UserData               map[string]string  `json:"userData,omitempty" yaml:"userData,omitempty"`
	InstallerArgs          string             `json:"installerArgs,omitempty" yaml:"installerArgs,omitempty"`
	IgnitionConfigOverride string             `json:"ignitionConfigOverride,omitempty" yaml:"ignitionConfigOverride,omitempty"`
	Role                   string             `json:"role,omitempty" yaml:"role,omitempty"`
}

// MachineNetwork
type MachineNetwork struct {
	Cidr string `json:"cidr,omitempty" yaml:"cidr,omitempty"`
}

// ClusterNetwork
type ClusterNetwork struct {
	Cidr       string `json:"cidr,omitempty" yaml:"cidr,omitempty"`
	HostPrefix int    `json:"hostPrefix,omitempty" yaml:"hostPrefix,omitempty"`
}

// ManifestsConfig
type ManifestsConfig struct {
	NtpServer string `json:"ntpServer,omitempty" yaml:"ntpServer,omitempty"`
}

// NodeNetwork
type NodeNetwork struct {
	Config     []byte      `json:"config,omitempty" yaml:"config,omitempty"`
	Interfaces []Interface `json:"interfaces,omitempty" yaml:"interfaces,omitempty"`
}

// BmcCredentialsName
type BmcCredentialsName struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}

// Interfaces
type Interface struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	MacAddress string `json:"macAddress,omitempty" yaml:"macAddress,omitempty"`
}

// SiteConfigStatus defines the observed state of SiteConfig
type SiteConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true

// SiteConfig is the Schema for the siteconfigs API

// +kubebuilder:subresource:status
type SiteConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SiteConfigSpec   `json:"spec,omitempty"`
	Status SiteConfigStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// SiteConfigList contains a list of SiteConfig
type SiteConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SiteConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SiteConfig{}, &SiteConfigList{})
}
