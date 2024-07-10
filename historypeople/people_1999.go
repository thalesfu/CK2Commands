package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1999][71999] = People_71999
	HistoryPeopleMap[1999][191999] = People_191999
	HistoryPeopleMap[1999][251999] = People_251999
}
