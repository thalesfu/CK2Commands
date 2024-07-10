package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[2700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[2700][32700] = People_32700
	HistoryPeopleMap[2700][72700] = People_72700
	HistoryPeopleMap[2700][142700] = People_142700
	HistoryPeopleMap[2700][212700] = People_212700
	HistoryPeopleMap[2700][222700] = People_222700
	HistoryPeopleMap[2700][232700] = People_232700
	HistoryPeopleMap[2700][262700] = People_262700
	HistoryPeopleMap[2700][462700] = People_462700
}
