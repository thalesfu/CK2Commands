package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5708] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5708][125708] = People_125708
	HistoryPeopleMap[5708][145708] = People_145708
}
