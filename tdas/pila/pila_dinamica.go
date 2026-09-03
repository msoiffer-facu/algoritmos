package pila

const CAPACIDAD_I = 5

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_I)
	pila.cantidad = 0
	return pila
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * 2)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	p.cantidad--
	dato := p.datos[p.cantidad]
	if p.cantidad > 0 && p.cantidad == cap(p.datos)/4 {
		p.redimensionar(cap(p.datos) / 2)
	}
	return dato
}

func (p *pilaDinamica[T]) redimensionar(tam int) {
	nuevosDatos := make([]T, tam)
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}
