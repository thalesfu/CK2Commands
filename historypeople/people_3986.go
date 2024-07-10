package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3986] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3986][73986] = People_73986
}
