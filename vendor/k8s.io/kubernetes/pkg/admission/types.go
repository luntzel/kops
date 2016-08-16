/*
Copyright 2016 The Kubernetes Authors.

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

package admission

import (
	"k8s.io/kubernetes/pkg/controller/framework/informers"
)

// Validator holds Validate functions, which are responsible for validation of initialized shared resources
// and should be implemented on admission plugins
type Validator interface {
	Validate() error
}

// WantsNamespaceInformer defines a function which sets NamespaceInformer for admission plugins that need it
type WantsInformerFactory interface {
	SetInformerFactory(informers.SharedInformerFactory)
	Validator
}