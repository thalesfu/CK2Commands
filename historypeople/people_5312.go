package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5312] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5312][125312] = People_125312
}
