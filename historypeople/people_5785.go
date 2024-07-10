package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5785] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5785][205785] = People_205785
	HistoryPeopleMap[5785][455785] = People_455785
}
