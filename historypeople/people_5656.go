package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[5656] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[5656][145656] = People_145656
	HistoryPeopleMap[5656][455656] = People_455656
}
