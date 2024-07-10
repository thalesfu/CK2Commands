package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5544] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5544][205544] = People_205544
	HistoryPeopleMap[5544][455544] = People_455544
}
