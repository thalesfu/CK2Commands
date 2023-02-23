package sens

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SensCounty interface {
	feud.County
	BChateaulandon朗东堡() feud.Barony
	BJoigny茹瓦尼() feud.Barony
	BMontargis蒙塔日() feud.Barony
	BMontereau蒙特罗() feud.Barony
	BNemours内穆尔() feud.Barony
	BNogentsurseine塞纳河畔诺让() feud.Barony
	BSens桑斯() feud.Barony
	BVilleneuveleroi鲁瓦新城() feud.Barony
}

type 桑斯SensCounty struct {
	feud.BaseCounty
	朗东堡Chateaulandon     feud.Barony
	茹瓦尼Joigny            feud.Barony
	蒙塔日Montargis         feud.Barony
	蒙特罗Montereau         feud.Barony
	内穆尔Nemours           feud.Barony
	塞纳河畔诺让Nogentsurseine feud.Barony
	桑斯Sens               feud.Barony
	鲁瓦新城Villeneuveleroi  feud.Barony
}

func (c *桑斯SensCounty) BChateaulandon朗东堡() feud.Barony {
	return c.朗东堡Chateaulandon
}

func (c *桑斯SensCounty) BJoigny茹瓦尼() feud.Barony {
	return c.茹瓦尼Joigny
}

func (c *桑斯SensCounty) BMontargis蒙塔日() feud.Barony {
	return c.蒙塔日Montargis
}

func (c *桑斯SensCounty) BMontereau蒙特罗() feud.Barony {
	return c.蒙特罗Montereau
}

func (c *桑斯SensCounty) BNemours内穆尔() feud.Barony {
	return c.内穆尔Nemours
}

func (c *桑斯SensCounty) BNogentsurseine塞纳河畔诺让() feud.Barony {
	return c.塞纳河畔诺让Nogentsurseine
}

func (c *桑斯SensCounty) BSens桑斯() feud.Barony {
	return c.桑斯Sens
}

func (c *桑斯SensCounty) BVilleneuveleroi鲁瓦新城() feud.Barony {
	return c.鲁瓦新城Villeneuveleroi
}

var CSens桑斯 SensCounty = &桑斯SensCounty{}

func init() {
	f := CSens桑斯.(*桑斯SensCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "131",
		Title:     "sens",
		TitleName: "桑斯",
		TitleCode: "c_sens",
		Baronies:  map[string]feud.Barony{},
	}

	f.朗东堡Chateaulandon = BChateaulandon朗东堡
	f.朗东堡Chateaulandon.SetParent(f)

	f.茹瓦尼Joigny = BJoigny茹瓦尼
	f.茹瓦尼Joigny.SetParent(f)

	f.蒙塔日Montargis = BMontargis蒙塔日
	f.蒙塔日Montargis.SetParent(f)

	f.蒙特罗Montereau = BMontereau蒙特罗
	f.蒙特罗Montereau.SetParent(f)

	f.内穆尔Nemours = BNemours内穆尔
	f.内穆尔Nemours.SetParent(f)

	f.塞纳河畔诺让Nogentsurseine = BNogentsurseine塞纳河畔诺让
	f.塞纳河畔诺让Nogentsurseine.SetParent(f)

	f.桑斯Sens = BSens桑斯
	f.桑斯Sens.SetParent(f)

	f.鲁瓦新城Villeneuveleroi = BVilleneuveleroi鲁瓦新城
	f.鲁瓦新城Villeneuveleroi.SetParent(f)

}
