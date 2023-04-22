type producto =
    struct
        val mutable nombre: string
        val mutable descripcion: string
        val mutable montoVenta: float
        new(x, y, z) = {nombre = x; descripcion = y; montoVenta = z }
    end

//1)	Haciendo uso de la función filter, implemente una función que a partir de una lista de cadenas
//representando frases de parámetro, filtre aquellas que contengan una palabra que el usuario indique
//en otro argumento (palabra completa contenida).

let rec miembro1 e l =
    if l = [] then
        false
    else if e = l.Head then
        true
    else 
        miembro1 e l.Tail
            
let palabras (completa:string list) (contenida:string)=
    lista
    |> List.filter (fun str ->
        let lista1 = str.Split ' ' |> List.ofArray
        miembro1 contenida completa)
       
// 2)	Construya una función que se llame general-sec A B. Esta función genera una secuencia de números
// en una lista de A hasta B, donde  A y B son números enteros.

let rec generar_sec A B =
    if A > B then []
    else A :: generar_sec (A + 1) B
    
//3)	Migrar el ejercicio realizado en Go sobre facturas con listas de productos haciendo énfasis
//en la implementación de las funciones para calcular el monto a pagar por la factura completa
//y el monto a pagar por concepto del 13% de impuesto de venta para aquellos productos que deban pagarlo según criterio de selección.

let factura = [new producto("tarjeta madre", "Asus", 54200.0); new producto("mouse", "alámbrico", 15000.0); new producto("teclado", "gamer con luces", 30000.0); new producto("memoria ssd", "2 gb", 65000.0); new producto("cable video", "display port 1m", 18000.0)]

let rangoPagoImpuestos = 20000.0
let porcentajeImpuesto = 0.13

let calcularImpuestoFactura = factura |> List.filter(fun p -> p.montoVenta > rangoPagoImpuestos) |> List.map(fun p -> p.montoVenta * porcentajeImpuesto) |> List.sum
    
let calcularMontoFactura = factura |> List.map(fun p -> p.montoVenta) |> List.sum

let frases = ["otra cosa";"la rueda";"mejor la otra";"no se cual pala"]

printfn "%A" (palabras frases "la")
printfn "%A" (generar_sec 1 8)

printfn "Monto de total de impuestos (IVA) %f" (calcularImpuestoFactura)
printfn "Monto total de la factura (sin IVA) %f" (calcularMontoFactura)
