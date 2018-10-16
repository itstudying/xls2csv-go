package xls2string

import (
	"fmt"
	"testing"
)

func TestXLS2String(t *testing.T) {
	const f = "/Users/itstudying/code/go/src/xls2csv-go/xls2csv/files/my.xls"
	var err error
	records := ""

	// Call XLS2CSV() to convert XLS and get all records.
	if records, err = XLS2String(f, 0); err != nil {
		fmt.Printf("XLS2CSV() error: %v\n", err)
	}

	fmt.Println(records)
}
