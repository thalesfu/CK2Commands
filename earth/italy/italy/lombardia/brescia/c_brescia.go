package brescia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BresciaCounty interface {
	feud.County
	BBergamo贝尔加莫() feud.Barony
	BBrescia布雷西亚() feud.Barony
	BCastiglione卡斯蒂廖内() feud.Barony
	BLumezzane卢梅扎内() feud.Barony
	BMontichiari蒙蒂基亚里() feud.Barony
	BPeschiera佩斯基耶拉() feud.Barony
	BSalo萨洛() feud.Barony
	BTreviglio特雷维利奥() feud.Barony
}

type 布雷西亚BresciaCounty struct {
	feud.BaseCounty
	贝尔加莫Bergamo      feud.Barony
	布雷西亚Brescia      feud.Barony
	卡斯蒂廖内Castiglione feud.Barony
	卢梅扎内Lumezzane    feud.Barony
	蒙蒂基亚里Montichiari feud.Barony
	佩斯基耶拉Peschiera   feud.Barony
	萨洛Salo           feud.Barony
	特雷维利奥Treviglio   feud.Barony
}

func (c *布雷西亚BresciaCounty) BBergamo贝尔加莫() feud.Barony {
	return c.贝尔加莫Bergamo
}

func (c *布雷西亚BresciaCounty) BBrescia布雷西亚() feud.Barony {
	return c.布雷西亚Brescia
}

func (c *布雷西亚BresciaCounty) BCastiglione卡斯蒂廖内() feud.Barony {
	return c.卡斯蒂廖内Castiglione
}

func (c *布雷西亚BresciaCounty) BLumezzane卢梅扎内() feud.Barony {
	return c.卢梅扎内Lumezzane
}

func (c *布雷西亚BresciaCounty) BMontichiari蒙蒂基亚里() feud.Barony {
	return c.蒙蒂基亚里Montichiari
}

func (c *布雷西亚BresciaCounty) BPeschiera佩斯基耶拉() feud.Barony {
	return c.佩斯基耶拉Peschiera
}

func (c *布雷西亚BresciaCounty) BSalo萨洛() feud.Barony {
	return c.萨洛Salo
}

func (c *布雷西亚BresciaCounty) BTreviglio特雷维利奥() feud.Barony {
	return c.特雷维利奥Treviglio
}

var CBrescia布雷西亚 BresciaCounty = &布雷西亚BresciaCounty{}

func init() {
	f := CBrescia布雷西亚.(*布雷西亚BresciaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "318",
		Title:     "brescia",
		TitleName: "布雷西亚",
		TitleCode: "c_brescia",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尔加莫Bergamo = BBergamo贝尔加莫
	f.贝尔加莫Bergamo.SetParent(f)

	f.布雷西亚Brescia = BBrescia布雷西亚
	f.布雷西亚Brescia.SetParent(f)

	f.卡斯蒂廖内Castiglione = BCastiglione卡斯蒂廖内
	f.卡斯蒂廖内Castiglione.SetParent(f)

	f.卢梅扎内Lumezzane = BLumezzane卢梅扎内
	f.卢梅扎内Lumezzane.SetParent(f)

	f.蒙蒂基亚里Montichiari = BMontichiari蒙蒂基亚里
	f.蒙蒂基亚里Montichiari.SetParent(f)

	f.佩斯基耶拉Peschiera = BPeschiera佩斯基耶拉
	f.佩斯基耶拉Peschiera.SetParent(f)

	f.萨洛Salo = BSalo萨洛
	f.萨洛Salo.SetParent(f)

	f.特雷维利奥Treviglio = BTreviglio特雷维利奥
	f.特雷维利奥Treviglio.SetParent(f)

}
