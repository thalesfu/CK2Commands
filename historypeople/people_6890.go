package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6890] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6890][6890] = People_6890
	HistoryPeopleMap[6890][166890] = People_166890
}
