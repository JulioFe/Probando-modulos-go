package main

import "fmt"

func main() {
	// var number uint8 = 12
	// var big_number uint32 = 1000000
	// valor := "Valor string"
	// var puntero *string = &valor
	// *puntero = "Valor cambiado desde el puntero"

	// var arry [1]string

	slice_animals := []string{"perro", "gato"}
	slice_animals = append(slice_animals, "ballena")
	//slice_cars := make([]string, 0, 3)
	//varo := []string{}

	map_animals := make(map[int]string)
	map_animals[1] = "gato"

	map_fruits := map[uint8]string{
		1: "manzana",
		2: "Pera",
	}

	delete(map_fruits, 1)

	if _, ok := map_animals[2]; !ok {
		map_animals[2] = "Gorilla"
	}

	type course struct {
		Name    string
		Country string
	}

	db := course{"Jose", "Canada"}

	db.Name = "Arnold"

	names := []string{"Juan", "Sofia", "Matias"}

	for i, v := range names {
		fmt.Printf("valor: %v, indice: %v \n", v, i)
		names[i] = "Arnold"

	}

	nums := []int{15, 12, 50, 20, 4, 6, 9}
	filter_nums := filter(nums, filter_callback)

	fmt.Println(filter_nums)

	saludo("Jose")("Alvarez")

	suma := sum(1, 2, 3, 4)
	fmt.Println(suma)

	x := func(num int) func(int) {
		return func(num2 int) {
			fmt.Println(num + num2)
		}
	}
	x(5)(20)

	result := num_secure(0)

	fmt.Println(result)

}

func filter_callback(number int) bool {
	if number > 10 {
		return true
	}
	return false
}

func filter(nums []int, filter_callback func(int) bool) []int {
	filter_nums := []int{}
	for _, v := range nums {
		if filter_callback(v) {
			filter_nums = append(filter_nums, v)
		}
	}
	return filter_nums
}

func saludo(name string) func(string) {
	return func(lastname string) {
		fmt.Println(name + " " + lastname)
	}
}

func sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func num_secure(num int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recuperandonos de %v", r)
		}
	}()
	return panic_func(num)
}

func panic_func(num int) int {
	if num == 0 {
		panic("0 no compa")
	}
	return num / 1
}
