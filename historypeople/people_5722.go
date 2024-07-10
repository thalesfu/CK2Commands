package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5722] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5722][125722] = People_125722
	HistoryPeopleMap[5722][145722] = People_145722
}
