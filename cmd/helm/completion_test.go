/*
Copyright The Helm Authors.

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

package main

import (
	"fmt"
	"strings"
	"testing"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"
)

// Check if file completion should be performed according to parameter 'shouldBePerformed'
func checkFileCompletion(t *testing.T, cmdName string, shouldBePerformed bool) {
	storage := storageFixture()
	storage.Create(&release.Release{
		Name:    "myrelease",
		Info:    &release.Info{Status: release.StatusDeployed},
		Chart:   &chart.Chart{},
		Version: 1,
	})

	testcmd := fmt.Sprintf("__complete %s ''", cmdName)
	_, out, err := executeActionCommandC(storage, testcmd)
	if err != nil {
		t.Errorf("unexpected error, %s", err)
	}
	if !strings.Contains(out, "ShellCompDirectiveNoFileComp") != shouldBePerformed {
		if shouldBePerformed {
			t.Error(fmt.Sprintf("Unexpected directive ShellCompDirectiveNoFileComp when completing '%s'", cmdName))
		} else {

			t.Error(fmt.Sprintf("Did not receive directive ShellCompDirectiveNoFileComp when completing '%s'", cmdName))
		}
		t.Log(out)
	}
}

func TestCompletionFileCompletion(t *testing.T) {
	checkFileCompletion(t, "completion", false)
	checkFileCompletion(t, "completion bash", false)
	checkFileCompletion(t, "completion zsh", false)
	checkFileCompletion(t, "completion fish", false)
}
