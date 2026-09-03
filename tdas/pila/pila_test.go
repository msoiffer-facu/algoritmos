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

func TestPilaStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("san")
	pila.Apilar("lorenzo")
	require.Equal(t, "lorenzo", pila.VerTope())
	require.Equal(t, "lorenzo", pila.Desapilar())
	require.Equal(t, "san", pila.VerTope())
	require.Equal(t, "san", pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaFloats(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	pila.Apilar(1.5)
	pila.Apilar(2.7)
	require.Equal(t, 2.7, pila.VerTope())
	require.Equal(t, 2.7, pila.Desapilar())
	require.Equal(t, 1.5, pila.VerTope())
	require.Equal(t, 1.5, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaBooleans(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()
	pila.Apilar(true)
	pila.Apilar(false)
	require.Equal(t, false, pila.VerTope())
	require.Equal(t, false, pila.Desapilar())
	require.Equal(t, true, pila.VerTope())
	require.Equal(t, true, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestPilaEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	require.Equal(t, 3, pila.VerTope())
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 2, pila.VerTope())
	require.Equal(t, 2, pila.Desapilar())
	require.Equal(t, 1, pila.VerTope())
	require.Equal(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

