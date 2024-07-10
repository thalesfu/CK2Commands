package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2880] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2880][72880] = People_72880
	HistoryPeopleMap[2880][142880] = People_142880
	HistoryPeopleMap[2880][262880] = People_262880
}
