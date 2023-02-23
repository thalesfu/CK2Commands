package messina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MessinaCounty interface {
	feud.County
	BCataratti卡塔拉蒂() feud.Barony
	BFurnari富尔纳里() feud.Barony
	BLipari利帕里() feud.Barony
	BMessina墨西拿() feud.Barony
	BSanmarcodalunzio圣马尔科达伦齐奥() feud.Barony
	BTaormina陶尔米纳() feud.Barony
	BTorregrota托雷格罗塔() feud.Barony
	BTroinia特罗伊纳() feud.Barony
}

type 墨西拿MessinaCounty struct {
	feud.BaseCounty
	卡塔拉蒂Cataratti            feud.Barony
	富尔纳里Furnari              feud.Barony
	利帕里Lipari                feud.Barony
	墨西拿Messina               feud.Barony
	圣马尔科达伦齐奥Sanmarcodalunzio feud.Barony
	陶尔米纳Taormina             feud.Barony
	托雷格罗塔Torregrota          feud.Barony
	特罗伊纳Troinia              feud.Barony
}

func (c *墨西拿MessinaCounty) BCataratti卡塔拉蒂() feud.Barony {
	return c.卡塔拉蒂Cataratti
}

func (c *墨西拿MessinaCounty) BFurnari富尔纳里() feud.Barony {
	return c.富尔纳里Furnari
}

func (c *墨西拿MessinaCounty) BLipari利帕里() feud.Barony {
	return c.利帕里Lipari
}

func (c *墨西拿MessinaCounty) BMessina墨西拿() feud.Barony {
	return c.墨西拿Messina
}

func (c *墨西拿MessinaCounty) BSanmarcodalunzio圣马尔科达伦齐奥() feud.Barony {
	return c.圣马尔科达伦齐奥Sanmarcodalunzio
}

func (c *墨西拿MessinaCounty) BTaormina陶尔米纳() feud.Barony {
	return c.陶尔米纳Taormina
}

func (c *墨西拿MessinaCounty) BTorregrota托雷格罗塔() feud.Barony {
	return c.托雷格罗塔Torregrota
}

func (c *墨西拿MessinaCounty) BTroinia特罗伊纳() feud.Barony {
	return c.特罗伊纳Troinia
}

var CMessina墨西拿 MessinaCounty = &墨西拿MessinaCounty{}

func init() {
	f := CMessina墨西拿.(*墨西拿MessinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "339",
		Title:     "messina",
		TitleName: "墨西拿",
		TitleCode: "c_messina",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡塔拉蒂Cataratti = BCataratti卡塔拉蒂
	f.卡塔拉蒂Cataratti.SetParent(f)

	f.富尔纳里Furnari = BFurnari富尔纳里
	f.富尔纳里Furnari.SetParent(f)

	f.利帕里Lipari = BLipari利帕里
	f.利帕里Lipari.SetParent(f)

	f.墨西拿Messina = BMessina墨西拿
	f.墨西拿Messina.SetParent(f)

	f.圣马尔科达伦齐奥Sanmarcodalunzio = BSanmarcodalunzio圣马尔科达伦齐奥
	f.圣马尔科达伦齐奥Sanmarcodalunzio.SetParent(f)

	f.陶尔米纳Taormina = BTaormina陶尔米纳
	f.陶尔米纳Taormina.SetParent(f)

	f.托雷格罗塔Torregrota = BTorregrota托雷格罗塔
	f.托雷格罗塔Torregrota.SetParent(f)

	f.特罗伊纳Troinia = BTroinia特罗伊纳
	f.特罗伊纳Troinia.SetParent(f)

}
