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

    log.Printf("AWS Region:     %s", configyaml.Cloud.Region)
    log.Printf("AWS VPC ID:     %s", configyaml.Cloud.VpcId)
    log.Printf("AWS Key ID::    %s", configyaml.Cloud.Secret)
    log.Printf("AWS Key Secret: %s \n", configyaml.Cloud.Key)
    log.Printf("Task Command:   %s", configyaml.Task.Cmd)
    log.Printf("Task Sub Cmd:   %s", configyaml.Task.Sub)
}

func init() {
	rootCmd.AddCommand(initCmd)
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Configuration struct {
	Cloud CloudConfiguration
	Task TaskConfiguration
}

type TaskConfiguration struct {
	Cmd string
	Sub string
}

type CloudConfiguration struct {
	Key string
	Secret string
	Region string
	VpcId string
}
