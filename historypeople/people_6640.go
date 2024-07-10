package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6640] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6640][6640] = People_6640
	HistoryPeopleMap[6640][106640] = People_106640
	HistoryPeopleMap[6640][166640] = People_166640
	HistoryPeopleMap[6640][206640] = People_206640
}
