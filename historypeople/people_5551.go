package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5551] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5551][205551] = People_205551
	HistoryPeopleMap[5551][455551] = People_455551
	HistoryPeopleMap[5551][465551] = People_465551
	HistoryPeopleMap[5551][475551] = People_475551
}
