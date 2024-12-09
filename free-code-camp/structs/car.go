package structs

import "fmt"

// CarPublic Car is a public
type CarPublic struct {
	Brand string
	Year  int64
}

// ** !! For make a struct public you have write in Mayuscula the frist word and comment
// ** !! For make a struct private you have write in minus the frist word and comment


func HelloWorld() {
	fmt.Println("Hello World!")
}
