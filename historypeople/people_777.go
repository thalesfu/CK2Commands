package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[777] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[777][20777] = People_20777
	HistoryPeopleMap[777][30777] = People_30777
	HistoryPeopleMap[777][70777] = People_70777
	HistoryPeopleMap[777][180777] = People_180777
	HistoryPeopleMap[777][260777] = People_260777
}
