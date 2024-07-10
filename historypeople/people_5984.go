package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5984] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5984][125984] = People_125984
	HistoryPeopleMap[5984][145984] = People_145984
	HistoryPeopleMap[5984][205984] = People_205984
}
