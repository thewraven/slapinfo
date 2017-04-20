package main

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func getSlapInfo() (byte, error) {
	id, err := getSlapID()
	var statusCode byte = 0xFF
	if err != nil {
		return statusCode, err
	}
	process, err := process.NewProcess(int32(id))
	if err != nil {
		return statusCode, err
	}
	sts, err := process.Status()
	if err != nil {
		return statusCode, err
	}
	switch sts {
	case "R":
		statusCode = 0x01
	case "S":
		statusCode = 0x02
	case "I":
		statusCode = 0x03
	case "T":
		statusCode = 0x04
	case "Z":
		statusCode = 0x05
	case "W":
		statusCode = 0x06
	case "L":
		statusCode = 0x07
	}
	return statusCode, nil
}

func getSlapID() (int, error) {
	cmd := exec.Command("pgrep", "slapd")
	result, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	content := string(result)
	content = strings.Replace(content, "\n", "", -1)
	return strconv.Atoi(content)
}
