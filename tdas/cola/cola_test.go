package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
}

func TestInvarianteCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	elementos := []int{10, 20, 30, 40, 50}
	for i := 0; i < len(elementos); i++ {
		cola.Encolar(elementos[i])
	}
	for i := 0; i < len(elementos); i++ {
		require.Equal(t, elementos[i], cola.VerPrimero())
		require.Equal(t, elementos[i], cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestColaNoVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 20; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 7; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.Equal(t, 7, cola.VerPrimero())
	require.False(t, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	n := 1000
	for i := 0; i < n; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < n; i++ {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func testColaGenerica[T any](t *testing.T, val1, val2 T) {
	cola := TDACola.CrearColaEnlazada[T]()
	cola.Encolar(val1)
	cola.Encolar(val2)
	require.Equal(t, val1, cola.VerPrimero())
	require.Equal(t, val1, cola.Desencolar())
	require.Equal(t, val2, cola.VerPrimero())
	require.Equal(t, val2, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestColaTipos(t *testing.T) {
	testColaGenerica(t, "san", "lorenzo")
	testColaGenerica(t, 1.5, 2.7)
	testColaGenerica(t, true, false)
	testColaGenerica(t, 1, 2)
}

func TestColaVaciada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 10; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })

	cola.Encolar(42)
	require.False(t, cola.EstaVacia())
	require.Equal(t, 42, cola.VerPrimero())
	require.Equal(t, 42, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestEncolarDesencolarAlternado(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	require.Equal(t, 1, cola.Desencolar())
	cola.Encolar(3)
	require.Equal(t, 2, cola.VerPrimero())
	cola.Encolar(4)
	require.Equal(t, 2, cola.Desencolar())
	require.Equal(t, 3, cola.VerPrimero())
	require.Equal(t, 3, cola.Desencolar())
	require.Equal(t, 4, cola.VerPrimero())
	require.Equal(t, 4, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}
