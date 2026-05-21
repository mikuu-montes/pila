package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const tamañoVolumen = 10000

// Pre Condicion: La pila debe haber sido creada previamente con CrearPilaDinamica()
// Post Condicion: Hace las pruebas que comprueban que una pila vacia se comporte como tal.
func pruebaPilaVacia[T any](t *testing.T, pila TDAPila.Pila[T]) {
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestPilaVacia(t *testing.T) {
	pilaBool := TDAPila.CrearPilaDinamica[bool]()
	pruebaPilaVacia(t, pilaBool)
}

func TestEsLifo(t *testing.T) {
	pilaString := TDAPila.CrearPilaDinamica[string]()
	elementosString := []string{"Hola", "Buenos", "Dias", "Como", "Esta"}

	for _, elemento := range elementosString {
		pilaString.Apilar(elemento)
		require.Equal(t, elemento, pilaString.VerTope())
	}

	for i := len(elementosString) - 1; i >= 0; i-- {
		require.Equal(t, elementosString[i], pilaString.Desapilar())
		if i > 0 {
			require.Equal(t, elementosString[i-1], pilaString.VerTope())
		}
	}
	require.True(t, pilaString.EstaVacia())
}

// Pre condicion: -
// Post Condicion: Rellena el array de forma ascendente (0, 1, 2, ...)
func rellenarArray(array []int) {
	for i := range array {
		array[i] = i
	}
}

func TestVolumenLifo(t *testing.T) {
	pilaInts := TDAPila.CrearPilaDinamica[int]()
	arrayInts := make([]int, tamañoVolumen)
	rellenarArray(arrayInts)

	for _, elemento := range arrayInts {
		pilaInts.Apilar(elemento)
		require.Equal(t, elemento, pilaInts.VerTope())
	}

	for i := len(arrayInts) - 1; i >= 0; i-- {
		require.Equal(t, arrayInts[i], pilaInts.Desapilar())
		if i > 0 {
			require.Equal(t, arrayInts[i-1], pilaInts.VerTope())
		}
	}
	require.True(t, pilaInts.EstaVacia())
}

func TestIntercalado(t *testing.T) {
	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	pilaFloat.Apilar(1.0)
	require.Equal(t, 1.0, pilaFloat.VerTope())

	pilaFloat.Apilar(2.0)
	require.Equal(t, 2.0, pilaFloat.VerTope())

	require.Equal(t, 2.0, pilaFloat.Desapilar())

	pilaFloat.Apilar(3.0)
	require.Equal(t, 3.0, pilaFloat.VerTope())

	require.Equal(t, 3.0, pilaFloat.Desapilar())
	require.Equal(t, 1.0, pilaFloat.Desapilar())

	require.True(t, pilaFloat.EstaVacia())
}
