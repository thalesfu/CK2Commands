package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4098] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4098][34098] = People_34098
	HistoryPeopleMap[4098][144098] = People_144098
	HistoryPeopleMap[4098][184098] = People_184098
	HistoryPeopleMap[4098][194098] = People_194098
	HistoryPeopleMap[4098][214098] = People_214098
}
