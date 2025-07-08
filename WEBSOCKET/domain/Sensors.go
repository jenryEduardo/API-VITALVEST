package domain

import (
	bme "API-VITALVEST/BME/domain"
	gsr "API-VITALVEST/GSR/domain"
	mlx "API-VITALVEST/MLX/domain"
	mpu "API-VITALVEST/MPU/domain"
)
type Sensors struct {
	BME *bme.Bme
	GSR *gsr.Gsr
	MLX *mlx.Mlx
	MPU *mpu.Mpu
}