package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1776] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1776][1776] = People_1776
	HistoryPeopleMap[1776][31776] = People_31776
	HistoryPeopleMap[1776][161776] = People_161776
}
