package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3890] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3890][73890] = People_73890
}
