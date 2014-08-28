/* Implementación en Go de un árbol de búsqueda binario rojinegro.
   Creado por L. Antonio Hidalgo Rodríguez, 201281845.
   Tarea programada 1 del curso de Lenguajes de Programación, grupo 1.
*/

package redBlackTree

type Value interface {}

type Item struct {
      value Value
}

type rbTree interface {
      NewTree()
      PrettyPrint()
      Clear()
      Insert()
      Delete()
      Find()
      InOrderPrint()
}
