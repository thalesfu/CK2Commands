package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5888] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5888][145888] = People_145888
	HistoryPeopleMap[5888][205888] = People_205888
}
