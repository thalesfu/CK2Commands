package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5665] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5665][145665] = People_145665
	HistoryPeopleMap[5665][205665] = People_205665
}
