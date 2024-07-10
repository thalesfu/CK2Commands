package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5923] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5923][125923] = People_125923
	HistoryPeopleMap[5923][145923] = People_145923
}
