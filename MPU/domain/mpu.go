package domain


// type Vector3D struct {
// 	X float64 `json:"x"`
// 	Y float64 `json:"y"`
// 	Z float64 `json:"z"`
// }

// type Mpu6050 struct {
// 	Aceleracion Vector3D `json:"aceleracion"`
// 	Giroscopio  Vector3D `json:"giroscopio"`
// }

type Mpu struct {
	Id int
	Pasos  int    `json:"pasos"`
}
