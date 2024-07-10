package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3800][73800] = People_73800
	HistoryPeopleMap[3800][93800] = People_93800
}
