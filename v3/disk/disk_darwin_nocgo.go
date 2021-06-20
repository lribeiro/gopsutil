// +build darwin
// +build !cgo

package disk

import (
	"context"

	"github.com/lribeiro/gopsutil/v3/internal/common"
)

func IOCountersWithContext(ctx context.Context, names ...string) (map[string]IOCountersStat, error) {
	return nil, common.ErrNotImplementedError
}
