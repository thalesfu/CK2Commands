package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8400][8400] = People_8400
	HistoryPeopleMap[8400][138400] = People_138400
	HistoryPeopleMap[8400][168400] = People_168400
	HistoryPeopleMap[8400][188400] = People_188400
}
