package query

import (
	corev1 "k8s.io/api/core/v1"
)

type Operand struct {
	Value     interface{}
	ValueFrom ValueReference
}

type ValueReference struct {
	// Selects a field of the pod: supports metadata.name, metadata.namespace, metadata.labels, metadata.annotations,
	// spec.nodeName, spec.serviceAccountName, status.hostIP, status.podIP.
	// +optional
	FieldRef *corev1.ObjectFieldSelector `json:"fieldRef,omitempty"`
	// Selects a resource of the container: only resources limits and requests
	// (limits.cpu, limits.memory, limits.ephemeral-storage, requests.cpu, requests.memory and requests.ephemeral-storage) are currently supported.
	// +optional
	ResourceFieldRef *corev1.ResourceFieldSelector `json:"resourceFieldRef,omitempty"`
	// Selects a key of a ConfigMap.
	// +optional
	ConfigMapKeyRef *corev1.ConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
	// Selects a key of a secret in the pod's namespace
	// +optional
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef,omitempty"`
}
