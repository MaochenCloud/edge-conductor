/*
* Copyright (c) 2022 Intel Corporation.
*
* SPDX-License-Identifier: Apache-2.0
*
 */

// Auto generated, do not modify.

package capiclusterdeploy

import (
	pluginapi "ep/pkg/api/plugins"
	eputils "ep/pkg/eputils"
	epplugin "ep/pkg/plugin"
)

var (
	Name   = "capi-cluster-deploy"
	Input  = eputils.NewSchemaMapData()
	Output = eputils.NewSchemaMapData()
)

//nolint:unparam
func __name(n string) string {
	return Name + "." + n
}

//nolint:deadcode,unused
func input_ep_params(in eputils.SchemaMapData) *pluginapi.EpParams {
	return in[__name("ep-params")].(*pluginapi.EpParams)
}

//nolint:deadcode,unused
func input_cluster_manifest(in eputils.SchemaMapData) *pluginapi.Clustermanifest {
	return in[__name("cluster-manifest")].(*pluginapi.Clustermanifest)
}

//nolint:deadcode,unused
func output_kubeconfig(outp *eputils.SchemaMapData) *pluginapi.Filecontent {
	return (*outp)[__name("kubeconfig")].(*pluginapi.Filecontent)
}

func init() {
	eputils.AddSchemaStruct(__name("ep-params"), func() eputils.SchemaStruct { return &pluginapi.EpParams{} })
	eputils.AddSchemaStruct(__name("cluster-manifest"), func() eputils.SchemaStruct { return &pluginapi.Clustermanifest{} })
	eputils.AddSchemaStruct(__name("kubeconfig"), func() eputils.SchemaStruct { return &pluginapi.Filecontent{} })

	Input[__name("ep-params")] = &pluginapi.EpParams{}
	Input[__name("cluster-manifest")] = &pluginapi.Clustermanifest{}
	Output[__name("kubeconfig")] = &pluginapi.Filecontent{}

	epplugin.RegisterPlugin(Name, &Input, &Output, PluginMain)
}
