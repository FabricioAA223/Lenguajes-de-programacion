package main

import (
	"fmt"
)

func ejercicio1(p string) {
	/*1)	Haga un programa que cuente e indique el número de caracteres,
	el número de palabras y el número de líneas de un texto cualquiera
	(obtenido de cualquier forma que considere). Considere hacer el cálculo
	de las tres variables en el mismo proceso.*/

	fmt.Println(p) //Texto a utilizar

	n_words := 1    //Variable para almacenar la cantidad de palabras (Empieza en uno por la palabra inicial)
	n_lines := 1    //Variable para almacenar la cantidad de lineas (Empieza en uno por la linea inicial)
	var n_chars int //Variable para almacenar la cantidad de caracteres

	ascii := []rune(p) //Arreglo obtener el ascii(int) de cada caracter del texto

	for i := 0; i < len(ascii); i++ {
		if ascii[i] == 32 { //El ascii 32 equivale a un espacio " "
			n_words = n_words + 1 //Si hay un espacio en el texto es porque termino una palabra, por lo que se aumenta
		} else if ascii[i] == 10 { //El ascii 10 equivale a un salto de linea "\n"
			n_lines = n_lines + 1 //Se aumenta la cantidad de lineas del texto
			n_words = n_words + 1 //La aumenta por la palabra al final de la linea
		} else { //El resto de ascii seran caracteres validos
			n_chars = n_chars + 1 //Se aumenta la variable
		}
	}
	fmt.Printf("Cantidad de lineas: %d\n", n_lines)
	fmt.Printf("Cantidad de palabras: %d\n", n_words)
	fmt.Printf("Cantidad de caracteres: %d\n", n_chars)
}

func ejercicio2(n int) {

	if n%2 == 0 || n < 0 { //Quiere decir que es un numero par o que es negativo
		println("N debe de ser un numero positivo e impar")
	} else {
		for linea := 1; linea <= n*2; linea++ { //linea que se esta imprimiendo
			l := linea
			if linea > n { //Si ya se imprimio la linea del centro
				l = (n * 2) - linea //Se multiplica la cantidad de * por dos y se resta la linea actual
			} //Para que sea al equivalente en el descenso
			if l%2 == 0 { //Si es una linea par
				fmt.Print("\n") //Solo imprime un salto de linea
			} else { //Si no
				tabs := (n - l) / 2 //La cantidad de espacios antes de empezar a pintar los *
				p := 0              //Indice para el for
				for p < tabs {      //Ciclo para imprimir los espacios en blanco
					print("    ") //Se imprime los espacios
					p = p + 1     //Se aumenta el indice
				}
				for m := 0; m < l; m++ { //Ciclo para imprimir los *
					print("*   ")
				}
				print("\n") //Se salta a la linea siguiente
			}
		}
	}
}

func ejercicio3(l *[]string, movimientos, dirrect int) {
	fmt.Println("Arreglo antes de la rotacion: ", *l)
	if dirrect == 0 {
		for i := 0; i < movimientos; i++ {
			*l = append((*l)[1:], (*l)[0])
		}
	} else if dirrect == 1 {
		for i := 0; i < movimientos; i++ {
			*l = append((*l)[len(*l)-1:], (*l)[:len(*l)-1]...)
		}
	}
	fmt.Println("Arreglo despues de la rotacion: ", *l)

}

func main() {
	/* //Ejercicio1
	p := "El martes recibo reque\n" +
		"el miercoles recibo lenguajes\n" +
		"el jueves recibo bases de datos"
	ejercicio1(p)
	*/

	/* //Ejercicio2
	n := 9                 //Cantidad de * por filas y columnas
	ejercicio2(n)
	*/

	/* //Ejercicio3
	arreglo := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	ejercicio3(&arreglo, 3, 0)
	*/
}
