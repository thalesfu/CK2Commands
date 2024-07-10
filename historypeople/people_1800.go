package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1800] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1800][31800] = People_31800
	HistoryPeopleMap[1800][71800] = People_71800
	HistoryPeopleMap[1800][191800] = People_191800
	HistoryPeopleMap[1800][451800] = People_451800
	HistoryPeopleMap[1800][461800] = People_461800
}
