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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// CinderPersistentVolumeSourceApplyConfiguration represents an declarative configuration of the CinderPersistentVolumeSource type for use
// with apply.
type CinderPersistentVolumeSourceApplyConfiguration struct {
	VolumeID  *string                            `json:"volumeID,omitempty"`
	FSType    *string                            `json:"fsType,omitempty"`
	ReadOnly  *bool                              `json:"readOnly,omitempty"`
	SecretRef *SecretReferenceApplyConfiguration `json:"secretRef,omitempty"`
}

// CinderPersistentVolumeSourceApplyConfiguration constructs an declarative configuration of the CinderPersistentVolumeSource type for use with
// apply.
func CinderPersistentVolumeSource() *CinderPersistentVolumeSourceApplyConfiguration {
	return &CinderPersistentVolumeSourceApplyConfiguration{}
}

// WithVolumeID sets the VolumeID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the VolumeID field is set to the value of the last call.
func (b *CinderPersistentVolumeSourceApplyConfiguration) WithVolumeID(value string) *CinderPersistentVolumeSourceApplyConfiguration {
	b.VolumeID = &value
	return b
}

// WithFSType sets the FSType field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FSType field is set to the value of the last call.
func (b *CinderPersistentVolumeSourceApplyConfiguration) WithFSType(value string) *CinderPersistentVolumeSourceApplyConfiguration {
	b.FSType = &value
	return b
}

// WithReadOnly sets the ReadOnly field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReadOnly field is set to the value of the last call.
func (b *CinderPersistentVolumeSourceApplyConfiguration) WithReadOnly(value bool) *CinderPersistentVolumeSourceApplyConfiguration {
	b.ReadOnly = &value
	return b
}

// WithSecretRef sets the SecretRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SecretRef field is set to the value of the last call.
func (b *CinderPersistentVolumeSourceApplyConfiguration) WithSecretRef(value *SecretReferenceApplyConfiguration) *CinderPersistentVolumeSourceApplyConfiguration {
	b.SecretRef = value
	return b
}
