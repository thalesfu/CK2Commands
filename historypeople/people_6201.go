package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[6201] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[6201][6201] = People_6201
	HistoryPeopleMap[6201][146201] = People_146201
	HistoryPeopleMap[6201][166201] = People_166201
}
