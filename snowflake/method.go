package snowflake

import (
	"strconv"
	"time"
)

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
