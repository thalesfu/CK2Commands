package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1990] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1990][71990] = People_71990
	HistoryPeopleMap[1990][191990] = People_191990
	HistoryPeopleMap[1990][451990] = People_451990
}
