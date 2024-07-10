package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1904] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1904][71904] = People_71904
	HistoryPeopleMap[1904][161904] = People_161904
	HistoryPeopleMap[1904][191904] = People_191904
	HistoryPeopleMap[1904][461904] = People_461904
}
