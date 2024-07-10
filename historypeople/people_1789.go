package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1789] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1789][31789] = People_31789
	HistoryPeopleMap[1789][191789] = People_191789
}
