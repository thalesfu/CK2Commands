package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6080] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6080][96080] = People_96080
	HistoryPeopleMap[6080][146080] = People_146080
	HistoryPeopleMap[6080][166080] = People_166080
	HistoryPeopleMap[6080][256080] = People_256080
}
