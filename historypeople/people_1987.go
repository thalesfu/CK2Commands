package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1987] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1987][71987] = People_71987
	HistoryPeopleMap[1987][191987] = People_191987
	HistoryPeopleMap[1987][451987] = People_451987
}
