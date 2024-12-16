package main

import (
	"bufio"
	"fmt"
	"os"
)

//Função que abre um arquivo externo e conta as linhas desse arquivo
func countRows(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	defer file.Close()

	//Cria um Scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)

	//Conta as linhas
	count := 0

	for scanner.Scan() {
		count++
	}

	return count, nil
}


func main() {
	files := []string {"file1.txt", "file2.txt", "file3.txt"}

	//Cria um canal (canal é um tipo como string, int ...)
	rowChannel := make(chan int, len(files))

	for _, file := range files {
		fileToRead := file

		//Inicializa a Go routine
		go func() {
			fmt.Println("Reading goroutine")

			numRows, err := countRows(fileToRead)
			if err != nil {
				panic("Ops!!") //O programa quebra(para) se cair aqui!
			}

			//Escreve no channel
			rowChannel <- numRows
		}()
	}

	var totalRows int

	for range files {
		//Lê o channel. O channel deve ser esvaziado, pois ele pode ser sobescrito posteriormente
		rows := <- rowChannel

		totalRows += rows
	}

	fmt.Println("Soma das linhas dos arquivos: ", totalRows)
}
