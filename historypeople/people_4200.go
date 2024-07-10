package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4200] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4200][34200] = People_34200
	HistoryPeopleMap[4200][144200] = People_144200
	HistoryPeopleMap[4200][194200] = People_194200
}
