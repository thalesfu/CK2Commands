package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[9123] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[9123][159123] = People_159123
	HistoryPeopleMap[9123][189123] = People_189123
}
