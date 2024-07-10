package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8585] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8585][168585] = People_168585
	HistoryPeopleMap[8585][188585] = People_188585
}
