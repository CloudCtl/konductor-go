/*
Copyright 2020 Kat Morgan <usrbinkat@braincraft.io>

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
package cmd

import (
    "os"
    "fmt"
    "log"
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    name string
    subcmd string
    home        = os.Getenv("HOME")
    configFile = ("konductor.yaml")
    configyaml Configuration
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Konductor init command to prepare and validate deploy config",
	Long: `
Konductor Init:
  Init provides deployment configuration file build guidance and 
  enables stowing the generated materials in encrypted S3 storage.
`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Starting Konductor Init....")
        CoreInit()
    },
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func CoreInit() {
    viper.Set("Verbose", true)
    viper.SetConfigType("yaml")
    viper.SetConfigName("konductor.yaml")

    viper.AddConfigPath(".")
    viper.AddConfigPath("${HOME}/konductor.yaml")


    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }
    err := viper.Unmarshal(&configyaml)
    if err != nil {
        log.Fatalf("Unable to decode into struct, %v", err)
    }

    fmt.Println( "INFO:" + "\n" +
"  Openshift Version:   " + configyaml.Openshift.Version      + "\n" +
"  AWS Secret:          " + configyaml.Auth.Secret            + "\n" +
"  AWS Secret:          " + configyaml.Cloud.CidrPrivate      + "\n" )

    fmt.Println(configyaml.Subnets.Private)
}

type Configuration struct {
    Auth          AuthConfiguration `mapstructure:"provider-auth"`
    Cloud         CloudConfiguration
    Redsord       RedSordConfiguration
    Subnets       SubnetsConfiguration
    Cluster       ClusterConfiguration
    Openshift     OpenshiftConfiguration
}

type OpenshiftConfiguration struct {
    Version       string
}

type ClusterConfiguration struct {
    Target        string
    VpcName       string `mapstructure:"vpc-name"`
    ClusterName   string `mapstructure:"cluster-name"`
    BaseDomain    string `mapstructure:"base-domain"`
    ClusterDomain string `mapstructure:"cluster-domain"`
    AmiId         string `mapstructure:"ami-id"`
}

type CloudConfiguration struct {
    Provider      string
    Region        string
    VpcId         string `mapstructure:"vpc-id"`
    CidrPrivate   string `mapstructure:"cidr-private"`
}

type SubnetsConfiguration struct {
    Private       interface{} `mapstructure:"private"`
/*
    Public        map[string]
    Private       map[string]
*/
}

type AuthConfiguration struct {
    Keys          bool
    Key           string
    Secret        string
}

type RedSordConfiguration struct {
    Redsord       bool
}
