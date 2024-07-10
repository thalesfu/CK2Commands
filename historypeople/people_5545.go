package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5545] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5545][455545] = People_455545
}
