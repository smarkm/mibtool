# SNMP MIB Parser for Go

[![Build Status](https://travis-ci.com/smarkm/mibtool.svg?branch=master)](https://travis-ci.com/smarkm/mibtool)
[![GoDoc](https://godoc.org/github.com/smarkm/mibtool/smi?status.svg)](https://godoc.org/github.com/smarkm/mibtool/smi)
![Version](https://img.shields.io/github/tag/smarkm/mibtool.svg?label=version)

The `mibtool` module contains packages for parsing SNMP MIBs and querying
the information contained in them.

The information that can currently be extracted from MIBs is limited to
symbol information and OIDs, but the intention is to extend the code
to make more information available.

## Installation

    go get -u github.com/smarkm/mibtool/smi

## Examples

	mib := smi.NewMIB("/usr/share/snmp/mibs/iana", "/usr/share/snmp/mibs/ietf")
	mib.Debug = true
	err := mib.LoadModules("IF-MIB")
	if err != nil {
		log.Fatal(err)
	}

	// Walk all symbols in MIB
	mib.VisitSymbols(func(sym *smi.Symbol, oid smi.OID) {
		fmt.Printf("%-40s %s\n", sym, oid)
	})

    // Look up OID for an OID string
    oid, err := mib.OID("IF-MIB::ifOperStatus.4")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(oid.String())
## More
this project is a hard fork from https://github.com/hallidave/mibtool, since the `issues` have long time not addressed, thanks hallidave's work