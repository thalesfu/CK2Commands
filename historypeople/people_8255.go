package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[8255] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[8255][168255] = People_168255
	HistoryPeopleMap[8255][188255] = People_188255
}
