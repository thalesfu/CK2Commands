package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4901] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4901][34901] = People_34901
	HistoryPeopleMap[4901][214901] = People_214901
}
