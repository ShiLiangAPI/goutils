package snowflake

import (
	"sync"
)

const (
	workerIdBits     = 5
	dataCenterIdBits = 5
	sequenceBits     = 12

	workerIdShift      = sequenceBits
	dataCenterIdShift  = sequenceBits + workerIdBits
	timestampLeftShift = sequenceBits + workerIdBits + dataCenterIdBits
	sequenceMask       = int64(-1) ^ (int64(-1) << sequenceBits)

	//twepoch = int64(1288834974657)
	twepoch = int64(1597075200000)
)

type IDGenerator struct {
	workerId     int64
	dataCenterId int64
	sequence     int64
	timestamp    int64

	lock *sync.Mutex
}

var flake *IDGenerator

func GetFlake(WorkerId, DataCenterId int64) *IDGenerator {
	if flake == nil {
		flake = NewIDGenerator(WorkerId, DataCenterId)
	}

	return flake
}

//fun func main() {
//	GetFlake(1, 1).GetNextID()
//}
