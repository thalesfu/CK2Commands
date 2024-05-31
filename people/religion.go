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

	if religion != nil {
		fmt.Printf("%s.%s(%d) set religion as %s(%s)\n", people.DynastyName, people.Name, people.ID, religion.Name, religion.Code)
		scriptGenerator.AddScriptGenerator(NewSetReligionScriptGenerator(religion))
	}

	return scriptGenerator
}
