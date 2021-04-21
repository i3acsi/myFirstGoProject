package main

import "fmt"

func main() {
	basicTypes()
}

func basicTypes() {
	var b bool = false
	var b1 bool
	var b2 bool = 0 == 0
	var s string = "new string"
	var i int = 2021
	var month uint = 0                     // uint >= 0, if less then 0 -> compilation error
	var i8 int8 = 127                      // -128...127  1b
	var ui8 uint8 = 0                      // 0...255	1b
	var i16 int16 = -32768                 // -32768...32767	2b
	var ui16 uint16 = 65535                // 0...65535	2b
	var i32 int32 = 2147483647             //-2 147 483 648...2 147 483 647		4b
	var ui32 uint32 = 4294967295           //0...4 294 967 295		4b
	var i64 int64 = 922337203685477580     //–9 223 372 036 854 775 808...9 223 372 036 854 775 80
	var ui64 uint64 = 18446744073709551615 //0...18 446 744 073 709 551 615		4b
	fmt.Println(b, b1, b2, s, i, month, i8, ui8, i16, ui16, i32, ui32, i64, ui64)
}
