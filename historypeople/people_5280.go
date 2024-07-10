package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5280] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5280][145280] = People_145280
	HistoryPeopleMap[5280][485280] = People_485280
}
