/*
* Copyright (c) 2022 Intel Corporation.
*
* SPDX-License-Identifier: Apache-2.0
*
 */

package app

import (
	cmapi "ep/pkg/api/certmgr"
	epapiplugins "ep/pkg/api/plugins"
	certmgr "ep/pkg/certmgr"
	"ep/pkg/eputils"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	cfgFile       string
	usercfg       epapiplugins.Ekconfig
	ekcfg         epapiplugins.Ekconfig
	initcerts     cmapi.Certificate
	registrycerts cmapi.Certificate
)

var (
	errCustomCfg = errors.New("Custom config not specified")
	errParameter = errors.New("Parameters not specified")
	errHome      = fmt.Errorf("Failed to get HOME dir")
)

func pre_check() error {
	if err := check_init_cmd(); err != nil {
		log.Errorln("Invalid Command line:", err)
		return err
	}

	if err := check_docker_configs(); err != nil {
		log.Warnln("failed to get docker cli configs for ", err)
	}

	return nil
}

func check_docker_configs() error {
	var cli_config_path_dir string
	if p := os.Getenv("DOCKER_CONFIG"); p != "" {
		cli_config_path_dir = p
	} else {
		if home, err := os.UserHomeDir(); err != nil {
			log.Errorf("Failed to get HOME dir for %s", err)
			return errHome
		} else {
			cli_config_path_dir = filepath.Join(home, ".docker")
		}
	}

	return eputils.CreateFolderIfNotExist(cli_config_path_dir)
}

func check_init_cmd() error {
	var files = []string{
		cfgFile,
	}
	if ekcfg.OS != nil {
		files = append(files, ekcfg.OS.Config)
		files = append(files, ekcfg.OS.Manifests...)
	}
	if ekcfg.Cluster != nil {
		files = append(files, ekcfg.Cluster.Config)
		files = append(files, ekcfg.Cluster.Manifests...)
	}
	if ekcfg.Components != nil {
		files = append(files, ekcfg.Components.Manifests...)
	}
	if ekcfg.Parameters != nil {
		files = append(files, ekcfg.Parameters.DefaultSSHKeyPath)
	}

	for _, file := range files {
		if len(strings.TrimSpace(file)) > 0 {
			if _, err := os.Stat(file); os.IsNotExist(err) {
				return err
			}
		}
	}

	if err := ekcfg.Validate(nil); err != nil {
		return err
	}

	return nil
}

func init_usercfg() {
	if usercfg.Cluster == nil {
		usercfg.Cluster = &epapiplugins.EkconfigCluster{}
	}
	if usercfg.Parameters == nil {
		usercfg.Parameters = &epapiplugins.EkconfigParameters{}
	}

	if usercfg.Parameters.GlobalSettings == nil {
		usercfg.Parameters.GlobalSettings = &epapiplugins.EkconfigParametersGlobalSettings{}
	}
	if usercfg.OS == nil {
		usercfg.OS = &epapiplugins.EkconfigOS{}
	}
	if usercfg.Components == nil {
		usercfg.Components = &epapiplugins.EkconfigComponents{}
	}
}

func init_ek_config() error {
	err := load_ek_config(&ekcfg, cfgFile)
	if err != nil {
		return err
	}
	if ekcfg.Cluster == nil {
		ekcfg.Cluster = usercfg.Cluster
		if ekcfg.Cluster.Provider == "" {
			log.Warnf("Cluster Provider not specified, use default value %s", DefaultClusterProvider)
		}
	}
	if len(ekcfg.Cluster.Manifests) <= 0 {
		ekcfg.Cluster.Manifests = append(ekcfg.Cluster.Manifests, DefaultClusterManifests)
	}
	if ekcfg.OS == nil {
		ekcfg.OS = usercfg.OS
	}
	if ekcfg.OS.Provider != "" && len(ekcfg.OS.Manifests) <= 0 {
		ekcfg.OS.Manifests = append(ekcfg.OS.Manifests, DefaultOSManifests)
	}
	if ekcfg.Components == nil {
		ekcfg.Components = usercfg.Components
		if ekcfg.Components.Selector == nil {
			log.Warnf("Components Selector not specified, use default value %s", DefaultComponentsSelector)
		}
	}
	if ekcfg.Parameters == nil {
		if usercfg.Parameters != nil {
			ekcfg.Parameters = usercfg.Parameters
			if ekcfg.Parameters.Customconfig == nil {
				log.Errorf("Custom config not specified")
				return errCustomCfg
			} else if ekcfg.Parameters.GlobalSettings == nil {
				log.Warnf("Global settings not specified, use default value")
				ekcfg.Parameters.GlobalSettings = usercfg.Parameters.GlobalSettings
			}
		} else {
			log.Errorf("Parameters not specified")
			return errParameter
		}
	}

	return nil
}

