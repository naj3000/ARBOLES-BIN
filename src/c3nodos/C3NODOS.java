/*
 * Click nbfs://nbhost/SystemFileSystem/Templates/Licenses/license-default.txt to change this license
 * Click nbfs://nbhost/SystemFileSystem/Templates/Classes/Main.java to edit this template
 */
package c3nodos;

/**
 *
 * @author USUARIO
 */
public class C3NODOS {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        // TODO code application logic here
        Arbol arbol = new Arbol();
        arbol.insertar(5);
        arbol.insertar(3);
        arbol.insertar(7);
        arbol.insertar(2);
        arbol.insertar(4);
        arbol.insertar(6);
        arbol.insertar(8);
        System.out.println(arbol.existe(5));
        System.out.println(arbol.existe(3));
        System.out.println(arbol.existe(123455));
        System.out.println("fin de la Ejecucion" );

    }
    }
    

