package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5702] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5702][125702] = People_125702
	HistoryPeopleMap[5702][145702] = People_145702
	HistoryPeopleMap[5702][205702] = People_205702
}
