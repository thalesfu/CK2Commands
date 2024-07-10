package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[729] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[729][729] = People_729
	HistoryPeopleMap[729][20729] = People_20729
	HistoryPeopleMap[729][30729] = People_30729
	HistoryPeopleMap[729][70729] = People_70729
	HistoryPeopleMap[729][260729] = People_260729
	HistoryPeopleMap[729][450729] = People_450729
}
