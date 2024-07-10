package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1925] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1925][71925] = People_71925
	HistoryPeopleMap[1925][161925] = People_161925
	HistoryPeopleMap[1925][191925] = People_191925
}
