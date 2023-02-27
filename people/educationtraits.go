package people

import (
	"bufio"
)

var educationTraits = []string{
	"naive_appeaser",         // 天真的外交家
	"underhanded_rogue",      // 狡猾的无赖
	"charismatic_negotiator", // 魅力非凡的说客
	"grey_eminence",          // 幕后操控人
	"misguided_warrior",      // 鲁莽的战士
	"tough_soldier",          // 坚强的战士
	"skilled_tactician",      // 优秀战术家
	"brilliant_strategist",   // 天才战略家
	"indulgent_wastrel",      // 放荡的浪子
	"thrifty_clerk",          // 节俭的职员
	"fortune_builder",        // 财富创造者
	"midas_touched",          // 点石成金者
	"amateurish_plotter",     // 业余阴谋家
	"flamboyant_schemer",     // 浮夸的谋士
	"intricate_webweaver",    // 暗中操纵着
	"elusive_shadow",         // 难以捉摸的影子
	"detached_priest",        // 清高的牧师
	"martial_cleric",         // 尽职的教士
	"scholarly_theologian",   // 渊博的神学家
	"mastermind_theologian",  // 神学的大师
}

func removeEducationTraits(writer *bufio.Writer, peopleId int) {
	for _, trait := range educationTraits {
		writeRemoveTrait(writer, trait, peopleId)
	}
}
