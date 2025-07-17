package domain

// import (
// 	bme "API-VITALVEST/BME/domain"
// 	gsr "API-VITALVEST/GSR/domain"
// 	mlx "API-VITALVEST/MLX/domain"
// 	mpu "API-VITALVEST/MPU/domain"
// )
type Sensors struct {
	BME280 struct {
		Temperatura float64 `json:"temperatura"`
		Presion     float64 `json:"presion"`
		Humedad     float64 `json:"humedad"`
	} `json:"bme280"`

	MPU6050 struct {
		Aceleracion struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"aceleracion"`
		Giroscopio struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
			Z float64 `json:"z"`
		} `json:"giroscopio"`
	} `json:"mpu6050"`

	MLX90614 struct {
		TempObjeto   float64 `json:"temp_objeto"`
		TempAmbiente float64 `json:"temp_ambiente"`
	} `json:"mlx90614"`
	// BME *bme.Bme
	// GSR *gsr.Gsr
	// MLX *mlx.Mlx
	// MPU *mpu.Mpu
}

