/*
* Copyright (c) 2022 Intel Corporation.
*
* SPDX-License-Identifier: Apache-2.0
*
 */
// Template auto-generated once, maintained by plugin owner.

package rkeparser

import (
	"bytes"
	papi "ep/pkg/api/plugins"
	eputils "ep/pkg/eputils"
	cutils "ep/pkg/eputils/conductorutils"
	"errors"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	errLoadJson  = errors.New("Failed to load json file!")
	errConfViper = errors.New("Could not config the viper!")
)

func PluginMain(in eputils.SchemaMapData, outp *eputils.SchemaMapData) error {
	input_ep_params := input_ep_params(in)
	input_eptopcfg := input_ep_params.Ekconfig
	input_cluster_manifest := input_cluster_manifest(in)

	output_docker_images := output_docker_images(outp)
	output_files := output_files(outp)

	provider, err := cutils.GetClusterManifest(input_cluster_manifest, "rke")
	if err != nil {
		log.Errorln("Failed to find manifest for RKE cluster.")
		return err
	}

	rkeBin, rkeSHA256, err := cutils.GetBinaryFromProvider(provider, "rketool")
	if err != nil {
		log.Errorln("Failed to find binary for RKE.")
		return err
	}
	rkeDir := eputils.GetBaseUrl(rkeBin)

	rkeCfgFile := input_eptopcfg.Cluster.Config
	log.Infof("Read cluster config file: %s", rkeCfgFile)
	rkeCfgFilebyteValue, err := eputils.LoadJsonFile(rkeCfgFile)
	if err != nil {
		log.Errorln("Failed to load json file!")
		return errLoadJson
	}

	rkeCfgViper := viper.New()
	rkeCfgViper.SetConfigType("yaml")
	if err := rkeCfgViper.ReadConfig(bytes.NewBuffer(rkeCfgFilebyteValue)); err != nil {
		log.Errorf("Could not config the viper!")
		return errConfViper
	}

	imageMap := rkeCfgViper.GetStringMapString("system_images")

	log.Infof("%s", imageMap)
	output_docker_images.Images = []*papi.ImagesItems0{}

	for key := range imageMap {
		output_docker_images.Images = append(output_docker_images.Images,
			&papi.ImagesItems0{Name: key, URL: imageMap[key]})
	}

	output_files.Files = []*papi.FilesItems0{
		{
			URL:      rkeBin,
			Hash:     rkeSHA256,
			Hashtype: "sha256",
			Urlreplacement: &papi.FilesItems0Urlreplacement{
				New:    "binary",
				Origin: rkeDir,
			},
		},
	}

	log.Debugf("%v", output_docker_images)
	log.Debugf("%v", output_files)

	return nil
}
