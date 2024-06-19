package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/CK2Commands/property"
	"github.com/thalesfu/CK2Commands/trait"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"math/rand"
)

var fatherTraits = []string{
	"mastermind_theologian", // 神学的大师
	"erudite",               // 博学
	"just",                  // 正直
	"content",               // 安于现状
	"theologian",            // 神学家
	"brave",                 // 勇敢
	"zealous",               // 狂热
}

func addFatherTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range fatherTraits {
		writeAddTrait(writer, trait, peopleId)
	}
}

func addFatherLearning(writer *bufio.Writer, peopleId int) {
	writeAddLearning(writer, peopleId, 5)
	writeAddMartial(writer, peopleId, 15)
}

func getFatherFunctions() []buildFunc {
	return []buildFunc{
		removeEducationTraits,
		RemoveBadTraits,
		addCommonGoodTraits,
		addFatherTraits,
		addFatherLearning,
	}
}

func BuildFather(peopleId ...int) {
	buildPeople("father", getFatherFunctions(), peopleId...)
}

func buildFatherPeopleScriptGenerator(space *nebulagolang.Space, people *ck2nebula.People) *PeopleScriptGenerator {
	result := NewPeopleScriptGenerator(people)

	modifiers := GetPropertyModifiers(space, people)

	var chance = 5

	switch modifiers[0].Talent {
	case TalentGenius:
		chance = 3
	case TalentQuick:
		chance = 4
	}

	fatherModifiers := make([]*PropertyModifier, 2)

	for _, modifier := range modifiers {
		if modifier.Property == property.PropertyTypeLearning {
			modifier.IsEducationTrait = true
			fatherModifiers[0] = modifier
		} else if modifier.Property == property.PropertyTypeMartial {
			fatherModifiers[1] = modifier
			modifier.ModifiedValue += 20
		}
	}

	traits, sgs := RemoveBadTraitsWithNebula(space, people)

	for _, sg := range sgs {
		result.AddScriptGenerator(sg)
	}

	traits, edusgs := BuildPeopleEductionTrait(people, traits, fatherModifiers)

	for _, edusg := range edusgs {
		result.AddScriptGenerator(edusg)
	}

	for _, tr := range trait.CommonGoodTraits {
		if rand.Intn(chance) == 0 {
			if traits[tr.Code] == nil {
				fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
				result.AddScriptGenerator(NewAddTraitScriptGenerator(tr))
				traits[tr.Code] = tr
			}
		}
	}

	for _, tr := range trait.VirtueTraits {
		if rand.Intn(2) == 0 {
			if traits[tr.Code] == nil {
				fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
				result.AddScriptGenerator(NewAddTraitScriptGenerator(tr))
				traits[tr.Code] = tr
			}
		}
	}

	lifeStyletraits := trait.LifeStyleTraitsByPropertyType[fatherModifiers[0].Property]

	for _, lifeStyleTrait := range lifeStyletraits {
		if rand.Intn(chance) == 0 {
			if traits[lifeStyleTrait.Code] == nil {
				fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, lifeStyleTrait.Name, lifeStyleTrait.Code)
				result.AddScriptGenerator(NewAddTraitScriptGenerator(lifeStyleTrait))
				traits[lifeStyleTrait.Code] = lifeStyleTrait
			}
		}
	}

	goodTraits := trait.GoodTraitsByPropertyType[fatherModifiers[0].Property]
	for _, tr := range goodTraits {
		if rand.Intn(chance) == 0 {
			if traits[tr.Code] == nil {
				result.AddScriptGenerator(NewAddTraitScriptGenerator(tr))
				fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
				traits[tr.Code] = tr
			}
		}
	}

	for _, modifier := range fatherModifiers {
		if modifier.ModifiedValue > 0 {
			fmt.Printf("%s.%s(%d) add %s %d\n", people.DynastyName, people.Name, people.ID, modifier.Property, modifier.ModifiedValue)
			result.AddScriptGenerator(NewModifyPropertyValueScriptGenerator(modifier))
		}
	}

	modifiedHealth := getModifiedFloat32(people.Health, 10, fatherModifiers[0].Talent)
	if modifiedHealth > 0 {
		fmt.Printf("%s.%s(%d) add health %f\n", people.DynastyName, people.Name, people.ID, modifiedHealth)
		result.AddScriptGenerator(NewModifyHealthScriptGenerator(modifiedHealth))
	}

	var money float32 = 2000
	fmt.Printf("%s.%s(%d) add money %f\n", people.DynastyName, people.Name, people.ID, money)
	result.AddScriptGenerator(NewModifyWealthScriptGenerator(money))

	result.AddScriptGenerator(NewSetBuildPeopleFlagScriptGenerator("father"))

	return result
}

func BuildFatherPeopleScript(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, len(people))

	for i, p := range people {
		if !p.IsBuilt {
			generators[i] = buildFatherPeopleScriptGenerator(space, p)
		}
	}

	CK2Commands.BuildScript("father", generators...)
}

func BuildMyFathers(space *nebulagolang.Space, people *ck2nebula.People) {
	fathers := make([]*ck2nebula.People, 0)

	vr := people.GetVassals(space)

	if vr.Ok {
		for _, v := range vr.Data {
			if v.IsReligionVassal(space) && v.Religion == people.Religion {
				fathers = append(fathers, v)
			}
		}
	}

	BuildFatherPeopleScript(space, fathers...)
}
