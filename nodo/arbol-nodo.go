package main

import "fmt"

type Nodo struct {
    dato interface{} // interface{} permite cualquier tipo de dato
    izdo *Nodo
    dcho *Nodo
}

func NewNodo(valor interface{}) *Nodo {
    return &Nodo{
        dato: valor,
        izdo: nil,
        dcho: nil,
    }
}

func NewNodoConRamas(ramaIzdo *Nodo, valor interface{}, ramaDcho *Nodo) *Nodo {
    return &Nodo{
        dato: valor,
        izdo: ramaIzdo,
        dcho: ramaDcho,
    }
}

func (n *Nodo) ValorNodo() interface{} {
    return n.dato
}

func (n *Nodo) GetSubarbolIzdo() *Nodo {
    return n.izdo
}

func (n *Nodo) GetSubarbolDcho() *Nodo {
    return n.dcho
}

func (n *Nodo) NuevoValor(d interface{}) {
    n.dato = d
}

func (n *Nodo) SetRamaIzdo(ramaIzdo *Nodo) {
    n.izdo = ramaIzdo
}

func (n *Nodo) SetRamaDcho(ramaDcho *Nodo) {
    n.dcho = ramaDcho
}

func main() {
    n := NewNodo(1)
    fmt.Println(n.ValorNodo()) // debería imprimir "1"
}
