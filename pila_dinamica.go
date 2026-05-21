package pila

const (
	capacidadInicial = 10
	cantCrecimiento  = 2
	cantReduccion    = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

// Pre Condicion: -
// Post Condicion: Devuelve una Pila Dinamica.
func CrearPilaDinamica[T any]() Pila[T] {
	pilaAux := new(pilaDinamica[T])
	pilaAux.datos = make([]T, capacidadInicial)
	pilaAux.cantidad = 0
	return pilaAux
}

// EstaVacia devuelve verdadero si la pila no tiene elementos apilados, false en caso contrario.
func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope.
// Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	return p.datos[p.cantidad-1]
}

// Pre Condicion: "tam" es el nuevo tamaño que desea que tenga su pila.
// Post Condicion: Redimensiona el vector de pila, manteniendo los datos preexistentes si corresponde.
func (p *pilaDinamica[T]) redimensionar(tam int) {
	nuevoVector := make([]T, tam)
	copy(nuevoVector, p.datos[:p.cantidad])
	p.datos = nuevoVector
}

// Apilar agrega un nuevo elemento a la pila.
func (p *pilaDinamica[T]) Apilar(valor T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(cap(p.datos) * cantCrecimiento)
	}
	p.datos[p.cantidad] = valor
	p.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y
// se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}

	valorTope := p.datos[p.cantidad-1]
	p.cantidad--

	if cap(p.datos) > capacidadInicial && p.cantidad <= cap(p.datos)/cantReduccion {
		p.redimensionar(cap(p.datos) / cantCrecimiento)
	}

	return valorTope
}
