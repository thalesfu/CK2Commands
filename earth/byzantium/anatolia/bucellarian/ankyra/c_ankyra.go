package ankyra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AnkyraCounty interface {
	feud.County
	BAnkyra安居拉() feud.Barony
	BGerma杰尔玛() feud.Barony
	BGordium戈尔迪乌姆() feud.Barony
	BGordoservon哥德瑟温() feud.Barony
	BHaymana哈伊马纳() feud.Barony
	BNakoleia纳科莱亚() feud.Barony
	BTectosagii特克托萨季() feud.Barony
}

type 安居拉AnkyraCounty struct {
	feud.BaseCounty
	安居拉Ankyra       feud.Barony
	杰尔玛Germa        feud.Barony
	戈尔迪乌姆Gordium    feud.Barony
	哥德瑟温Gordoservon feud.Barony
	哈伊马纳Haymana     feud.Barony
	纳科莱亚Nakoleia    feud.Barony
	特克托萨季Tectosagii feud.Barony
}

func (c *安居拉AnkyraCounty) BAnkyra安居拉() feud.Barony {
	return c.安居拉Ankyra
}

func (c *安居拉AnkyraCounty) BGerma杰尔玛() feud.Barony {
	return c.杰尔玛Germa
}

func (c *安居拉AnkyraCounty) BGordium戈尔迪乌姆() feud.Barony {
	return c.戈尔迪乌姆Gordium
}

func (c *安居拉AnkyraCounty) BGordoservon哥德瑟温() feud.Barony {
	return c.哥德瑟温Gordoservon
}

func (c *安居拉AnkyraCounty) BHaymana哈伊马纳() feud.Barony {
	return c.哈伊马纳Haymana
}

func (c *安居拉AnkyraCounty) BNakoleia纳科莱亚() feud.Barony {
	return c.纳科莱亚Nakoleia
}

func (c *安居拉AnkyraCounty) BTectosagii特克托萨季() feud.Barony {
	return c.特克托萨季Tectosagii
}

var CAnkyra安居拉 AnkyraCounty = &安居拉AnkyraCounty{}

func init() {
	f := CAnkyra安居拉.(*安居拉AnkyraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "753",
		Title:     "ankyra",
		TitleName: "安居拉",
		TitleCode: "c_ankyra",
		Baronies:  map[string]feud.Barony{},
	}

	f.安居拉Ankyra = BAnkyra安居拉
	f.安居拉Ankyra.SetParent(f)

	f.杰尔玛Germa = BGerma杰尔玛
	f.杰尔玛Germa.SetParent(f)

	f.戈尔迪乌姆Gordium = BGordium戈尔迪乌姆
	f.戈尔迪乌姆Gordium.SetParent(f)

	f.哥德瑟温Gordoservon = BGordoservon哥德瑟温
	f.哥德瑟温Gordoservon.SetParent(f)

	f.哈伊马纳Haymana = BHaymana哈伊马纳
	f.哈伊马纳Haymana.SetParent(f)

	f.纳科莱亚Nakoleia = BNakoleia纳科莱亚
	f.纳科莱亚Nakoleia.SetParent(f)

	f.特克托萨季Tectosagii = BTectosagii特克托萨季
	f.特克托萨季Tectosagii.SetParent(f)

}
