package xls2csv

/*
#cgo darwin CFLAGS: -I${SRCDIR}/libxls_darwin/include
#cgo darwin LDFLAGS: -L${SRCDIR}/libxls_darwin/lib -l xlsreader
#cgo linux CFLAGS: -I${SRCDIR}/libxls_linux/include
#cgo linux LDFLAGS: -L${SRCDIR}/libxls_linux/lib -l xlsreader
#include <stdio.h>
#include <stdlib.h>
#include <libxls/xls.h>
#include "xls2csv.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func XLS2String(xlsFile string, sheetID int) (records string, err error) {
	var buf *C.char

	f := C.CString(xlsFile)
	// C string should be free after use.
	defer C.free(unsafe.Pointer(f))

	id := C.int(sheetID)

	buf = C.xls2csv(f, id)
	if buf == nil {
		err = fmt.Errorf("xls2csv() error")
		return records, nil
	}

	// Free memory block after use.
	defer C.free(unsafe.Pointer(buf))

	records = C.GoString(buf)
	return records, nil
}
