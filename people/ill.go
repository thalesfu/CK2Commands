package people

import (
	"bufio"
	"fmt"
	"github.com/thalesfu/ck2nebula"
	"github.com/thalesfu/nebulagolang"
)

var IllTraitMap = map[string]string{
	"severely_injured":    "身受重伤",
	"lunatic":             "精神错乱",
	"leper":               "麻风病人",
	"incapable":           "无能",
	"fatigue":             "疲乏",
	"one_legged":          "独腿",
	"sick_incapable":      "",
	"depressed":           "沮丧消沉",
	"infirm":              "体弱多病",
	"stressed":            "紧张不安",
	"maimed":              "残废",
	"mangled":             "血肉模糊",
	"hard_pregnancy":      "危险妊娠",
	"pregnancy_finishing": "晚期妊娠",
	"disfigured":          "毁容",
	"troubled_pregnancy":  "高危妊娠",
	"possessed":           "恶魔附身",
	"wounded":             "受伤",
	"drunkard":            "酒鬼",
	"one_eyed":            "独眼",
	"one_handed":          "独手",
	"food_poisoning":      "食物中毒",
	"has_bubonic_plague":  "黑死病",
	"has_aztec_disease":   "阿兹特克梅毒",
	"cancer":              "癌症",
	"syphilitic":          "梅毒",
	"has_typhus":          "露营热",
	"cramps":              "抽筋",
	"fever":               "发热",
	"gout":                "痛风",
	"dancing_plague":      "舞蹈瘟疫",
	"pneumonic":           "肺炎",
	"has_typhoid_fever":   "慢性热",
	"flu":                 "流感",
	"diarrhea":            "腹泻",
	"dysentery":           "痢疾",
	"ill":                 "身患疾病",
	"has_tuberculosis":    "肺痨",
	"has_small_pox":       "天花",
	"sickly":              "多病",
	"malaise":             "不适",
	"vomiting":            "呕吐",
	"abdominal_pain":      "腹痛",
	"cough":               "咳嗽",
	"headache":            "头痛",
	"has_measles":         "麻疹",
	"rabies":              "狂犬病",
	"chest_pain":          "胸痛",
	"infection":           "感染",
	"rash":                "皮疹",
	"scurvy":              "坏血病",
}

func cureIlls(writer *bufio.Writer, peopleId int) {
	for trait, _ := range IllTraitMap {
		writeRemoveTrait(writer, trait, peopleId)
	}
}

func CurePeopleIll(peopleId ...int) {
	var functions = []buildFunc{
		cureIlls,
	}

	buildPeople("cureill", functions, peopleId...)
}

func cureIllsWithNebula(writer *bufio.Writer, space *nebulagolang.Space, people *ck2nebula.People) {
	result := people.GetTraits(space)

	if result.Ok {
		for _, trait := range result.Data {
			if trait.IsHealth || trait.IsIllness || trait.Blinding {
				fmt.Printf("%s.%s remove trait %s\n", people.DynastyName, people.Name, trait.Name)
				writeRemoveTrait(writer, trait.Code, people.ID)
			}
		}
	}
}

func CurePeopleWithNebula(space *nebulagolang.Space, people ...*ck2nebula.People) {
	var functions = []buildFuncWithNebula{
		cureIllsWithNebula,
	}

	buildPeopleWithNebula("cure", functions, space, people...)
}
