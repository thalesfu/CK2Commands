package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2099] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2099][12099] = People_12099
	HistoryPeopleMap[2099][72099] = People_72099
	HistoryPeopleMap[2099][252099] = People_252099
}
