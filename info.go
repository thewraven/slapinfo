package main

import (
	"errors"
	"fmt"
	"math/big"
	"os/exec"
	"strconv"

	"github.com/k-sone/snmpgo"
	"github.com/shirou/gopsutil/process"
)

type v struct {
	info string
	err  error
}

func (v *v) BigInt() (*big.Int, error) {
	return nil, errors.New("the value hasn't an bigInt repr")
}
func (v *v) String() string {
	if v.err != nil {
		return v.err.Error()
	}
	return v.info
}

func (v *v) Type() string {
	return "slapd process status"
}

func (v *v) Marshal() ([]byte, error) {
	var marshal []byte
	if v.err != nil {
		marshal = []byte(v.err.Error())
	} else {
		marshal = []byte(v.info)
	}
	return marshal, nil
}

func (v *v) Unmarshal(info []byte) (rest []byte, err error) {
	v.info = string(info)
	return nil, nil
}

func wrapSlapInfo() snmpgo.Variable {
	info, err := getSlapInfo()
	fmt.Println("unwrapping variable", info, err)
	return &v{
		info: info,
		err:  err,
	}
}

func getSlapInfo() (string, error) {
	id, err := getSlapID()
	if err != nil {
		return "", err
	}
	process, err := process.NewProcess(int32(id))
	if err != nil {
		return "", err
	}
	return process.Status()
}

func getSlapID() (int, error) {
	cmd := exec.Command("pgrep", "slapd")
	result, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(result))
}
