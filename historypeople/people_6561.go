package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6561] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6561][166561] = People_166561
}
