package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9042] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9042][129042] = People_129042
	HistoryPeopleMap[9042][159042] = People_159042
	HistoryPeopleMap[9042][189042] = People_189042
}
