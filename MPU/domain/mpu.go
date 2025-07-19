package domain


type Vector3D struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type Mpu6050 struct {
	Aceleracion Vector3D `json:"aceleracion"`
	Giroscopio  Vector3D `json:"giroscopio"`
}

type Mpu struct {
	Id     int    `json:"id"`
	Mpu6050 Mpu6050 `json:"mpu6050"`
	Pasos  int    `json:"pasos"`
	NivelActividad string `json:"nivel_actividad"`
	Fecha  string `json:"fecha"`
}
