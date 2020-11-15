/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"os/user"
	"path/filepath"

	"github.com/smarkm/mibtool/smi"
	"github.com/spf13/cobra"
)

// args
var (
	MibPath   string
	Mib       string
	Syntax    bool
	MaxAccess bool
	Status    bool
)

// glanceCmd represents the glance command
var glanceCmd = &cobra.Command{
	Use:   "glance",
	Short: "Display oid information",
	Long:  "Display oid information",
	Run: func(cmd *cobra.Command, args []string) {
		mib := smi.NewMIB(MibPath)
		err := mib.LoadModules(Mib)
		if err != nil {
			fmt.Println(err)
		}

		if Mib == "" {
			for _, m := range mib.Modules {
				fmt.Printf("%s\n", m.Name)
			}
			return
		}
		dumpModule(mib, Mib)
	},
}

func init() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	defaultDir, err := filepath.Abs(filepath.Join(usr.HomeDir, ".snmp", "mibs"))
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(glanceCmd)
	glanceCmd.Flags().StringVarP(&MibPath, "path", "p", defaultDir, "Mib  path")
	glanceCmd.Flags().StringVarP(&Mib, "mib", "m", "", "Only dispaly oid in this mib, if empty will only list mib under path ")
	glanceCmd.Flags().BoolVarP(&Syntax, "syntax", "t", false, "Show syntax of object")
	glanceCmd.Flags().BoolVarP(&MaxAccess, "access", "a", false, "Show max-access")
	glanceCmd.Flags().BoolVarP(&Status, "status", "s", false, "Show status of object")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// glanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// glanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func dumpModule(mib *smi.MIB, modName string) {
	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		if sym.Module.Name == modName {
			fmt.Printf("%-40s %s\n", sym, oid)
		}
	})
}
