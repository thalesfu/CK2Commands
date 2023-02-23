package orvieto

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OrvietoCounty interface {
	feud.County
	BAlviano阿尔维亚诺() feud.Barony
	BAmelia阿梅利亚() feud.Barony
	BBaschi巴斯基() feud.Barony
	BCiconia奇科尼亚() feud.Barony
	BMontecastrilli蒙特卡斯特里利() feud.Barony
	BNarni纳尔尼() feud.Barony
	BOrvieto奥尔维耶托() feud.Barony
	BOtricoli奥特里科利() feud.Barony
}

type 奥尔维耶托OrvietoCounty struct {
	feud.BaseCounty
	阿尔维亚诺Alviano          feud.Barony
	阿梅利亚Amelia            feud.Barony
	巴斯基Baschi             feud.Barony
	奇科尼亚Ciconia           feud.Barony
	蒙特卡斯特里利Montecastrilli feud.Barony
	纳尔尼Narni              feud.Barony
	奥尔维耶托Orvieto          feud.Barony
	奥特里科利Otricoli         feud.Barony
}

func (c *奥尔维耶托OrvietoCounty) BAlviano阿尔维亚诺() feud.Barony {
	return c.阿尔维亚诺Alviano
}

func (c *奥尔维耶托OrvietoCounty) BAmelia阿梅利亚() feud.Barony {
	return c.阿梅利亚Amelia
}

func (c *奥尔维耶托OrvietoCounty) BBaschi巴斯基() feud.Barony {
	return c.巴斯基Baschi
}

func (c *奥尔维耶托OrvietoCounty) BCiconia奇科尼亚() feud.Barony {
	return c.奇科尼亚Ciconia
}

func (c *奥尔维耶托OrvietoCounty) BMontecastrilli蒙特卡斯特里利() feud.Barony {
	return c.蒙特卡斯特里利Montecastrilli
}

func (c *奥尔维耶托OrvietoCounty) BNarni纳尔尼() feud.Barony {
	return c.纳尔尼Narni
}

func (c *奥尔维耶托OrvietoCounty) BOrvieto奥尔维耶托() feud.Barony {
	return c.奥尔维耶托Orvieto
}

func (c *奥尔维耶托OrvietoCounty) BOtricoli奥特里科利() feud.Barony {
	return c.奥特里科利Otricoli
}

var COrvieto奥尔维耶托 OrvietoCounty = &奥尔维耶托OrvietoCounty{}

func init() {
	f := COrvieto奥尔维耶托.(*奥尔维耶托OrvietoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "243",
		Title:     "orvieto",
		TitleName: "奥尔维耶托",
		TitleCode: "c_orvieto",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔维亚诺Alviano = BAlviano阿尔维亚诺
	f.阿尔维亚诺Alviano.SetParent(f)

	f.阿梅利亚Amelia = BAmelia阿梅利亚
	f.阿梅利亚Amelia.SetParent(f)

	f.巴斯基Baschi = BBaschi巴斯基
	f.巴斯基Baschi.SetParent(f)

	f.奇科尼亚Ciconia = BCiconia奇科尼亚
	f.奇科尼亚Ciconia.SetParent(f)

	f.蒙特卡斯特里利Montecastrilli = BMontecastrilli蒙特卡斯特里利
	f.蒙特卡斯特里利Montecastrilli.SetParent(f)

	f.纳尔尼Narni = BNarni纳尔尼
	f.纳尔尼Narni.SetParent(f)

	f.奥尔维耶托Orvieto = BOrvieto奥尔维耶托
	f.奥尔维耶托Orvieto.SetParent(f)

	f.奥特里科利Otricoli = BOtricoli奥特里科利
	f.奥特里科利Otricoli.SetParent(f)

}
