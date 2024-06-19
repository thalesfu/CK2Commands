package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/CK2Commands/property"
	"github.com/thalesfu/CK2Commands/trait"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"math/rand"
)

func GetPropertyTypeByEductionTrait(tr *ck2nebula.Trait) property.PropertyType {
	if tr == nil || !tr.Education {
		return property.PropertyTypeNone
	}

	if t, ok := trait.EducationsTrait[tr.Code]; ok {
		return t
	}

	return property.PropertyTypeNone
}

func BuildPeopleTrait(space *nebulagolang.Space, people *ck2nebula.People, modifiers []*PropertyModifier) (map[string]*ck2nebula.Trait, []CK2Commands.ScriptGenerator) {
	traits, scriptGenerators := RemoveBadTraitsWithNebula(space, people)

	traits, edusg := BuildPeopleEductionTrait(people, traits, modifiers)

	scriptGenerators = append(scriptGenerators, edusg...)

	var chance = 5

	switch modifiers[0].Talent {
	case TalentGenius:
		if traits[trait.Trait_56_genius_天才.Code] == nil {
			scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(trait.Trait_56_genius_天才))
			fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, trait.Trait_56_genius_天才.Name, trait.Trait_56_genius_天才.Code)
			traits[trait.Trait_56_genius_天才.Code] = trait.Trait_56_genius_天才
		}
		chance = 3
	case TalentQuick:
		if traits[trait.Trait_57_quick_敏锐.Code] == nil {
			scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(trait.Trait_57_quick_敏锐))
			fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, trait.Trait_57_quick_敏锐.Name, trait.Trait_57_quick_敏锐.Code)
			traits[trait.Trait_57_quick_敏锐.Code] = trait.Trait_57_quick_敏锐
		}
		chance = 4
	default:
	}

	for _, tr := range trait.CommonGoodTraits {
		if rand.Intn(chance) == 0 {
			if traits[tr.Code] == nil {
				scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(tr))
				fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
				traits[tr.Code] = tr
			}
		}
	}

	for _, tr := range trait.VirtueTraits {
		if rand.Intn(chance) == 0 {
			if traits[tr.Code] == nil {
				scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(tr))
				fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
				traits[tr.Code] = tr
			}
		}
	}

	if modifiers[0].IsEducationTrait {
		lifeStyletraits := trait.LifeStyleTraitsByPropertyType[modifiers[0].Property]

		lifeStyleTrait := lifeStyletraits[rand.Intn(len(lifeStyletraits))]

		if lifeStyleTrait.Code == "seducer" {
			if people.IsFemale {
				lifeStyleTrait = trait.Trait_179_seductress_魅惑大师
			}
		} else if lifeStyleTrait.Code == "seductress" {
			if !people.IsFemale {
				lifeStyleTrait = trait.Trait_178_seducer_勾引大师
			}
		}

		if traits[lifeStyleTrait.Code] == nil {
			scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(lifeStyleTrait))
			fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, lifeStyleTrait.Name, lifeStyleTrait.Code)
			traits[lifeStyleTrait.Code] = lifeStyleTrait
		}

		goodTraits := trait.GoodTraitsByPropertyType[modifiers[0].Property]
		for _, tr := range goodTraits {
			if rand.Intn(chance) == 0 {
				if traits[tr.Code] == nil {
					scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(tr))
					fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
					traits[tr.Code] = tr
				}
			}
		}

		if modifiers[0].Property == property.PropertyTypeMartial {
			i := 0
			for _, tr := range trait.LeaderTraits {
				if rand.Intn(chance+i*2) == 0 {
					if traits[tr.Code] == nil {
						scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(tr))
						fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
						traits[tr.Code] = tr
					}
				}
			}
		}
	} else {
		for _, tr := range trait.ChildhoodTraits {
			if rand.Intn(chance) == 0 {
				if traits[tr.Code] == nil {
					scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(tr))
					fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, tr.Name, tr.Code)
					traits[tr.Code] = tr
				}
			}
		}
	}

	return traits, scriptGenerators
}

