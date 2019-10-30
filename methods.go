package main

import "fmt"

type Family struct{
	father Father
	mother Mother
	engaged bool
}

type Father struct{
	name string
	age int
	humor string
	cash float64
}

type Mother struct{
	name string
	age int
	humor string
	cash float64
}

type Lover struct{
	humor string
}

func (family *Family) canBuyAShowTicket() bool{
	if family.father.cash + family.mother.cash < 15000{
		return false
	} else {
		return true
	}
}

func (family *Family) totalFamilyMoney() float64{
	return family.father.cash + family.mother.cash
}

func (family *Family) desengage(loverino Lover) {
	if loverino.humor == "bad"{
		family.engaged = false
	}
}

func main(){

	father1 := Father{
		"Nick",
		50,
		"good",
		5632.00,
	}

	mother1 := Mother{
		"Laura", 
		45,
		"bad",
		9156.80,
	}

	family1 := Family{
		father1,
		mother1,
		true,
	}

	lover1 := Lover{
		"good",
	}

	ticket := family1.canBuyAShowTicket()
	totalMoney := family1.totalFamilyMoney()
	
	nome, _ := nomeEidade()

	fmt.Println("Nome:", nome)
	fmt.Println("Podemos comprar um ticket? ", ticket)
	fmt.Println("Temos o total de: ", totalMoney)

	family1.desengage(lover1)

	fmt.Println("Lover with good humor, engaged? ", family1.engaged)

	lover1.humor = "bad"
	
	family1.desengage(lover1)
	
	fmt.Println("Lover with bad humor, engaged? ", family1.engaged)

}

func nomeEidade() (string, int){
	nome := "Luan"
	idade := 24

	return nome, idade
}