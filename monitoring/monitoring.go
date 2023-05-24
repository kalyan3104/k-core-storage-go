package monitoring

import (
	logger "github.com/kalyan3104/k-core-logger-go"
	"github.com/kalyan3104/k-core/core"
	"github.com/kalyan3104/k-core/core/atomic"
)

var log = logger.GetOrCreate("storage")

var cumulatedSizeInBytes atomic.Counter

// MonitorNewCache adds the size in the global cumulated size variable
func MonitorNewCache(tag string, sizeInBytes uint64) {
	cumulatedSizeInBytes.Add(int64(sizeInBytes))
	log.Debug("MonitorNewCache", "name", tag, "capacity", core.ConvertBytes(sizeInBytes), "cumulated", core.ConvertBytes(cumulatedSizeInBytes.GetUint64()))
}
