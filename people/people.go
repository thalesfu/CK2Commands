package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"math/rand"
	"sort"
	"time"
)

type PropertyType string

const (
	PropertyTypeNone        PropertyType = "none"
	PropertyTypeDiplomacy                = "diplomacy"
	PropertyTypeMartial                  = "martial"
	PropertyTypeStewardship              = "stewardship"
	PropertyTypeIntrigue                 = "intrigue"
	PropertyTypeLearning                 = "learning"
)

type Talent int

const (
	TalentNormal Talent = iota
	TalentQuick
	TalentGenius
)

func (t Talent) GetPropertyTopValueLimit() int {
	switch t {
	case TalentGenius:
		return 20
	case TalentQuick:
		return 15
	default:
		return 10
	}
}

type PropertyModifier struct {
	Property         PropertyType
	OriginalValue    int
	ModifiedValue    int
	IsEducationTrait bool
	Talent           Talent
}

func NewPropertyModifier(property PropertyType, originalValue int, isEdu bool) *PropertyModifier {
	return &PropertyModifier{
		Property:         property,
		OriginalValue:    originalValue,
		IsEducationTrait: isEdu,
	}
}

func GetPropertyModifiers(space *nebulagolang.Space, people *ck2nebula.People) []*PropertyModifier {
	var result []*PropertyModifier

	eductionResult := map[PropertyType]bool{
		PropertyTypeDiplomacy:   false,
		PropertyTypeMartial:     false,
		PropertyTypeStewardship: false,
		PropertyTypeIntrigue:    false,
		PropertyTypeLearning:    false,
	}

	educationResult := people.GetEductionTraits(space)

	if educationResult.Ok {
		for _, trait := range educationResult.Data {
			if trait.Education {
				eductionResult[GetPropertyTypeByEductionTrait(trait)] = true
			}
		}
	}

	result = append(result, NewPropertyModifier(PropertyTypeDiplomacy, people.Diplomacy, eductionResult[PropertyTypeDiplomacy]))
	result = append(result, NewPropertyModifier(PropertyTypeMartial, people.Martial, eductionResult[PropertyTypeMartial]))
	result = append(result, NewPropertyModifier(PropertyTypeStewardship, people.Stewardship, eductionResult[PropertyTypeStewardship]))
	result = append(result, NewPropertyModifier(PropertyTypeIntrigue, people.Intrigue, eductionResult[PropertyTypeIntrigue]))
	result = append(result, NewPropertyModifier(PropertyTypeLearning, people.Learning, eductionResult[PropertyTypeLearning]))

	sort.Slice(result, func(i, j int) bool {
		if result[i].IsEducationTrait && !result[j].IsEducationTrait {
			return true
		}

		if !result[i].IsEducationTrait && result[j].IsEducationTrait {
			return false
		}

		return result[i].OriginalValue > result[j].OriginalValue
	})

	topTalent := getTopTalent(space, people)

	for _, modifier := range result {
		modifier.Talent = topTalent
		top := modifier.Talent.GetPropertyTopValueLimit()
		age := people.Age
		if age <= 6 {
			age = 6
		}
		if age < 16 {
			top = top * age / 16
		}
		v := top - modifier.OriginalValue
		if v > 0 {
			modifier.ModifiedValue = rand.Intn(v)
			if topTalent > TalentNormal {
				topTalent = Talent(rand.Intn(int(topTalent) + 1))
			}
		}
	}

	return result
}

func getTopTalent(space *nebulagolang.Space, people *ck2nebula.People) Talent {
	tr := people.GetTraits(space)

	if !tr.Ok {
		return Talent(rand.Intn(3))
	}

	if _, ok := tr.Data[GeniusTrait.Code]; ok {
		return TalentGenius
	}

	if _, ok := tr.Data[QuickTrait.Code]; ok {
		t := Talent(rand.Intn(3))
		if t == TalentGenius {
			return TalentQuick
		} else {
			return TalentQuick
		}
	}

	return Talent(rand.Intn(3))
}

func BuildPeople(space *nebulagolang.Space, people ...*ck2nebula.People) {
	generators := make([]CK2Commands.ScriptGenerator, 0)

	for _, p := range people {
		if !p.IsBuilt {
			generators = append(generators, BuildSinglePeople(space, p))
		}
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
			fmt.Printf("%s.%s(%d) add fertility %f\n", p.DynastyName, p.Name, p.ID, modifiedFertility)
			pg.AddScriptGenerator(NewModifyFertilityScriptGenerator(modifiedFertility))
		}

		modifiedHealth := getModifiedFloat32(p.Health, 10, talent)
		if modifiedHealth > 0 {
			fmt.Printf("%s.%s(%d) add health %f\n", p.DynastyName, p.Name, p.ID, modifiedHealth)
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
			fmt.Printf("%s.%s(%d) add %s %d\n", people.DynastyName, people.Name, people.ID, modifier.Property, modifier.ModifiedValue)
			result.AddScriptGenerator(NewModifyPropertyValueScriptGenerator(modifier))
		}
	}

	modifiedFertility := getModifiedFloat32(people.Fertility, 1, modifiers[0].Talent)
	if modifiedFertility > 0 {
		fmt.Printf("%s.%s(%d) add fertility %f\n", people.DynastyName, people.Name, people.ID, modifiedFertility)
		result.AddScriptGenerator(NewModifyFertilityScriptGenerator(modifiedFertility))
	}

	modifiedHealth := getModifiedFloat32(people.Health, 10, modifiers[0].Talent)
	if modifiedHealth > 0 {
		fmt.Printf("%s.%s(%d) add health %f\n", people.DynastyName, people.Name, people.ID, modifiedHealth)
		result.AddScriptGenerator(NewModifyHealthScriptGenerator(modifiedHealth))
	}

	var money float32 = 2000
	fmt.Printf("%s.%s(%d) add money %f\n", people.DynastyName, people.Name, people.ID, money)
	result.AddScriptGenerator(NewModifyWealthScriptGenerator(money))

	peopleType := "common"
	if people.Age < 16 {
		peopleType = "child"
	}

	result.AddScriptGenerator(NewSetBuildPeopleFlagScriptGenerator(peopleType))

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
