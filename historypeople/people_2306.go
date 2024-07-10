package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2306] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2306][12306] = People_12306
	HistoryPeopleMap[2306][32306] = People_32306
	HistoryPeopleMap[2306][72306] = People_72306
	HistoryPeopleMap[2306][82306] = People_82306
	HistoryPeopleMap[2306][142306] = People_142306
	HistoryPeopleMap[2306][242306] = People_242306
	HistoryPeopleMap[2306][252306] = People_252306
	HistoryPeopleMap[2306][462306] = People_462306
	HistoryPeopleMap[2306][472306] = People_472306
}
