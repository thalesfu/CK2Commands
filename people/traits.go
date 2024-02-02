package people

import (
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
	"math/rand"
)

var AllTraits map[string]*ck2nebula.Trait
var CommonGoodTraits map[string]*ck2nebula.Trait
var LifeStyleTraitsByPropertyType map[PropertyType][]*ck2nebula.Trait
var GoodTraitsByPropertyType map[PropertyType][]*ck2nebula.Trait
var GeniusTrait *ck2nebula.Trait
var QuickTrait *ck2nebula.Trait
var LeaderTraits map[string]*ck2nebula.Trait
var VirtueTraits map[string]*ck2nebula.Trait
var ChildhoodTraits map[string]*ck2nebula.Trait

func init() {
	tr := ck2nebula.GetAllTraits(ck2nebula.SPACE)

	if tr.Ok {
		AllTraits = tr.Data
		CommonGoodTraits = map[string]*ck2nebula.Trait{
			"fair":    AllTraits["fair"],    // 魅力非凡
			"groomed": AllTraits["groomed"], // 整洁
			"shrewd":  AllTraits["shrewd"],  // 精明
			"robust":  AllTraits["robust"],  // 健壮
			"strong":  AllTraits["strong"],  // 强壮
		}
		LifeStyleTraitsByPropertyType = map[PropertyType][]*ck2nebula.Trait{
			PropertyTypeDiplomacy: {
				AllTraits["socializer"], // 社交者 d:3
				AllTraits["hedonist"],   // 享乐主义者 d:2,i:1
				AllTraits["gamer"],      // 游戏大师 d:2,m:1
				AllTraits["seducer"],    // 勾引大师 d:1,i:2
				AllTraits["seductress"], // 魅惑大师 d:1,i:2
			},
			PropertyTypeMartial: {
				AllTraits["duelist"],    // 决斗者 m:3
				AllTraits["hunter"],     // 猎人 m:2,d1
				AllTraits["strategist"], // 策略家 m:2:s:1
				AllTraits["gamer"],      // 游戏大师 m:1,d:2
				AllTraits["architect"],  // 建筑师 m:1,s:2
			},
			PropertyTypeStewardship: {
				AllTraits["administrator"], // 管理者 s:3
				AllTraits["architect"],     // 建筑师 s:2,m:1
				AllTraits["gardener"],      // 园丁 s:2,l:1
				AllTraits["strategist"],    // 策略家 s:1,m:2
				AllTraits["mystic"],        // 神秘主义者 s:1,l:2
			},
			PropertyTypeIntrigue: {
				AllTraits["schemer"],    // 密谋大师 i:3
				AllTraits["seducer"],    // 勾引大师 i:2,d:1
				AllTraits["seductress"], // 魅惑大师 i:2,d:1
				AllTraits["impaler"],    // 拷打爱好者 i:2,l:1
				AllTraits["hedonist"],   // 享乐主义者 i:1,d:2
				AllTraits["theologian"], // 神学家 i:1,l:2
			},
			PropertyTypeLearning: {
				AllTraits["scholar"],    // 学者 l:3
				AllTraits["mystic"],     // 神秘主义者 l:2,s:1
				AllTraits["theologian"], // 神学家 l:2,i:1
				AllTraits["gardener"],   // 园丁 l:1,s:2
				AllTraits["impaler"],    // 拷打爱好者 l:1,i:2
			},
		}
		GeniusTrait = AllTraits["genius"]
		QuickTrait = AllTraits["quick"]

		GoodTraitsByPropertyType = map[PropertyType][]*ck2nebula.Trait{
			PropertyTypeDiplomacy: {
				AllTraits["honest"],     // 诚实 d:3,i:-2
				AllTraits["shrewd"],     // 精明 d:2,m:2,s:2,i:2,l:2
				AllTraits["ambitious"],  // 野心勃勃 d:2,m:2,s:2,i:2,l:2
				AllTraits["gregarious"], // 合群 d:2
				AllTraits["robust"],     // 健壮 d:1,m:2
				AllTraits["strong"],     // 强壮 d:1,m:2
				AllTraits["physician"],  // 著名医师 d:1,l:2
				AllTraits["falconer"],   // 驯鹰师 d:1
				AllTraits["poet"],       // 诗人 d:1
				AllTraits["trusting"],   // 轻信他人 d:1,i:-2
			},
			PropertyTypeMartial: {
				AllTraits["berserker"],  // 狂战士 m:4,d:-2
				AllTraits["adventurer"], // 冒险者 m:1,d:-1
				AllTraits["shrewd"],     // 精明 m:2,d:2,s:2,i:2,l:2
				AllTraits["ambitious"],  // 野心勃勃 m:2,d:2,s:2,i:2,l:2
				AllTraits["robust"],     // 健壮 m2,d:1
				AllTraits["strong"],     // 强壮 m:2,d:1
				AllTraits["zealous"],    // 狂热 m:2
				AllTraits["brave"],      // 勇敢 m:2
			},
			PropertyTypeStewardship: {
				AllTraits["shrewd"],    // 精明 s:2,d:2,m:2,i:2,l:2
				AllTraits["ambitious"], // 野心勃勃 s:2,d:2,m:2,i:2,l:2
				AllTraits["just"],      // 正直 s:2,l:1
				AllTraits["stubborn"],  // 固执 s:1,d:-1
			},
			PropertyTypeIntrigue: {
				AllTraits["shrewd"],    // 精明 i:2,d:2,m:2,s:2,l:2
				AllTraits["deceitful"], // 狡诈 i:3,d:-2
				AllTraits["ambitious"], // 野心勃勃 i:2,d:2,m:2,s:2,l:2
				AllTraits["paranoid"],  // 多疑 i:2,d:-1
				AllTraits["cynical"],   // 愤世嫉俗 i:2
				AllTraits["cruel"],     // 残暴 i:1,d:-1
				AllTraits["arbitrary"], // 专断 i:1,s:-2,l:-1
			},
			PropertyTypeLearning: {
				AllTraits["shrewd"],    // 精明 l:2,d:2,m:2,s:2,i:2
				AllTraits["ambitious"], // 野心勃勃 l:2,d:2,m:2,s:2,i:2
				AllTraits["physician"], // 著名医师 l:2,d:1
				AllTraits["erudite"],   // 博学 l:2
				AllTraits["just"],      // 正直 l:1,s:2
			},
		}
	}

	lr := ck2nebula.GetAllLeaderTraits(ck2nebula.SPACE)

	if lr.Ok {
		LeaderTraits = lr.Data
	}

	vr := ck2nebula.GetAllVirtueTraits(ck2nebula.SPACE)

	if vr.Ok {
		VirtueTraits = vr.Data
	}

	cr := ck2nebula.GetAllChildHoodTraits(ck2nebula.SPACE)

	if cr.Ok {
		ChildhoodTraits = cr.Data
	}
}

