package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5909] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5909][125909] = People_125909
	HistoryPeopleMap[5909][145909] = People_145909
}
