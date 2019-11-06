package main

import "fmt"

func main(){

	// make é para criar um slice vazio

	var dict map[string]int
	fmt.Println(dict)

	dict2 := make(map[string]int)
	fmt.Println(dict2)

	dict3 := map[string]int{"teste": 2}
	fmt.Println(dict3)
}


// dict[key]: valor


// map[tipo_da_key]tipo_do_valor