# xls2csv-go

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

* Install `xls2csv` package

        CGO_CFLAGS=-I/usr/local/libxls/include CGO_LDFLAGS="-L/usr/local/libxls/lib -l xlsreader" go get github.com/itstudying/xls2csv-go/xls2csv

#### Usage

    func main() {
    	const f  = "my.xls"
    	var err error
    	records := [][]string{}
    
    	// Call XLS2CSV() to convert XLS and get all records.
    	if records, err = xls2csv.XLS2CSV(f, 0); err != nil {
    		log.Printf("XLS2CSV() error: %v\n", err)
    		goto end
    	}
    
    	for i, row := range records {
    		fmt.Printf("%v", row)
    		if i != len(records)-1 {
    			fmt.Printf("\n")
    		}
    	}
    
    end:
    }

#### Build and Test Your App
  * Do not forget to add `CGO_CFLAGS=-I/usr/local/libxls/include CGO_LDFLAGS="-L/usr/local/libxls/lib -l xlsreader"` before `go build` or `go test`

          go build
          
  > 在编译时需要加上libxls库中头文件的检索目录CGO_CFLAGS，链接运行时已经将darwin和linux环境下的依赖文件加入至项目中，因此不再需要CGO_LDFLAGS，见/xls2csv.go:3

