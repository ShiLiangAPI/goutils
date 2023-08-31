package snowflake

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

var flake *IDGenerator

func GetFlake() *IDGenerator {
	if flake == nil {
		panic("Please run the InitFlake method first")
	}

	return flake
}

//fun func main() {
//	GetFlake(1, 1).GetNextID()
//}
