package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[4799] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[4799][74799] = People_74799
	HistoryPeopleMap[4799][204799] = People_204799
}
