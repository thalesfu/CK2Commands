package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8765] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8765][168765] = People_168765
	HistoryPeopleMap[8765][188765] = People_188765
	HistoryPeopleMap[8765][248765] = People_248765
}
