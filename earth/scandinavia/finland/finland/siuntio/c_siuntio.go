package siuntio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SiuntioCounty interface {
	feud.County
	BBodom博多姆() feud.Barony
	BHelsinki赫尔辛基() feud.Barony
	BKerava凯拉瓦() feud.Barony
	BLohja洛赫亚() feud.Barony
	BRaseborg拉塞博里() feud.Barony
	BSiuntio锡温蒂奥() feud.Barony
	BVantaa万塔() feud.Barony
}

type 锡温蒂奥SiuntioCounty struct {
	feud.BaseCounty
	博多姆Bodom     feud.Barony
	赫尔辛基Helsinki feud.Barony
	凯拉瓦Kerava    feud.Barony
	洛赫亚Lohja     feud.Barony
	拉塞博里Raseborg feud.Barony
	锡温蒂奥Siuntio  feud.Barony
	万塔Vantaa     feud.Barony
}

func (c *锡温蒂奥SiuntioCounty) BBodom博多姆() feud.Barony {
	return c.博多姆Bodom
}

func (c *锡温蒂奥SiuntioCounty) BHelsinki赫尔辛基() feud.Barony {
	return c.赫尔辛基Helsinki
}

func (c *锡温蒂奥SiuntioCounty) BKerava凯拉瓦() feud.Barony {
	return c.凯拉瓦Kerava
}

func (c *锡温蒂奥SiuntioCounty) BLohja洛赫亚() feud.Barony {
	return c.洛赫亚Lohja
}

func (c *锡温蒂奥SiuntioCounty) BRaseborg拉塞博里() feud.Barony {
	return c.拉塞博里Raseborg
}

func (c *锡温蒂奥SiuntioCounty) BSiuntio锡温蒂奥() feud.Barony {
	return c.锡温蒂奥Siuntio
}

func (c *锡温蒂奥SiuntioCounty) BVantaa万塔() feud.Barony {
	return c.万塔Vantaa
}

var CSiuntio锡温蒂奥 SiuntioCounty = &锡温蒂奥SiuntioCounty{}

func init() {
	f := CSiuntio锡温蒂奥.(*锡温蒂奥SiuntioCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1606",
		Title:     "siuntio",
		TitleName: "锡温蒂奥",
		TitleCode: "c_siuntio",
		Baronies:  map[string]feud.Barony{},
	}

	f.博多姆Bodom = BBodom博多姆
	f.博多姆Bodom.SetParent(f)

	f.赫尔辛基Helsinki = BHelsinki赫尔辛基
	f.赫尔辛基Helsinki.SetParent(f)

	f.凯拉瓦Kerava = BKerava凯拉瓦
	f.凯拉瓦Kerava.SetParent(f)

	f.洛赫亚Lohja = BLohja洛赫亚
	f.洛赫亚Lohja.SetParent(f)

	f.拉塞博里Raseborg = BRaseborg拉塞博里
	f.拉塞博里Raseborg.SetParent(f)

	f.锡温蒂奥Siuntio = BSiuntio锡温蒂奥
	f.锡温蒂奥Siuntio.SetParent(f)

	f.万塔Vantaa = BVantaa万塔
	f.万塔Vantaa.SetParent(f)

}
