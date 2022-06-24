/*
* Copyright (c) 2022 Intel Corporation.
*
* SPDX-License-Identifier: Apache-2.0
*
 */
package app

import (
	"ep/pkg/api/plugins"
	epapiplugins "ep/pkg/api/plugins"
	"testing"

	"github.com/spf13/cobra"
	"github.com/undefinedlabs/go-mpatch"
)

func TestOsDeployCmd(t *testing.T) {
	cases := []struct {
		funcBeforeTest func() []*mpatch.Patch
		wantError      error
		funcAfterTest  func()
	}{
		{
			funcBeforeTest: func() []*mpatch.Patch {
				ekcfg = epapiplugins.Ekconfig{}
				patch := patchEpWfPreInit(t, nil, testError)
				ekcfg.Cluster = &plugins.EkconfigCluster{}
				ekcfg.Components = &plugins.EkconfigComponents{}
				return []*mpatch.Patch{patch}
			},
			wantError: testError,
		},
		{
			funcBeforeTest: func() []*mpatch.Patch {
				ekcfg = epapiplugins.Ekconfig{}
				patchEpWfPreInit := patchEpWfPreInit(t, nil, nil)
				patchEpWfStart := patchEpWfStart(t, testError)
				ekcfg.Cluster = &plugins.EkconfigCluster{}
				ekcfg.Components = &plugins.EkconfigComponents{}
				return []*mpatch.Patch{patchEpWfPreInit, patchEpWfStart}
			},
			wantError: testError,
		},
		{
			funcBeforeTest: func() []*mpatch.Patch {
				ekcfg = epapiplugins.Ekconfig{}
				patchEpWfPreInit := patchEpWfPreInit(t, nil, nil)
				patchEpWfStart := patchEpWfStart(t, nil)
				ekcfg.Cluster = &plugins.EkconfigCluster{}
				ekcfg.Components = &plugins.EkconfigComponents{}
				return []*mpatch.Patch{patchEpWfPreInit, patchEpWfStart}
			},
			wantError: nil,
		},
	}
	testCmdList := []*cobra.Command{
		osDeployBuildCmd,
		osDeployStartCmd,
		osDeployStopCmd,
		osDeployCleanupCmd,
	}

	for _, testCmd := range testCmdList {
		for n, testCase := range cases {
			t.Logf("test deploy %s command case %d start", testCmd.Use, n)
			func() {
				if testCase.funcBeforeTest != nil {
					pList := testCase.funcBeforeTest()
					defer unpatchAll(t, pList)
				}

				err := testCmd.RunE(&cobra.Command{}, []string{})
				if !isWantedError(err, testCase.wantError) {
					t.Errorf("Unexpected error: %v", err)
				}
			}()
			t.Logf("test deploy %s command case %d End", testCmd.Use, n)
		}
	}

	t.Log("Done")
}
