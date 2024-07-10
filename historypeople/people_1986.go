package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1986] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1986][71986] = People_71986
	HistoryPeopleMap[1986][191986] = People_191986
	HistoryPeopleMap[1986][451986] = People_451986
}
