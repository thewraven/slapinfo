package stats

import (
	"encoding/json"
	"strings"
)

const (
	Running     = 0x01
	Sleep       = 0x02
	Idle        = 0x03
	Stopped     = 0x04
	Zombie      = 0x05
	Wait        = 0x06
	Lock        = 0x07
	Unavailable = 0xFF
)

type TrapReq struct {
	Type        string `json:"Type"`
	RequestID   string `json:"RequestId"`
	ErrorStatus string `json:"ErrorStatus"`
	ErrorIndex  string `json:"ErrorIndex"`
	VarBinds    []struct {
		Oid      string `json:"Oid"`
		Variable struct {
			Type  string `json:"Type"`
			Value string `json:"Value"`
		} `json:"Variable"`
	} `json:"VarBinds"`
}

func TranslateMessage(msg string) string {
	decoder := json.NewDecoder(strings.NewReader(msg))
	var t TrapReq
	decoder.Decode(&t)
	firstrap := t.VarBinds[0]
	print := ""
	switch firstrap.Variable.Value {
	case "01":
		print = "Running"
	case "02":
		print = "Sleep"
	case "03":
		print = "Idle"
	case "04":
		print = "Stopped"
	case "05":
		print = "Zombie"
	case "06":
		print = "Wait"
	case "07":
		print = "Lock"
	case "ff":
		print = "Unavailable"
	default:
		print = "Unknown status"
	}
	return print
}
