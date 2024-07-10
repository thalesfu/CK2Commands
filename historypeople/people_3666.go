package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3666] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3666][33666] = People_33666
	HistoryPeopleMap[3666][73666] = People_73666
	HistoryPeopleMap[3666][83666] = People_83666
	HistoryPeopleMap[3666][93666] = People_93666
}
