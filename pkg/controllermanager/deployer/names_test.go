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

import "testing"

func TestStableBaseName(t *testing.T) {
	name := stableBaseName("demo-subscription", "b6422900-42ab-4cb2-b05f-b13047af9f72", "cls-bbb-cluster-b")

	if name != "demo-subscription-ivakdtbq0s" {
		t.Fatalf("stableBaseName() = %q, want %q", name, "demo-subscription-ivakdtbq0s")
	}

	if got := stableBaseName("demo-subscription", "b6422900-42ab-4cb2-b05f-b13047af9f72", "cls-bbb-cluster-b"); got != name {
		t.Fatalf("stableBaseName() is not stable, got %q then %q", name, got)
	}

	if got := stableBaseName("demo-subscription", "b6422900-42ab-4cb2-b05f-b13047af9f72", "cls-aaa-cluster-a"); got == name {
		t.Fatalf("stableBaseName() should vary by target namespace, got %q for both namespaces", got)
	}
}