var EducationsTrait = map[string]PropertyType{
	"grey_eminence":          PropertyTypeDiplomacy,   //幕后操控人9
	"charismatic_negotiator": PropertyTypeDiplomacy,   //魅力非凡的说客6
	"underhanded_rogue":      PropertyTypeDiplomacy,   //狡猾的无赖3
	"naive_appeaser":         PropertyTypeDiplomacy,   //天真的外交家1
	"brilliant_strategist":   PropertyTypeMartial,     //天才战略家9
	"skilled_tactician":      PropertyTypeMartial,     //优秀战术家6
	"tough_soldier":          PropertyTypeMartial,     //坚强的战士3
	"misguided_warrior":      PropertyTypeMartial,     //鲁莽的战士1
	"fortune_builder":        PropertyTypeStewardship, //财富创造者6
	"thrifty_clerk":          PropertyTypeStewardship, //节俭的职员3
	"indulgent_wastrel":      PropertyTypeStewardship, //放荡的浪子1
	"midas_touched":          PropertyTypeStewardship, //点石成金者9
	"elusive_shadow":         PropertyTypeIntrigue,    //难以捉摸的影子9
	"intricate_webweaver":    PropertyTypeIntrigue,    //暗中策划者6
	"flamboyant_schemer":     PropertyTypeIntrigue,    //浮夸的谋士3
	"amateurish_plotter":     PropertyTypeIntrigue,    //业余阴谋家1
	"mastermind_theologian":  PropertyTypeLearning,    //神学的大师9
	"scholarly_theologian":   PropertyTypeLearning,    //渊博的神学家6
	"martial_cleric":         PropertyTypeLearning,    //尽职的教士3
	"detached_priest":        PropertyTypeLearning,    //清高的圣职1
}