func ep_init() error {
	var err error
	log.Infoln("Init", PROJECTNAME)
	log.Infoln("==")
	log.Infoln("Top Config File:", cfgFile)

	if err = init_ek_config(); err != nil {
		log.Errorln("Failed to init top config:", err)
		return err
	}

	if err = pre_check(); err != nil {
		log.Errorln("failed pre check:", err)
		return err
	}

	// check and gen certs
	if err := certmgr.GenCertAndConfig(initcerts, ekcfg.Parameters.GlobalSettings.ProviderIP); err != nil {
		log.Error(err)
		return err
	}
	if err := certmgr.GenCertAndConfig(registrycerts, ekcfg.Parameters.GlobalSettings.ProviderIP); err != nil {
		log.Error(err)
		return err
	}
	epparams := InitEpParams(ekcfg)
	epparams_runtime_file, err := FileNameofRuntime(fnRuntimeInitParams)
	if err != nil {
		log.Errorln("Failed to get runtime file path:", err)
		return err
	}

	paramsInject := map[string]string{
		EKConfigPath: cfgFile,
	}

	if epparams, err = EpWfPreInit(epparams, paramsInject); err != nil {
		log.Errorln("Failed to init workflow:", err)
		return err
	}

	defer func() {
		err := EpWfTearDown(epparams, epparams_runtime_file)
		if err != nil {
			log.Errorln("Workflow Tear Down Error:", err)
		}
	}()

	if err = EpWfStart(epparams, "init"); err != nil {
		log.Errorln("Failed to start workflow:", err)
		return err
	}

	registry := fmt.Sprintf("%s:%s", epparams.Ekconfig.Parameters.GlobalSettings.ProviderIP, epparams.Ekconfig.Parameters.GlobalSettings.RegistryPort)
	if err = copyCaRuntimeDataDir(registry, epparams.Workspace, epparams.Runtimedata, epparams.Registrycert.Ca.Cert); err != nil {
		log.Errorln("Failed to copy CA:", err)
		return err
	}

	log.Infoln("==")
	log.Infoln("Done")
	return nil
}

// initCmd represents init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init.",
	Long:  `Init configurations and base services.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := ep_init(); err != nil {
			log.Errorln(err)
			return err
		}
		return nil
	},
}

type ReplaceMsgHook struct {
	logScrub map[string]string
}

func (hook *ReplaceMsgHook) Fire(entry *log.Entry) error {
	for k, v := range hook.logScrub {
		re := regexp.MustCompile(k)
		entry.Message = re.ReplaceAllString(entry.Message, v)
	}
	return nil
}

func (hook *ReplaceMsgHook) Levels() []log.Level {
	return log.AllLevels
}

var logreplace = map[string]string{
	`(?m)(?i)password\s*(:|=)\s*(.*)`:                          `password $1 ********`,
	`(?m)(?i)client-certificate-data\s*(:|=)\s*(.*)`:           `client-certificate-data $1 ********`,
	`(?m)(?i)client-key-data\s*(:|=)\s*(.*)`:                   `client-key-data $1 ********`,
	`(?m)(?i)(.*)HARBORADMINPASSWD\\\\\$#(.*)#'\s*(.*)`:        `${1}HARBORADMINPASSWD\\$#******#' $3`,
	`(?s)(?m)(?i)password\s*(.*)-----END RSA PRIVATE KEY-----`: `password: ********\n *******-----END RSA PRIVATE KEY-----`,
	`(?m)(?i)ssh_passwd\s*(:|=)\s*(.*)ssh_port(.*s)`:           `ssh_passwd $1 ******** ssh_port $3`,
	`(?m)(?i)ssh_passwd\s*(:|=)\s*(.*)`:                        `ssh_passwd $1 ********`,
}

