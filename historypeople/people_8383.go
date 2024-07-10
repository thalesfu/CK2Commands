package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8383] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8383][168383] = People_168383
	HistoryPeopleMap[8383][188383] = People_188383
}
