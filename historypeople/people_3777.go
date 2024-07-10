package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[3777] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[3777][33777] = People_33777
	HistoryPeopleMap[3777][73777] = People_73777
	HistoryPeopleMap[3777][83777] = People_83777
	HistoryPeopleMap[3777][93777] = People_93777
}
