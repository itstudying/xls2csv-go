# xls2csv-go

[![Go Report Card](https://goreportcard.com/badge/github.com/itstudying/xls2csv-go)](https://goreportcard.com/report/github.com/itstudying/xls2csv-go)

[![GoDoc](https://godoc.org/github.com/itstudying/xls2csv-go/xls2string?status.svg)](https://godoc.org/github.com/itstudying/xls2csv-go/xls2string)

此包以cgo的方式调用c的libxls库，用于解析xls和xlsx文件，以及将xls文件转换至csv

#### Install `xls2csv` package
* xls2csv requires [libxls](http://libxls.sourceforge.net/) to be installed.

  * Install libxls

          wget http://downloads.sourceforge.net/libxls/libxls-0.2.0.tar.gz
          tar -xzvf libxls-0.2.0.tar.gz
          cd libxls-0.2.0
          ./configure
          make
          sudo make install
  * Add libxls lib path to `LD_LIBRARY_PATH` （mac环境下不需要）
    * Create a new `/etc/ld.so.conf.d/libxls.conf`

              sudo vi /etc/ld.so.conf.d/libxls.conf

              // Add path of libxls to this file
              /usr/local/libxls/lib

    * Update `LD_LIBRARY_PATH`

            sudo ldconfig
            // Check libxlsreader.so
            sudo ldconfig -p | grep libxls

#### Build and Test Your App
  * Do not forget to add `CGO_CFLAGS=-I/usr/local/libxls/include CGO_LDFLAGS="-L/usr/local/libxls/lib -l xlsreader"` before `go build` or `go test`

          go build
          

