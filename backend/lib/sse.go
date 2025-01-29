package lib

type ActiveSSE map[int32]chan string

var activeSSE ActiveSSE

func init() {
	activeSSE = make(ActiveSSE)
}

func GetActiveSSE(uid int32) chan string {
	ch, ok := activeSSE[uid]
	if !ok {
		activeSSE[uid] = make(chan string, 10)
		return activeSSE[uid]
	}
	return ch
}

func DeactivateSSE(uid int32) {
	delete(activeSSE, uid)
}
