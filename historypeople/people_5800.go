package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5800][125800] = People_125800
	HistoryPeopleMap[5800][145800] = People_145800
	HistoryPeopleMap[5800][205800] = People_205800
	HistoryPeopleMap[5800][465800] = People_465800
}
