package snowflake

import (
	"errors"
	"sync"
)

type IDGenerator struct {
	workerId     int64
	dataCenterId int64
	sequence     int64
	timestamp    int64

	lock *sync.Mutex
}

func NewFlake(WorkerId, DataCenterId int64) *IDGenerator {
	if flake == nil {
		flake = NewIDGenerator(WorkerId, DataCenterId)
	}

	return flake
}

func NewIDGenerator(workerId, dataCenterId int64) *IDGenerator {
	if workerId > ((1<<workerIdBits)-1) || workerId < 0 {
		panic(errors.New("worker id error"))
	}

	if dataCenterId > ((1<<dataCenterIdBits)-1) || dataCenterId < 0 {
		panic(errors.New("data center id error"))
	}

	return &IDGenerator{
		workerId:     workerId,
		dataCenterId: dataCenterId,
		lock:         &sync.Mutex{},
	}
}
