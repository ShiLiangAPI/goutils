package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

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

func (g *IDGenerator) GetNextID() int64 {

	g.lock.Lock()
	defer g.lock.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if timestamp == g.timestamp {
		g.sequence = (g.sequence + 1) & sequenceMask
		if g.sequence == 0 {
			for timestamp <= g.timestamp {
				timestamp = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		g.sequence = 0
	}

	g.timestamp = timestamp

	id := ((timestamp - twepoch) << timestampLeftShift) |
		(g.dataCenterId << dataCenterIdShift) |
		(g.workerId << workerIdShift) |
		g.sequence

	return id
}

func (g *IDGenerator) GetNextStringID() string {

	g.lock.Lock()
	defer g.lock.Unlock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if timestamp == g.timestamp {
		g.sequence = (g.sequence + 1) & sequenceMask
		if g.sequence == 0 {
			for timestamp <= g.timestamp {
				timestamp = time.Now().UnixNano() / int64(time.Millisecond)
			}
		}
	} else {
		g.sequence = 0
	}

	g.timestamp = timestamp

	id := ((timestamp - twepoch) << timestampLeftShift) |
		(g.dataCenterId << dataCenterIdShift) |
		(g.workerId << workerIdShift) |
		g.sequence

	return strconv.FormatInt(id, 10)
}
