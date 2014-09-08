/* Implementación en Go de un árbol de búsqueda binario rojinegro.
   Creado por L. Antonio Hidalgo Rodríguez, 201281845.
   Tarea programada 1 del curso de Lenguajes de Programación, grupo 1.
*/

package redBlackTree

import (
      "fmt"
      "strings"
)

const (
      NEGRO, ROJO Color = true, false
)

// Interfaz para el árbol rojinegro. Cualquier implementación del árbol
// requiere estos métodos (según se indican en la tarea programada).
type RBTreer interface {
      NewTree(Cmp)
      PrettyPrint()
      Clear()
      Insert(interface{})
      Delete(interface{})
      Find(interface{})
      String()
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

func (pNode *Node) grandpa() *Node {
      return pNode.parent.parent
}

func (pNode *Node) uncle() *Node {
      if pNode.parent.isRight() {
            return pNode.grandpa().left
      }
      return pNode.grandpa().right
}

// Función para desplegar el valor y el color del nodo mediante print.
func (pNode *Node) String() string {
      // Se le da formato al valor y al color. %v despliega la interfaz
      // de value como {sam {12345 67890}} (para nombre y números de teléfono
      // por ejemplo.
      return fmt.Sprintf("(%v : %s)", pNode.value, pNode.Color())
}

// Función para borrar todos los campos de un nodo, para que pueda eliminarse del 
// árbol al llamar tree.Clear()
func(pNode *Node) clear() {
      pNode.parent = nil
      pNode.right = nil
      pNode.left = nil
      pNode.color = false
      pNode.value = nil
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
func NewTree(pCmp Cmp) *RBTree {
      tree := &RBTree{root: nil, cmp: pCmp, count: 0}
      return tree
}

// Método de inserción en el árbol, que introduce el nodo con valor pValue.
// Este método solamente inserta el valor como en un árbol de búsqueda binario
// y no se usa directamente.
func (tree *RBTree) insertValue(pValue interface{}) *Node {
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
      node := tree.insertValue(pValue)

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
            case node.uncle() != nil && node.uncle().color == ROJO:
                  node.parent.color = NEGRO
                  node.uncle().color = NEGRO
                  node.grandpa().color = ROJO
                  node = node.grandpa()
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


// El iterador recorre todo el árbol y devuelve todos los nodos, en el orden requerido.
type Iterator interface {
      Iterate(*Node)
}

// El iterador in-order es una estructura que tiene un canal de nodos
type InorderIterator struct {
}

// Mediante un canal se recorre el árbol en in-order.
func (iter *InorderIterator) Iterate(node *Node) <-chan *Node {
      // Canal de tipo *Node.
      ch := make(chan *Node)

      // Función de visita, que se necesita solamente para
      // el iterador, en in-order.
      var visit func(*Node)
      visit = func(visitNode *Node) {
            if visitNode.left != nil {
                  visit(visitNode.left)
            }

            ch <- visitNode

            if visitNode.right != nil {
                  visit(visitNode.right)
            }
      }

      // Se llama a una goroutine que se encarga de visitar el resto de los nodos.
      go func() {
            if node != nil {
                  visit(node)
            }
            close(ch)
      }()

      return ch
}

// Determina si la llave pKey se encuentra entre los valores de los nodos del árbol, 
// si lo encuentra devuelve true y si no, false, además del nodo que contiene el valor.
// Find utiliza el iterador para obtener el valor solicitado.
func (tree *RBTree) Find(pKey interface{}) (bool, *Node) {
      // Recorre el árbol, obtiene los valores en un channel
      // y compara el valor de pKey con el de cada nodo
      iter := &InorderIterator{}
      for node := range iter.Iterate(tree.root) {
            if pKey == node.value {
                  return true, node
            }
      }
      return false, nil
}

// FindKey determina si el valor de pKey es parte de los valores en el árbol. Es un
// wrapper que devuelve solamente el primer argumento de Find.
func (tree *RBTree) FindKey(pKey interface{}) bool {
      found,_ := tree.Find(pKey)
      return found
}

// Despliega los elementos del árbol en in-order, para mostrarlos mediante fmt.Print().
// La hilera resultante se obtiene de recorrer el árbol mediante el iterador.
func (tree *RBTree) String() string {
      iter := &InorderIterator{}
      s := "{"
      for node := range iter.Iterate(tree.root) {
            s += fmt.Sprintf("%v ", node)
      }
      s = strings.TrimSpace(s)
      s += "}"
      return s
}

// Clear borra completamente el árbol mediante deleteAll. 
func (tree *RBTree) Clear() {
      deleteAll(tree.root)
      tree.root = nil
}

// deleteAll elimina los nodos recursivamente, mediante un recorrido en postorder
func deleteAll(node *Node) {
      if node != nil {
            deleteAll(node.left)
            deleteAll(node.right)
            node.clear()
            node = nil
      }
}

// Delete elimina el nodo que coincide con el valor dado. No hace nada
// si la llave no existe
func (tree *RBTree) Delete(pKey interface{}) {
      _, nodeToDel := tree.Find(pKey)
      nodeCopy := nodeToDel
      copyColor := nodeCopy.color
}
