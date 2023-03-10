package main

import (
	"github.com/thalesfu/CK2Commands/family/bi"
	"github.com/thalesfu/CK2Commands/family/circinn"
	"github.com/thalesfu/CK2Commands/family/fu"
	"github.com/thalesfu/CK2Commands/family/lin"
	"github.com/thalesfu/CK2Commands/family/lou"
	"github.com/thalesfu/CK2Commands/family/pictish"
	"github.com/thalesfu/CK2Commands/family/wu"
	"github.com/thalesfu/CK2Commands/family/wuyibu"
	"github.com/thalesfu/CK2Commands/family/yin"
	"github.com/thalesfu/CK2Commands/family/yuan"
	"github.com/thalesfu/CK2Commands/feudal"
	"github.com/thalesfu/CK2Commands/job"
	"github.com/thalesfu/CK2Commands/people"
)

func main() {

	feudal.BuildTitle()
	makefriends()
	pollinate()
	buildjobs()

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

}

func buildjobs() {
	people.BuildMyLoad(job.GetMyloads()...)
	people.BuildAmbassador(job.GetAmbassadors()...)
	people.BuildGeneral(job.GetGenerals()...)
	people.BuildManager(job.GetManagers()...)
	people.BuildSpy(job.GetSpys()...)
	people.BuildFather(job.GetFathers()...)
	people.BuildDoctor(job.GetDoctors()...)
	people.BuildRandom(job.GetRandoms()...)
}

func makefriends() {
	people.MakeFriends(getMyBrothersAndSisters(),
		getMySonsAndDaughters(),
		getMyBigFamiliesFriends(),
		getMyBigLoads(),
		circinn.GetFriends(),
		wu.GetFriends(),
		lou.GetLouFriends(),
		bi.GetBiFriends(),
		getGermanFriends(),
		getYiwubuFriends(),
		lin.GetFriends(),
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
	peoples = append(peoples, yuan.F??????_??????II)
	peoples = append(peoples, yuan.M??????_??????II)
	peoples = append(peoples, wuyibu.M?????????_??????II)
	peoples = append(peoples, wuyibu.F??????_??????II)
	peoples = append(peoples, circinn.M??????_??????II)
	peoples = append(peoples, circinn.F???_??????II)
	peoples = append(peoples, lou.M??????_??????II)
	peoples = append(peoples, lou.F??????_??????II)
	peoples = append(peoples, wu.M??????_??????II)
	peoples = append(peoples, wu.F?????????_??????II)
	peoples = append(peoples, yin.M??????_??????II)
	peoples = append(peoples, yin.F??????_??????II)
	peoples = append(peoples, lin.F??????_??????II)
	peoples = append(peoples, lin.M???_??????II)
	peoples = append(peoples, lin.M?????????_??????II)
	peoples = append(peoples, lin.M???_??????II)
	peoples = append(peoples, bi.M??????_??????II)
	peoples = append(peoples, bi.F??????_??????II)

	return peoples
}

func getMyBigLoads() []int {
	var peoples []int

	peoples = append(peoples, fu.Me)
	peoples = append(peoples, 2830311)
	peoples = append(peoples, 2840690)
	peoples = append(peoples, lin.M?????????10960227)

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

	peoples = append(peoples, wuyibu.F??????_??????II)
	peoples = append(peoples, wuyibu.M?????????_??????II)
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

func pollinate() {
	yuan.Hy()
	pictish.Hy()
	fu.Hy()
	lin.Pollinate()
	wu.Pollinate()
	bi.Pollinate()
	lou.Pollinate()
}
