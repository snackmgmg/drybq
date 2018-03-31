package utils_test

import (
	"os/exec"
	"reflect"
	"testing"

	"strconv"

	"github.com/snackmgmg/drybq/utils"
)

func TestCombineStrings(t *testing.T) {
	// test for unuse separator
	target := []string{"hoge", "fuga", "piyo"}
	result := utils.CombineStrings(target, "")
	if r := result; r != "hoge fuga piyo" {
		t.Fatalf("unexpected result: %s, want: %s", r, "hoge fuga piyo")
	}

	// test for use separator
	target2 := []string{"hoge", "fuga", "piyo"}
	result2 := utils.CombineStrings(target2, ":")
	if r := result2; r != "hoge:fuga:piyo" {
		t.Fatalf("unexpected result: %s, want: %s", r, "hoge:fuga:piyo")
	}
}

func TestMakeCmd(t *testing.T) {
	target := []string{"ls", "-l", "-a", "-h"}
	result, err := utils.MakeCmd(target)
	if err != nil {
		t.Fatalf("error raised: %v", err)
	}
	expected := exec.Command("ls", "-l", "-a", "-h")
	if r := result; !reflect.DeepEqual(r, expected) {
		t.Fatalf("made command and expected command are unmatched")
	}
}

func TestFloatEquals(t *testing.T) {
	target := 0.1

	target *= 3.0
	target /= 3.0
	target *= 3.0

	if target == 0.3 {
		t.Fatalf("%f not equal %f with calculation error", target, 0.3)
	}

	if !utils.FloatEquals(target, 0.3) {
		t.Fatalf("unexpected result: target is %f", target)
	}
}

func TestConvertByteToTByte(t *testing.T) {
	target := 1099511627776
	estimate := 1.0
	if r := utils.ConvertByteToTByte(float64(target)); !utils.FloatEquals(r, estimate) {
		t.Fatalf("get: %f, want: %f", r, estimate)
	}
}

func TestGetCost(t *testing.T) {
	target := strconv.Itoa(1099511627776 * 2)
	estimate := 10.0

	cost, err := utils.GetCost(target)
	if err != nil {
		t.Fatalf("error raised: %v", err)
	}
	if r := cost; !utils.FloatEquals(r, estimate) {
		t.Fatalf("get: %f, want: %f", r, estimate)
	}
}

func TestGetQueryBytes(t *testing.T) {
	targetStr := `running this query will process 12345678 bytes of data.`
	estimate := "12345678"

	qBytes, err := utils.GetQueryBytes(targetStr)
	if err != nil {
		t.Fatalf("error raised: %v", err)
	}
	if r := qBytes; r != estimate {
		t.Fatalf("get: %s, want: %s", r, estimate)
	}
}
