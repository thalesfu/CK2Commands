package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/CK2Commands/culture"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"strings"
)

func HanPictish() {
	peopleIds := make([]int, 0)
	peopleIds = append(peopleIds, 2749220)

	BuildCultureScript(ck2nebula.SPACE, culture.Culture_中华_汉_chinese_group_han, culture.Culture_凯尔特_皮克特_celtic_pictish, peopleIds...)
}

func BuildCultureScript(space *nebulagolang.Space, culture *ck2nebula.Culture, ethnicity *ck2nebula.Culture, people ...int) {
	storyResult := ck2nebula.GetLatestStory(space)
	if !storyResult.Ok {
		fmt.Println(storyResult.Err.Error())
		return
	}

	generators := make([]CK2Commands.ScriptGenerator, 0)

	for _, peopleId := range people {
		people := ck2nebula.GetPeopleByID(ck2nebula.SPACE, storyResult.Data.StoryId, peopleId)

		generator := buildCultureScriptGenerator(people.Data, culture, ethnicity)

		generators = append(generators, generator)
	}

	CK2Commands.BuildScript(strings.ToLower(fmt.Sprintf("%s%s", culture.Code, ethnicity.Code)), generators...)

}

func buildCultureScriptGenerator(people *ck2nebula.People, culture *ck2nebula.Culture, ethnicity *ck2nebula.Culture) *PeopleScriptGenerator {
	scriptGenerator := NewPeopleScriptGenerator(people)

	if culture != nil {
		fmt.Printf("%s.%s(%d) set culture as %s(%s)\n", people.DynastyName, people.Name, people.ID, culture.Name, culture.Code)
		scriptGenerator.AddScriptGenerator(NewSetCultureScriptGenerator(culture))
	}

	if ethnicity != nil {
		fmt.Printf("%s.%s(%d) set gfx_culture as %s(%s)\n", people.DynastyName, people.Name, people.ID, ethnicity.Name, ethnicity.Code)
		scriptGenerator.AddScriptGenerator(NewSetEthnicityScriptGenerator(ethnicity))
	}

	return scriptGenerator
}
