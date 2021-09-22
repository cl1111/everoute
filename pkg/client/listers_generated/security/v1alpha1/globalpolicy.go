/*
Copyright 2021 The Everoute Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/everoute/everoute/pkg/apis/security/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GlobalPolicyLister helps list GlobalPolicies.
type GlobalPolicyLister interface {
	// List lists all GlobalPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GlobalPolicy, err error)
	// Get retrieves the GlobalPolicy from the index for a given name.
	Get(name string) (*v1alpha1.GlobalPolicy, error)
	GlobalPolicyListerExpansion
}

// globalPolicyLister implements the GlobalPolicyLister interface.
type globalPolicyLister struct {
	indexer cache.Indexer
}

// NewGlobalPolicyLister returns a new GlobalPolicyLister.
func NewGlobalPolicyLister(indexer cache.Indexer) GlobalPolicyLister {
	return &globalPolicyLister{indexer: indexer}
}

// List lists all GlobalPolicies in the indexer.
func (s *globalPolicyLister) List(selector labels.Selector) (ret []*v1alpha1.GlobalPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GlobalPolicy))
	})
	return ret, err
}

// Get retrieves the GlobalPolicy from the index for a given name.
func (s *globalPolicyLister) Get(name string) (*v1alpha1.GlobalPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("globalpolicy"), name)
	}
	return obj.(*v1alpha1.GlobalPolicy), nil
}
