package main

import (
	"fmt"
	"sort"
)

type producto struct {
	nombre   string
	cantidad int
	precio   int
}
type listaProductos []producto

var lProductos listaProductos
var lProductosMinima listaProductos

const existenciaMinima int = 10 //la existencia mínima es el número mínimo debajo de el cual se deben tomar eventuales desiciones

func (l *listaProductos) agregarProducto(nombre string, cantidad int, precio int) {
	finded := l.buscarProducto(nombre)
	if finded != -1 {
		(*l)[finded].precio = precio
		(*l)[finded].cantidad = (*l)[finded].cantidad + cantidad
	} else {
		*l = append(*l, producto{nombre: nombre, cantidad: cantidad, precio: precio})
	}
	// modificar el código para que cuando se agregue un producto, si este ya se encuentra, incrementar la cantidad
	// de elementos del producto y eventualmente el precio si es que es diferente
}

func (l *listaProductos) buscarProducto(nombre string) int { //el retorno es el índice del producto encontrado y -1 si no existe
	var result = -1
	var i int
	for i = 0; i < len(*l); i++ {
		if (*l)[i].nombre == nombre {
			result = i
		}
	}
	return result
}

func (l *listaProductos) venderProducto(nombre string, n int) {
	var prod = l.buscarProducto(nombre)
	if prod != -1 {
		if (*l)[prod].cantidad-n < 0 {
			println("La cantidad solicitada sobrepasa la cantidad en inventario")
		} else {
			(*l)[prod].cantidad = (*l)[prod].cantidad - n
			if (*l)[prod].cantidad == 0 {
				*l = append((*l)[:prod], (*l)[prod+1:]...)
			}
		}
		//modificar para que cuando no haya existencia de cantidad de productos, el producto se elimine de "la lista" y notificar
	}
}
func llenarDatos() {
	lProductos.agregarProducto("arroz", 15, 2500)
	lProductos.agregarProducto("frijoles", 4, 2000)
	lProductos.agregarProducto("leche", 8, 1200)
	lProductos.agregarProducto("café", 12, 4500)
	lProductos.agregarProducto("cereal", 9, 3000)
	lProductos.agregarProducto("huevos", 18, 2500)
	lProductos.agregarProducto("pollo", 6, 5000)
	lProductos.agregarProducto("pescado", 13, 3500)
}

func (l *listaProductos) listarProductosMínimos() {
	for i := 0; i < len(*l); i++ {
		if (*l)[i].cantidad < existenciaMinima {
			lProductosMinima.agregarProducto((*l)[i].nombre, (*l)[i].cantidad, (*l)[i].precio)
		}
	}
	// debe retornar una nueva lista con productos con existencia mínima
}

// //a.	A partir de la lista de productos con mínimas existencias de inventario, ampliar el inventario
// con la compra de más unidades de cada producto de esta lista hasta que cumplan con el mínimo establecido.
func (l *listaProductos) aumentarInventarioDeMinimos() {
	for i := 0; i < len(lProductosMinima); i++ {
		index := (*l).buscarProducto(lProductosMinima[i].nombre)
		(*l)[index].cantidad = existenciaMinima
	}
}

// //c.	Crear una función que ordene la lista de productos usando como llave para el ordenamiento cualquiera
// de los elementos de la estructura producto. La lista/slice debe quedar modificada al finalizar el método.
// Se solicita investigar y hacer uso de alguna función predefinida de alguna librería del lenguaje Go que ayude a resolver el problema.
func (l *listaProductos) ordenarProductosXPrecio() {
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].precio > (*l)[j].precio //Ordena el slice de mayor a menor respecto al precio del producto
	})
}

func main() {
	llenarDatos()
	fmt.Println("Datos iniciales: \n", lProductos)
	lProductos.venderProducto("café", 12)
	fmt.Println("\nLuego de vender todas las unidades de cafe: \n", lProductos)
	lProductos.listarProductosMínimos()
	fmt.Println("\nProductos por debajo de la cantidad minima: \n", lProductosMinima)
	lProductos.aumentarInventarioDeMinimos()
	fmt.Println("\nLista con el inventario aumentado por encima de la cMin: \n", lProductos)
	lProductos.ordenarProductosXPrecio()
	fmt.Println("\nLista ordenada de mayor a menor respecto al precio: \n", lProductos)
}
