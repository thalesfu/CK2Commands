package oulu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OuluCounty interface {
	feud.County
	BJakobstad皮耶塔尔萨里() feud.Barony
	BKiiminki基明基() feud.Barony
	BKokkola科科拉() feud.Barony
	BLapuan拉普阿() feud.Barony
	BNykarleby新卡勒比() feud.Barony
	BOulu奥卢() feud.Barony
	BRaahe拉赫() feud.Barony
	BSiuruan休鲁阿() feud.Barony
}

type 奥卢OuluCounty struct {
	feud.BaseCounty
	皮耶塔尔萨里Jakobstad feud.Barony
	基明基Kiiminki     feud.Barony
	科科拉Kokkola      feud.Barony
	拉普阿Lapuan       feud.Barony
	新卡勒比Nykarleby   feud.Barony
	奥卢Oulu          feud.Barony
	拉赫Raahe         feud.Barony
	休鲁阿Siuruan      feud.Barony
}

func (c *奥卢OuluCounty) BJakobstad皮耶塔尔萨里() feud.Barony {
	return c.皮耶塔尔萨里Jakobstad
}

func (c *奥卢OuluCounty) BKiiminki基明基() feud.Barony {
	return c.基明基Kiiminki
}

func (c *奥卢OuluCounty) BKokkola科科拉() feud.Barony {
	return c.科科拉Kokkola
}

func (c *奥卢OuluCounty) BLapuan拉普阿() feud.Barony {
	return c.拉普阿Lapuan
}

func (c *奥卢OuluCounty) BNykarleby新卡勒比() feud.Barony {
	return c.新卡勒比Nykarleby
}

func (c *奥卢OuluCounty) BOulu奥卢() feud.Barony {
	return c.奥卢Oulu
}

func (c *奥卢OuluCounty) BRaahe拉赫() feud.Barony {
	return c.拉赫Raahe
}

func (c *奥卢OuluCounty) BSiuruan休鲁阿() feud.Barony {
	return c.休鲁阿Siuruan
}

var COulu奥卢 OuluCounty = &奥卢OuluCounty{}

func init() {
	f := COulu奥卢.(*奥卢OuluCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1600",
		Title:     "oulu",
		TitleName: "奥卢",
		TitleCode: "c_oulu",
		Baronies:  map[string]feud.Barony{},
	}

	f.皮耶塔尔萨里Jakobstad = BJakobstad皮耶塔尔萨里
	f.皮耶塔尔萨里Jakobstad.SetParent(f)

	f.基明基Kiiminki = BKiiminki基明基
	f.基明基Kiiminki.SetParent(f)

	f.科科拉Kokkola = BKokkola科科拉
	f.科科拉Kokkola.SetParent(f)

	f.拉普阿Lapuan = BLapuan拉普阿
	f.拉普阿Lapuan.SetParent(f)

	f.新卡勒比Nykarleby = BNykarleby新卡勒比
	f.新卡勒比Nykarleby.SetParent(f)

	f.奥卢Oulu = BOulu奥卢
	f.奥卢Oulu.SetParent(f)

	f.拉赫Raahe = BRaahe拉赫
	f.拉赫Raahe.SetParent(f)

	f.休鲁阿Siuruan = BSiuruan休鲁阿
	f.休鲁阿Siuruan.SetParent(f)

}
