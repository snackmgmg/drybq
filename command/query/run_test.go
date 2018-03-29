package query_test

import (
	"testing"

	"strconv"

	"github.com/snackmgmg/drybq/command/query"
	"github.com/snackmgmg/drybq/utils"
)

func TestConvertByteToTByte(t *testing.T) {
	target := 1099511627776
	estimate := 1.0
	if r := query.ConvertByteToTByte(float64(target)); !utils.FloatEquals(r, estimate) {
		t.Fatalf("get: %f, want: %f", r, estimate)
	}
}

func TestGetCost(t *testing.T) {
	target := strconv.Itoa(1099511627776 * 2)
	estimate := 10.0

	cost, err := query.GetCost(target)
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

	qBytes, err := query.GetQueryBytes(targetStr)
	if err != nil {
		t.Fatalf("error raised: %v", err)
	}
	if r := qBytes; r != estimate {
		t.Fatalf("get: %s, want: %s", r, estimate)
	}
}
