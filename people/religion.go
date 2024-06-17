package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/CK2Commands/religion"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"strings"
)

func Taoist() {
	peopleIds := make([]int, 0)
	peopleIds = append(peopleIds, 2698888)

	BuildReligionScript(ck2nebula.SPACE, religion.Religion_东方宗教_道教_taoist, peopleIds...)
}

func BuildReligionScript(space *nebulagolang.Space, religion *ck2nebula.Religion, people ...int) {
	storyResult := ck2nebula.GetLatestStory(space)
	if !storyResult.Ok {
		fmt.Println(storyResult.Err.Error())
		return
	}

	generators := make([]CK2Commands.ScriptGenerator, 0)

	for _, peopleId := range people {
		people := ck2nebula.GetPeopleByID(ck2nebula.SPACE, storyResult.Data.StoryId, peopleId)

		generator := buildReligionScriptGenerator(people.Data, religion)

		generators = append(generators, generator)
	}

	CK2Commands.BuildScript(strings.ToLower(religion.Code), generators...)
}

func buildReligionScriptGenerator(people *ck2nebula.People, religion *ck2nebula.Religion) *PeopleScriptGenerator {
	scriptGenerator := NewPeopleScriptGenerator(people)

	if people.Religion == "" || people.Religion != religion.Code {
		fmt.Printf("%s.%s(%d) set religion as %s(%s)\n", people.DynastyName, people.Name, people.ID, religion.Name, religion.Code)
		scriptGenerator.AddScriptGenerator(NewSetReligionScriptGenerator(religion))
	}

	if people.SecretReligion != "" {
		fmt.Printf("%s.%s(%d) clear secret religion %s\n", people.DynastyName, people.Name, people.ID, people.SecretReligionName)
		scriptGenerator.AddScriptGenerator(NewClearSecretReligionScriptGenerator(people.SecretReligionName))
	}

	return scriptGenerator
}

func BuildWashReligionScript(space *nebulagolang.Space, religion *ck2nebula.Religion) {
	storyResult := ck2nebula.GetLatestStory(space)
	if !storyResult.Ok {
		fmt.Println(storyResult.Err.Error())
		return
	}

	pr := storyResult.Data.GetPlayer(space)

	if !pr.Ok {
		fmt.Println(pr.Err.Error())
		return
	}

	generators := buildWashVassalReligionsScriptGenerators(space, pr.Data, religion, nil)

	CK2Commands.BuildScript("washreligion", generators...)
}

func buildWashVassalReligionsScriptGenerators(space *nebulagolang.Space, people *ck2nebula.People, religion *ck2nebula.Religion, scannedPeople map[int]*ck2nebula.People) []CK2Commands.ScriptGenerator {
	vr := people.GetVassals(space)

	result := make([]CK2Commands.ScriptGenerator, 0)

	if scannedPeople == nil {
		scannedPeople = make(map[int]*ck2nebula.People)
	}

	if vr.Ok {
		for _, v := range vr.Data {
			if _, ok := scannedPeople[v.ID]; !ok {
				if v.Religion != religion.Code || v.SecretReligion != religion.Code {
					prsg := buildReligionScriptGenerator(v, religion)
					result = append(result, prsg)
					scannedPeople[v.ID] = v
				}

				subPrsgs := buildWashVassalReligionsScriptGenerators(space, v, religion, scannedPeople)
				result = append(result, subPrsgs...)
			}
		}
	}

	er := people.GetEmpirePeople(space)

	if er.Ok {
		for _, e := range er.Data {
			if _, ok := scannedPeople[e.ID]; !ok {
				if e.Religion != religion.Code || e.SecretReligion != religion.Code {
					prsg := buildReligionScriptGenerator(e, religion)
					result = append(result, prsg)
					scannedPeople[e.ID] = e
				}
			}
		}
	}

	return result
}
