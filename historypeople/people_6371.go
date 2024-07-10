package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6371] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6371][166371] = People_166371
}
