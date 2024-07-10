package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4455] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4455][34455] = People_34455
	HistoryPeopleMap[4455][74455] = People_74455
}
