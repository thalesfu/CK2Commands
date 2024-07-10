package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6899] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6899][6899] = People_6899
	HistoryPeopleMap[6899][166899] = People_166899
}
