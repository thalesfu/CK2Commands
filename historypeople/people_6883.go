package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6883] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6883][6883] = People_6883
	HistoryPeopleMap[6883][166883] = People_166883
}
