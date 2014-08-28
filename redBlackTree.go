/* Implementación en Go de un árbol de búsqueda binario rojinegro.
   Creado por L. Antonio Hidalgo Rodríguez, 201281845.
   Tarea programada 1 del curso de Lenguajes de Programación, grupo 1.
*/

package redBlackTree

type rbTree interface {
      NewTree()
      PrettyPrint()
      Clear()
      Insert()
      Delete()
      Find()
      InOrderPrint()
}

type Node struct {
      value interface{}
      color  Color
      left   *Node
      right  *Node
      parent *Node
}
