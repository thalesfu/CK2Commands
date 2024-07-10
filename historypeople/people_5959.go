package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5959] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5959][125959] = People_125959
	HistoryPeopleMap[5959][145959] = People_145959
}
