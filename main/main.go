package main

import (
	"CK2Commands/family/bi"
	"CK2Commands/family/fu"
	"CK2Commands/family/lin"
	"CK2Commands/family/lou"
	"CK2Commands/family/pictish"
	"CK2Commands/family/sesende"
	"CK2Commands/family/wu"
	"CK2Commands/family/wuyibu"
	"CK2Commands/family/yin"
	"CK2Commands/family/yuan"
	"CK2Commands/people"
	"fmt"
	"github.com/thalesfu/paradoxtools/CK2/earth"
)

func main() {

	fmt.Println(earth.Germany日耳曼尼亚.KSaxony萨克森().DBremen不来梅().CCelle不来梅().BBremen不来梅().GetTitleName())

	people.MakeFriends(getMyBrothersAndSisters(),
		getMySonsAndDaughters(),
		getMyBigFamiliesFriends(),
		getMyBigLoads(),
		getScotlandFriends(),
		getWuFriends(),
		getLouFriends(),
		getBiFriends(),
		getGermanFriends(),
		getYiwubuFriends(),
	)

	yuan.Hy()
	pictish.Hy()
	fu.Hy()
	lin.Hy()

	people.BuildMyLoad(0,
		fu.Me,
		fu.MyWife,
	)
	people.BuildAmbassador(0,
		2805833,
	)
	people.BuildGeneral(0,
		2841364,
	)
	people.BuildManager(0,
		wu.M方庆_守礼II,
		wu.F埃莱内_守礼II,
	)
	people.BuildSpy(0,
		fu.MyBrother2,
	)
	people.BuildFather(0,
		bi.M君集_守礼II,
		bi.F素节_守礼II,
	)
	people.BuildDoctor(0,
		yuan.M待举_守礼II,
		yuan.F湘兰_守礼II,
	)
	people.BuildRandom(0,
		2853577,
	)

	var curedIllPeople []int
	curedIllPeople = append(curedIllPeople, getMyBrothersAndSisters()...)
	curedIllPeople = append(curedIllPeople, getMySonsAndDaughters()...)
	curedIllPeople = append(curedIllPeople, getMyBigFamiliesFriends()...)
	people.CurePeopleIll(curedIllPeople...)
	var removePeople []int
	removePeople = append(removePeople, getMyBrothersAndSisters()...)
	removePeople = append(removePeople, getMySonsAndDaughters()...)
	removePeople = append(removePeople, getMyBigFamiliesFriends()...)
	people.RemoveBad(removePeople...)

	people.ReligionIsTaoist(0,
		2827199,
		2830449,
		2845776,
		2849855,
		2846929,
		2853147,
	)

	people.CultureIsHanPictish(2804979,
		2847861,
		2848344,
		2864693,
	)

	people.CultureIsHanScottish(2798738,
		2798653,
		2798658,
		2798660,
		2803320,
		2769336,
		2801182,
		2766527,
		2798635,
		2798637,
		2798638,
		2790843,
	)

	people.CultureIsHanWelsh(2793945,
		2803978,
		2803979,
		2803980,
		2803981,
		2803982,
		2803983,
	)

}

func getMySonsAndDaughters() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, fu.MySon1)
	peoples = append(peoples, fu.MySon2)
	peoples = append(peoples, fu.MySon3)
	peoples = append(peoples, fu.MySon4)
	peoples = append(peoples, fu.MySon5)
	peoples = append(peoples, fu.MySon6)
	peoples = append(peoples, fu.MySon7)
	peoples = append(peoples, fu.MySon8)
	peoples = append(peoples, fu.MySon9)
	peoples = append(peoples, fu.MySon10)
	peoples = append(peoples, fu.MyDaughter1)
	peoples = append(peoples, fu.MyDaughter2)
	peoples = append(peoples, fu.MyDaughter3)
	peoples = append(peoples, fu.MyDaughter4)
	peoples = append(peoples, fu.MyDaughter5)
	peoples = append(peoples, fu.MyDaughter6)
	peoples = append(peoples, fu.MyDaughter7)
	peoples = append(peoples, fu.MyDaughter8)
	peoples = append(peoples, fu.MyDaughter9)
	peoples = append(peoples, fu.MyDaughter10)

	return peoples
}

