package main

import (
	"flag"
	"fmt"

	"github.com/k-sone/snmpgo"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	addr := flag.String("addr", "127.0.0.1:162", "address to bind")
	flag.Parse()
	snmp, err := snmpgo.NewSNMP(
		snmpgo.SNMPArguments{
			Version:   snmpgo.V2c,
			Address:   *addr,
			Retries:   3,
			Community: "public",
		})
	checkErr(err)
	fmt.Println(snmp)
	//	oid, err := snmpgo.NewOid("1.3.6.1.6.3.1.1.5.3")
	//	checkErr(err)
	//	fmt.Println(oid.String())
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
	fmt.Println("OK")
}
