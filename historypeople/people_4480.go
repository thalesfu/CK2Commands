package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4480] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4480][34480] = People_34480
	HistoryPeopleMap[4480][74480] = People_74480
}
