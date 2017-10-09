// Package v1alpha1 implements all the required types and methods for parsing
// resources for v1alpha1 versioned ClusterServiceVersions.
package v1alpha1

import (
	"encoding/json"

	"github.com/coreos/go-semver/semver"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	GroupVersion = "v1alpha1" // used in registering ClusterServiceVersion scheme

	ClusterServiceVersionCRDName       = "clusterserviceversion-v1s.app.coreos.com"
	ClusterServiceVersionCRDAPIVersion = "apiextensions.k8s.io/v1beta1" // API version w/ CRD support

)

// NamedInstallStrategy represents the block of an ClusterServiceVersion resource
// where the install strategy is specified.
type NamedInstallStrategy struct {
	StrategyName    string          `json:"strategy"`
	StrategySpecRaw json.RawMessage `json:"spec"`
}

// CustomResourceDefinitions declares all of the CRDs managed or required by
// an operator being ran by ClusterServiceVersion.
//
// If the CRD is present in the Owned list, it is implicitly required.
type CustomResourceDefinitions struct {
	Owned    []string `json:"owned"`
	Required []string `json:"required"`
}

// ClusterServiceVersionSpec declarations tell the ALM how to install an operator
// that can manage apps for given version and AppType.
type ClusterServiceVersionSpec struct {
	InstallStrategy           NamedInstallStrategy      `json:"install"`
	Version                   semver.Version            `json:"version"`
	Maturity                  string                    `json:"maturity"`
	CustomResourceDefinitions CustomResourceDefinitions `json:"customresourcedefinitions"`
	DisplayName               string                    `json:"displayName"`
	Description               string                    `json:"description"`
	Keywords                  []string                  `json:"keywords"`
	Maintainers               []Maintainer              `json:"maintainers"`
	Links                     []AppLink                 `json:"links"`
	Icon                      Icon                      `json:"icon"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,11,rep,name=labels"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,12,rep,name=annotations"`

	// Label selector for pods. Existing ReplicaSets whose pods are
	// selected by this will be the ones affected by this deployment.
	// +optional
	Selector *metav1.LabelSelector `json:"selector,omitempty" protobuf:"bytes,2,opt,name=selector"`
}

type Maintainer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AppLink struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Icon struct {
	Data      string `json:"base64data"`
	MediaType string `json:"mediatype"`
}

// ClusterServiceVersionPhase is a label for the condition of a ClusterServiceVersion at the current time.
type ClusterServiceVersionPhase string

// These are the valid phases of ClusterServiceVersion
const (
	CSVPhaseNone = ""
	// CSVPending means the csv has been accepted by the system, but the install strategy has not been attempted.
	// This is likely because there are unmet requirements.
	CSVPhasePending ClusterServiceVersionPhase = "Pending"
	// CSVRunning means that the requirements are met but the install strategy has not been run.
	CSVPhaseInstalling ClusterServiceVersionPhase = "Installing"
	// CSVSucceeded means that the resources in the CSV were created successfully.
	CSVPhaseSucceeded ClusterServiceVersionPhase = "Succeeded"
	// CSVFailed means that the install strategy could not be successfully completed.
	CSVPhaseFailed ClusterServiceVersionPhase = "Failed"
	// CSVUnknown means that for some reason the state of the csv could not be obtained.
	CSVPhaseUnknown ClusterServiceVersionPhase = "Unknown"
)

// ConditionReason is a camelcased reason for the state transition
type ConditionReason string

const (
	CSVReasonRequirementsUnkown ConditionReason = "RequirementsUnknown"
	CSVReasonRequirementsNotMet ConditionReason = "RequirementsNotMet"
	CSVReasonRequirementsMet    ConditionReason = "AllRequirementsMet"
	CSVReasonComponentFailed    ConditionReason = "InstallComponentFailed"
	CSVReasonInstallSuccessful  ConditionReason = "InstallSucceeded"
	CSVReasonInstallCheckFailed ConditionReason = "InstallCheckFailed"
)

// Conditions appear in the status as a record of state transitions on the ClusterServiceVersion
type ClusterServiceVersionCondition struct {
	// Condition of the ClusterServiceVersion
	Phase ClusterServiceVersionPhase `json:"phase,omitempty"`
	// A human readable message indicating details about why the ClusterServiceVersion is in this condition.
	// +optional
	Message string `json:"message,omitempty"`
	// A brief CamelCase message indicating details about why the ClusterServiceVersion is in this state.
	// e.g. 'RequirementsNotMet'
	// +optional
	Reason ConditionReason `json:"reason,omitempty"`
	// Last time we updated the status
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
	// Last time the status transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

type RequirementStatus struct {
	Group   string `json:"group"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	UUID    string `json:"uuid,omitempty"`
}

// ClusterServiceVersionStatus represents information about the status of a pod. Status may trail the actual
// state of a system.
type ClusterServiceVersionStatus struct {
	// Current condition of the ClusterServiceVersion
	Phase ClusterServiceVersionPhase `json:"phase,omitempty"`
	// A human readable message indicating details about why the ClusterServiceVersion is in this condition.
	// +optional
	Message string `json:"message,omitempty"`
	// A brief CamelCase message indicating details about why the ClusterServiceVersion is in this state.
	// e.g. 'RequirementsNotMet'
	// +optional
	Reason ConditionReason `json:"reason,omitempty"`
	// Last time we updated the status
	// +optional
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`
	// Last time the status transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// List of conditions, a history of state transitions
	Conditions []ClusterServiceVersionCondition `json:"conditions,omitempty"`
	// The status of each requirement for this CSV
	RequirementStatus []RequirementStatus `json:"requirementStatus,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ClusterServiceVersion is a Custom Resource of type `ClusterServiceVersionSpec`.
type ClusterServiceVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Spec   ClusterServiceVersionSpec   `json:"spec"`
	Status ClusterServiceVersionStatus `json:"status"`
}

// ClusterServiceVersionList represents a list of ClusterServiceVersions.
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterServiceVersionList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`

	Items []ClusterServiceVersion `json:"items"`
}
