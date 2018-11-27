package main

import (
	"log"
	"os"
)

func ExampleGss() {
	f := func(x float64) float64 {
		tmp := x - 2
		return tmp * tmp
	}
	logger := log.New(os.Stdout, "", 0)
	Gss(f, 1, 5, 1e-6, logger)
	// Output:
	// 0	        1	        5
	// 1	        1	  3.47214
	// 2	        1	  2.52786
	// 3	  1.58359	  2.52786
	// 4	  1.58359	  2.16718
	// 5	   1.8065	  2.16718
	// 6	  1.94427	  2.16718
	// 7	  1.94427	  2.08204
	// 8	  1.94427	  2.02942
	// 9	  1.97679	  2.02942
	// 10	  1.97679	  2.00932
	// 11	  1.98922	  2.00932
	// 12	  1.99689	  2.00932
	// 13	  1.99689	  2.00457
	// 14	  1.99689	  2.00164
	// 15	  1.99871	  2.00164
	// 16	  1.99871	  2.00052
	// 17	   1.9994	  2.00052
	// 18	  1.99983	  2.00052
	// 19	  1.99983	  2.00025
	// 20	  1.99983	  2.00009
	// 21	  1.99993	  2.00009
	// 22	  1.99993	  2.00003
	// 23	  1.99997	  2.00003
	// 24	  1.99999	  2.00003
	// 25	  1.99999	  2.00001
	// 26	  1.99999	  2.00001
	// 27	        2	  2.00001
	// 28	        2	        2
	// 29	        2	        2
	// 30	        2	        2
	// 31	        2	        2
	// 32	        2	        2

}