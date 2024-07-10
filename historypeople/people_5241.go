package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5241] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5241][125241] = People_125241
	HistoryPeopleMap[5241][145241] = People_145241
	HistoryPeopleMap[5241][155241] = People_155241
	HistoryPeopleMap[5241][205241] = People_205241
}
