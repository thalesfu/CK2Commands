package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2112] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2112][12112] = People_12112
	HistoryPeopleMap[2112][72112] = People_72112
	HistoryPeopleMap[2112][142112] = People_142112
	HistoryPeopleMap[2112][252112] = People_252112
	HistoryPeopleMap[2112][462112] = People_462112
}
