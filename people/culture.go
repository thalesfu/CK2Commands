package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
)

var CultureMap map[string]*ck2nebula.Culture

func init() {
	culturesResult := ck2nebula.GetAllCulture(ck2nebula.SPACE)
	if culturesResult.Ok {
		CultureMap = make(map[string]*ck2nebula.Culture)
		for _, culture := range culturesResult.Data {
			CultureMap[culture.Name] = culture
		}
	}
}

func setCultureToHanPictish(writer *bufio.Writer, peopleId int) {
	writeChangeCulture(writer, peopleId, "han", "pictish")
}

func CultureIsHanPictish(peopleId ...int) {
	var functions = []buildFunc{
		setCultureToHanPictish,
	}

	buildPeople("hanpictish", functions, peopleId...)
}

func HanPictish() {
	storyResult := ck2nebula.GetLatestStory(ck2nebula.SPACE)
	if !storyResult.Ok {
		fmt.Println(storyResult.Err.Error())
		return
	}

	han := CultureMap["汉"]
	pictish := CultureMap["皮克特"]

	peopleIds := make([]int, 0)
	peopleIds = append(peopleIds, 2717663)
	peopleIds = append(peopleIds, 2729781)
	peopleIds = append(peopleIds, 2729779)
	peopleIds = append(peopleIds, 2729774)
	peopleIds = append(peopleIds, 2712711)
	peopleIds = append(peopleIds, 2726856)
	peopleIds = append(peopleIds, 2726855)
	peopleIds = append(peopleIds, 2726776)

	generators := make([]CK2Commands.ScriptGenerator, 0)

	for _, peopleId := range peopleIds {
		people := ck2nebula.GetPeopleByID(ck2nebula.SPACE, storyResult.Data.StoryId, peopleId)

		generator := buildCultureScriptGenerator(people.Data, han, pictish)

		generators = append(generators, generator)
	}

	CK2Commands.BuildScript("hanpictish", generators...)
}

func setCultureToHanScottish(writer *bufio.Writer, peopleId int) {
	writeChangeCulture(writer, peopleId, "han", "scottish")
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

func CultureIsHanScottish(peopleId ...int) {
	var functions = []buildFunc{
		setCultureToHanScottish,
	}

	buildPeople("hanscottish", functions, peopleId...)
}

func setCultureToHanWelsh(writer *bufio.Writer, peopleId int) {
	writeChangeCulture(writer, peopleId, "han", "welsh")
}

func CultureIsHanWelsh(peopleId ...int) {
	var functions = []buildFunc{
		setCultureToHanWelsh,
	}

	buildPeople("hanwelsh", functions, peopleId...)
}
