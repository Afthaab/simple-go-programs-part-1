package main

import "fmt"

func main() {
	fmt.Println("Structure In Golang")
	type User struct {
		Name    string
		Age     int
		City    string
		Pincode int
	}
	u := User{"Afthab", 18, "Bangalore", 571232}

	//To Print All
	fmt.Println()

	//To Print purticular value in the struct
	fmt.Println(u.Pincode)

	f := User{"Afthab", 22, "Kodagu", 571233}
	fmt.Println(f, u)

	//Pointers to a struct

	ptr := &u
	fmt.Printf("Without Using Astring %p \n", &ptr.Name)
	fmt.Println("With Using Asring", (*ptr).Name)

	//Anonymous Struct
	Element := struct {
		name      string
		branch    string
		language  string
		Particles int
	}{"afthab", "BCA", "Malayalam", 45}
	fmt.Println(Element)
	fmt.Println(Element.name)

	//Nested Struct
	type Author struct {
		name   string
		branch string
		year   int
	}

	// Creating nested structure
	type HR struct {
		//New struct Fields
		school string

		// structure as a field
		details Author
	}
	nested := HR{
		school:  "Angela",
		details: Author{"afthab", "Bca", 2015},
	}
	fmt.Println(nested)

	//Anonymous Structure with anonymous fields
}
