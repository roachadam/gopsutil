//go:build darwin && !cgo
// +build darwin,!cgo

package host

import (
	"context"

	"github.com/roachadam/gopsutil/v3/internal/common"
)

func SensorsTemperaturesWithContext(ctx context.Context) ([]TemperatureStat, error) {
	return []TemperatureStat{}, common.ErrNotImplementedError
}
