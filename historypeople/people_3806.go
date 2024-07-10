package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3806] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3806][73806] = People_73806
	HistoryPeopleMap[3806][93806] = People_93806
}
