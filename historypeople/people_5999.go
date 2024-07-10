package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5999] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5999][5999] = People_5999
	HistoryPeopleMap[5999][125999] = People_125999
	HistoryPeopleMap[5999][145999] = People_145999
}
