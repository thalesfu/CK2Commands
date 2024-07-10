package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5700][125700] = People_125700
	HistoryPeopleMap[5700][145700] = People_145700
	HistoryPeopleMap[5700][205700] = People_205700
}
