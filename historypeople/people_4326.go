package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4326] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4326][34326] = People_34326
	HistoryPeopleMap[4326][194326] = People_194326
}
