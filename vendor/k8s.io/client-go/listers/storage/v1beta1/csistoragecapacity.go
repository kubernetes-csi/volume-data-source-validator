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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "k8s.io/api/storage/v1beta1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// CSIStorageCapacityLister helps list CSIStorageCapacities.
// All objects returned here must be treated as read-only.
type CSIStorageCapacityLister interface {
	// List lists all CSIStorageCapacities in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.CSIStorageCapacity, err error)
	// CSIStorageCapacities returns an object that can list and get CSIStorageCapacities.
	CSIStorageCapacities(namespace string) CSIStorageCapacityNamespaceLister
	CSIStorageCapacityListerExpansion
}

// cSIStorageCapacityLister implements the CSIStorageCapacityLister interface.
type cSIStorageCapacityLister struct {
	listers.ResourceIndexer[*v1beta1.CSIStorageCapacity]
}

// NewCSIStorageCapacityLister returns a new CSIStorageCapacityLister.
func NewCSIStorageCapacityLister(indexer cache.Indexer) CSIStorageCapacityLister {
	return &cSIStorageCapacityLister{listers.New[*v1beta1.CSIStorageCapacity](indexer, v1beta1.Resource("csistoragecapacity"))}
}

// CSIStorageCapacities returns an object that can list and get CSIStorageCapacities.
func (s *cSIStorageCapacityLister) CSIStorageCapacities(namespace string) CSIStorageCapacityNamespaceLister {
	return cSIStorageCapacityNamespaceLister{listers.NewNamespaced[*v1beta1.CSIStorageCapacity](s.ResourceIndexer, namespace)}
}

// CSIStorageCapacityNamespaceLister helps list and get CSIStorageCapacities.
// All objects returned here must be treated as read-only.
type CSIStorageCapacityNamespaceLister interface {
	// List lists all CSIStorageCapacities in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.CSIStorageCapacity, err error)
	// Get retrieves the CSIStorageCapacity from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.CSIStorageCapacity, error)
	CSIStorageCapacityNamespaceListerExpansion
}

// cSIStorageCapacityNamespaceLister implements the CSIStorageCapacityNamespaceLister
// interface.
type cSIStorageCapacityNamespaceLister struct {
	listers.ResourceIndexer[*v1beta1.CSIStorageCapacity]
}
