/*
Copyright 2023 The KubeAdmiral Authors.

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
	pkgruntime "k8s.io/apimachinery/pkg/runtime"
)

type GenericOverridePolicy interface {
	metav1.Object
	pkgruntime.Object
	GetSpec() *GenericOverridePolicySpec
	GetKey() string
	GenericRefCountedPolicy
}

type GenericRefCountedPolicy interface {
	metav1.Object
	pkgruntime.Object
	GetRefCountedStatus() *GenericRefCountedStatus
}

type GenericPropagationPolicy interface {
	metav1.Object
	pkgruntime.Object
	GetSpec() *PropagationPolicySpec
	GenericRefCountedPolicy
}

type GenericFederatedObject interface {
	metav1.Object
	pkgruntime.Object
	GetSpec() *GenericFederatedObjectSpec
	GetStatus() *GenericFederatedObjectStatus
	DeepCopyGenericFederatedObject() GenericFederatedObject
}

type GenericCollectedStatusObject interface {
	metav1.Object
	pkgruntime.Object
	GetGenericCollectedStatus() *GenericCollectedStatus
	GetLastUpdateTime() *metav1.Time
	DeepCopyGenericCollectedStatusObject() GenericCollectedStatusObject
}
