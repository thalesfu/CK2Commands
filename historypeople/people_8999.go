package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8999][188999] = People_188999
}
