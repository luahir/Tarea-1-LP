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
func (pColor Color) String() string {
      switch pColor {
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

// Getters y setters - como value puede ser más complicado no se indica su getter

func (pNode *Node) Color() Color {
      return pNode.color
}

func (pNode *Node) SetColor(pColor Color) {
      pNode.color = pColor
}

func (pNode *Node) Parent() *Node {
      return pNode.parent
}

// Función para desplegar el valor y el color del nodo mediante print
func (pNode *Node) String() string {
      // Se le da formato al valor y al color. %v despliega la interfaz
      // de value como {sam {12345 67890}} (para nombre y números de teléfono
      // por ejemplo.
      return fmt.Sprint("%v : %s)", pNode.value, pNode.Color())
}

// Es necesario que se defina un método de comparación entre los contenidos del árbol rojinegro,
// por lo que se define comparación entre enteros e hileras.
// El comparador genérico permite que el árbol reciba el comparador como un tipo, para que pueda
// comparar los valores dentro de los nodos.
type Comp func (o1, o2 interface{}) int

func IntComp(o1, o2 interface{}) int {
      int1, int2 := o1.(int), o2.(int)

      switch {
      case int1 > int2:
            return 1
      case int1 < int2:
            return -1
      default:
            return 0
      }
}

func StringComp(o1, o2 string) int {
      st1, st2 := o1.(string), o2.(string)

      switch {
      case st1 > st2:
            return 1
      case st1 < st2:
            return -1
      default:
            return 0
      }
}

