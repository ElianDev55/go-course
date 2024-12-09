package main // always you have to call this
import "fmt" // if you want to print something so  important this
func main() {

	// String

	var name string = "Elian"
	var lastName = "Villamarin"
	lastNameTwo := "Urrutia"
	
	// Number

	var age int8 = 19
	year := 2005

	// list

	var listFruts = [4]string{"Apple", "Orange"}  // limit the number  and type 
	listCountries := [4]string{}
	listCountries[0] = "Colombia"
	listCountries[1] = "Ecuador"
	listCountries[1] = "Usa"

	// list without limit
	var listCountriesUnlimited []string
  listCountriesUnlimited = append(listCountriesUnlimited, "Colombia")
  listCountriesUnlimited = append(listCountriesUnlimited, "Ecuador")
  listCountriesUnlimited = append(listCountriesUnlimited, "Usa")

	listCountriesTwo := listCountriesUnlimited[0:2] // range with initial and end
	listCountriesTheree := listCountriesUnlimited[1:] // range with initial but without end
	listCountriesFour := listCountriesUnlimited[:2] // range with without initial but end


		fmt.Println("hello world i am " + name + " " + lastName + " " + lastNameTwo) // we print all strings
		fmt.Println("I am " , age , " years old") // we print the number 
		fmt.Println("I was born in " , year) // we print the number
		fmt.Println ("List of", listFruts) // we print the list
		fmt.Println ("List of", listFruts[1]) // we print the value numer 1 ont the list
		fmt.Println ("List of countries", listCountries) // we print the list
		fmt.Println ("List of countries unlimited", listCountriesUnlimited) // we print the list
		fmt.Println ("List of countries two", listCountriesTwo) // we print the list
		fmt.Println ("List of countries theree", listCountriesTheree) // we print the list
		fmt.Println ("List of countries four", listCountriesFour) // we print the list
}
