package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/CK2Commands/utils"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"log"
	"math/rand"
)

type Couple struct {
	Husband int
	Wife    int
}

func Pollinate(couples []Couple, fileName string) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for i, couple := range couples {
		if i != 0 {
			writer.WriteString("\n")
		}

		writer.WriteString(fmt.Sprintf("pollinate %d %d", couple.Wife, couple.Husband))
	}

	writer.Flush()
}

func buildPollinatePeopleScriptGenerator(space *nebulagolang.Space, people *ck2nebula.People) *PeopleScriptGenerator {
	scriptGenerator := NewPeopleScriptGenerator(people)

	spouses := make([]*ck2nebula.People, 0)

	if people.IsFemale {
		if people.IsSuccessionMarriage(space) {
			sr := people.GetAliveSpouses(space)

			if sr.Ok {
				for _, spouse := range sr.Data {
					if !spouse.IsDead {
						spouses = append(spouses, spouse)
					}
				}
			}
		}
	} else {
		if !people.IsSuccessionMarriage(space) {
			sr := people.GetAliveSpouses(space)

			if sr.Ok {
				for _, spouse := range sr.Data {
					if !spouse.IsDead && !spouse.IsPregnant && spouse.Age <= 45 {
						spouses = append(spouses, spouse)
					}
				}
			}
		}

		cr := people.GetAliveConsorts(space)

		if cr.Ok {
			for _, consort := range cr.Data {
				if !consort.IsDead && !consort.IsPregnant && consort.Age <= 45 {
					spouses = append(spouses, consort)
				}
			}
		}
	}

	if len(spouses) > 0 {
		spouse := spouses[rand.Intn(len(spouses))]

		fmt.Printf("%s.%s(%d) and %s.%s(%d) make a baby\n", people.DynastyName, people.Name, people.ID, spouse.DynastyName, spouse.Name, spouse.ID)
		scriptGenerator.AddScriptGenerator(NewImpregnateScriptGenerator(spouse))
	}

	return scriptGenerator
}

func NeedPollinate(space *nebulagolang.Space, people *ck2nebula.People) bool {
	if people.Age < 16 {
		return false
	}

	if people.IsFemale {
		if !people.IsSuccessionMarriage(space) {
			return false
		}
	}

	sonCount := 0

	sr := people.GetAliveSon(space)

	if sr.Ok {
		for _, son := range sr.Data {
			if son.Age < 45 {
				sonCount++
			}
		}
	}

	if sonCount >= 2 {
		return false
	}

	grandsonCount := 0

	gsr := people.GetAliveGrandSon(space)

	if gsr.Ok {
		for _, grandson := range gsr.Data {
			if grandson.Age < 45 {
				grandsonCount++
			}
		}
	}

	if grandsonCount+sonCount >= 4 {
		return false
	}

	return true
}

func BuildPollinatePeopleScript(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, 0)

	for _, p := range people {
		if NeedPollinate(space, p) {
			generators = append(generators, buildPollinatePeopleScriptGenerator(space, p))
		}
	}

	CK2Commands.BuildScript("pollinate", generators...)
}

func PollinateFriendsAndFriendsFriends(space *nebulagolang.Space, people *ck2nebula.People) {
	fr := people.GetFriends(space)

	if fr.Ok {
		friends := make([]*ck2nebula.People, 0)
		friends = append(friends, people)

		for _, friend := range fr.Data {
			friends = append(friends, friend)

			sfr := friend.GetFriends(space)

			if sfr.Ok {
				for _, sfriend := range sfr.Data {
					friends = append(friends, sfriend)
				}
			}
		}

		BuildPollinatePeopleScript(space, friends...)
	}
}
