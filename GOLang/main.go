package main

import "fmt"

//Declaração de variáveis
var boolean bool
var myString string
var x, y int

//STRUCTS
type marioState struct {
	coins int
	life int
	blood int
	chararacter string
}
type kirbyState struct {
	character string
	blood int
	power int
}

//MÉTODO
func (g marioState) GetCharacterAndBlood() (string, int) {
	return g.chararacter, g.blood
}

func (k kirbyState) GetCharacterAndBlood() (string, int) {
	return k.character, k.blood
}

//INTERFACE
type gameStateGetter interface {
	GetCharacterAndBlood() (string, int)
}

//Função principal do pacote
func main() {
	//Usando o método
	stateMario := marioState {
		coins: 4,
		life: 2,
		blood: 50,
		chararacter: "Mario",
	}

	cha, bd := stateMario.GetCharacterAndBlood()
	fmt.Println("===Mario===")
	fmt.Printf("character: %s blood: %d \n", cha, bd)

	//Usando a interface
	var kirbyState gameStateGetter = kirbyState {
		character: "Kirby",
		blood: 75,
		power: 94,
	}

	chaK, bdK := kirbyState.GetCharacterAndBlood()
	fmt.Println("===Kirby===")
	fmt.Printf("character: %s blood: %d \n", chaK, bdK)


	// hello := "Hello"
	// greet := hello + " World"
	// fmt.Printf("%s World of nice devs!!", hello)

	//Conversão int to string
	// value := 123
	// s := strconv.Itoa(value)
	// s += "xxx" 
	// fmt.Println(s)

	//Conversão string to boolean
	// value := "true"
	// s, err := strconv.ParseBool(value)
	// if err != nil {
	// 	panic("Problema!!")
	// }

	//ARRAY
	// var array [2]string
	// array[0] = "a"
	// array[1] = "b"

	// array1 := [2]string {"a", "b"}

	// fmt.Println("Array: ", array1)

	//SLICES
	// slice := []string {"a", "b", "c", "d", "e"}
	// fmt.Println("Slice: ", slice)


	//SLICES-RANGE - cortar um slice
	// slice2 := slice[:3]
	// slice2 := slice[1:3]
	// fmt.Println("Range: ", slice2)

	// //SLICES-APPEND
	// slice3 := append(slice, "f")
	// fmt.Println("Adicionado o f: ", slice3)

	//MAPS
	// gameState := map[string]int {
	// 	"coins": 22,
	// 	"life": 2,
	// 	"blood": 50,
	// }

	// delete(gameState, "life")
	// fmt.Println("Map: ", gameState)

	//Consumnido o STRUCTS
	// players := map[string]gameState {
	// 	"lais": {
	// 		Coins: 22,
	// 		Life: 1,
	// 		Blood: 100,
	// 		Chararacter: "Princess",
	// 	},
	// 	"gui": {
	// 		Coins: 50,
	// 		Life: 2,
	// 		Blood: 80,
	// 		Chararacter: "Wizard",
	// 	},
	// }

	// fmt.Println(players["gui"].Chararacter)

	//IF / ELSE IF
	// if players["lais"].Chararacter == "Princess" {
	// 	fmt.Println("Oh, Mario!")
	// } else if players["gui"].Chararacter == "Wizard" {
	// 	fmt.Println("I'm a Wizard!!")
	// } else {
	// 	fmt.Println("Who am I?")
	// }

	//SWITCH CASE
	// character := players["lais"].Chararacter

	// switch character {
	// case "Princess":
	// 	fmt.Println("Oh, Mario!")
	// case "Wizard":
	// 	fmt.Println("I'm a Wizard!!")
	// default:
	// 	fmt.Println("Who am I?")
	// }

	//FOR
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// letters := []string {"a", "b", "c"}
	
	// for index, value := range letters {
	// 	fmt.Println(index, value)
	// }

	// for player, state := range players {
	// 	fmt.Println(player, state)
	// }
	
	//Chama a função abaixo
	// message := whatTheCharacterSay(players)
	// fmt.Println(message)
	

}

//FUNCTION (exemplo)
// func whatTheCharacterSay(players map[string]gameState) string {
// 	for player, state := range players {
// 		fmt.Println("player: ", player)
// 		switch state.Chararacter {
// 		case "Princess":
// 			return "Its me, Princess!"
// 		case "Wizard":
// 			return "I'm a Wizard!!"
// 		default:
// 			return "Who am I?"
// 		}
// 	}

// 	return ""
// }
