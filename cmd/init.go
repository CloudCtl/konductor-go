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
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

var (
    name string
    home        = os.Getenv("HOME")
    configFile = ("config.yaml")
//  dirPlatform = (home + "/" + "deploy")
//  configFile  = (dirPlatform + "/" + "config.yaml")
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
    viper.SetConfigName("konductor.yaml")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("${HOME}/konductor.yaml")
    viper.AddConfigPath(".")
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }
    task := viper.Get("task")
    if task == "" { // Handle errors reading the value 
        fmt.Println(fmt.Errorf("Fatal error! 'task' has no value. \n"))
    }
    fmt.Printf("Sparta is running: %s \n", task)
}

func init() {
	rootCmd.AddCommand(initCmd)
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

