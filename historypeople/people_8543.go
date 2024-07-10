package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8543] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8543][168543] = People_168543
	HistoryPeopleMap[8543][188543] = People_188543
}
