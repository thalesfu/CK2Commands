package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1250] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1250][1250] = People_1250
	HistoryPeopleMap[1250][41250] = People_41250
	HistoryPeopleMap[1250][71250] = People_71250
	HistoryPeopleMap[1250][91250] = People_91250
	HistoryPeopleMap[1250][161250] = People_161250
	HistoryPeopleMap[1250][191250] = People_191250
	HistoryPeopleMap[1250][251250] = People_251250
	HistoryPeopleMap[1250][261250] = People_261250
}
