package main

import "fmt"

type Nodo struct {
    dato int
    izdo *Nodo
    dcho *Nodo
}

func NewNodo(valor int) *Nodo {
    return &Nodo{
        dato: valor,
        izdo: nil,
        dcho: nil,
    }
}

func NewNodoConRamas(ramaIzdo *Nodo, valor int, ramaDcho *Nodo) *Nodo {
    return &Nodo{
        dato: valor,
        izdo: ramaIzdo,
        dcho: ramaDcho,
    }
}

func (n *Nodo) ValorNodo() int {
    return n.dato
}

func (n *Nodo) GetSubarbolIzdo() *Nodo {
    return n.izdo
}

func (n *Nodo) GetSubarbolDcho() *Nodo {
    return n.dcho
}

func (n *Nodo) NuevoValor(d int) {
    n.dato = d
}

func (n *Nodo) SetRamaIzdo(izdo *Nodo) {
    n.izdo = izdo
}

func (n *Nodo) SetRamaDcho(dcho *Nodo) {
    n.dcho = dcho
}

type Arbol struct {
    raiz *Nodo
}

func NewArbol() *Arbol {
    return &Arbol{}
}

func (a *Arbol) existe(busqueda int) bool {
    return existe(a.raiz, busqueda)
}

func existe(n *Nodo, busqueda int) bool {
    if n == nil {
        return false
    }
    if n.ValorNodo() == busqueda {
        return true
    } else if busqueda < n.ValorNodo() {
        return existe(n.GetSubarbolIzdo(), busqueda)
    } else {
        return existe(n.GetSubarbolDcho(), busqueda)
    }
}

func (a *Arbol) Insertar(dato int) {
    if a.raiz == nil {
        a.raiz = NewNodo(dato)
    } else {
        insertar(a.raiz, dato)
    }
}

func insertar(padre *Nodo, dato int) {
    if dato > padre.ValorNodo() {
        if padre.GetSubarbolDcho() == nil {
            padre.SetRamaDcho(NewNodo(dato))
        } else {
            insertar(padre.GetSubarbolDcho(), dato)
        }
    } else {
        if padre.GetSubarbolIzdo() == nil {
            padre.SetRamaIzdo(NewNodo(dato))
        } else {
            insertar(padre.GetSubarbolIzdo(), dato)
        }
    }
}

func main() {
    arbol := NewArbol()
    arbol.Insertar(5)
    arbol.Insertar(3)
    arbol.Insertar(7)
    arbol.Insertar(2)
    arbol.Insertar(4)
    arbol.Insertar(6)
    arbol.Insertar(8)
    fmt.Println(arbol.existe(5))
    fmt.Println(arbol.existe(3))
    fmt.Println(arbol.existe(123455))
    fmt.Println("fin de la Ejecucion")
}