func getMyBrothersAndSisters() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, fu.MyWife)
	peoples = append(peoples, fu.MyBrother1)
	peoples = append(peoples, fu.MyBrother2)
	peoples = append(peoples, fu.MyBrother3)
	peoples = append(peoples, fu.MyBrother4)
	peoples = append(peoples, fu.MyBrother5)
	peoples = append(peoples, fu.MyBrother6)
	peoples = append(peoples, fu.MyBrother7)
	peoples = append(peoples, fu.MyBrother8)
	peoples = append(peoples, fu.MyBrother9)
	peoples = append(peoples, fu.MyBrother10)
	peoples = append(peoples, fu.MySister1)
	peoples = append(peoples, fu.MySister2)
	peoples = append(peoples, fu.MySister3)
	peoples = append(peoples, fu.MySister4)
	peoples = append(peoples, fu.MySister5)
	peoples = append(peoples, fu.MySister6)
	peoples = append(peoples, fu.MySister7)
	peoples = append(peoples, fu.MySister8)
	peoples = append(peoples, fu.MySister9)
	peoples = append(peoples, fu.MySister10)

	return peoples
}

func getMyBigFamiliesFriends() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, yuan.F湘兰_守礼II)
	peoples = append(peoples, yuan.M待举_守礼II)
	peoples = append(peoples, wuyibu.M多梅希_守礼II)
	peoples = append(peoples, wuyibu.F善理_守礼II)
	peoples = append(peoples, sesende.M安石_守礼II)
	peoples = append(peoples, sesende.F侯_守礼II)
	peoples = append(peoples, lou.M鸿渐_守礼II)
	peoples = append(peoples, lou.F仙惠_守礼II)
	peoples = append(peoples, wu.M方庆_守礼II)
	peoples = append(peoples, wu.F埃莱内_守礼II)
	peoples = append(peoples, yin.M日用_守礼II)
	peoples = append(peoples, yin.F蕣华_守礼II)
	peoples = append(peoples, lin.F超超_守礼II)
	peoples = append(peoples, lin.M确_守礼II)
	peoples = append(peoples, lin.M伍斯特_守礼II)
	peoples = append(peoples, lin.M瓌_守礼II)
	peoples = append(peoples, bi.M君集_守礼II)
	peoples = append(peoples, bi.F素节_守礼II)

	return peoples
}

func getMyBigLoads() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, 2830311)
	peoples = append(peoples, 2840690)

	return peoples
}

func getScotlandFriends() []int {
	var peoples []int

	peoples = append(peoples, sesende.M安石_守礼II)
	peoples = append(peoples, sesende.F侯_守礼II)
	peoples = append(peoples, 2811089)
	peoples = append(peoples, 2830311)
	peoples = append(peoples, 2822231)
	peoples = append(peoples, 2830384)
	peoples = append(peoples, 2850246)

	return peoples
}

func getWuFriends() []int {
	var peoples []int

	peoples = append(peoples, wu.M方庆_守礼II)
	peoples = append(peoples, wu.F埃莱内_守礼II)
	peoples = append(peoples, 2830314)
	peoples = append(peoples, 2811649)
	peoples = append(peoples, 2822928)
	peoples = append(peoples, 2831544)
	peoples = append(peoples, 2815827)
	peoples = append(peoples, 2846817)
	peoples = append(peoples, 2849641)

	return peoples
}

func getLouFriends() []int {
	var peoples []int

	peoples = append(peoples, lou.F仙惠_守礼II)
	peoples = append(peoples, lou.M鸿渐_守礼II)
	peoples = append(peoples, 2813163)
	peoples = append(peoples, 2816546)
	peoples = append(peoples, 2832575)
	peoples = append(peoples, 2833850)
	peoples = append(peoples, 2835430)
	peoples = append(peoples, 2825502)
	peoples = append(peoples, 2842863)
	peoples = append(peoples, 2844558)
	peoples = append(peoples, 2850317)
	peoples = append(peoples, 2851236)
	peoples = append(peoples, 2844939)
	peoples = append(peoples, 2840690)

	return peoples
}

func getBiFriends() []int {
	var peoples []int

	peoples = append(peoples, bi.F素节_守礼II)
	peoples = append(peoples, bi.M君集_守礼II)
	peoples = append(peoples, 2810790)
	peoples = append(peoples, 2825235)
	peoples = append(peoples, 2846549)
	peoples = append(peoples, 2830070)
	peoples = append(peoples, 2830070)
	peoples = append(peoples, 2849704)
	peoples = append(peoples, 2851140)

	return peoples
}

func getGermanFriends() []int {
	var peoples []int

	peoples = append(peoples, 2845721)
	peoples = append(peoples, 2840086)

	return peoples
}

func getYiwubuFriends() []int {
	var peoples []int

	peoples = append(peoples, wuyibu.F善理_守礼II)
	peoples = append(peoples, wuyibu.M多梅希_守礼II)
	peoples = append(peoples, 2843524)
	peoples = append(peoples, 2841943)
	peoples = append(peoples, 2838684)
	peoples = append(peoples, 2847861)
	peoples = append(peoples, 2847861)
	peoples = append(peoples, 2864693)
	peoples = append(peoples, 2853577)
	peoples = append(peoples, 2856195)
	peoples = append(peoples, 2850307)

	return peoples
}
