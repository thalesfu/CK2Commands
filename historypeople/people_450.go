package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[450] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[450][20450] = People_20450
	HistoryPeopleMap[450][30450] = People_30450
	HistoryPeopleMap[450][70450] = People_70450
	HistoryPeopleMap[450][190450] = People_190450
	HistoryPeopleMap[450][260450] = People_260450
}