func GetPropertyTypeByEductionTrait(trait *ck2nebula.Trait) PropertyType {
	if trait == nil || !trait.Education {
		return PropertyTypeNone
	}

	if t, ok := EducationsTrait[trait.Code]; ok {
		return t
	}

	return PropertyTypeNone
}

func BuildPeopleTrait(space *nebulagolang.Space, people *ck2nebula.People, modifiers []*PropertyModifier) (map[string]*ck2nebula.Trait, []CK2Commands.ScriptGenerator) {
	traits, scriptGenerators := RemoveBadTraitsWithNebula(space, people)

	traits, edusg := BuildPeopleEductionTrait(people, traits, modifiers)

	scriptGenerators = append(scriptGenerators, edusg...)

	var chance = 5

	switch modifiers[0].Talent {
	case TalentGenius:
		if traits[GeniusTrait.Code] == nil {
			scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(GeniusTrait))
			fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, GeniusTrait.Name)
			traits[GeniusTrait.Code] = GeniusTrait
		}
		chance = 3
	case TalentQuick:
		if traits[QuickTrait.Code] == nil {
			scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(QuickTrait))
			fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, QuickTrait.Name)
			traits[QuickTrait.Code] = QuickTrait
		}
		chance = 4
	}

	for _, trait := range CommonGoodTraits {
		if rand.Intn(chance) == 0 {
			if traits[trait.Code] == nil {
				scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(trait))
				fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, trait.Name)
				traits[trait.Code] = trait
			}
		}
	}

	for _, trait := range VirtueTraits {
		if rand.Intn(chance) == 0 {
			if traits[trait.Code] == nil {
				scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(trait))
				fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, trait.Name)
				traits[trait.Code] = trait
			}
		}
	}

	if modifiers[0].IsEducationTrait {
		lifeStyletraits := LifeStyleTraitsByPropertyType[modifiers[0].Property]

		lifeStyleTrait := lifeStyletraits[rand.Intn(len(lifeStyletraits))]

		if lifeStyleTrait.Code == "seducer" {
			if people.IsFemale {
				lifeStyleTrait = AllTraits["seductress"]
			}
		} else if lifeStyleTrait.Code == "seductress" {
			if !people.IsFemale {
				lifeStyleTrait = AllTraits["seducer"]
			}
		}

		if traits[lifeStyleTrait.Code] == nil {
			scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(lifeStyleTrait))
			fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, lifeStyleTrait.Name)
			traits[lifeStyleTrait.Code] = lifeStyleTrait
		}

		goodTraits := GoodTraitsByPropertyType[modifiers[0].Property]
		for _, trait := range goodTraits {
			if rand.Intn(chance) == 0 {
				if traits[trait.Code] == nil {
					scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(trait))
					fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, trait.Name)
					traits[trait.Code] = trait
				}
			}
		}

		if modifiers[0].Property == PropertyTypeMartial {
			i := 0
			for _, trait := range LeaderTraits {
				if rand.Intn(chance+i*2) == 0 {
					if traits[trait.Code] == nil {
						scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(trait))
						fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, trait.Name)
						traits[trait.Code] = trait
					}
				}
			}
		}
	} else {
		for _, trait := range ChildhoodTraits {
			if rand.Intn(chance) == 0 {
				if traits[trait.Code] == nil {
					scriptGenerators = append(scriptGenerators, NewAddTraitScriptGenerator(trait))
					fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, trait.Name)
					traits[trait.Code] = trait
				}
			}
		}
	}

	return traits, scriptGenerators
}

