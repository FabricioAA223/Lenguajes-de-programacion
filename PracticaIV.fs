module Ejemplos3.GrafosPesos.Program

open Microsoft.FSharp.Data.UnitSystems.SI.UnitNames

// Ejercicios de rutas en Grafos con búsqueda en profundidad

//        a --- c ---- x
//      /   \  /
//     /     \/
//    i      X
//     \    / \
//      \  /   \
//        b --- d ---- f
let grafo = [
    [('i', 0);('a', 1);('b', 2)];
    [('a', 1);('i', 0);('c', 3);('d', 2)];
    [('b', 2);('i', 0);('c', 1);('d', 4)];
    [('c', 3);('a', 1);('b', 1);('x', 2)];
    [('d', 2);('a', 4);('b', 1);('f', 3)];
    [('f', 3);('d', 2)];
    [('x', 2);('c', 3)]
]

//1)	Modifique el ejercicio de búsqueda en profundidad visto en clase para cálculo de rutas en un grafo,
//para ingresar pesos a los vértices y con esta nueva información, calcular el camino más corto en la búsqueda en profundidad
//basándose en la sumatoria de pesos de los vértices de la ruta. Una opción sin mucho cambio al código hecho es generar todos
//los caminos posibles (con su información de pesos) y luego a partir de esto, filtrar el de menos tamaño en función a la información de pesos.

let rec menorLista lista (menor:int) (index:int)=
    if lista = [] then
        index
    elif lista.Head < menor then
        menorLista lista.Tail lista.Head lista.Tail.Length
    else menorLista lista.Tail menor index

let getShorterWay rutas =
    let rutasTemp = List.map (fun ruta -> List.map(fun (vertice:(char*int)) -> snd(vertice)) ruta) rutas
    let intList = List.map (fun ruta -> List.sum ruta) rutasTemp
    printf "\nDistancia de las rutas: %A" intList
    rutas[rutas.Length - (menorLista intList 1000000 0) - 1]
    
    
let miembro e (lista: ('a * 'b) list) =
    lista
    |> List.map (fun (x,_) -> x = e)
    |> List.reduce (fun x y -> x || y)

let rec vecinos nodo grafo =
    match grafo with
    | [] -> []
    | ((head, peso)::adyacencias)::tail when nodo = head -> adyacencias
    | _::tail -> vecinos nodo tail 

let extender (ruta: _ list) grafo =
    (vecinos (fst ruta.Head) grafo)
    |> List.map (fun (x,peso) -> if (miembro x ruta) then [] else (x,peso)::ruta )
    |> List.filter (fun x -> x <> [])

let rec prof_aux fin (rutas: _ list list) grafo =
    if rutas = [] then
        []
    elif fin = (fst rutas.Head.Head) then
        List.append
            ([List.rev rutas.Head])
            (prof_aux fin rutas.Tail grafo)
    else
        prof_aux fin (List.append (extender rutas.Head grafo) rutas.Tail) grafo
        
let prof ini fin grafo =
    printf "Rutas: %A" (prof_aux fin [[(ini,0)]] grafo)
    printf "\nRuta de menor distancia: %A" (getShorterWay (prof_aux fin [[(ini,0)]] grafo))
    
prof 'i' 'f' grafo
