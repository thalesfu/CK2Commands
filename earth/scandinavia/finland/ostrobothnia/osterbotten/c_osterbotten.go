package osterbotten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OsterbottenCounty interface {
	feud.County
	BKalajoki卡拉约基() feud.Barony
	BKaskinen卡斯基宁() feud.Barony
	BKorsholm穆斯塔萨里() feud.Barony
	BKristinestad克里斯蒂南考蓬基() feud.Barony
	BLiminka利明卡() feud.Barony
	BVaasa瓦萨() feud.Barony
	BVeteli韦泰利() feud.Barony
}

type 东博滕OsterbottenCounty struct {
	feud.BaseCounty
	卡拉约基Kalajoki         feud.Barony
	卡斯基宁Kaskinen         feud.Barony
	穆斯塔萨里Korsholm        feud.Barony
	克里斯蒂南考蓬基Kristinestad feud.Barony
	利明卡Liminka           feud.Barony
	瓦萨Vaasa              feud.Barony
	韦泰利Veteli            feud.Barony
}

func (c *东博滕OsterbottenCounty) BKalajoki卡拉约基() feud.Barony {
	return c.卡拉约基Kalajoki
}

func (c *东博滕OsterbottenCounty) BKaskinen卡斯基宁() feud.Barony {
	return c.卡斯基宁Kaskinen
}

func (c *东博滕OsterbottenCounty) BKorsholm穆斯塔萨里() feud.Barony {
	return c.穆斯塔萨里Korsholm
}

func (c *东博滕OsterbottenCounty) BKristinestad克里斯蒂南考蓬基() feud.Barony {
	return c.克里斯蒂南考蓬基Kristinestad
}

func (c *东博滕OsterbottenCounty) BLiminka利明卡() feud.Barony {
	return c.利明卡Liminka
}

func (c *东博滕OsterbottenCounty) BVaasa瓦萨() feud.Barony {
	return c.瓦萨Vaasa
}

func (c *东博滕OsterbottenCounty) BVeteli韦泰利() feud.Barony {
	return c.韦泰利Veteli
}

var COsterbotten东博滕 OsterbottenCounty = &东博滕OsterbottenCounty{}

func init() {
	f := COsterbotten东博滕.(*东博滕OsterbottenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "385",
		Title:     "osterbotten",
		TitleName: "东博滕",
		TitleCode: "c_osterbotten",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡拉约基Kalajoki = BKalajoki卡拉约基
	f.卡拉约基Kalajoki.SetParent(f)

	f.卡斯基宁Kaskinen = BKaskinen卡斯基宁
	f.卡斯基宁Kaskinen.SetParent(f)

	f.穆斯塔萨里Korsholm = BKorsholm穆斯塔萨里
	f.穆斯塔萨里Korsholm.SetParent(f)

	f.克里斯蒂南考蓬基Kristinestad = BKristinestad克里斯蒂南考蓬基
	f.克里斯蒂南考蓬基Kristinestad.SetParent(f)

	f.利明卡Liminka = BLiminka利明卡
	f.利明卡Liminka.SetParent(f)

	f.瓦萨Vaasa = BVaasa瓦萨
	f.瓦萨Vaasa.SetParent(f)

	f.韦泰利Veteli = BVeteli韦泰利
	f.韦泰利Veteli.SetParent(f)

}
