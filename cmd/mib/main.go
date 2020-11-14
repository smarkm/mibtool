package main

import (
	"fmt"
	"github.com/smarkm/mibtool/smi"
	"os"
	"os/user"
	"path/filepath"
)

func userMibDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	defaultDir, err := filepath.Abs(filepath.Join(usr.HomeDir, ".snmp", "mibs"))
	if err != nil {
		panic(err)
	}
	return defaultDir
}

func dumpModule(mib *smi.MIB, modName string) {
	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		if sym.Module.Name == modName {
			fmt.Printf("%-40s %s\n", sym, oid)
		}
	})
}

func main() {
	if len(os.Args) != 3 || os.Args[1] != "dump" {
		fmt.Printf("Usage: %v dump [module]\n", os.Args[0])
		os.Exit(1)
	}

	mib := smi.NewMIB(userMibDir())
	err := mib.LoadModules(os.Args[2])
	if err != nil {
		fmt.Println(err)
	}
	dumpModule(mib, os.Args[2])
}
