package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[350] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[350][350] = People_350
	HistoryPeopleMap[350][30350] = People_30350
	HistoryPeopleMap[350][40350] = People_40350
	HistoryPeopleMap[350][70350] = People_70350
	HistoryPeopleMap[350][170350] = People_170350
	HistoryPeopleMap[350][200350] = People_200350
	HistoryPeopleMap[350][260350] = People_260350
}
