## GoLang

**Definições do Go**

A função "main" é a principal função do GO. É ela que chama todas as outrs funções. Essa função não tem retorno por padrão.

- import - importa os pacotes do go. Assim como nas outras linguagens.

<br />

**Tipos de variáveis**

- Boleano = bool (false se não inicializada)
- String = string
- Inteiro = int
- Inteiro 8bits = int8
- Inteiro 16bits = int16
- Inteiro 32bits = int32
- Inteiro 64bits = int64
- Float 32bits = float32
- Float 64bits = float64

declaração de uma variável:

```go
var boolean bool
```

inicialização de uma variável:

```go
var myText = "Hello text"
```

- forma rápida (só funciona dentro de um escopo):

```go
 myText2 := "Hello Fast"
```

Obs.: variáveis iniciadas com letra "maiuscula" podem ser usadas em outros arquivos (pacotes).

**Declarando constante**

```go
const STATE = "done"
```

<br />

**Operadores aritméticos**

Os de sempre: + - \* / = > < OK?

Exemplos de uso:

```go
x = 10
x = x / 10
x += 1

y = 4

boolean = x <= y
```

<br />

**Concatenação**

Primeiro método: usando o sinal de "+".

```go
hello := "Hello"
greet := hello + " World"
```

Segundo método: usando o pacote FMT.

```go
fmt.Printf("%s World of nice devs!!", hello)
```

<br />

**Conversão string <-> int**

- convertendo "string" em "int"

É usado o pacote strconv e a função Atoi. Exemplo:

```go
value := "123"
s, err := strconv.Atoi(value)
if err != nil {
	panic("Problema!!")
}
```

- convertendo "int" em "string"

É usado o pacote strconv e a função Itoa. Exemplo:

```go
value := 123
s := strconv.Itoa(value)
s += "xxx"
fmt.Println(s)
```

- conventendo "string" em "boleano"

É usado o pacote strconv e a função ParseBool. Exemplo:

```go
value := "true"
s, err := strconv.ParseBool(value)
if err != nil {
	panic("Problema!!")
}

fmt.Println(s)
```

**ARRAY**

- primeiro método:

```go
var array [2]string
array[0] = "a"
array[1] = "b"
```

- segundo método:

```go
array1 := [2]string {"a", "b"}
```

**SLICE**

É um array com tamanho indeterminado.

```go
slice := []string {"a", "b", "c", "d", "e"}
fmt.Println("Slice: ", slice)
```

Podemos cortar um slice (RANGE)

```go
slice2 := slice[:3]
//saída: [a b c]

slice3 := slice[1:3]
//saída: [b c]
```

Podemos adicionar um valor ao slice (APPEND)

```go
slice3 := append(slice, "f")
//saída: [a b c d e f]
```

**MAP**

É um array, porém seus itens são compostos por chave e valor. (ex.: "id": 1).

```go
gameState := map[string]int {
		"coins": 22,
		"life": 2,
		"blood": 50,
	}

fmt.Println("Map: ", gameState["coins"])
fmt.Println("Map: ", gameState["coins"] - 1)
```

Podemos deletar um item do map.

```go
delete(gameState, "life")
fmt.Println("Map: ", gameState)
```

**STRUCTS**

É um grupo de informações. Quando os tipos são declarados com letra maiúsculas, podem ser usados fora do pacote.

Uma Structs pode ser usado dentro de outra Structs.

```go
type Character struct {
	Name string
	Color string
}

type GameState struct {
	Coins int
	Life int
	Blood int
	Char Character
	Characters map[string]Character
}


func main() {} //É declarado fora da função como um type no TS!!
```

Consumnido o STRUCTS

```go
players := map[string]gameState {
	"lais": {
		Coins: 22,
		Life: 1,
		Blood: 100,
		Chararacter: "Princess",
	},
	"gui": {
		Coins: 50,
		Life: 2,
		Blood: 80,
		Chararacter: "Wizard",
	},
}
fmt.Println(players["gui"].Chararacter)
```

**Fluxos de controles**

- IF / ELSE IF

```go
if players["lais"].Chararacter == "Princess" {
	fmt.Println("Oh, Mario!")
} else if players["gui"].Chararacter =="Wizard" {
	fmt.Println("I'm a Wizard!!")
} else {
	fmt.Println("Who am I?")
}
```

- SWITCH CASE

```go
character := players["lais"].Chararacter
switch character {
case "Princess":
	fmt.Println("Oh, Mario!")
case "Wizard":
	fmt.Println("I'm a Wizard!!")
default:
	fmt.Println("Who am I?")
}
```

**Estrutura de repetição**

- FOR

```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}


letters := []string {"a", "b", "c"}

for index, value := range letters {
	fmt.Println(index, value)
}


for player, state := range players {
	fmt.Println(player, state)
}
```

Obs.: não existe a estrutura de repetição WHILE no GO.

<br />

**Function**

declaração:

```go
func nomeDaFunção(propridade tipoProp) tipoReturno {
	corpo da função
	return ("retorno obrigatório se não for void")
}
```

Exemplo:

```go
func whatTheCharacterSay(players map[string]gameState) string {
	for player, state := range players {
		fmt.Println("player: ", player)
		switch state.Chararacter {
		case "Princess":
			return "Its me, Princess!"
		case "Wizard":
			return "I'm a Wizard!!"
		default:
			return "Who am I?"
		}
	}

	return ""
}
```

**Método**

Funciona como uma função, porém trabalha em cima de uma estrutura específica (STRUCTS).

Declaração:

```go
func (structs tipoDaStructs) nomeDoMetodo(propridade tipoProp) tipoReturno {
	corpo do método
	return ("retorno obrigatório se não for void")
}
```

Exemplo:

```go
func (g marioState) GetCharacterAndBlood() (string, int) {
	return g.chararacter, g.blood
}
```

Usando o método:

```go
state := marioState {
	coins: 4,
	life: 2,
	blood: 50,
	chararacter: "Mario",
}
cha, bd := state.GetCharacterAndBlood()
fmt.Printf("character: %s blood: %d", cha, bd)
```

**Interface**

Exemplo:

```go
type gameStateGetter interface {
	GetCharacterAndBlood() (string, int)
}
```

Usando a interface:

```go
var kirbyState gameStateGetter = kirbyState {
	character: "Kirby",
	blood: 75,
	power: 94,
}
```

A interface é usada como o tipo da variável e é obrigatório implementar, nesse caso, o método que está dentro dela. Pois a interface é um contrato que exirge certas obrigatoriedades, assim como em outras linguagens orientada a objetos.
