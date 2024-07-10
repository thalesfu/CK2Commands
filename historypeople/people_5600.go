package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5600] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5600][125600] = People_125600
	HistoryPeopleMap[5600][145600] = People_145600
	HistoryPeopleMap[5600][205600] = People_205600
	HistoryPeopleMap[5600][465600] = People_465600
}
