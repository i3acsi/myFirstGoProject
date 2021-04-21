package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var r rune = 'Я'          // store unicode symbol
	fmt.Println(r, string(r)) // 1071 Я
	booleans()
	integers()
	floats()
	compositeTypes()
	//task()
}

func booleans() {
	var b bool = false
	var b1 bool
	var b2 bool = 0 == 0
	fmt.Println(b, b1, b2)
}

func integers() {
	var i0 int
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
	fmt.Println(i0, i, month, i8, ui8, i16, ui16, i32, ui32, i64, ui64)
	someInt := 184
	fmt.Printf("Type %T for %v\n", someInt, someInt)
}

func floats() {
	// float32 = 4b (32 bit) ; float64 = 8b
	var f0 float64
	fmt.Println(f0)
	var f032 float32
	fmt.Println(f032)
	f064 := 0.0
	fmt.Println(f064)
	fmt.Printf("%4.2f\n", f064)
	f := 3.14
	fmt.Printf("Type %T for %v\n", f, f)
	f32 := float32(f)
	fmt.Printf("Type %T for %v\n", f32, f32)
	var answer float64 = 42
	fmt.Printf("Type %T for %v\n", answer, answer)
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi64) // Выводит: 3.141592653589793
	fmt.Println(pi32) // Выводит: 3.1415927

	// Ошибки округления
	third := 1.0 / 3.0
	fmt.Println(third + third + third) // Выводит: 1

	dot1 := 0.1
	dot3 := 0.2 + dot1
	fmt.Println(dot3 == 0.3)                             // false
	fmt.Println(dot3)                                    // Выводит: 0.30000000000000004
	fmt.Println(math.Abs(dot3-0.3) < 0.0001)             // true
	fmt.Println(math.Abs(dot3-0.3) * 100000000000000000) // Выводит: 5.551115123125783

	celsius := 21.0
	fmt.Print((celsius/5.0*9.0)+32, "° F\n")
	fmt.Print((9.0/5.0*celsius)+32, "° F\n")
	// В выводе: 69.80000000000001° F

	fahrenheit := (celsius * 9.0 / 5.0) + 32.0
	fmt.Println(fahrenheit, "° F") // Выводит: 69.8° F

}

func compositeTypes() {
	nums := [3]int{1, 2, 3}
	fmt.Println(nums, nums[1])

	three := "thre"
	three += "e"
	words := [4]string{"one", "two", three, "four"}
	wordList := []string{"one", "two", three}
	wordList = append(wordList, "four")
	fmt.Println(words, words[3], wordList)

	dict := map[int]string{}
	dict[1] = "one"
	dict[2] = "two"
	fmt.Println(dict, dict[0], dict[2])

	type person struct {
		name string
		age  int
	}

	persons := map[string]person{}
	persons["ya"] = person{name: "ya", age: 34}
	persons["ya"] = person{name: "ya", age: 35}
	fmt.Println(persons)

}

func task() {
	//	Представьте, что вам нужно накопить денег на подарок другу.
	//	Напишите программу, которая случайным образом размещает монеты пять ($0.05), десять ($0.10) и двадцать пять ($0.25) центов
	//	в пустую копилку до тех пор, пока внутри не будет хотя бы двадцать долларов ($20.00).
	//	Пускай после каждого пополнения копилки текущий баланс отображается на экране,
	//	отформатированный с нужной шириной и точностью.
	var cash float64 = 0.0
	fmt.Println("\nInitial cash is 0.0")
	for cash < 20.0 {
		coin := randomCoin()
		cash += coin
		fmt.Printf("Received %2.2f. Total cash is %2.2f\n", coin, cash)
	}
}

func randomCoin() float64 {
	switch rand.Intn(3) {
	case 0:
		return 0.05
	case 1:
		return 0.10
	case 2:
		return 0.25
	}
	return 0
}
