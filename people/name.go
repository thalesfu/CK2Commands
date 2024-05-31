package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
)

func ChangePeopleName() {
	people := make(map[int]string)
	people[2721442] = "继荣"

	BuildChangePeopleNameScript(ck2nebula.SPACE, people)
}

func BuildChangePeopleNameScript(space *nebulagolang.Space, people map[int]string) {
	storyResult := ck2nebula.GetLatestStory(space)
	if !storyResult.Ok {
		fmt.Println(storyResult.Err.Error())
		return
	}

	generators := make([]CK2Commands.ScriptGenerator, 0)

	for peopleId, name := range people {
		people := ck2nebula.GetPeopleByID(ck2nebula.SPACE, storyResult.Data.StoryId, peopleId)

		generator := buildChangePeopleNameScriptGenerator(people.Data, name)

		generators = append(generators, generator)
	}

	CK2Commands.BuildScript("name", generators...)

}

func buildChangePeopleNameScriptGenerator(people *ck2nebula.People, name string) *PeopleScriptGenerator {
	scriptGenerator := NewPeopleScriptGenerator(people)

	if name != "" {
		fmt.Printf("%s.%s(%d) change name to %s\n", people.DynastyName, people.Name, people.ID, name)
		scriptGenerator.AddScriptGenerator(NewChangeNameScriptGenerator(name))
	}

	return scriptGenerator
}
