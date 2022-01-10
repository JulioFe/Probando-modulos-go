package main

import "fmt"

func main() {

	CourseGo := Course{
		"Go",
		[]uint{12, 2, 4},
		map[uint]string{
			1: "Variables",
			2: "Condicionales",
			3: "Ciclos",
		},
		12000,
	}
	//GoPointer := &CourseGo
	fmt.Println(CourseGo.name)
	CourseGo.AllClasess()
	fmt.Println(CourseGo.Price)
	CourseGo.SetPrice(15000)
	fmt.Println(CourseGo.Price)

}
