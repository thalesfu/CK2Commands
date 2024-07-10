package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5780] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5780][145780] = People_145780
	HistoryPeopleMap[5780][205780] = People_205780
}
