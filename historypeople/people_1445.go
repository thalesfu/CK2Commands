package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1445] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1445][31445] = People_31445
	HistoryPeopleMap[1445][71445] = People_71445
	HistoryPeopleMap[1445][91445] = People_91445
	HistoryPeopleMap[1445][191445] = People_191445
}
