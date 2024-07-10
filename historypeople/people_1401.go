package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1401] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1401][71401] = People_71401
	HistoryPeopleMap[1401][91401] = People_91401
	HistoryPeopleMap[1401][161401] = People_161401
	HistoryPeopleMap[1401][191401] = People_191401
}
