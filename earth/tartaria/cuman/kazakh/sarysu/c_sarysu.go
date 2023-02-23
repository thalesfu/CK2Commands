package sarysu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SarysuCounty interface {
	feud.County
	BAktam阿克塔姆() feud.Barony
	BKumkent库姆肯特() feud.Barony
	BSaudakent萨乌达肯特() feud.Barony
	BZhaiylma扎伊尔马() feud.Barony
	BZhanatas札纳塔斯() feud.Barony
	BZhuldyz茹尔德兹() feud.Barony
}

type 萨雷苏SarysuCounty struct {
	feud.BaseCounty
	阿克塔姆Aktam      feud.Barony
	库姆肯特Kumkent    feud.Barony
	萨乌达肯特Saudakent feud.Barony
	扎伊尔马Zhaiylma   feud.Barony
	札纳塔斯Zhanatas   feud.Barony
	茹尔德兹Zhuldyz    feud.Barony
}

func (c *萨雷苏SarysuCounty) BAktam阿克塔姆() feud.Barony {
	return c.阿克塔姆Aktam
}

func (c *萨雷苏SarysuCounty) BKumkent库姆肯特() feud.Barony {
	return c.库姆肯特Kumkent
}

func (c *萨雷苏SarysuCounty) BSaudakent萨乌达肯特() feud.Barony {
	return c.萨乌达肯特Saudakent
}

func (c *萨雷苏SarysuCounty) BZhaiylma扎伊尔马() feud.Barony {
	return c.扎伊尔马Zhaiylma
}

func (c *萨雷苏SarysuCounty) BZhanatas札纳塔斯() feud.Barony {
	return c.札纳塔斯Zhanatas
}

func (c *萨雷苏SarysuCounty) BZhuldyz茹尔德兹() feud.Barony {
	return c.茹尔德兹Zhuldyz
}

var CSarysu萨雷苏 SarysuCounty = &萨雷苏SarysuCounty{}

func init() {
	f := CSarysu萨雷苏.(*萨雷苏SarysuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1868",
		Title:     "sarysu",
		TitleName: "萨雷苏",
		TitleCode: "c_sarysu",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克塔姆Aktam = BAktam阿克塔姆
	f.阿克塔姆Aktam.SetParent(f)

	f.库姆肯特Kumkent = BKumkent库姆肯特
	f.库姆肯特Kumkent.SetParent(f)

	f.萨乌达肯特Saudakent = BSaudakent萨乌达肯特
	f.萨乌达肯特Saudakent.SetParent(f)

	f.扎伊尔马Zhaiylma = BZhaiylma扎伊尔马
	f.扎伊尔马Zhaiylma.SetParent(f)

	f.札纳塔斯Zhanatas = BZhanatas札纳塔斯
	f.札纳塔斯Zhanatas.SetParent(f)

	f.茹尔德兹Zhuldyz = BZhuldyz茹尔德兹
	f.茹尔德兹Zhuldyz.SetParent(f)

}
