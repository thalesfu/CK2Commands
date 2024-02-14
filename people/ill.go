package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
)

var IllTraitMap map[string]*ck2nebula.Trait

func init() {
	tr := ck2nebula.GetAllTraits(ck2nebula.SPACE)
	if tr.Ok {
		IllTraitMap = make(map[string]*ck2nebula.Trait)
		for _, trait := range tr.Data {
			if trait.Code == "pregnancy_finishing" {
				continue
			}
			if trait.IsHealth || trait.IsIllness || trait.Blinding {
				IllTraitMap[trait.Code] = trait
			}
		}
	}
}

func cureIlls(writer *bufio.Writer, peopleId int) {
	for trait, _ := range IllTraitMap {
		writeRemoveTrait(writer, trait, peopleId)
	}
}

func CurePeopleIll(peopleId ...int) {
	var functions = []buildFunc{
		cureIlls,
	}

	buildPeople("cureill", functions, peopleId...)
}

func buildForceCurePeopleScriptGenerator(space *nebulagolang.Space, people *ck2nebula.People) *PeopleScriptGenerator {
	scriptGenerator := NewPeopleScriptGenerator(people)

	for _, trait := range IllTraitMap {
		if trait.Code == "pregnancy_finishing" {
			continue
		}
		if trait.IsHealth || trait.IsIllness || trait.Blinding {
			scriptGenerator.AddScriptGenerator(NewRemoveTraitScriptGenerator(trait))
		}
	}

	return scriptGenerator
}

func BuildForceCurePeopleScript(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, len(people))

	for i, p := range people {
		generators[i] = buildForceCurePeopleScriptGenerator(space, p)
	}

	CK2Commands.BuildScript("cure", generators...)
}

func ForceCurePeople(space *nebulagolang.Space, player *ck2nebula.People, coreFamily map[int]string) {
	groups := GetFriendsGroup(space, player, coreFamily)

	people := make([]*ck2nebula.People, 0)

	for _, group := range groups {
		for _, p := range group {
			people = append(people, p)
		}
	}

	people = append(people, player)

	BuildForceCurePeopleScript(space, people...)
}

func GetAllReligionVassal(space *nebulagolang.Space, people *ck2nebula.People, religion string) []*ck2nebula.People {
	religionVassals := make([]*ck2nebula.People, 0)

	vr := people.GetVassals(space)

	if vr.Ok {
		for _, v := range vr.Data {
			if v.IsReligionVassal(space) && religion == v.Religion {
				religionVassals = append(religionVassals, v)
			}

			religionVassals = append(religionVassals, GetAllReligionVassal(space, v, religion)...)
		}
	}

	return religionVassals
}

func buildCurePeopleScriptGenerator(space *nebulagolang.Space, people *ck2nebula.People) *PeopleScriptGenerator {
	scriptGenerator := NewPeopleScriptGenerator(people)

	tr := people.GetTraits(space)

	if tr.Ok {
		for _, trait := range tr.Data {
			if trait.Code == "pregnancy_finishing" {
				continue
			}
			if trait.IsHealth || trait.IsIllness || trait.Blinding {
				fmt.Printf("%s.%s(%d) remove trait %s(%s)\n", people.DynastyName, people.Name, people.ID, trait.Name, trait.Code)
				scriptGenerator.AddScriptGenerator(NewRemoveTraitScriptGenerator(trait))
			}
		}
	}

	return scriptGenerator
}

func BuildCurePeopleScript(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, len(people))

	for i, p := range people {
		generators[i] = buildCurePeopleScriptGenerator(space, p)
	}

	CK2Commands.BuildScript("cure", generators...)
}

func CureFriends(space *nebulagolang.Space, people *ck2nebula.People) {
	fr := people.GetFriends(space)

	if fr.Ok {
		friends := make([]*ck2nebula.People, len(fr.Data)+1)
		friends[0] = people

		i := 1
		for _, friend := range fr.Data {
			friends[i] = friend
			i++
		}

		BuildCurePeopleScript(space, friends...)
	}
}
