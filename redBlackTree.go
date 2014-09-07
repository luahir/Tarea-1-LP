/* Implementación en Go de un árbol de búsqueda binario rojinegro.
   Creado por L. Antonio Hidalgo Rodríguez, 201281845.
   Tarea programada 1 del curso de Lenguajes de Programación, grupo 1.
*/

package redBlackTree

import (
      "fmt"
)

const (
      NEGRO, ROJO Color = true, false
)

// Interfaz para el árbol rojinegro. Cualquier implementación del árbol
// requiere estos métodos (según se indican en la tarea programada).
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
      value   interface{}
      color   Color
      left   *Node
      right  *Node
      parent *Node
}

// Getters y setters - como value puede ser más complicado no se indica su getter.

func (pNode *Node) Color() Color {
      return pNode.color
}

func (pNode *Node) SetColor(pColor Color) {
      pNode.color = pColor
}

func (pNode *Node) Parent() *Node {
      return pNode.parent
}

func (pNode *Node) Left() *Node {
      return pNode.left
}

func (pNode *Node) Right() *Node {
      return pNode.right
}

func (pNode *Node) isLeft() bool {
      return pNode == pNode.parent.left
}

func (pNode *Node) isRight() bool {
      return pNode == pNode.parent.right
}

// Función para desplegar el valor y el color del nodo mediante print.
func (pNode *Node) String() string {
      // Se le da formato al valor y al color. %v despliega la interfaz
      // de value como {sam {12345 67890}} (para nombre y números de teléfono
      // por ejemplo.
      return fmt.Sprintf("(%v : %s)", pNode.value, pNode.Color())
}

// Es necesario que se defina un método de comparación entre los contenidos del árbol rojinegro,
// por lo que se define comparación entre enteros e hileras.
// El comparador genérico permite que el árbol reciba el comparador como un tipo, para que pueda
// comparar los valores dentro de los nodos.
type Cmp func (o1, o2 interface{}) int