func init() {
	if err := Utils_Init(); err != nil {
		log.Errorln("Failed to initialize workspace:", err)
		return
	}

	// Set log files.
	logdir := filepath.Join(GetWorkspacePath(), "log")
	if err := MakeDir(logdir); err != nil {
		log.Errorln("Failed to MakeDir", logdir)
		return
	}
	logScrubHook := &ReplaceMsgHook{
		logScrub: logreplace,
	}
	logfile := filepath.Join(logdir, "log.txt")
	logMap := lfshook.PathMap{
		log.TraceLevel: logfile,
		log.DebugLevel: logfile,
		log.InfoLevel:  logfile,
		log.WarnLevel:  logfile,
		log.ErrorLevel: logfile,
		log.FatalLevel: logfile,
		log.PanicLevel: logfile,
	}
	log.AddHook(logScrubHook)
	log.AddHook(lfshook.NewHook(
		logMap,
		&log.TextFormatter{},
	))

	initcerts = cmapi.Certificate{
		Name: "workflow",
		Ca: &cmapi.CertificateCa{
			Csr: certmgr.CACSR,
		},
		Server: &cmapi.CertificateServer{
			Csr: certmgr.WFSERVERCSR,
		},
		Client: &cmapi.CertificateClient{
			Csr: certmgr.WFCLIENTCSR,
		},
	}
	registrycerts = cmapi.Certificate{
		Name: "registry",
		Ca:   initcerts.Ca,
		Server: &cmapi.CertificateServer{
			Csr: certmgr.REGISTRYCSR,
		},
	}
	rootCmd.AddCommand(initCmd)

	init_usercfg()

	// Top Config
	initCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", GetDefaultTopConfigName(),
		PROJECTNAME+" top config file")
	if err := initCmd.MarkPersistentFlagRequired("config"); err != nil {
		log.Error(err)
		return
	}

	// Certifications
	initCmd.PersistentFlags().StringVar(&initcerts.Ca.Cert, "cacert", ROOTCACERTFILE,
		PROJECTNAME+" root ca cert file")
	initCmd.PersistentFlags().StringVar(&initcerts.Ca.Key, "cakey", certmgr.ROOTCAKEYFILE,
		PROJECTNAME+" root ca key file, for signing server and client certificates")
	initCmd.PersistentFlags().StringVar(&initcerts.Server.Cert, "servercert", certmgr.WFSERVERCERTFILE,
		PROJECTNAME+" workflow server certificate file")
	initCmd.PersistentFlags().StringVar(&initcerts.Server.Key, "serverkey", certmgr.WFSERVERKEYFILE,
		PROJECTNAME+" workflow server certificate key file")
	initCmd.PersistentFlags().StringVar(&initcerts.Client.Cert, "clientcert", certmgr.WFCLIENTCERTFILE,
		PROJECTNAME+" workflow client certificate file")
	initCmd.PersistentFlags().StringVar(&initcerts.Client.Key, "clientkey", certmgr.WFCLIENTKEYFILE,
		PROJECTNAME+" workflow client certificate key file")
	initCmd.PersistentFlags().StringVar(&registrycerts.Server.Cert, "registrycert", certmgr.REGSERVERCERTFILE,
		PROJECTNAME+" registry certificate file")
	initCmd.PersistentFlags().StringVar(&registrycerts.Server.Key, "registrykey", certmgr.REGSERVERKEYFILE,
		PROJECTNAME+" registry certificate key file")
}
