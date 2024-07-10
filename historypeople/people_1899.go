package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1899] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1899][191899] = People_191899
}
