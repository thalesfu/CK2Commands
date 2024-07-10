package historypeople

import "github.com/thalesfu/ck2nebula"

func init() {
	HistoryPeopleMap[1700] = make(map[int]*ck2nebula.People)
	HistoryPeopleMap[1700][1700] = People_1700
	HistoryPeopleMap[1700][31700] = People_31700
	HistoryPeopleMap[1700][71700] = People_71700
	HistoryPeopleMap[1700][131700] = People_131700
	HistoryPeopleMap[1700][161700] = People_161700
	HistoryPeopleMap[1700][191700] = People_191700
	HistoryPeopleMap[1700][221700] = People_221700
	HistoryPeopleMap[1700][461700] = People_461700
	HistoryPeopleMap[1700][471700] = People_471700
}
