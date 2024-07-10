package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4400] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4400][4400] = People_4400
	HistoryPeopleMap[4400][34400] = People_34400
	HistoryPeopleMap[4400][74400] = People_74400
	HistoryPeopleMap[4400][194400] = People_194400
}
