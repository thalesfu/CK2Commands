package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[890] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[890][20890] = People_20890
	HistoryPeopleMap[890][30890] = People_30890
	HistoryPeopleMap[890][70890] = People_70890
	HistoryPeopleMap[890][260890] = People_260890
}
