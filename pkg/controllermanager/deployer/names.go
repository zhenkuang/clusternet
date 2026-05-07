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
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"strings"

	utilvalidation "k8s.io/apimachinery/pkg/util/validation"
)

const stableBaseNameHashLength = 10

func stableBaseName(subscriptionName, subscriptionUID, targetNamespace string) string {
	hash := sha256.New()
	hash.Write([]byte(subscriptionUID))
	hash.Write([]byte("/"))
	hash.Write([]byte(targetNamespace))
	suffix := strings.ToLower(base32.HexEncoding.WithPadding(base32.NoPadding).EncodeToString(hash.Sum(nil)))[:stableBaseNameHashLength]

	prefix := subscriptionName
	maxPrefixLength := utilvalidation.DNS1123SubdomainMaxLength - stableBaseNameHashLength - 1
	if len(prefix) > maxPrefixLength {
		prefix = strings.TrimRight(prefix[:maxPrefixLength], "-.")
	}
	if prefix == "" {
		prefix = "base"
	}

	return fmt.Sprintf("%s-%s", prefix, suffix)
}
