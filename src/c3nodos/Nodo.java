/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Class.java to edit this template
 */
package c3nodos;

/**
 *
 * @author USUARIO
 */
public class Nodo {
   
       protected int dato;
    protected Nodo izdo;
    protected Nodo dcho;

    public Nodo(int valor) {
        dato = valor;
        izdo = dcho = null;
    }

    public Nodo(Nodo ramaIzdo, int valor, Nodo ramaDcho) {
        //    this(dato);
        dato = valor;
        izdo = ramaIzdo;
        dcho = ramaDcho;
    }
// operaciones de acceso

    public Object valorNodo() {
        //    return valor;
        return dato;
    }

    public Nodo GetSubarbolIzdo() {
        return izdo;
    }

    public Nodo GetSubarbolDcho() {
        return dcho;
    }

    public void nuevoValor(int d) {
        dato = d;
    }

    public void SetRamaIzdo(Nodo n) {
        izdo = n;
    }

    public void SetRamaDcho(Nodo n) {
        dcho = n;
    }

}




//siguiente clase







////clase principal



//neo4j