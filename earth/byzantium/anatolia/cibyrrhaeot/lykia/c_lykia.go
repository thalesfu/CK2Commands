package lykia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LykiaCounty interface {
	feud.County
	BHalikarnassos哈利卡尔那索斯() feud.Barony
	BKibyra基比拉() feud.Barony
	BLimyra利米拉() feud.Barony
	BMylasa米拉萨() feud.Barony
	BMyra米拉() feud.Barony
	BPatara帕塔拉() feud.Barony
	BPhaselis法塞利斯() feud.Barony
	BTelmissos特尔梅索斯() feud.Barony
}

type 吕基亚LykiaCounty struct {
	feud.BaseCounty
	哈利卡尔那索斯Halikarnassos feud.Barony
	基比拉Kibyra            feud.Barony
	利米拉Limyra            feud.Barony
	米拉萨Mylasa            feud.Barony
	米拉Myra               feud.Barony
	帕塔拉Patara            feud.Barony
	法塞利斯Phaselis         feud.Barony
	特尔梅索斯Telmissos       feud.Barony
}

func (c *吕基亚LykiaCounty) BHalikarnassos哈利卡尔那索斯() feud.Barony {
	return c.哈利卡尔那索斯Halikarnassos
}

func (c *吕基亚LykiaCounty) BKibyra基比拉() feud.Barony {
	return c.基比拉Kibyra
}

func (c *吕基亚LykiaCounty) BLimyra利米拉() feud.Barony {
	return c.利米拉Limyra
}

func (c *吕基亚LykiaCounty) BMylasa米拉萨() feud.Barony {
	return c.米拉萨Mylasa
}

func (c *吕基亚LykiaCounty) BMyra米拉() feud.Barony {
	return c.米拉Myra
}

func (c *吕基亚LykiaCounty) BPatara帕塔拉() feud.Barony {
	return c.帕塔拉Patara
}

func (c *吕基亚LykiaCounty) BPhaselis法塞利斯() feud.Barony {
	return c.法塞利斯Phaselis
}

func (c *吕基亚LykiaCounty) BTelmissos特尔梅索斯() feud.Barony {
	return c.特尔梅索斯Telmissos
}

var CLykia吕基亚 LykiaCounty = &吕基亚LykiaCounty{}

func init() {
	f := CLykia吕基亚.(*吕基亚LykiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "747",
		Title:     "lykia",
		TitleName: "吕基亚",
		TitleCode: "c_lykia",
		Baronies:  map[string]feud.Barony{},
	}

	f.哈利卡尔那索斯Halikarnassos = BHalikarnassos哈利卡尔那索斯
	f.哈利卡尔那索斯Halikarnassos.SetParent(f)

	f.基比拉Kibyra = BKibyra基比拉
	f.基比拉Kibyra.SetParent(f)

	f.利米拉Limyra = BLimyra利米拉
	f.利米拉Limyra.SetParent(f)

	f.米拉萨Mylasa = BMylasa米拉萨
	f.米拉萨Mylasa.SetParent(f)

	f.米拉Myra = BMyra米拉
	f.米拉Myra.SetParent(f)

	f.帕塔拉Patara = BPatara帕塔拉
	f.帕塔拉Patara.SetParent(f)

	f.法塞利斯Phaselis = BPhaselis法塞利斯
	f.法塞利斯Phaselis.SetParent(f)

	f.特尔梅索斯Telmissos = BTelmissos特尔梅索斯
	f.特尔梅索斯Telmissos.SetParent(f)

}
