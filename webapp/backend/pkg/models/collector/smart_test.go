package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmartInfo_Capacity(t *testing.T) {
	t.Run("should report nvme capacity", func(t *testing.T) {
		smartInfo := SmartInfo{
			UserCapacity: UserCapacity{
				Bytes: 1234,
			},
			NvmeTotalCapacity: 5678,
		}
		assert.Equal(t, int64(5678), smartInfo.CapacityDetermine())
	})

	t.Run("should report user capacity", func(t *testing.T) {
		smartInfo := SmartInfo{
			UserCapacity: UserCapacity{
				Bytes: 1234,
			},
		}
		assert.Equal(t, int64(1234), smartInfo.CapacityDetermine())
	})

	t.Run("should report 0 for unknown capacities", func(t *testing.T) {
		var smartInfo SmartInfo
		assert.Zero(t, smartInfo.CapacityDetermine())
	})
}

func TestSmartInfo_ModelName(t *testing.T) {
	t.Run("should report SCSI model name", func(t *testing.T) {
		smartInfo := SmartInfo{
			ModelName:     "ATA",
			ScsiModelName: "SCSI",
		}
		assert.Equal(t, string("SCSI"), smartInfo.ModelNameDetermine())
	})

	t.Run("should report ATA model name", func(t *testing.T) {
		smartInfo := SmartInfo{
			ModelName: "ATA",
		}
		assert.Equal(t, string("ATA"), smartInfo.ModelNameDetermine())
	})

	t.Run("should report blank for unknown model name", func(t *testing.T) {
		var smartInfo SmartInfo
		assert.Equal(t, string(""), smartInfo.ModelNameDetermine())
	})
}

func TestSmartInfo_FirmwareVersion(t *testing.T) {
	t.Run("should report SCSI firmware version", func(t *testing.T) {
		smartInfo := SmartInfo{
			FirmwareVersion: "ATA",
			ScsiRevision:    "SCSI",
		}
		assert.Equal(t, string("SCSI"), smartInfo.FirmwareVersionDetermine())
	})

	t.Run("should report ATA firmware version", func(t *testing.T) {
		smartInfo := SmartInfo{
			FirmwareVersion: "ATA",
		}
		assert.Equal(t, string("ATA"), smartInfo.FirmwareVersionDetermine())
	})

	t.Run("should report blank for unknown firmware version", func(t *testing.T) {
		var smartInfo SmartInfo
		assert.Equal(t, string(""), smartInfo.FirmwareVersionDetermine())
	})
}

func TestSmartInfo_PowerCycle(t *testing.T) {
	t.Run("should report SCSI power cycle count", func(t *testing.T) {
		smartInfo := SmartInfo{
			PowerCycleCount: 1234,
			ScsiStartStopCycleCounter: ScsiStartStopCycleCounter{
				AccumulatedStartStopCycles: 5678,
			},
		}
		assert.Equal(t, int64(5678), smartInfo.PowerCycleDetermine())
	})

	t.Run("should report ATA power cycle count", func(t *testing.T) {
		smartInfo := SmartInfo{
			PowerCycleCount: 1234,
		}
		assert.Equal(t, int64(1234), smartInfo.PowerCycleDetermine())
	})

	t.Run("should report 0 for unknown power cycle count", func(t *testing.T) {
		var smartInfo SmartInfo
		assert.Zero(t, smartInfo.CapacityDetermine())
	})
}
