package main

import (
	"fmt"

	course "github.com/JulioFe/Probando-modulos-go/tree/poo/Course"
)

func main() {
	CourseGo := course.New("Go", 12)

	// CourseGo := Course{
	// 	"Go",
	// 	[]uint{12, 2, 4},
	// 	map[uint]string{
	// 		1: "Variables",
	// 		2: "Condicionales",
	// 		3: "Ciclos",
	// 	},
	// 	12000,
	// }
	//GoPointer := &CourseGo
	//fmt.Println(CourseGo.name)
	fmt.Println(CourseGo)

}
