package main

import "fmt"

func main() {
	var cont int
	for {
		fmt.Println("Cuantas encuestas se realizaran? ")
		fmt.Scanln(&cont)

		if cont > 0 {
			break
		} else {
			fmt.Println("Ingrese un valor valido")
			var basura string
			fmt.Scanln(&basura)
			continue
		}
	}

	var firstTen []int
	var allVotos []int
	for i := 0; i < cont; i++ {
		puntaje := encuesta()
		if i < 10 {
			firstTen = append(firstTen, puntaje)
		}
		allVotos = append(allVotos, puntaje)
	}

	fmt.Println("Total primeros 10: ")
	contarVotos(firstTen)
	fmt.Println("Total de todos los votos:")
	contarVotos(allVotos)
}

func encuesta() int {
	var puntaje int
	for {
		fmt.Println("Indique un puntaje entre 1-5: ")
		fmt.Scanln(&puntaje)
		if puntaje > 0 && puntaje <= 5 {
			break
		} else {
			fmt.Println("Ingrese un valor valido")
			var basura string
			fmt.Scanln(&basura)
			continue
		}
	}
	return puntaje
}

func contarVotos(listaVotos []int) {
	mapa := make(map[int]int)
	mapa[1] = 0
	mapa[2] = 0
	mapa[3] = 0
	mapa[4] = 0
	mapa[5] = 0

	for _, v := range listaVotos {
		switch v {
		case 1:
			mapa[1]++
		case 2:
			mapa[2]++
		case 3:
			mapa[3]++
		case 4:
			mapa[4]++
		case 5:
			mapa[5]++
		}
	}

	for k, v := range mapa {
		fmt.Printf("Opción %d: %d votos\n", k, v)
	}

	if mapa[1]+mapa[2] > mapa[5]+mapa[4] {
		fmt.Println("Resultado mejorable!")
	} else {
		fmt.Println("¡Buen Resultado!")
	}
}
