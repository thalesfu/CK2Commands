package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5882] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5882][145882] = People_145882
	HistoryPeopleMap[5882][205882] = People_205882
}
