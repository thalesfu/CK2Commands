package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1762] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1762][1762] = People_1762
	HistoryPeopleMap[1762][31762] = People_31762
	HistoryPeopleMap[1762][71762] = People_71762
	HistoryPeopleMap[1762][161762] = People_161762
	HistoryPeopleMap[1762][191762] = People_191762
}
