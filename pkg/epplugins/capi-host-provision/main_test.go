/*
* Copyright (c) 2021 Intel Corporation.
*
* SPDX-License-Identifier: Apache-2.0
*
 */

// Template auto-generated once, maintained by plugin owner.

//nolint: dupl
package capihostprovision

import (
	"ep/pkg/api/plugins"
	eputils "ep/pkg/eputils"
	capiutils "ep/pkg/eputils/capiutils"
	kubeutils "ep/pkg/eputils/kubeutils"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/undefinedlabs/go-mpatch"
	// TODO: Add Plugin Unit Test Imports Here
)

var (
	errProvider = errors.New("Please select one provider")
	errTest_    = errors.New("test_error")
)

func TestPluginMain(t *testing.T) {
	func_LoadSchemaStructFromYamlFile_err := func(ctrl *gomock.Controller) []*mpatch.Patch {
		pathchLoadSchemaStructFromYamlFile, err := mpatch.PatchMethod(eputils.LoadSchemaStructFromYamlFile, func(v eputils.SchemaStruct, file string) error {
			return errTest_
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}

		return []*mpatch.Patch{pathchLoadSchemaStructFromYamlFile}
	}
	func_GetCapiTemplate_err := func(ctrl *gomock.Controller) []*mpatch.Patch {
		pathchLoadSchemaStructFromYamlFile, err := mpatch.PatchMethod(eputils.LoadSchemaStructFromYamlFile, func(v eputils.SchemaStruct, file string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiSetting, err := mpatch.PatchMethod(capiutils.GetCapiSetting, func(epparams *plugins.EpParams, clusterManifest *plugins.Clustermanifest, clusterConfig *plugins.CapiClusterConfig, setting *plugins.CapiSetting) {
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiTemplate, err := mpatch.PatchMethod(capiutils.GetCapiTemplate, func(epparams *plugins.EpParams, setting plugins.CapiSetting, cp *capiutils.CapiTemplate) error {
			return errTest_
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		return []*mpatch.Patch{pathchLoadSchemaStructFromYamlFile, pathchGetCapiSetting, pathchGetCapiTemplate}
	}
	func_CreateNamespace_err := func(ctrl *gomock.Controller) []*mpatch.Patch {
		pathchLoadSchemaStructFromYamlFile, err := mpatch.PatchMethod(eputils.LoadSchemaStructFromYamlFile, func(v eputils.SchemaStruct, file string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiSetting, err := mpatch.PatchMethod(capiutils.GetCapiSetting, func(epparams *plugins.EpParams, clusterManifest *plugins.Clustermanifest, clusterConfig *plugins.CapiClusterConfig, setting *plugins.CapiSetting) {
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiTemplate, err := mpatch.PatchMethod(capiutils.GetCapiTemplate, func(epparams *plugins.EpParams, setting plugins.CapiSetting, cp *capiutils.CapiTemplate) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchCreateNamespace, err := mpatch.PatchMethod(kubeutils.CreateNamespace, func(kubeconfig string, namespace string) error {
			return errTest_
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		return []*mpatch.Patch{pathchLoadSchemaStructFromYamlFile, pathchGetCapiSetting, pathchGetCapiTemplate, pathchCreateNamespace}
	}
	func_metal3HostProvision_err := func(ctrl *gomock.Controller) []*mpatch.Patch {
		pathchLoadSchemaStructFromYamlFile, err := mpatch.PatchMethod(eputils.LoadSchemaStructFromYamlFile, func(v eputils.SchemaStruct, file string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiSetting, err := mpatch.PatchMethod(capiutils.GetCapiSetting, func(epparams *plugins.EpParams, clusterManifest *plugins.Clustermanifest, clusterConfig *plugins.CapiClusterConfig, setting *plugins.CapiSetting) {
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiTemplate, err := mpatch.PatchMethod(capiutils.GetCapiTemplate, func(epparams *plugins.EpParams, setting plugins.CapiSetting, cp *capiutils.CapiTemplate) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchCreateNamespace, err := mpatch.PatchMethod(kubeutils.CreateNamespace, func(kubeconfig string, namespace string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchmetal3HostProvision, err := mpatch.PatchMethod(metal3HostProvision, func(ep_params *plugins.EpParams, workFolder string, management_kubeconfig string, clusterConfig *plugins.CapiClusterConfig, tmpl *capiutils.CapiTemplate) error {
			return errTest_
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		return []*mpatch.Patch{pathchLoadSchemaStructFromYamlFile, pathchGetCapiSetting, pathchGetCapiTemplate, pathchCreateNamespace, pathchmetal3HostProvision}
	}
	func_byohHostProvision_err := func(ctrl *gomock.Controller) []*mpatch.Patch {
		pathchLoadSchemaStructFromYamlFile, err := mpatch.PatchMethod(eputils.LoadSchemaStructFromYamlFile, func(v eputils.SchemaStruct, file string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiSetting, err := mpatch.PatchMethod(capiutils.GetCapiSetting, func(epparams *plugins.EpParams, clusterManifest *plugins.Clustermanifest, clusterConfig *plugins.CapiClusterConfig, setting *plugins.CapiSetting) {
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiTemplate, err := mpatch.PatchMethod(capiutils.GetCapiTemplate, func(epparams *plugins.EpParams, setting plugins.CapiSetting, cp *capiutils.CapiTemplate) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchCreateNamespace, err := mpatch.PatchMethod(kubeutils.CreateNamespace, func(kubeconfig string, namespace string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchbyohHostProvision, err := mpatch.PatchMethod(byohHostProvision, func(ep_params *plugins.EpParams, workFolder string, management_kubeconfig string, clusterConfig *plugins.CapiClusterConfig, tmpl *capiutils.CapiTemplate) error {
			return errTest_
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		return []*mpatch.Patch{pathchLoadSchemaStructFromYamlFile, pathchGetCapiSetting, pathchGetCapiTemplate, pathchCreateNamespace, pathchbyohHostProvision}
	}
	func_byohHostProvision_ok := func(ctrl *gomock.Controller) []*mpatch.Patch {
		pathchLoadSchemaStructFromYamlFile, err := mpatch.PatchMethod(eputils.LoadSchemaStructFromYamlFile, func(v eputils.SchemaStruct, file string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiSetting, err := mpatch.PatchMethod(capiutils.GetCapiSetting, func(epparams *plugins.EpParams, clusterManifest *plugins.Clustermanifest, clusterConfig *plugins.CapiClusterConfig, setting *plugins.CapiSetting) {
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchGetCapiTemplate, err := mpatch.PatchMethod(capiutils.GetCapiTemplate, func(epparams *plugins.EpParams, setting plugins.CapiSetting, cp *capiutils.CapiTemplate) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchCreateNamespace, err := mpatch.PatchMethod(kubeutils.CreateNamespace, func(kubeconfig string, namespace string) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		pathchbyohHostProvision, err := mpatch.PatchMethod(byohHostProvision, func(ep_params *plugins.EpParams, workFolder string, management_kubeconfig string, clusterConfig *plugins.CapiClusterConfig, tmpl *capiutils.CapiTemplate) error {
			return nil
		})
		if err != nil {
			t.Errorf("patch error: %v", err)
		}
		return []*mpatch.Patch{pathchLoadSchemaStructFromYamlFile, pathchGetCapiSetting, pathchGetCapiTemplate, pathchCreateNamespace, pathchbyohHostProvision}
	}
	cases := []struct {
		name                  string
		input, expectedOutput map[string][]byte
		expectError           bool
		expectErrorContent    error
		funcBeforeTest        func(*gomock.Controller) []*mpatch.Patch
	}{
		{
			name: "providers_err",
			input: map[string][]byte{
				"ep-params": []byte(`{
					"ekconfig": {
						"Cluster": {
							"provider": "clusterapi"
						}, 
						"Parameters": {
							"Extensions": ["test_err"]
						}
					}
				}`),
				"cluster-manifest": []byte(`{"clusterapi":{"configs":[{"name": "metal3", "bin_url": "aaa", "images": ["test:test"]}]}}`),
			},
			expectError:        true,
			expectErrorContent: errProvider,
			//funcBeforeTest:     func_all_cases_tanzu_inject_ca,
		},
		{
			name: "LoadSchemaStructFromYamlFile_err",
			input: map[string][]byte{
				"ep-params": []byte(`{
					"ekconfig": {
						"Cluster": {
							"provider": "clusterapi"
						}, 
						"Parameters": {
							"Extensions": ["capi-metal3"]
						}
					}
				}`),
				"cluster-manifest": []byte(`{"clusterapi":{"configs":[{"name": "metal3", "bin_url": "aaa", "images": ["test:test"]}]}}`),
			},
			expectError:        true,
			expectErrorContent: errTest_,
			funcBeforeTest:     func_LoadSchemaStructFromYamlFile_err,
		},
		{
			name: "GetCapiTemplate_err",
			input: map[string][]byte{
				"ep-params": []byte(`{
					"ekconfig": {
						"Cluster": {
							"provider": "clusterapi"
						}, 
						"Parameters": {
							"Extensions": ["capi-metal3"]
						}
					}
				}`),
				"cluster-manifest": []byte(`{"clusterapi":{"configs":[{"name": "metal3", "bin_url": "aaa", "images": ["test:test"]}]}}`),
			},
			expectError:        true,
			expectErrorContent: errTest_,
			funcBeforeTest:     func_GetCapiTemplate_err,
		},
		{
			name: "CreateNamespace_err",
			input: map[string][]byte{
				"ep-params": []byte(`{
					"ekconfig": {
						"Cluster": {
							"provider": "clusterapi"
						}, 
						"Parameters": {
							"Extensions": ["capi-metal3"]
						}
					}
				}`),
				"cluster-manifest": []byte(`{"clusterapi":{"configs":[{"name": "metal3", "bin_url": "aaa", "images": ["test:test"]}]}}`),
			},
			expectError:        true,
			expectErrorContent: errTest_,
			funcBeforeTest:     func_CreateNamespace_err,
		},
		{
			name: "metal3HostProvision_err",
			input: map[string][]byte{
				"ep-params": []byte(`{
					"ekconfig": {
						"Cluster": {
							"provider": "clusterapi"
						}, 
						"Parameters": {
							"Extensions": ["capi-metal3"]
						}
					}
				}`),
				"cluster-manifest": []byte(`{"clusterapi":{"configs":[{"name": "metal3", "bin_url": "aaa", "images": ["test:test"]}]}}`),
			},
			expectError:        true,
			expectErrorContent: errTest_,
			funcBeforeTest:     func_metal3HostProvision_err,
		},
		{
			name: "byohHostProvision_err",
			input: map[string][]byte{
				"ep-params": []byte(`{
					"ekconfig": {
						"Cluster": {
							"provider": "clusterapi"
						}, 
						"Parameters": {
							"Extensions": ["capi-byoh"],
							"Nodes": [{
                                                                "Role": ["worker"]
                                                        },
                                                        {
                                                                "Role": ["controlplane"]
                                                        }]
						}
					}
				}`),
				"cluster-manifest": []byte(`{"clusterapi":{"configs":[{"name": "metal3", "bin_url": "aaa", "images": ["test:test"]}]}}`),
			},
			expectError:        true,
			expectErrorContent: errTest_,
			funcBeforeTest:     func_byohHostProvision_err,
		},
		{
			name: "byohHostProvision_ok",
			input: map[string][]byte{
				"ep-params": []byte(`{
					"ekconfig": {
						"Cluster": {
							"provider": "clusterapi"
						}, 
						"Parameters": {
							"Extensions": ["capi-byoh"],
							"Nodes": [{
                                                                "Role": ["worker"]
                                                        },
                                                        {
                                                                "Role": ["controlplane"]
                                                        }]
						}
					}
				}`),
				"cluster-manifest": []byte(`{"clusterapi":{"configs":[{"name": "metal3", "bin_url": "aaa", "images": ["test:test"]}]}}`),
			},
			expectError:    false,
			funcBeforeTest: func_byohHostProvision_ok,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// Run test cases in parallel if necessary.
			// t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			t.Logf("Run Test Case %s", tc.name)

			input := generateInput(tc.input)

			if tc.funcBeforeTest != nil {
				plist := tc.funcBeforeTest(ctrl)
				for _, p := range plist {
					defer unpatch(t, p)
				}

			}

			if input == nil {
				t.Fatalf("Failed to generateInput %s", tc.input)
			}
			testOutput := generateOutput(nil)

			if tc.name == "err_ekconfig" {
				input_ep_params := input_ep_params(input)
				input_ep_params.Ekconfig.Parameters.GlobalSettings = nil
			}

			err := PluginMain(input, &testOutput)

			if (tc.expectError == false && err != nil) || (tc.expectError == true && err == nil) {
				t.Error(err)
			} else {
				if tc.expectError == false && err == nil {
					t.Logf("Test Case %s Pass", tc.name)
				}
				if tc.expectError == true && err != nil {
					if fmt.Sprint(err) == fmt.Sprint(tc.expectErrorContent) {
						t.Logf("Test Case %s Pass", tc.name)
					} else {
						t.Error(err)
					}
				}
			}
			_ = testOutput
		})
	}
}
