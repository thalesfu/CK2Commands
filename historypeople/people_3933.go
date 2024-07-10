package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3933] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3933][3933] = People_3933
	HistoryPeopleMap[3933][73933] = People_73933
}
