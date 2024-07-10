package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2800][32800] = People_32800
	HistoryPeopleMap[2800][72800] = People_72800
	HistoryPeopleMap[2800][142800] = People_142800
	HistoryPeopleMap[2800][212800] = People_212800
	HistoryPeopleMap[2800][262800] = People_262800
	HistoryPeopleMap[2800][472800] = People_472800
}