func BuildPeopleEductionTrait(people *ck2nebula.People, traits map[string]*ck2nebula.Trait, modifiers []*PropertyModifier) (map[string]*ck2nebula.Trait, []CK2Commands.ScriptGenerator) {
	var newEducationTrait *ck2nebula.Trait
	var oldEducationTrait *ck2nebula.Trait

	for _, trait := range traits {
		if trait.Education {
			oldEducationTrait = trait
			break
		}
	}

	if oldEducationTrait == nil {
		return traits, nil
	}

	generators := make([]CK2Commands.ScriptGenerator, 0)

	switch modifiers[0].Property {
	case PropertyTypeDiplomacy:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "grey_eminence" {
				return traits, generators
			}

			newEducationTrait = AllTraits["grey_eminence"]
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "grey_eminence" || oldEducationTrait.Code == "charismatic_negotiator" {
				return traits, generators
			}

			newEducationTrait = AllTraits["charismatic_negotiator"]
		} else {
			if oldEducationTrait.Code == "grey_eminence" || oldEducationTrait.Code == "charismatic_negotiator" || oldEducationTrait.Code == "underhanded_rogue" {
				return traits, generators
			}

			newEducationTrait = AllTraits["underhanded_rogue"]
		}
	case PropertyTypeMartial:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "brilliant_strategist" {
				return traits, generators
			}

			newEducationTrait = AllTraits["brilliant_strategist"]
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "brilliant_strategist" || oldEducationTrait.Code == "skilled_tactician" {
				return traits, generators
			}

			newEducationTrait = AllTraits["skilled_tactician"]
		} else {
			if oldEducationTrait.Code == "brilliant_strategist" || oldEducationTrait.Code == "skilled_tactician" || oldEducationTrait.Code == "tough_soldier" {
				return traits, generators
			}

			newEducationTrait = AllTraits["tough_soldier"]
		}
	case PropertyTypeStewardship:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "midas_touched" {
				return traits, generators
			}

			newEducationTrait = AllTraits["midas_touched"]
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "midas_touched" || oldEducationTrait.Code == "thrifty_clerk" {
				return traits, generators
			}

			newEducationTrait = AllTraits["thrifty_clerk"]

			return traits, generators
		} else {
			if oldEducationTrait.Code == "midas_touched" || oldEducationTrait.Code == "thrifty_clerk" || oldEducationTrait.Code == "fortune_builder" {
				return traits, generators
			}

			newEducationTrait = AllTraits["fortune_builder"]
		}
	case PropertyTypeIntrigue:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "elusive_shadow" {
				return traits, generators
			}

			newEducationTrait = AllTraits["elusive_shadow"]
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "elusive_shadow" || oldEducationTrait.Code == "intricate_webweaver" {
				return traits, generators
			}

			newEducationTrait = AllTraits["intricate_webweaver"]
		} else {
			if oldEducationTrait.Code == "elusive_shadow" || oldEducationTrait.Code == "intricate_webweaver" || oldEducationTrait.Code == "flamboyant_schemer" {
				return traits, generators
			}

			newEducationTrait = AllTraits["flamboyant_schemer"]
		}
	case PropertyTypeLearning:
		if modifiers[0].Talent == TalentGenius {
			if oldEducationTrait.Code == "mastermind_theologian" {
				return traits, generators
			}

			newEducationTrait = AllTraits["mastermind_theologian"]
		} else if modifiers[0].Talent == TalentQuick {
			if oldEducationTrait.Code == "mastermind_theologian" || oldEducationTrait.Code == "scholarly_theologian" {
				return traits, generators
			}

			newEducationTrait = AllTraits["scholarly_theologian"]
		} else {
			if oldEducationTrait.Code == "mastermind_theologian" || oldEducationTrait.Code == "scholarly_theologian" || oldEducationTrait.Code == "martial_cleric" {
				return traits, generators
			}

			newEducationTrait = AllTraits["martial_cleric"]
		}

		fmt.Printf("%s.%s remove trait %s\n", people.DynastyName, people.Name, oldEducationTrait.Name)
		generators = append(generators, NewRemoveTraitScriptGenerator(oldEducationTrait))
		delete(traits, oldEducationTrait.Code)

		fmt.Printf("%s.%s add trait %s\n", people.DynastyName, people.Name, newEducationTrait.Name)
		generators = append(generators, NewAddTraitScriptGenerator(newEducationTrait))
		traits[newEducationTrait.Code] = newEducationTrait
	}

	return traits, generators
}
