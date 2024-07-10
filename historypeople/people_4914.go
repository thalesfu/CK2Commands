package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4914] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4914][34914] = People_34914
	HistoryPeopleMap[4914][214914] = People_214914
}
