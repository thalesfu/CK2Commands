package cairo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CairoCounty interface {
	feud.County
	BCairo开罗() feud.Barony
	BFustat福斯塔特() feud.Barony
	BHeliopolis赫利奥波利斯() feud.Barony
	BHelwan赫勒万() feud.Barony
	BMaadi马阿迪() feud.Barony
	BMemphis孟菲斯() feud.Barony
	BMerimdabenisalama蒙尔马迪贝尼所罗蒙() feud.Barony
	BTekkekyahudiyya泰勒阿呼蒂亚() feud.Barony
}

type 开罗CairoCounty struct {
	feud.BaseCounty
	开罗Cairo                    feud.Barony
	福斯塔特Fustat                 feud.Barony
	赫利奥波利斯Heliopolis           feud.Barony
	赫勒万Helwan                  feud.Barony
	马阿迪Maadi                   feud.Barony
	孟菲斯Memphis                 feud.Barony
	蒙尔马迪贝尼所罗蒙Merimdabenisalama feud.Barony
	泰勒阿呼蒂亚Tekkekyahudiyya      feud.Barony
}

func (c *开罗CairoCounty) BCairo开罗() feud.Barony {
	return c.开罗Cairo
}

func (c *开罗CairoCounty) BFustat福斯塔特() feud.Barony {
	return c.福斯塔特Fustat
}

func (c *开罗CairoCounty) BHeliopolis赫利奥波利斯() feud.Barony {
	return c.赫利奥波利斯Heliopolis
}

func (c *开罗CairoCounty) BHelwan赫勒万() feud.Barony {
	return c.赫勒万Helwan
}

func (c *开罗CairoCounty) BMaadi马阿迪() feud.Barony {
	return c.马阿迪Maadi
}

func (c *开罗CairoCounty) BMemphis孟菲斯() feud.Barony {
	return c.孟菲斯Memphis
}

func (c *开罗CairoCounty) BMerimdabenisalama蒙尔马迪贝尼所罗蒙() feud.Barony {
	return c.蒙尔马迪贝尼所罗蒙Merimdabenisalama
}

func (c *开罗CairoCounty) BTekkekyahudiyya泰勒阿呼蒂亚() feud.Barony {
	return c.泰勒阿呼蒂亚Tekkekyahudiyya
}

var CCairo开罗 CairoCounty = &开罗CairoCounty{}

func init() {
	f := CCairo开罗.(*开罗CairoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "796",
		Title:     "cairo",
		TitleName: "开罗",
		TitleCode: "c_cairo",
		Baronies:  map[string]feud.Barony{},
	}

	f.开罗Cairo = BCairo开罗
	f.开罗Cairo.SetParent(f)

	f.福斯塔特Fustat = BFustat福斯塔特
	f.福斯塔特Fustat.SetParent(f)

	f.赫利奥波利斯Heliopolis = BHeliopolis赫利奥波利斯
	f.赫利奥波利斯Heliopolis.SetParent(f)

	f.赫勒万Helwan = BHelwan赫勒万
	f.赫勒万Helwan.SetParent(f)

	f.马阿迪Maadi = BMaadi马阿迪
	f.马阿迪Maadi.SetParent(f)

	f.孟菲斯Memphis = BMemphis孟菲斯
	f.孟菲斯Memphis.SetParent(f)

	f.蒙尔马迪贝尼所罗蒙Merimdabenisalama = BMerimdabenisalama蒙尔马迪贝尼所罗蒙
	f.蒙尔马迪贝尼所罗蒙Merimdabenisalama.SetParent(f)

	f.泰勒阿呼蒂亚Tekkekyahudiyya = BTekkekyahudiyya泰勒阿呼蒂亚
	f.泰勒阿呼蒂亚Tekkekyahudiyya.SetParent(f)

}
