package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1909] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1909][71909] = People_71909
	HistoryPeopleMap[1909][191909] = People_191909
	HistoryPeopleMap[1909][461909] = People_461909
}
