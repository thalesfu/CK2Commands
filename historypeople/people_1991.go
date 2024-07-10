package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1991] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1991][71991] = People_71991
	HistoryPeopleMap[1991][191991] = People_191991
	HistoryPeopleMap[1991][451991] = People_451991
}
