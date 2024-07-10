package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6250] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6250][166250] = People_166250
}
