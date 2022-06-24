/*
* Copyright (c) 2021 Intel Corporation.
*
* SPDX-License-Identifier: Apache-2.0
*
 */

// Template auto-generated once, maintained by plugin owner.

package capihostprovision

import (
	pluginapi "ep/pkg/api/plugins"
	eputils "ep/pkg/eputils"
	capiutils "ep/pkg/eputils/capiutils"
	kubeutils "ep/pkg/eputils/kubeutils"
	"errors"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

var (
	errSelectProvider = errors.New("Please select one provider")
)

func PluginMain(in eputils.SchemaMapData, outp *eputils.SchemaMapData) error {
	input_ep_params := input_ep_params(in)
	input_cluster_manifest := input_cluster_manifest(in)

	var err error
	var provider string
	providers := make([]string, 0)
	for _, p := range input_ep_params.Ekconfig.Parameters.Extensions {
		for _, i := range capiutils.InfraProviderList {
			if p == i {
				providers = append(providers, p)
			}
		}
	}

	if len(providers) != 1 {
		return errSelectProvider
	} else {
		provider = providers[0]
	}

	workFolder := filepath.Join(input_ep_params.Runtimedir, provider)
	var clusterConfig pluginapi.CapiClusterConfig
	clusterConfig.WorkloadCluster = new(pluginapi.CapiClusterConfigWorkloadCluster)
	clusterConfig.BaremetelOperator = new(pluginapi.CapiClusterConfigBaremetelOperator)
	err = eputils.LoadSchemaStructFromYamlFile(&clusterConfig, input_ep_params.Ekconfig.Cluster.Config)
	if err != nil {
		log.Errorf("Load capi cluster config failed, %v", err)
		return err
	}

	var capiSetting pluginapi.CapiSetting
	capiSetting.Provider = provider
	capiSetting.InfraProvider = new(pluginapi.CapiSettingInfraProvider)
	capiSetting.ByohConfig = new(pluginapi.CapiSettingByohConfig)
	capiSetting.IronicConfig = new(pluginapi.CapiSettingIronicConfig)

	var tmpl capiutils.CapiTemplate
	capiutils.GetCapiSetting(input_ep_params, input_cluster_manifest, &clusterConfig, &capiSetting)
	err = capiutils.CheckCapiSetting(&capiSetting)
	if err != nil {
		log.Errorln("CapiHostProvision, failed to pass CapiSetting checking.")
		return err
	}

	err = capiutils.GetCapiTemplate(input_ep_params, capiSetting, &tmpl)
	if err != nil {
		log.Errorf("CapiHostProvision, get CapiTemplate failed, %v", err)
		return err
	}

	management_kubeconfig := capiutils.GetManagementClusterKubeconfig(input_ep_params)
	namespace := clusterConfig.WorkloadCluster.Namespace
	err = kubeutils.CreateNamespace(management_kubeconfig, namespace)
	if err != nil {
		log.Errorf("Create workload ns %s fail, %v", namespace, err)
		return err
	}

	if provider == capiutils.CAPI_METAL3 {
		err = metal3HostProvision(input_ep_params, workFolder, management_kubeconfig, &clusterConfig, &tmpl)
		if err != nil {
			return err
		}
	} else if provider == capiutils.CAPI_BYOH {
		err = byohHostProvision(input_ep_params, workFolder, management_kubeconfig, &clusterConfig, &tmpl)
		if err != nil {
			return err
		}
	}

	return nil
}
