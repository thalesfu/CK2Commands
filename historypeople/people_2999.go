package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2999][32999] = People_32999
	HistoryPeopleMap[2999][72999] = People_72999
	HistoryPeopleMap[2999][142999] = People_142999
	HistoryPeopleMap[2999][202999] = People_202999
}
