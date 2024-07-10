package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2775] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2775][32775] = People_32775
	HistoryPeopleMap[2775][72775] = People_72775
	HistoryPeopleMap[2775][142775] = People_142775
	HistoryPeopleMap[2775][222775] = People_222775
}
