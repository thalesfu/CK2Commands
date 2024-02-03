package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/CK2Commands"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
)

var badTraits = map[string]string{
	"is_fat":          "肥胖",
	"is_malnourished": "营养不良",
	"clubfooted":      "足内翻",
	"dwarf":           "侏儒",
	"harelip":         "兔唇",
	"hunchback":       "驼背",
	"imbecile":        "痴呆",
	"inbred":          "近亲缺陷",
	"lisp":            "口齿不清",
	"slow":            "迟钝",
	"stutter":         "口吃",
	"ugly":            "丑陋",
	"weak":            "虚弱",
	"feeble":          "脆弱",
	"dull":            "迟钝",
	"uncouth":         "邋遢",
	"gluttonous":      "暴食",
	"greedy":          "贪婪",
	"slothful":        "懒惰",
	"wroth":           "易怒",
	"envious":         "嫉妒",
	"proud":           "骄傲",
	"chaste":          "忠贞",
	"ambitious":       "野心勃勃",
	"arbitrary":       "专断",
	"craven":          "怯懦",
	"cruel":           "残暴",
	"cynical":         "愤世嫉俗",
	"deceitful":       "狡诈",
	"paranoid":        "多疑",
	"shy":             "害羞",
	"stubborn":        "固执",
	"trusting":        "轻信他人",
	"homosexual":      "同性恋",
	"celibate":        "独身主义",
}

func RemoveBadTraits(writer *bufio.Writer, peopleId int) {
	for trait, _ := range badTraits {
		writeRemoveTrait(writer, trait, peopleId)
	}

	cureIlls(writer, peopleId)
	cureWound(writer, peopleId)
}

func RemoveBadTraitsWithNebula(space *nebulagolang.Space, people *ck2nebula.People) (map[string]*ck2nebula.Trait, []CK2Commands.ScriptGenerator) {
	r := people.GetTraits(space)

	traitResult := make(map[string]*ck2nebula.Trait)
	scriptsGenerator := make([]CK2Commands.ScriptGenerator, 0)

	if r.Ok {
		for _, trait := range r.Data {

			if trait.Code == "pregnancy_finishing" {
				continue
			}

			_, exist := badTraits[trait.Code]
			if trait.IsHealth || trait.IsIllness || trait.Blinding || trait.Inbred || trait.Vice || trait.SexAppealOpinion < 0 || exist {
				fmt.Printf("%s.%s(%d) remove trait %s(%s)\n", people.DynastyName, people.Name, people.ID, trait.Name, trait.Code)
				scriptsGenerator = append(scriptsGenerator, NewRemoveTraitScriptGenerator(trait))
			} else {
				traitResult[trait.Code] = trait
			}
		}
	}

	return traitResult, scriptsGenerator
}

func RemoveBad(peopleId ...int) {
	var functions = []buildFunc{
		cureIlls,
		cureWound,
	}

	buildPeople("removebad", functions, peopleId...)
}
