class Nodo:
    def __init__(self, valor, izdo=None, dcho=None):
        self.dato = valor
        self.izdo = izdo
        self.dcho = dcho

    def valorNodo(self):
        return self.dato

    def GetSubarbolIzdo(self):
        return self.izdo

    def GetSubarbolDcho(self):
        return self.dcho

    def nuevoValor(self, d):
        self.dato = d

    def SetRamaIzdo(self, n):
        self.izdo = n

    def SetRamaDcho(self, n):
        self.dcho = n


class Arbol:
    def __init__(self):
        self.raiz = None

    def existe(self, busqueda):
        return self._existe(self.raiz, busqueda)

    def _existe(self, n, busqueda):
        if n is None:
            return False
        if n.valorNodo() == busqueda:
            return True
        elif busqueda < n.valorNodo():
            return self._existe(n.GetSubarbolIzdo(), busqueda)
        else:
            return self._existe(n.GetSubarbolDcho(), busqueda)

    def insertar(self, dato):
        if self.raiz is None:
            self.raiz = Nodo(dato)
        else:
            self._insertar(self.raiz, dato)

    def _insertar(self, padre, dato):
        if dato > padre.valorNodo():
            if padre.GetSubarbolDcho() is None:
                padre.SetRamaDcho(Nodo(dato))
            else:
                self._insertar(padre.GetSubarbolDcho(), dato)
        else:
            if padre.GetSubarbolIzdo() is None:
                padre.SetRamaIzdo(Nodo(dato))
            else:
                self._insertar(padre.GetSubarbolIzdo(), dato)


if __name__ == "__main__":
    arbol = Arbol()
    arbol.insertar(5)
    arbol.insertar(3)
    arbol.insertar(7)
    arbol.insertar(2)
    arbol.insertar(4)
    arbol.insertar(6)
    arbol.insertar(8)
    print(arbol.existe(5))
    print(arbol.existe(3))
    print(arbol.existe(123455))
    print("fin de la Ejecucion")
