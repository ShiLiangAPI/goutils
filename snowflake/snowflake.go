package snowflake

import (
	"errors"
	"sync"
	"time"
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

func NewIDGenerator(workerId, dataCenterId int64) (*IDGenerator, error) {
	if workerId > ((1<<workerIdBits)-1) || workerId < 0 {
		return nil, errors.New("worker id error")
	}

	if dataCenterId > ((1<<dataCenterIdBits)-1) || dataCenterId < 0 {
		return nil, errors.New("data center id error")
	}

	return &IDGenerator{
		workerId:     workerId,
		dataCenterId: dataCenterId,
		lock:         &sync.Mutex{},
	}, nil
}

func GetNextID(workerId, dataCenterId int64) int64 {

	g, _ := NewIDGenerator(workerId, dataCenterId)

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

//func main() {
//	//workId, _ := NewIDGenerator(1, 1)
//	//Id := workId.getNextID()
//	Id := GetNextID()
//	fmt.Println(Id)
//}
