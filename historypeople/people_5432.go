package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5432] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5432][455432] = People_455432
}
