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
	"time"
)

type buildFuncWithNebula func(*bufio.Writer, *nebulagolang.Space, *ck2nebula.People)

func buildPeopleWithNebula(fileName string, buildFuncs []buildFuncWithNebula, space *nebulagolang.Space, people ...*ck2nebula.People) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for _, p := range people {
		for _, bf := range buildFuncs {
			bf(writer, space, p)
		}
	}

	writer.Flush()
}

func buildPeopleRandomWithNebula(fileName string, buildFuncs [][]buildFuncWithNebula, space *nebulagolang.Space, people ...*ck2nebula.People) {
	file, err := utils.OpenFile(fileName)

	if err != nil {
		log.Println(err)
		return
	}

	defer utils.CloseFile(file)

	writer := bufio.NewWriter(file)

	for _, p := range people {
		group := getRandomBuildFuncWithNebula(buildFuncs)
		for _, bf := range group {
			bf(writer, space, p)
		}
	}

	writer.Flush()
}

func getRandomBuildFuncWithNebula(buildFuncs [][]buildFuncWithNebula) []buildFuncWithNebula {
	rand.Seed(time.Now().UnixNano())
	return buildFuncs[rand.Intn(len(buildFuncs))]
}

func AutoBuild(space *nebulagolang.Space, player *ck2nebula.People, coreFamily map[int]string) {

	group, newFriends := GetFriendsGroup(space, player, coreFamily)

	peopleGeneratorMap := MakeFriendsWithNebula(space, group)

	for _, d := range newFriends {
		for _, p := range d {
			spg := BuildSinglePeople(space, p)
			if pg, ok := peopleGeneratorMap[spg.Me.ID]; ok {
				for _, g := range spg.ScriptGenerators {
					pg.AddScriptGenerator(g)
				}
			} else {
				peopleGeneratorMap[spg.Me.ID] = spg
			}
		}
	}

	peopleGenerators := make([]CK2Commands.ScriptGenerator, 0)

	for _, p := range peopleGeneratorMap {
		peopleGenerators = append(peopleGenerators, p)
	}

	CK2Commands.BuildScript("autobuild", peopleGenerators...)
}

func BuildPeople(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, 0)

	for _, p := range people {
		generators = append(generators, BuildSinglePeople(space, p))
	}

	CK2Commands.BuildScript("buildpeople", generators...)
}

func BuildHealthAndFertility(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, 0)

	for _, p := range people {
		pg := NewPeopleScriptGenerator(p)
		talent := getTopTalent(space, p)
		modifiedFertility := getModifiedFloat32(p.Fertility, 1, talent)
		if modifiedFertility > 0 {
			fmt.Printf("%s.%s add fertility %f\n", p.DynastyName, p.Name, modifiedFertility)
			pg.AddScriptGenerator(NewModifyFertilityScriptGenerator(modifiedFertility))
		}

		modifiedHealth := getModifiedFloat32(p.Health, 10, talent)
		if modifiedHealth > 0 {
			fmt.Printf("%s.%s add health %f\n", p.DynastyName, p.Name, modifiedHealth)
			pg.AddScriptGenerator(NewModifyHealthScriptGenerator(modifiedHealth))
		}

		generators = append(generators, pg)
	}

	CK2Commands.BuildScript("buildpeople", generators...)
}

func BuildSinglePeople(space *nebulagolang.Space, people *ck2nebula.People) *PeopleScriptGenerator {
	result := NewPeopleScriptGenerator(people)

	modifiers := GetPropertyModifiers(space, people)

	_, generators := BuildPeopleTrait(space, people, modifiers)
	result.ScriptGenerators = append(result.ScriptGenerators, generators...)

	for _, modifier := range modifiers {
		if modifier.ModifiedValue > 0 {
			fmt.Printf("%s.%s add %s %d\n", people.DynastyName, people.Name, modifier.Property, modifier.ModifiedValue)
			result.AddScriptGenerator(NewModifyPropertyValueScriptGenerator(modifier))
		}
	}

	modifiedFertility := getModifiedFloat32(people.Fertility, 1, modifiers[0].Talent)
	if modifiedFertility > 0 {
		fmt.Printf("%s.%s add fertility %f\n", people.DynastyName, people.Name, modifiedFertility)
		result.AddScriptGenerator(NewModifyFertilityScriptGenerator(modifiedFertility))
	}

	modifiedHealth := getModifiedFloat32(people.Health, 10, modifiers[0].Talent)
	if modifiedHealth > 0 {
		fmt.Printf("%s.%s add health %f\n", people.DynastyName, people.Name, modifiedHealth)
		result.AddScriptGenerator(NewModifyHealthScriptGenerator(modifiedHealth))
	}

	return result
}

func getModifiedFloat32(base float32, top float32, talent Talent) float32 {
	var v float32
	repeatCount := (int(talent) + 1) * 5
	for i := 0; i <= repeatCount; i++ {
		rand.Seed(time.Now().UnixNano())
		r := rand.Float32() * (top - base)
		if r > v {
			v = r
		}
	}

	return v
}
