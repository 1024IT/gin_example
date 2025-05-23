package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	// Time to practice what you learned!

	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := []string{"Sports", "Cooking", "Reading"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//	and last element of the original array.

	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5) Create a "dynamic array(ダイナミックアレイ)" that contains your course goals (at least 2 goals)

	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array

	courseGoals[1] = "Learn all the details!"
	courseGoals = append(courseGoals, "Learn all the basice!")
	fmt.Println(courseGoals)

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//dynamic list of products (at least 2 products).
	//Then add a third product to the existing list of products.

	products := []Product{
		{
			"first-product",
			"A First Product",
			12.99,
		},
		{
			"second-product",
			"A Second Product",
			129.99,
		},
	}

	fmt.Println(products)

	newProduct := Product{
		"third-product",
		"A Third Product",
		15.99,
	}

	products = append(products, newProduct)
	fmt.Println(products)

}
