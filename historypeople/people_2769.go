package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2769] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2769][32769] = People_32769
	HistoryPeopleMap[2769][72769] = People_72769
	HistoryPeopleMap[2769][142769] = People_142769
	HistoryPeopleMap[2769][232769] = People_232769
}
