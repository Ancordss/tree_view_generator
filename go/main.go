package main

import (
"fmt"
"regexp"
)

// Definir la estructura Nodo que representa un nodo del árbol
type Nodo struct {
	valor string
	izquierda *Nodo
	derecha *Nodo
}

// Función para crear un árbol de expresiones a partir de una expresión dada
func crearArbolExpresiones(expresion string) *Nodo {
// Eliminar espacios en blanco de la expresión
expresion = regexp.MustCompile([\s\(\)]).ReplaceAllString(expresion, "")

//go

// Separar la expresión en una lista de símbolos
simbolos := regexp.MustCompile(`\d+\.?\d*|\(|\)|\+|\-|\*|\/`).FindAllString(expresion, -1)

// Función auxiliar para convertir una lista de símbolos en un árbol de expresiones
var crearSubarbol func([]string, int, int) *Nodo
crearSubarbol = func(simbolos []string, inicio, fin int) *Nodo {
	// Encontrar el índice del operador con la menor precedencia
	nivelOperador := 0
	indiceOperador := -1
	for i := inicio; i <= fin; i++ {
		if simbolos[i] == "(" {
			nivelOperador++
		} else if simbolos[i] == ")" {
			nivelOperador--
		} else if nivelOperador == 0 {
			if simbolos[i] == "+" || simbolos[i] == "-" {
				if indiceOperador == -1 || simbolos[i] == "*" || simbolos[i] == "/" {
					indiceOperador = i
				}
			} else if simbolos[i] == "*" || simbolos[i] == "/" {
				if indiceOperador == -1 {
					indiceOperador = i
				}
			}
		}
	}

	// Si no se encontró ningún operador, entonces la expresión es simplemente un número
	if indiceOperador == -1 {
		return &Nodo{valor: simbolos[inicio]}
	}

	// Crear el nodo actual y recursivamente crear los subárboles de la izquierda y la derecha
	nodo := &Nodo{valor: simbolos[indiceOperador]}
	nodo.izquierda = crearSubarbol(simbolos, inicio, indiceOperador-1)
	nodo.derecha = crearSubarbol(simbolos, indiceOperador+1, fin)
	return nodo
}

// Crear el árbol de expresiones a partir de la lista de símbolos
raiz := crearSubarbol(simbolos, 0, len(simbolos)-1)

return raiz

}

// Función para imprimir el árbol de expresiones en la terminal
func imprimirArbolExpresiones(nodo *Nodo, prefijo string, esIzquierdo *bool) {
	if nodo == nil {
		return
	}
		fmt.Print(prefijo)
	if *esIzquierdo {
		fmt.Print("├───")
	} else {
		fmt.Print("└───")
	}
	fmt.Println(nodo.valor).
	imprimirArbolExpresiones(nodo.izquierda,prefijo+func(esIzquierdo bool) string {
		if esIzquierdo {
			return "| "
		} else {
			return " "
		}
		}	(*esIzquierdo)+"", true)

	imprimirArbolExpresiones(nodo.derecha, prefijo+func(esIzquierdo bool) string {
		if esIzquierdo {
		return " "
		} else {
		return "| "
		}
		}(*esIzquierdo)+"", false)
	}
	
	// Ejemplo de uso del programa
func main() {
	expresion := "3 * (9 - 3 * 4)"
	arbol := crearArbolExpresiones(expresion)
	var esIzquierdo *bool = nil
	imprimirArbolExpresiones(arbol, "", esIzquierdo)
}