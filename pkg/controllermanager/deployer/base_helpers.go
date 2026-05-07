/*
Copyright 2026 The Clusternet Authors.

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

package deployer

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsapi "github.com/clusternet/clusternet/pkg/apis/apps/v1alpha1"
	appsclient "github.com/clusternet/clusternet/pkg/generated/clientset/versioned/typed/apps/v1alpha1"
	applisters "github.com/clusternet/clusternet/pkg/generated/listers/apps/v1alpha1"
)

func getBaseForUpdate(ctx context.Context, appsClient appsclient.AppsV1alpha1Interface, baseLister applisters.BaseLister,
	namespace, name string) (*appsapi.Base, error) {
	base, err := baseLister.Bases(namespace).Get(name)
	if err == nil || !apierrors.IsNotFound(err) {
		return base, err
	}

	return appsClient.Bases(namespace).Get(ctx, name, metav1.GetOptions{})
}
