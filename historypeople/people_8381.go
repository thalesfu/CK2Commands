package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8381] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8381][138381] = People_138381
	HistoryPeopleMap[8381][168381] = People_168381
	HistoryPeopleMap[8381][188381] = People_188381
}
