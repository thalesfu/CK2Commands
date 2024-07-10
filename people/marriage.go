package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
)

func buildMarriageScriptGenerator(husband *ck2nebula.People, wife *ck2nebula.People, isMatrilineal bool) (*PeopleScriptGenerator, *PeopleScriptGenerator) {
	husbandGenerator := NewPeopleScriptGenerator(husband)
	wifeGenerator := NewPeopleScriptGenerator(wife)

	if isMatrilineal {
		fmt.Printf("%s.%s(%d) marry with %s.%s(%d) matrilineal\n", husband.DynastyName, husband.Name, husband.ID, wife.DynastyName, wife.Name, wife.ID)
		husbandMarriage := NewMarriageScriptGenerator(husband, wife, isMatrilineal)
		husbandGenerator.AddScriptGenerator(husbandMarriage)
	} else {
		fmt.Printf("%s.%s(%d) marry with %s.%s(%d)\n", husband.DynastyName, husband.Name, husband.ID, wife.DynastyName, wife.Name, wife.ID)
		wifeMarriage := NewMarriageScriptGenerator(husband, wife, isMatrilineal)
		wifeGenerator.AddScriptGenerator(wifeMarriage)
	}

	modifiedHusbandFertility := getModifiedFloat32(husband.Fertility, 1, TalentQuick)
	if modifiedHusbandFertility > 0 {
		fmt.Printf("%s.%s(%d) add fertility %f\n", husband.DynastyName, husband.Name, husband.ID, modifiedHusbandFertility)
		husbandGenerator.AddScriptGenerator(NewModifyFertilityScriptGenerator(modifiedHusbandFertility))
	}

	modifiedHusbandHealth := getModifiedFloat32(husband.Health, 10, TalentQuick)
	if modifiedHusbandHealth > 0 {
		fmt.Printf("%s.%s(%d) add health %f\n", husband.DynastyName, husband.Name, husband.ID, modifiedHusbandHealth)
		husbandGenerator.AddScriptGenerator(NewModifyHealthScriptGenerator(modifiedHusbandHealth))
	}

	modifiedWifeFertility := getModifiedFloat32(wife.Fertility, 1, TalentQuick)
	if modifiedWifeFertility > 0 {
		fmt.Printf("%s.%s(%d) add fertility %f\n", wife.DynastyName, wife.Name, wife.ID, modifiedWifeFertility)
		wifeGenerator.AddScriptGenerator(NewModifyFertilityScriptGenerator(modifiedWifeFertility))
	}

	modifiedWifeHealth := getModifiedFloat32(wife.Health, 10, TalentQuick)
	if modifiedWifeHealth > 0 {
		fmt.Printf("%s.%s(%d) add health %f\n", wife.DynastyName, wife.Name, wife.ID, modifiedWifeHealth)
		wifeGenerator.AddScriptGenerator(NewModifyHealthScriptGenerator(modifiedWifeHealth))
	}

	return husbandGenerator, wifeGenerator
}

func BuildMarriageScript(space *nebulagolang.Space, coreFamily map[int]string) {
	story := ck2nebula.GetLatestStory(space)

	generators := make([]CK2Commands.ScriptGenerator, 0)

	allPeople := make(map[int]*PeopleScriptGenerator)

	BuildMarriage(space, story.Data.StoryId, coreFamily, allPeople)

	for _, p := range allPeople {
		generators = append(generators, p)
	}

	CK2Commands.BuildScript("marriage", generators...)
}
