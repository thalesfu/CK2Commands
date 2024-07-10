package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1942] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1942][71942] = People_71942
	HistoryPeopleMap[1942][191942] = People_191942
	HistoryPeopleMap[1942][451942] = People_451942
}
