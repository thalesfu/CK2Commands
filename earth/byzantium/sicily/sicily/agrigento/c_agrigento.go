package agrigento

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AgrigentoCounty interface {
	feud.County
	BAgricento吉尔真蒂() feud.Barony
	BButera布泰拉() feud.Barony
	BCaltabellotta卡尔塔贝洛塔() feud.Barony
	BGela杰拉() feud.Barony
	BLicata利卡塔() feud.Barony
	BMontallegro蒙塔莱格罗() feud.Barony
	BRaffadali拉法达利() feud.Barony
	BSanbiagioplatani圣比亚焦普拉塔尼() feud.Barony
}

type 吉尔真蒂AgrigentoCounty struct {
	feud.BaseCounty
	吉尔真蒂Agricento            feud.Barony
	布泰拉Butera                feud.Barony
	卡尔塔贝洛塔Caltabellotta      feud.Barony
	杰拉Gela                   feud.Barony
	利卡塔Licata                feud.Barony
	蒙塔莱格罗Montallegro         feud.Barony
	拉法达利Raffadali            feud.Barony
	圣比亚焦普拉塔尼Sanbiagioplatani feud.Barony
}

func (c *吉尔真蒂AgrigentoCounty) BAgricento吉尔真蒂() feud.Barony {
	return c.吉尔真蒂Agricento
}

func (c *吉尔真蒂AgrigentoCounty) BButera布泰拉() feud.Barony {
	return c.布泰拉Butera
}

func (c *吉尔真蒂AgrigentoCounty) BCaltabellotta卡尔塔贝洛塔() feud.Barony {
	return c.卡尔塔贝洛塔Caltabellotta
}

func (c *吉尔真蒂AgrigentoCounty) BGela杰拉() feud.Barony {
	return c.杰拉Gela
}

func (c *吉尔真蒂AgrigentoCounty) BLicata利卡塔() feud.Barony {
	return c.利卡塔Licata
}

func (c *吉尔真蒂AgrigentoCounty) BMontallegro蒙塔莱格罗() feud.Barony {
	return c.蒙塔莱格罗Montallegro
}

func (c *吉尔真蒂AgrigentoCounty) BRaffadali拉法达利() feud.Barony {
	return c.拉法达利Raffadali
}

func (c *吉尔真蒂AgrigentoCounty) BSanbiagioplatani圣比亚焦普拉塔尼() feud.Barony {
	return c.圣比亚焦普拉塔尼Sanbiagioplatani
}

var CAgrigento吉尔真蒂 AgrigentoCounty = &吉尔真蒂AgrigentoCounty{}

func init() {
	f := CAgrigento吉尔真蒂.(*吉尔真蒂AgrigentoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "342",
		Title:     "agrigento",
		TitleName: "吉尔真蒂",
		TitleCode: "c_agrigento",
		Baronies:  map[string]feud.Barony{},
	}

	f.吉尔真蒂Agricento = BAgricento吉尔真蒂
	f.吉尔真蒂Agricento.SetParent(f)

	f.布泰拉Butera = BButera布泰拉
	f.布泰拉Butera.SetParent(f)

	f.卡尔塔贝洛塔Caltabellotta = BCaltabellotta卡尔塔贝洛塔
	f.卡尔塔贝洛塔Caltabellotta.SetParent(f)

	f.杰拉Gela = BGela杰拉
	f.杰拉Gela.SetParent(f)

	f.利卡塔Licata = BLicata利卡塔
	f.利卡塔Licata.SetParent(f)

	f.蒙塔莱格罗Montallegro = BMontallegro蒙塔莱格罗
	f.蒙塔莱格罗Montallegro.SetParent(f)

	f.拉法达利Raffadali = BRaffadali拉法达利
	f.拉法达利Raffadali.SetParent(f)

	f.圣比亚焦普拉塔尼Sanbiagioplatani = BSanbiagioplatani圣比亚焦普拉塔尼
	f.圣比亚焦普拉塔尼Sanbiagioplatani.SetParent(f)

}
