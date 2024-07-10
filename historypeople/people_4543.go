package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4543] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4543][34543] = People_34543
	HistoryPeopleMap[4543][74543] = People_74543
	HistoryPeopleMap[4543][144543] = People_144543
}
