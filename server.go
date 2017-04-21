package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/k-sone/snmpgo"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}

var addr *string

func init() {
	addr = flag.String("addr", "127.0.0.1:162", "address to bind")
	flag.Parse()
}

func sendTrap() {
	snmp, err := snmpgo.NewSNMP(snmpgo.SNMPArguments{
		Version:   snmpgo.V2c,
		Address:   *addr,
		Retries:   3,
		Community: "public",
	})
	checkErr(err)
	var binds snmpgo.VarBinds
	info, err := getSlapInfo()
	checkErr(err)
	binds = append(binds,
		snmpgo.NewVarBind(snmpgo.OidSnmpTrap,
			snmpgo.NewOctetString([]byte{info})))
	err = snmp.Open()
	checkErr(err)
	defer snmp.Close()
	err = snmp.V2Trap(binds)
	checkErr(err)
}

func main() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case _ = <-ticker.C:
			sendTrap()
		}
	}
}
