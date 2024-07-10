package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2345] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2345][12345] = People_12345
	HistoryPeopleMap[2345][72345] = People_72345
	HistoryPeopleMap[2345][142345] = People_142345
	HistoryPeopleMap[2345][242345] = People_242345
	HistoryPeopleMap[2345][252345] = People_252345
}
