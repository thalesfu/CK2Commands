package people

import (
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"math/rand"
	"sort"
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

	topTalent := Talent(rand.Intn(3))

	for _, modifier := range result {
		modifier.Talent = topTalent
		top := modifier.Talent.GetPropertyTopValueLimit()
		if people.Age < 16 {
			top = top * people.Age / 16
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
