/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

This file may have been modified by The KubeAdmiral Authors
("KubeAdmiral Modifications"). All KubeAdmiral Modifications
are Copyright 2023 The KubeAdmiral Authors.
*/

package common

import (
	"fmt"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// QualifiedName comprises a resource name with an optional namespace.
// If namespace is provided, a QualifiedName will be rendered as
// "<namespace>/<name>".  If not, it will be rendered as "name".  This
// is intended to allow the FederatedTypeAdapter interface and its
// consumers to operate on both namespaces and namespace-qualified
// resources.
type QualifiedName struct {
	Namespace string
	Name      string
}

func NewQualifiedName(obj metav1.Object) QualifiedName {
	return QualifiedName{
		Namespace: obj.GetNamespace(),
		Name:      obj.GetName(),
	}
}

// String returns the general purpose string representation
func (n QualifiedName) String() string {
	if len(n.Namespace) == 0 {
		return n.Name
	}
	return fmt.Sprintf("%s/%s", n.Namespace, n.Name)
}

func NewQualifiedFromString(s string) QualifiedName {
	items := strings.Split(s, "/")
	qualifiedName := QualifiedName{}
	if len(items) == 1 {
		qualifiedName.Name = items[0]
	} else {
		qualifiedName.Namespace = items[0]
		qualifiedName.Name = items[1]
	}
	return qualifiedName
}
