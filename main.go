package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/the-palace-company/cli-github-api-go/layout"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		layout.Menu()

		fmt.Print("Seleccione una opción: ")
		scanner.Scan()
		opcion := scanner.Text()

		opcionInt, err := strconv.Atoi(opcion)
		if err != nil {
			fmt.Println("Error: entrada no válida, por favor ingrese un número.")
			continue
		}

		layout.SeleccionaOpcion(opcionInt)
		fmt.Println()
	}
}
