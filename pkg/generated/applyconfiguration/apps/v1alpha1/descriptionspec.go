/*
Copyright The Clusternet Authors.

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

package v1alpha1

import (
	appsv1alpha1 "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
)

// DescriptionSpecApplyConfiguration represents a declarative configuration of the DescriptionSpec type for use
// with apply.
type DescriptionSpecApplyConfiguration struct {
	Deployer *appsv1alpha1.DescriptionDeployer  `json:"deployer,omitempty"`
	Charts   []ChartReferenceApplyConfiguration `json:"charts,omitempty"`
	Raw      [][]byte                           `json:"raw,omitempty"`
	ChartRaw [][]byte                           `json:"chartRaw,omitempty"`
}

// DescriptionSpecApplyConfiguration constructs a declarative configuration of the DescriptionSpec type for use with
// apply.
func DescriptionSpec() *DescriptionSpecApplyConfiguration {
	return &DescriptionSpecApplyConfiguration{}
}

// WithDeployer sets the Deployer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Deployer field is set to the value of the last call.
func (b *DescriptionSpecApplyConfiguration) WithDeployer(value appsv1alpha1.DescriptionDeployer) *DescriptionSpecApplyConfiguration {
	b.Deployer = &value
	return b
}

// WithCharts adds the given value to the Charts field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Charts field.
func (b *DescriptionSpecApplyConfiguration) WithCharts(values ...*ChartReferenceApplyConfiguration) *DescriptionSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCharts")
		}
		b.Charts = append(b.Charts, *values[i])
	}
	return b
}

// WithRaw adds the given value to the Raw field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Raw field.
func (b *DescriptionSpecApplyConfiguration) WithRaw(values ...[]byte) *DescriptionSpecApplyConfiguration {
	for i := range values {
		b.Raw = append(b.Raw, values[i])
	}
	return b
}

// WithChartRaw adds the given value to the ChartRaw field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ChartRaw field.
func (b *DescriptionSpecApplyConfiguration) WithChartRaw(values ...[]byte) *DescriptionSpecApplyConfiguration {
	for i := range values {
		b.ChartRaw = append(b.ChartRaw, values[i])
	}
	return b
}
