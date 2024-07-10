package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4222] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4222][34222] = People_34222
	HistoryPeopleMap[4222][144222] = People_144222
	HistoryPeopleMap[4222][194222] = People_194222
}
