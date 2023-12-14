package main

import "fmt"

func main() {
	/*arr := [8]float64{1, 2, 3, 4, 5, 6, 7, 8}
	fatia1 := arr[2:5]
	fatia2 := append(fatia1, 0, 8)
	fmt.Println(fatia1, fatia2)*/

	cliente := make(map[string]string)
	cliente["Joãozinho"] = "Rua nove n° 9"
	cliente["Pedrinho"] = "Rua dez n° 20"
	cliente["Luizinho"] = "Rua seis n° 136"
	fmt.Println(cliente["Pedrinho"])
}
