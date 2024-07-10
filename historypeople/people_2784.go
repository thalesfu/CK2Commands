package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2784] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2784][32784] = People_32784
	HistoryPeopleMap[2784][72784] = People_72784
	HistoryPeopleMap[2784][142784] = People_142784
	HistoryPeopleMap[2784][472784] = People_472784
}