func BuildPeopleEductionTrait(people *ck2nebula.People, traits map[string]*ck2nebula.Trait, modifiers []*PropertyModifier) (map[string]*ck2nebula.Trait, []CK2Commands.ScriptGenerator) {
	var newEducationTrait *ck2nebula.Trait
	var oldEducationTrait *ck2nebula.Trait

	for _, tr := range traits {
		if tr.Education {
			oldEducationTrait = tr
			break
		}
	}

	if oldEducationTrait == nil {
		return traits, nil
	}

	generators := make([]CK2Commands.ScriptGenerator, 0)

	switch modifiers[0].Property {
	case property.PropertyTypeDiplomacy:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "grey_eminence" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_8_grey_eminence_幕后操控人
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "grey_eminence" || oldEducationTrait.Code == "charismatic_negotiator" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_7_charismatic_negotiator_魅力非凡的说客
		} else {
			if oldEducationTrait.Code == "grey_eminence" || oldEducationTrait.Code == "charismatic_negotiator" || oldEducationTrait.Code == "underhanded_rogue" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_6_underhanded_rogue_狡猾的无赖
		}
	case property.PropertyTypeMartial:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "brilliant_strategist" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_16_brilliant_strategist_天才战略家
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "brilliant_strategist" || oldEducationTrait.Code == "skilled_tactician" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_15_skilled_tactician_优秀战术家
		} else {
			if oldEducationTrait.Code == "brilliant_strategist" || oldEducationTrait.Code == "skilled_tactician" || oldEducationTrait.Code == "tough_soldier" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_14_tough_soldier_坚强的战士
		}
	case property.PropertyTypeStewardship:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "midas_touched" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_12_midas_touched_点石成金者
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "midas_touched" || oldEducationTrait.Code == "thrifty_clerk" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_10_thrifty_clerk_节俭的职员

			return traits, generators
		} else {
			if oldEducationTrait.Code == "midas_touched" || oldEducationTrait.Code == "thrifty_clerk" || oldEducationTrait.Code == "fortune_builder" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_11_fortune_builder_财富创造者
		}
	case property.PropertyTypeIntrigue:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "elusive_shadow" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_4_elusive_shadow_难以捉摸的影子
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "elusive_shadow" || oldEducationTrait.Code == "intricate_webweaver" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_3_intricate_webweaver_暗中策划者
		} else {
			if oldEducationTrait.Code == "elusive_shadow" || oldEducationTrait.Code == "intricate_webweaver" || oldEducationTrait.Code == "flamboyant_schemer" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_2_flamboyant_schemer_浮夸的谋士
		}
	case property.PropertyTypeLearning:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "mastermind_theologian" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_20_mastermind_theologian_神学的大师
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "mastermind_theologian" || oldEducationTrait.Code == "scholarly_theologian" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_19_scholarly_theologian_渊博的神学家
		} else {
			if oldEducationTrait.Code == "mastermind_theologian" || oldEducationTrait.Code == "scholarly_theologian" || oldEducationTrait.Code == "martial_cleric" {
				return traits, generators
			}

			newEducationTrait = trait.Trait_18_martial_cleric_尽职的教士
		}

		fmt.Printf("%s.%s(%d) remove trait %s(%s)\n", people.DynastyName, people.Name, people.ID, oldEducationTrait.Name, oldEducationTrait.Code)
		generators = append(generators, NewRemoveTraitScriptGenerator(oldEducationTrait))
		delete(traits, oldEducationTrait.Code)

		fmt.Printf("%s.%s(%d) add trait %s(%s)\n", people.DynastyName, people.Name, people.ID, newEducationTrait.Name, newEducationTrait.Code)
		generators = append(generators, NewAddTraitScriptGenerator(newEducationTrait))
		traits[newEducationTrait.Code] = newEducationTrait
	}

	return traits, generators
}
