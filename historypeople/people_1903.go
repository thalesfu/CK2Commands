package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1903] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1903][71903] = People_71903
	HistoryPeopleMap[1903][161903] = People_161903
	HistoryPeopleMap[1903][191903] = People_191903
	HistoryPeopleMap[1903][461903] = People_461903
}
