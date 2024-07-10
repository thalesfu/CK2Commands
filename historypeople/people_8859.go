package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8859] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8859][188859] = People_188859
}
