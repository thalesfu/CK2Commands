package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[256] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[256][20256] = People_20256
	HistoryPeopleMap[256][40256] = People_40256
	HistoryPeopleMap[256][70256] = People_70256
	HistoryPeopleMap[256][160256] = People_160256
	HistoryPeopleMap[256][170256] = People_170256
	HistoryPeopleMap[256][180256] = People_180256
	HistoryPeopleMap[256][190256] = People_190256
	HistoryPeopleMap[256][200256] = People_200256
	HistoryPeopleMap[256][260256] = People_260256
}
