/* Implementación en Go de un árbol de búsqueda binario rojinegro.
   Creado por L. Antonio Hidalgo Rodríguez, 201281845.
   Tarea programada 1 del curso de Lenguajes de Programación, grupo 1.
*/

package redBlackTree

import (
      "fmt"
)

// Interfaz para el árbol rojinegro. Cualquier implementación del árbol
// requiere estos métodos (según se indican en la tarea programada)
type rbTreer interface {
      NewTree()
      PrettyPrint()
      Clear()
      Insert()
      Delete()
      Find()
      InOrderPrint()
}

// Se define el color como rojo o negro mediante una variable booleana.
// Se usa true como negro y false como rojo.
type Color bool

// Función para desplegar el color como un string
func (color Color) String() string {
      switch color {
      case true:
            return "Negro"
      default:
            return "Rojo"
      }
}

// El nodo indica sus hijos (izquierdo y derecho) y su padre.
type Node struct {
      // Se utiliza value como interface{} para que pueda ser de 
      // cualquier tipo.
      value  interface{}
      color  Color
      left   *Node
      right  *Node
      parent *Node
}

// Getters - como value puede ser más complicado no se indica su getter

func (node *Node) Color() Color {
      return node.Color
}

func (node *Node) Parent() *Node {
      return node.parent
}

// Función para desplegar el valor y el color del nodo mediante print
func (node *Node) String() string {
      // Se le da formato al valor y al color. %v despliega la interfaz
      // de value como {sam {12345 67890}} (para nombre y números de teléfono
      // por ejemplo.
      return fmt.Sprint("%v : %s)", node.value, n.Color())
}
