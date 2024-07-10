package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8640] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8640][8640] = People_8640
	HistoryPeopleMap[8640][168640] = People_168640
	HistoryPeopleMap[8640][188640] = People_188640
}
