package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4541] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4541][34541] = People_34541
	HistoryPeopleMap[4541][74541] = People_74541
	HistoryPeopleMap[4541][144541] = People_144541
}
