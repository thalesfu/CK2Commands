package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4001] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4001][34001] = People_34001
	HistoryPeopleMap[4001][54001] = People_54001
	HistoryPeopleMap[4001][74001] = People_74001
	HistoryPeopleMap[4001][144001] = People_144001
	HistoryPeopleMap[4001][154001] = People_154001
	HistoryPeopleMap[4001][184001] = People_184001
	HistoryPeopleMap[4001][214001] = People_214001
	HistoryPeopleMap[4001][224001] = People_224001
	HistoryPeopleMap[4001][244001] = People_244001
}
