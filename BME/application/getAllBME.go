package application

import (
	"context"
	"time"
)

func (uc *GetAllBME_UC) Run(ctx context.Context) ([]domain.Bme, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	resultChan := make(chan []domain.Bme)
	errChan := make(chan error)

	go func() {
		data, err := uc.db.FindAll()
		if err != nil {
			errChan <- err
			return
		}
		resultChan <- data
	}()

	select {
	case res := <-resultChan:
		return res, nil

	case err := <-errChan:
		return nil, err

	case <-ctx.Done():
		return nil, ctx.Err() // timeout exceeded
	}
}
