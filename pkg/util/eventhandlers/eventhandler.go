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

package eventhandlers

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
)

// NewTriggerOnAllChanges returns a cache.ResourceEventHandlerFuncs that will call the given triggerFunc on all object
// changes. triggerFunc is also called for add or delete events.
func NewTriggerOnAllChanges[Source any](triggerFunc func(Source)) *cache.ResourceEventHandlerFuncs {
	return &cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(old interface{}) {
			if deleted, ok := old.(cache.DeletedFinalStateUnknown); ok {
				old = deleted.Obj
				if old == nil {
					return
				}
			}
			oldSource := old.(Source)
			triggerFunc(oldSource)
		},
		AddFunc: func(cur interface{}) {
			curObj := cur.(Source)
			triggerFunc(curObj)
		},
		UpdateFunc: func(old, cur interface{}) {
			if !reflect.DeepEqual(old, cur) {
				curObj := cur.(Source)
				triggerFunc(curObj)
			}
		},
	}
}

// NewTriggerOnAllChangesWithTransform returns a cache.ResourceEventHandler that will call the given triggerFunc on all
// object changes. triggerFunc is also called for add or delete events. transformFunc will first be used to transform
// the original object into the target type.
func NewTriggerOnAllChangesWithTransform[Source any, Target any](
	transformFunc func(Source) Target,
	triggerFunc func(Target),
) cache.ResourceEventHandler {
	return &cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(old interface{}) {
			if deleted, ok := old.(cache.DeletedFinalStateUnknown); ok {
				old = deleted.Obj
				if old == nil {
					return
				}
			}
			oldSource := transformFunc(old.(Source))
			triggerFunc(oldSource)
		},
		AddFunc: func(cur interface{}) {
			curObj := transformFunc(cur.(Source))
			triggerFunc(curObj)
		},
		UpdateFunc: func(old, cur interface{}) {
			if !reflect.DeepEqual(old, cur) {
				curObj := transformFunc(cur.(Source))
				triggerFunc(curObj)
			}
		},
	}
}

// NewTriggerOnChanges returns a cache.ResourceEventHandlerFuncs that will call the given triggerFunc on object changes
// that passes the given predicate. triggerFunc is also called for add and delete events.
func NewTriggerOnChanges[Source any](
	predicate func(old, cur Source) bool,
	triggerFunc func(Source),
) *cache.ResourceEventHandlerFuncs {
	return &cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(old interface{}) {
			if deleted, ok := old.(cache.DeletedFinalStateUnknown); ok {
				// This object might be stale but ok for our current usage.
				old = deleted.Obj
				if old == nil {
					return
				}
			}
			oldObj := old.(Source)
			triggerFunc(oldObj)
		},
		AddFunc: func(cur interface{}) {
			curObj := cur.(Source)
			triggerFunc(curObj)
		},
		UpdateFunc: func(old, cur interface{}) {
			oldObj := old.(Source)
			curObj := cur.(Source)
			if predicate(oldObj, curObj) {
				triggerFunc(curObj)
			}
		},
	}
}

// NewTriggerOnGenerationChanges returns a cache.ResourceEventHandlerFuncs that will call the given triggerFunc on
// object generation changes. triggerFunc is also called for add and delete events.
func NewTriggerOnGenerationChanges[Source any](triggerFunc func(Source)) *cache.ResourceEventHandlerFuncs {
	return &cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(old interface{}) {
			if deleted, ok := old.(cache.DeletedFinalStateUnknown); ok {
				// This object might be stale but ok for our current usage.
				old = deleted.Obj
				if old == nil {
					return
				}
			}
			oldObj := old.(Source)
			triggerFunc(oldObj)
		},
		AddFunc: func(cur interface{}) {
			curObj := cur.(Source)
			triggerFunc(curObj)
		},
		UpdateFunc: func(old, cur interface{}) {
			oldObj := old.(metav1.Object)
			curObj := cur.(metav1.Object)

			if oldObj.GetGeneration() != curObj.GetGeneration() {
				triggerFunc(cur.(Source))
			}
		},
	}
}
