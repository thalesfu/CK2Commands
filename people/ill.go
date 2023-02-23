package people

import "bufio"

var illTraits = []string{
	"depressed",          // 沮丧消沉
	"stressed",           // 紧张不安
	"lunatic",            // 精神错乱
	"has_aztec_disease",  // 阿兹特克梅毒
	"has_typhus",         // 露营热
	"cancer",             // 癌症
	"has_tuberculosis",   // 肺痨
	"dancing_plague",     // 舞蹈瘟疫
	"dysentery",          // 痢疾
	"flu",                // 流感
	"food_poisoning",     // 食物中毒
	"gout",               // 痛风
	"syphilitic",         // 梅毒
	"has_typhoid_fever",  // 慢性热
	"ill",                // 身患疾病
	"leper",              // 麻风病人
	"has_measles",        // 麻疹
	"pneumonic",          // 肺炎
	"rabies",             // 狂犬病
	"scurvy",             // 坏血病
	"sickly",             // 多病
	"ill",                // 身患疾病
	"has_small_pox",      // 天花
	"has_bubonic_plague", // 黑死病
	"abdominal_pain",     // 腹痛
	"chest_pain",         // 胸痛
	"cough",              // 咳嗽
	"cramps",             // 抽筋
	"diarrhea",           // 腹泻
	"fatigue",            // 疲乏
	"fever",              // 发热
	"headache",           // 头痛
	"infection",          // 感染
	"malaise",            // 不适
	"rash",               // 皮疹
	"vomiting",           // 呕吐
	"lovers_pox",         // 情人的疱疹
}

func cureIlls(writer *bufio.Writer, peopleId int) {
	for _, trait := range illTraits {
		removeTrait(writer, trait, peopleId)
	}
}

func CurePeopleIll(peopleId ...int) {
	var functions = []buildFunc{
		cureIlls,
	}

	buildPeople("cureill", functions, peopleId...)
}
