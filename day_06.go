package main

import "fmt"

func Day06Main() {
	fmt.Println("Day )> 6")
	someOrange()
}

func someOrange() {
	fmt.Println("Just a test. 🦥")
	printSomething("This is intersting")
	printSomething(1994)
	op := addGeneric(10, 20)
	fmt.Println(op)

	newProduct := NewProduct("some product", 1001, 199.99, []string{"some", "things"})
	sample_array := []int{1, 20, 3}
	fmt.Println(*newProduct)
	fmt.Println(len(sample_array))

	for i := 0; i < len(sample_array); i++ {
		fmt.Println(sample_array[i])
	}
}

// Generic in GoLang
func addGeneric[T int | float64](a, b T) T {
	return a + b
}

// Arrays -> if you are comming from python its a list of data, if you are comming from any static typed languages you already know what Arrays are.

type Product struct {
	title string
	id    int
	price float64
	tags  []string
}

func (product *Product) UpdateProduct(title string, id int, price float64, tags []string) {
	product.title = title
	product.id = id
	product.price = price
	product.tags = tags
}

func NewProduct(title string, id int, price float64, tags []string) *Product {
	return &Product{
		title,
		id,
		price,
		tags,
	}
}
