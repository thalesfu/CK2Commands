package thomond

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ThomondCounty interface {
	feud.County
	BAdare阿代尔() feud.Barony
	BAskeaton阿斯基顿() feud.Barony
	BBunratty班拉蒂() feud.Barony
	BEmly紫杉境() feud.Barony
	BEnnis恩尼斯() feud.Barony
	BKilfenora基尔费诺拉() feud.Barony
	BKillaloe基拉卢() feud.Barony
	BLimerick利默里克() feud.Barony
}

type 托蒙德ThomondCounty struct {
	feud.BaseCounty
	阿代尔Adare       feud.Barony
	阿斯基顿Askeaton   feud.Barony
	班拉蒂Bunratty    feud.Barony
	紫杉境Emly        feud.Barony
	恩尼斯Ennis       feud.Barony
	基尔费诺拉Kilfenora feud.Barony
	基拉卢Killaloe    feud.Barony
	利默里克Limerick   feud.Barony
}

func (c *托蒙德ThomondCounty) BAdare阿代尔() feud.Barony {
	return c.阿代尔Adare
}

func (c *托蒙德ThomondCounty) BAskeaton阿斯基顿() feud.Barony {
	return c.阿斯基顿Askeaton
}

func (c *托蒙德ThomondCounty) BBunratty班拉蒂() feud.Barony {
	return c.班拉蒂Bunratty
}

func (c *托蒙德ThomondCounty) BEmly紫杉境() feud.Barony {
	return c.紫杉境Emly
}

func (c *托蒙德ThomondCounty) BEnnis恩尼斯() feud.Barony {
	return c.恩尼斯Ennis
}

func (c *托蒙德ThomondCounty) BKilfenora基尔费诺拉() feud.Barony {
	return c.基尔费诺拉Kilfenora
}

func (c *托蒙德ThomondCounty) BKillaloe基拉卢() feud.Barony {
	return c.基拉卢Killaloe
}

func (c *托蒙德ThomondCounty) BLimerick利默里克() feud.Barony {
	return c.利默里克Limerick
}

var CThomond托蒙德 ThomondCounty = &托蒙德ThomondCounty{}

func init() {
	f := CThomond托蒙德.(*托蒙德ThomondCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "13",
		Title:     "thomond",
		TitleName: "托蒙德",
		TitleCode: "c_thomond",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿代尔Adare = BAdare阿代尔
	f.阿代尔Adare.SetParent(f)

	f.阿斯基顿Askeaton = BAskeaton阿斯基顿
	f.阿斯基顿Askeaton.SetParent(f)

	f.班拉蒂Bunratty = BBunratty班拉蒂
	f.班拉蒂Bunratty.SetParent(f)

	f.紫杉境Emly = BEmly紫杉境
	f.紫杉境Emly.SetParent(f)

	f.恩尼斯Ennis = BEnnis恩尼斯
	f.恩尼斯Ennis.SetParent(f)

	f.基尔费诺拉Kilfenora = BKilfenora基尔费诺拉
	f.基尔费诺拉Kilfenora.SetParent(f)

	f.基拉卢Killaloe = BKillaloe基拉卢
	f.基拉卢Killaloe.SetParent(f)

	f.利默里克Limerick = BLimerick利默里克
	f.利默里克Limerick.SetParent(f)

}
