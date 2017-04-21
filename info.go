package main

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/thewraven/slapinfo/stats"

	"github.com/shirou/gopsutil/process"
)

func getSlapInfo() (byte, error) {
	id, err := getSlapID()
	var statusCode byte = stats.Unavailable
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
		statusCode = stats.Running
	case "S":
		statusCode = stats.Sleep
	case "I":
		statusCode = stats.Idle
	case "T":
		statusCode = stats.Stopped
	case "Z":
		statusCode = stats.Zombie
	case "W":
		statusCode = stats.Wait
	case "L":
		statusCode = stats.Lock
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
