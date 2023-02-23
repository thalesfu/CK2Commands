package varmland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VarmlandCounty interface {
	feud.County
	BArvika阿尔维卡() feud.Barony
	BFryksdal弗吕克斯达尔() feud.Barony
	BGillberg伊尔贝里() feud.Barony
	BJosse约瑟() feud.Barony
	BKil希尔() feud.Barony
	BNordmark努德马克() feud.Barony
	BSaxholm萨克斯霍尔姆() feud.Barony
	BVase韦瑟() feud.Barony
}

type 韦姆兰VarmlandCounty struct {
	feud.BaseCounty
	阿尔维卡Arvika     feud.Barony
	弗吕克斯达尔Fryksdal feud.Barony
	伊尔贝里Gillberg   feud.Barony
	约瑟Josse        feud.Barony
	希尔Kil          feud.Barony
	努德马克Nordmark   feud.Barony
	萨克斯霍尔姆Saxholm  feud.Barony
	韦瑟Vase         feud.Barony
}

func (c *韦姆兰VarmlandCounty) BArvika阿尔维卡() feud.Barony {
	return c.阿尔维卡Arvika
}

func (c *韦姆兰VarmlandCounty) BFryksdal弗吕克斯达尔() feud.Barony {
	return c.弗吕克斯达尔Fryksdal
}

func (c *韦姆兰VarmlandCounty) BGillberg伊尔贝里() feud.Barony {
	return c.伊尔贝里Gillberg
}

func (c *韦姆兰VarmlandCounty) BJosse约瑟() feud.Barony {
	return c.约瑟Josse
}

func (c *韦姆兰VarmlandCounty) BKil希尔() feud.Barony {
	return c.希尔Kil
}

func (c *韦姆兰VarmlandCounty) BNordmark努德马克() feud.Barony {
	return c.努德马克Nordmark
}

func (c *韦姆兰VarmlandCounty) BSaxholm萨克斯霍尔姆() feud.Barony {
	return c.萨克斯霍尔姆Saxholm
}

func (c *韦姆兰VarmlandCounty) BVase韦瑟() feud.Barony {
	return c.韦瑟Vase
}

var CVarmland韦姆兰 VarmlandCounty = &韦姆兰VarmlandCounty{}

func init() {
	f := CVarmland韦姆兰.(*韦姆兰VarmlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "288",
		Title:     "varmland",
		TitleName: "韦姆兰",
		TitleCode: "c_varmland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔维卡Arvika = BArvika阿尔维卡
	f.阿尔维卡Arvika.SetParent(f)

	f.弗吕克斯达尔Fryksdal = BFryksdal弗吕克斯达尔
	f.弗吕克斯达尔Fryksdal.SetParent(f)

	f.伊尔贝里Gillberg = BGillberg伊尔贝里
	f.伊尔贝里Gillberg.SetParent(f)

	f.约瑟Josse = BJosse约瑟
	f.约瑟Josse.SetParent(f)

	f.希尔Kil = BKil希尔
	f.希尔Kil.SetParent(f)

	f.努德马克Nordmark = BNordmark努德马克
	f.努德马克Nordmark.SetParent(f)

	f.萨克斯霍尔姆Saxholm = BSaxholm萨克斯霍尔姆
	f.萨克斯霍尔姆Saxholm.SetParent(f)

	f.韦瑟Vase = BVase韦瑟
	f.韦瑟Vase.SetParent(f)

}
