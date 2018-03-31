package utils

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

var EPSILON = 0.00000001

// ToDo: read from config file
const COSTPERTB = 5.0

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

// GetQueryBytes is get query byte size
func GetQueryBytes(str string) (string, error) {
	regex := regexp.MustCompile(`running this query will process (\d+) bytes of data.`)
	queryBytes := regex.FindStringSubmatch(string(str))
	if len(queryBytes) != 2 {
		return "", fmt.Errorf("unexpected result: bytes count is %d, must be %d", len(queryBytes), 1)
	}
	return queryBytes[1], nil
}

// GetCost is get query cost
func GetCost(size string) (float64, error) {
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		return 0.0, err
	}
	tByte := convertByteToTByte(float64(sizeInt))
	return COSTPERTB * tByte, nil
}

func convertByteToTByte(b float64) float64 {
	return b / (1024 * 1024 * 1024 * 1024)
}
