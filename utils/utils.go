package utils

import (
	"fmt"
	"os/exec"
)

var EPSILON = 0.00000001

// CombineStrings is combine slice strings
func CombineStrings(strs []string, sep string) string {
	if sep == "" {
		sep = " "
	}
	combined := ""
	for i, s := range strs {
		combined += s
		if i != len(strs)-1 {
			combined += sep
		}
	}
	return combined
}

// MakeCmd is make command from args
func MakeCmd(strs []string) (*exec.Cmd, error) {
	switch len(strs) {
	case 0:
		return nil, fmt.Errorf("executable command not found")
	case 1:
		return exec.Command(strs[0]), nil
	default:
		return exec.Command(strs[0], strs[1:]...), nil
	}
	// not reached
	return nil, fmt.Errorf("unexpected error raised")
}

// FloatEquals is compare two float value
func FloatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}
