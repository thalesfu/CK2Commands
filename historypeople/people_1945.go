package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1945] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1945][71945] = People_71945
	HistoryPeopleMap[1945][161945] = People_161945
	HistoryPeopleMap[1945][191945] = People_191945
	HistoryPeopleMap[1945][451945] = People_451945
}
