package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodo[T any](dato T) *nodoCola[T] {
	return &nodoCola[T]{dato: dato, prox: nil}
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(dato T) {
	nuevoNodo := crearNodo(dato)
	if c.EstaVacia() {
		c.primero = nuevoNodo
	} else {
		c.ultimo.prox = nuevoNodo
	}
	c.ultimo = nuevoNodo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := c.primero.dato
	c.primero = c.primero.prox
	if c.primero == nil {
		c.ultimo = nil
	}
	return dato
}
