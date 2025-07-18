package control

import (
	"math"
)

// Esta función simula el cálculo de pasos a partir de aceleraciones.
// Puedes reemplazar la lógica con una más avanzada si lo necesitas.
func ConvertirDatosEnPasos(acX, acY, acZ float64) (int, string) {
	// Magnitud del vector de aceleración
	magnitud := math.Sqrt(acX*acX + acY*acY + acZ*acZ)

	// Definir umbrales para contar pasos
	var pasos int
	var nivel string

	if magnitud > 1.2 {
		pasos = 1
		nivel = "Alta"
	} else if magnitud > 0.8 {
		pasos = 0
		nivel = "Media"
	} else {
		pasos = 0
		nivel = "Baja"
	}

	return pasos, nivel
}

