package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5760] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5760][145760] = People_145760
	HistoryPeopleMap[5760][205760] = People_205760
	HistoryPeopleMap[5760][455760] = People_455760
}
