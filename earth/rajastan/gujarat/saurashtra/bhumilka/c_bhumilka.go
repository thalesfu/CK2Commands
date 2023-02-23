package bhumilka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BhumilkaCounty interface {
	feud.County
	BBhavnat般那特() feud.Barony
	BBhumilka菩弥伽() feud.Barony
	BGirnar吉尔纳尔() feud.Barony
	BMangrol门罗尔() feud.Barony
	BShrinagar室利那揭罗() feud.Barony
	BUperkot优波罗拘吒() feud.Barony
	BVamanasthali筏摩那悉他利() feud.Barony
}

type 菩弥伽BhumilkaCounty struct {
	feud.BaseCounty
	般那特Bhavnat         feud.Barony
	菩弥伽Bhumilka        feud.Barony
	吉尔纳尔Girnar         feud.Barony
	门罗尔Mangrol         feud.Barony
	室利那揭罗Shrinagar     feud.Barony
	优波罗拘吒Uperkot       feud.Barony
	筏摩那悉他利Vamanasthali feud.Barony
}

func (c *菩弥伽BhumilkaCounty) BBhavnat般那特() feud.Barony {
	return c.般那特Bhavnat
}

func (c *菩弥伽BhumilkaCounty) BBhumilka菩弥伽() feud.Barony {
	return c.菩弥伽Bhumilka
}

func (c *菩弥伽BhumilkaCounty) BGirnar吉尔纳尔() feud.Barony {
	return c.吉尔纳尔Girnar
}

func (c *菩弥伽BhumilkaCounty) BMangrol门罗尔() feud.Barony {
	return c.门罗尔Mangrol
}

func (c *菩弥伽BhumilkaCounty) BShrinagar室利那揭罗() feud.Barony {
	return c.室利那揭罗Shrinagar
}

func (c *菩弥伽BhumilkaCounty) BUperkot优波罗拘吒() feud.Barony {
	return c.优波罗拘吒Uperkot
}

func (c *菩弥伽BhumilkaCounty) BVamanasthali筏摩那悉他利() feud.Barony {
	return c.筏摩那悉他利Vamanasthali
}

var CBhumilka菩弥伽 BhumilkaCounty = &菩弥伽BhumilkaCounty{}

func init() {
	f := CBhumilka菩弥伽.(*菩弥伽BhumilkaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1134",
		Title:     "bhumilka",
		TitleName: "菩弥伽",
		TitleCode: "c_bhumilka",
		Baronies:  map[string]feud.Barony{},
	}

	f.般那特Bhavnat = BBhavnat般那特
	f.般那特Bhavnat.SetParent(f)

	f.菩弥伽Bhumilka = BBhumilka菩弥伽
	f.菩弥伽Bhumilka.SetParent(f)

	f.吉尔纳尔Girnar = BGirnar吉尔纳尔
	f.吉尔纳尔Girnar.SetParent(f)

	f.门罗尔Mangrol = BMangrol门罗尔
	f.门罗尔Mangrol.SetParent(f)

	f.室利那揭罗Shrinagar = BShrinagar室利那揭罗
	f.室利那揭罗Shrinagar.SetParent(f)

	f.优波罗拘吒Uperkot = BUperkot优波罗拘吒
	f.优波罗拘吒Uperkot.SetParent(f)

	f.筏摩那悉他利Vamanasthali = BVamanasthali筏摩那悉他利
	f.筏摩那悉他利Vamanasthali.SetParent(f)

}
