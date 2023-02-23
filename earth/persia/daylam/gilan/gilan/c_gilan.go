package gilan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GilanCounty interface {
	feud.County
	BAstara阿斯塔拉() feud.Barony
	BGilan吉兰() feud.Barony
	BLahijan拉希詹() feud.Barony
	BMasal马萨拉() feud.Barony
	BMasouleh马苏莱() feud.Barony
	BRudkhan鲁德罕() feud.Barony
	BTalysh塔雷什() feud.Barony
}

type 吉兰GilanCounty struct {
	feud.BaseCounty
	阿斯塔拉Astara  feud.Barony
	吉兰Gilan     feud.Barony
	拉希詹Lahijan  feud.Barony
	马萨拉Masal    feud.Barony
	马苏莱Masouleh feud.Barony
	鲁德罕Rudkhan  feud.Barony
	塔雷什Talysh   feud.Barony
}

func (c *吉兰GilanCounty) BAstara阿斯塔拉() feud.Barony {
	return c.阿斯塔拉Astara
}

func (c *吉兰GilanCounty) BGilan吉兰() feud.Barony {
	return c.吉兰Gilan
}

func (c *吉兰GilanCounty) BLahijan拉希詹() feud.Barony {
	return c.拉希詹Lahijan
}

func (c *吉兰GilanCounty) BMasal马萨拉() feud.Barony {
	return c.马萨拉Masal
}

func (c *吉兰GilanCounty) BMasouleh马苏莱() feud.Barony {
	return c.马苏莱Masouleh
}

func (c *吉兰GilanCounty) BRudkhan鲁德罕() feud.Barony {
	return c.鲁德罕Rudkhan
}

func (c *吉兰GilanCounty) BTalysh塔雷什() feud.Barony {
	return c.塔雷什Talysh
}

var CGilan吉兰 GilanCounty = &吉兰GilanCounty{}

func init() {
	f := CGilan吉兰.(*吉兰GilanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "666",
		Title:     "gilan",
		TitleName: "吉兰",
		TitleCode: "c_gilan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿斯塔拉Astara = BAstara阿斯塔拉
	f.阿斯塔拉Astara.SetParent(f)

	f.吉兰Gilan = BGilan吉兰
	f.吉兰Gilan.SetParent(f)

	f.拉希詹Lahijan = BLahijan拉希詹
	f.拉希詹Lahijan.SetParent(f)

	f.马萨拉Masal = BMasal马萨拉
	f.马萨拉Masal.SetParent(f)

	f.马苏莱Masouleh = BMasouleh马苏莱
	f.马苏莱Masouleh.SetParent(f)

	f.鲁德罕Rudkhan = BRudkhan鲁德罕
	f.鲁德罕Rudkhan.SetParent(f)

	f.塔雷什Talysh = BTalysh塔雷什
	f.塔雷什Talysh.SetParent(f)

}
