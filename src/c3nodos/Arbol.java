/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package c3nodos;

/**
 *
 * @author USUARIO
 */
public class Arbol {
 
    private Nodo raiz;

    public Arbol(){

    }

    public boolean existe(int busqueda) {
        return existe(this.raiz, busqueda);
    }

    private boolean existe(Nodo n, int busqueda) {
        if (n == null) {
            return false;
        }
        if ((int)n.valorNodo() == busqueda) {
            return true;
        } else if (busqueda < (int)n.valorNodo()) {
            return existe(n.GetSubarbolIzdo(), busqueda);
        } else {
            return existe(n.GetSubarbolDcho(), busqueda);
        }

    }

    public void insertar(int dato) {
        if (this.raiz == null) {
            this.raiz = new Nodo(dato);
        } else {
            this.insertar(this.raiz, dato);
        }
    }

    private void insertar(Nodo padre, int dato) {
        if (dato > (int)padre.valorNodo()) {
            if (padre.GetSubarbolDcho() == null) {
                padre.SetRamaDcho(new Nodo(dato));
            } else {
                this.insertar(padre.GetSubarbolDcho(), dato);
            }
        } else {
            if (padre.GetSubarbolIzdo() == null) {
                padre.SetRamaIzdo(new Nodo(dato));
            } else {
                this.insertar(padre.GetSubarbolIzdo(), dato);
            }
        }
    }

    

}

