package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5990] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5990][125990] = People_125990
	HistoryPeopleMap[5990][145990] = People_145990
	HistoryPeopleMap[5990][205990] = People_205990
}
