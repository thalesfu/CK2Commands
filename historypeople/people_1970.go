package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1970] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1970][71970] = People_71970
	HistoryPeopleMap[1970][161970] = People_161970
	HistoryPeopleMap[1970][191970] = People_191970
	HistoryPeopleMap[1970][451970] = People_451970
}
