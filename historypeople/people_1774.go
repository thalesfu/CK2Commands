package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1774] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1774][1774] = People_1774
	HistoryPeopleMap[1774][31774] = People_31774
	HistoryPeopleMap[1774][71774] = People_71774
	HistoryPeopleMap[1774][161774] = People_161774
}
