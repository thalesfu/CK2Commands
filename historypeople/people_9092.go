package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9092] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9092][159092] = People_159092
	HistoryPeopleMap[9092][189092] = People_189092
}
