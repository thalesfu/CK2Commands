package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[389] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[389][190389] = People_190389
	HistoryPeopleMap[389][200389] = People_200389
	HistoryPeopleMap[389][260389] = People_260389
}
