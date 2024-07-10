package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5989] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5989][125989] = People_125989
	HistoryPeopleMap[5989][145989] = People_145989
}