func IntCmp(o1, o2 interface{}) int {
      // Se comprueba que tanto o1 como o2 son enteros y por ende comparables.
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

func StringCmp(o1, o2 interface{}) int {
      // Se comprueba que tanto o1 como o2 son hileras y por ende comparables.
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

// La estructura de árbol requiere una raíz (en forma de nodo) y un comparador
// para los diferentes tipos permitidos. Si se quisiera utilizar otro tipo
// sería necesario escribir un comparador.
type RBTree struct {
      root *Node
      cmp   Cmp
      count int
}

// Devuelve la raíz del árbol.
func (tree *RBTree) Root() *Node {
      return tree.root
}
// Se define un nuevo árbol con un comparador y raíz nula.
func NewRBTree(pCmp Cmp) *RBTree {
      tree := &RBTree{root: nil, cmp: pCmp, count: 0}
      return tree
}

// Método de inserción en el árbol, que introduce el nodo con valor pValue.
// Este método solamente inserta el valor como en un árbol de búsqueda binario
// y no se usa directamente.
func (tree *RBTree) InsertValue(pValue interface{}) *Node {
      // Si la raíz no existe, se inserta una nueva y se aumenta el contador
      // de nodos.
      if tree.root == nil {
            node := &Node{value: pValue, color: NEGRO}
            tree.root = node
            tree.count++
            return node
      }

      // Se define el primer padre como la raíz.
      parentNode := tree.root

      for true {
            // Se compara el valor de entrada (int o hilera) para saber qué
            // lado se debe seguir (mayor o menor que la raíz).
            compare := tree.cmp(pValue, parentNode.value)

            switch {
            // Si la comparación coincide, no se inserta el valor y se devuelve
            // el nodo nulo.
            case compare == 0:
                  return nil
            // Si es menor, se va por el lado izquierdo del árbol. Si el nodo actual
            // no tiene hijos, se inserta de inmediato, si no se hace a este nodo el
            // nuevo padre.
            case compare == -1 && parentNode.left == nil:
                  n := &Node{value: pValue, parent: parentNode}
                  parentNode.left = n
                  tree.count++
                  return n
            case compare == -1 && parentNode.left != nil:
                  parentNode = parentNode.left
            // Análogamente para la rama derecha.
            case compare == 1 && parentNode.right == nil:
                  n := &Node{value: pValue, parent: parentNode}
                  parentNode.right = n
                  tree.count++
                  return n
            case compare == 1 && parentNode.right != nil:
                  parentNode = parentNode.right
            }

      }
      panic("Inserción fallida")
}

// Método de inserción en el árbol que revisa las condiciones de árbol rojinegro
// este método utiliza insertValue y devuelve falso si el valor ya se encuentra
// en el árbol. Si no está en el árbol se inserta y devuelve true.
func (tree *RBTree) Insert(pValue interface{}) bool {
      node := tree.InsertValue(pValue)

      // Si el método devuelve nil, significa que no se insertó nada (el valor ya
      // existe en el árbol).
      if node == nil {
            return false
      }

      // Cada nodo nuevo que se inserta debe ser rojo (más fácil revisar las violaciones
      // de las condiciones).
      node.color = ROJO

      for true {
            // Casos como en wikipedia: http://en.wikipedia.org/wiki/Red-black_tree
            switch {
            // Caso 1: N es la nueva raíz del árbol.
            case node.parent == nil:
                  node.color = NEGRO
                  return true
            // Caso 2: el padre de N debe ser negro.
            case node.parent.color == NEGRO:
                  return true
            // Caso 3: tanto padre como tío son rojos, ambos deben repintarse.
            // negro y el abuelo se vuelve rojo.
            case node.parent.parent.left != nil && node.parent.parent.left.color == ROJO:
                  grandpa := node.parent.parent
                  uncle := grandpa.left
                  node.parent.color = NEGRO
                  uncle.color = NEGRO
                  grandpa.color = ROJO
                  node = grandpa
            case node.parent.parent.right != nil && node.parent.parent.right.color == ROJO:
                  grandpa := node.parent.parent
                  uncle := grandpa.right
                  node.parent.color = NEGRO
                  uncle.color = NEGRO
                  grandpa.color = ROJO
                  node = grandpa
            // Caso 4: padre rojo, tío negro.
            case node.isRight() && node.parent.isLeft():
                  tree.rotLeft(node.parent)
                  node = node.left
            case node.isLeft() && node.parent.isRight():
                  tree.rotRight(node.parent)
                  node = node.right
            // Caso 5: padre rojo, tío negro.
            case node.isRight():
                  node.parent.color = NEGRO
                  node.parent.parent.color = ROJO
                  tree.rotLeft(node.parent.parent)
                  return true
            case node.isLeft():
                  node.parent.color = NEGRO
                  node.parent.parent.color = ROJO
                  tree.rotRight(node.parent.parent)
                  return true
            }
      }
      panic("Inserción fallida")
}

// Rotación a la derecha:
/*
      Q             P
    P   C   ->    A   Q
  A   B             B   C
*/
func (tree *RBTree) rotRight(Q *Node) {
      P := Q.left
      Q.left = P.right
      // Si P tiene hijo derecho, se lo pasa a Q.
      if P.right!= nil {
            P.right.parent = Q
      }

      P.parent = Q.parent

      switch {
      // Si el padre de Q era nulo, ahora P es la raíz.
      case Q.parent == nil:
            tree.root = P
      // Si era hijo derecho, ahora P es hijo derecho y
      // viceversa.
      case Q.isLeft():
            Q.parent.left = P
      case Q.isRight():
            Q.parent.right = P
      }
      // El hijo derecho de P es ahora Q y P es el padre de Q
      P.right = Q
      Q.parent = P
}

// Rotación a la izquierda:
/*
      P             Q
    A   Q   ->    P   C
      B   C     A   B
*/
func (tree *RBTree) rotLeft(P *Node) {
      Q := P.right
      P.right = Q.left
      // Si Q tiene hijo izquierdo, se lo pasa a P.
      if Q.left != nil {
            Q.left.parent = P
      }

      Q.parent = P.parent

      switch {
      // Si el padre de P era nulo, ahora Q es la raíz.
      case P.parent == nil:
            tree.root = Q
      // Si era hijo izquierdo, ahora Q es hijo izquierdo y
      // viceversa.
      case P.isLeft():
            P.parent.left = Q
      case P.isRight():
            P.parent.right = Q
      }
      // El hijo izquierdo de Q es ahora P y Q es el padre de P
      Q.left = P
      P.parent = Q
}

// Delete elimina el nodo que coincide con el valor dado. No hace nada
// si la llave no existe
