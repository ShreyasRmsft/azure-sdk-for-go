//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

const (
	module  = "armportal"
	version = "v0.1.0"
)

type ConfigurationName string

const (
	ConfigurationNameDefault ConfigurationName = "default"
)

// PossibleConfigurationNameValues returns the possible values for the ConfigurationName const type.
func PossibleConfigurationNameValues() []ConfigurationName {
	return []ConfigurationName{
		ConfigurationNameDefault,
	}
}

// ToPtr returns a *ConfigurationName pointing to the current value.
func (c ConfigurationName) ToPtr() *ConfigurationName {
	return &c
}