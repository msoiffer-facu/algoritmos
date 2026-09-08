package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaCreada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
}

func TestInvariantePila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	elementos := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(elementos); i++ {
		pila.Apilar(elementos[i])
	}
	for i := len(elementos) - 1; i >= 0; i-- {
		require.Equal(t, elementos[i], pila.VerTope())
		require.Equal(t, elementos[i], pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaNoVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 20; i++ {
		pila.Apilar(i)
	}
	for i := 19; i >= 13; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	require.Equal(t, 12, pila.VerTope())
	require.False(t, pila.EstaVacia())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	n := 1000
	for i := 0; i < n; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := n - 1; i >= 0; i-- {
		require.Equal(t, i, pila.VerTope())
		require.Equal(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

func testPilaGenerica[T any](t *testing.T, val1, val2 T) {
	pila := TDAPila.CrearPilaDinamica[T]()
	pila.Apilar(val1)
	pila.Apilar(val2)
	require.Equal(t, val2, pila.VerTope())
	require.Equal(t, val2, pila.Desapilar())
	require.Equal(t, val1, pila.VerTope())
	require.Equal(t, val1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaTipos(t *testing.T) {
	testPilaGenerica(t, "san", "lorenzo")
	testPilaGenerica(t, 1.5, 2.7)
	testPilaGenerica(t, true, false)
	testPilaGenerica(t, 1, 2)
}

func TestPilaVaciada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10; i++ {
		pila.Apilar(i)
	}
	for i := 9; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(42)
	require.False(t, pila.EstaVacia())
	require.Equal(t, 42, pila.VerTope())
	require.Equal(t, 42, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}
