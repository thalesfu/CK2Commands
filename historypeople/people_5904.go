package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5904] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5904][125904] = People_125904
	HistoryPeopleMap[5904][145904] = People_145904
}
