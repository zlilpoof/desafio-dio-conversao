package main

import "fmt"

const valorK int64 = 273

func main() {

	var tempAgua float64 = 373
	var conversaoValor = tempAgua - float64(valorK)
	fmt.Printf("O valor convertido é: %g", conversaoValor)

}
