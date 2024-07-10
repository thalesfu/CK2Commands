package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5263] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5263][145263] = People_145263
	HistoryPeopleMap[5263][205263] = People_205263
}
