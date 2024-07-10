package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[942] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[942][942] = People_942
	HistoryPeopleMap[942][70942] = People_70942
	HistoryPeopleMap[942][260942] = People_260942
}
