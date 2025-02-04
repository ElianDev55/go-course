package main

import (
	"fmt"
	"strconv"
	"unsafe"
)


func main () {

	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable ---->", myIntVar)

	var myUnit uint
	myUnit = 12
	fmt.Println("My variable Unit ---->",myUnit) // Always positive

	var myStringVar string
	myStringVar = "This is variable string"
	fmt.Println("My variable String ---->", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My variable boolean ---->", myBoolVar)

	fmt.Println("My Addres variable from myBoolVar --->", &myBoolVar)

	myIntVarTwo := 12
	fmt.Println("Differnt form to create variable with := --->", myIntVarTwo)

	const myIntConst = 100
	fmt.Print("My variable const ---> ", myIntConst)

	fmt.Println()
	const myStringConst string = "Hello"
	fmt.Println("My variable const with type", myStringConst)

	/*
	int8    8 bits      -128 to 127
	int16   16 bits     -2^15 to 2^15 - 1
	int32   32 bits     -2^31 to 2^31 - 1
	int64   64 bits     -2^63 to 2^63 - 1
	int     Platform dependent    Platform dependent
	*/

	fmt.Println()
	
	var my8Bits int8
	fmt.Printf("Default value : %d\n", my8Bits)
	fmt.Printf("type: %T, bytes %d, bits %d\n", my8Bits, unsafe.Sizeof(my8Bits),unsafe.Sizeof(my8Bits) * 8 )

	fmt.Println()
	var my16Bits int16
	fmt.Printf("Default value : %d\n", my16Bits)
	fmt.Printf("type: %T, bytes %d, bits %d\n", my16Bits, unsafe.Sizeof(my16Bits),unsafe.Sizeof(my16Bits) * 8 )

	fmt.Println()
	var my32Bits int32
	fmt.Printf("Default value : %d\n", my32Bits)
	fmt.Printf("type: %T, bytes %d, bits %d\n", my32Bits, unsafe.Sizeof(my32Bits),unsafe.Sizeof(my32Bits) * 8 )

	fmt.Println()
	var my64Bits int64
	fmt.Printf("Default value : %d\n", my64Bits)
	fmt.Printf("type: %T, bytes %d, bits %d\n", my64Bits, unsafe.Sizeof(my64Bits),unsafe.Sizeof(my64Bits) * 8 )

	/*
	Type    Size                Range
	uint8   8 bits              0 to 255
	uint16  16 bits             0 to 2^16 - 1
	uint32  32 bits             0 to 2^32 - 1
	uint64  64 bits             0 to 2^64 - 1
	uint    Platform dependent  Platform dependent
	*/

	fmt.Println()

	var myUint8 uint8
	fmt.Printf("Default value : %d\n", myUint8)
	fmt.Printf("type: %T, bytes %d, bits %d\n", myUint8, unsafe.Sizeof(myUint8), unsafe.Sizeof(myUint8) * 8 )

	fmt.Println()
	var myUint16 uint16
	fmt.Printf("Default value : %d\n", myUint16)
	fmt.Printf("type: %T, bytes %d, bits %d\n", myUint16, unsafe.Sizeof(myUint16), unsafe.Sizeof(myUint16) * 8 )

	fmt.Println()
	var myUint32 uint32
	fmt.Printf("Default value : %d\n", myUint32)
	fmt.Printf("type: %T, bytes %d, bits %d\n", myUint32, unsafe.Sizeof(myUint32), unsafe.Sizeof(myUint32) * 8 )

	fmt.Println()
	var myUint64 uint64
	fmt.Printf("Default value : %d\n", myUint64)
	fmt.Printf("type: %T, bytes %d, bits %d\n", myUint64, unsafe.Sizeof(myUint64), unsafe.Sizeof(myUint64) * 8 )

		//Floats

		fmt.Println()
		var myFloat32 float32
		fmt.Printf("Default value : %f\n", myFloat32)
		fmt.Printf("type: %T, bytes %d, bits %d\n", myFloat32, unsafe.Sizeof(myFloat32), unsafe.Sizeof(myUint64) * 8 )

		fmt.Println()
		var myFloat64 float64
		fmt.Printf("Default value : %f\n", myFloat64)
		fmt.Printf("type: %T, bytes %d, bits %d\n", myFloat64, unsafe.Sizeof(myFloat64), unsafe.Sizeof(myUint64) * 8 )

		// Strings
		fmt.Println()
		var myString string
		var stringDif string = `
		Hello this is a text with jump
		like this
		`
		fmt.Println(stringDif)
		fmt.Printf("Default value : `%s`\n", myString)
		fmt.Printf("type: %T, bytes %d\n", myString, unsafe.Sizeof(myString))

		{
			// Scope

			instVal1, err := strconv.ParseInt("123", 0, 64)
			fmt.Println(err)
			fmt.Println("Value of instVal1", instVal1)

			instVal2, _ := strconv.ParseFloat("123.122",64)
			fmt.Println("Value of instVal2", instVal2)

		}


}
