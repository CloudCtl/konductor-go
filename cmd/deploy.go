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
    "log"
    "fmt"
    "sync"
    "os/exec"
//  "flag"
//  "strings"
//  "path/filepath"

    "github.com/spf13/cobra"
    kcorelog "github.com/CodeSparta/konductor-go/plugins/log"
//  kpullsecret "github.com/CodeSparta/konductor-go/plugins/auth"
//  "github.com/CodeSparta/konductor-go/plugins/err"
)

// deployCmd represents the deploy command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "A brief description of your command",
	Long: `
Konductor Engine Init:
  Init is designed to orchestrate cloud deployment & operation
  "maestro" automation plugins to deliver seamlessly in
  restricted & airgap environment capabilities.
`,
	Run: func(cmd *cobra.Command, args []string) {
		core()
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
}

func core() {
    CmdPluginRun()
}

// Run Konductor Plugin from site.yml
func CmdPluginRun() {

    // Run Plugin
    fmtPrintf("  >> Running Plugin ./site.yml")
    cmd := exec.Command("./site.yml")
    cmd.Dir = "/root/dev"
    //cmd.Dir = "/root/deploy/ansible/deploy"
    var stdout, stderr []byte
    var errStdout, errStderr error
    stdoutIn, _ := cmd.StdoutPipe()
    stderrIn, _ := cmd.StderrPipe()
    err := cmd.Start()
    if err != nil {
        log.Fatalf("cmd.Start() failed with '%s'\n", err)
    }
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        stdout, errStdout = kcorelog.CopyAndCapture(os.Stdout, stdoutIn)
        wg.Done()
    }()
    stderr, errStderr = kcorelog.CopyAndCapture(os.Stderr, stderrIn)
    wg.Wait()
    err = cmd.Wait()
    if err != nil {
        log.Fatalf("cmd.Run() failed with %s\n", err)
    }
    if errStdout != nil || errStderr != nil {
        log.Fatal("failed to capture stdout \n")
    }
    errStr := string(stderr)
    if stderr != nil {
        fmt.Printf("\nerr:\n%s\n", errStr)
    }
}

// Info Log to screen running plugin
func konductorLoop(repo string) {
    fmt.Println(" >>  Running Plugin: ", repo)
}
