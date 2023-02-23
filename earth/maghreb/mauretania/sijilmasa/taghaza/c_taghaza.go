package taghaza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaghazaCounty interface {
	feud.County
	BAdrar阿德拉尔() feud.Barony
	BAlouarta瓦尔塔() feud.Barony
	BAougrout阿乌格鲁特() feud.Barony
	BCharouine沙尔温() feud.Barony
	BKerzaz克尔扎兹() feud.Barony
	BMetarta梅达塔() feud.Barony
	BReggane雷甘() feud.Barony
	BTaghaza泰加扎() feud.Barony
}

type 泰加扎TaghazaCounty struct {
	feud.BaseCounty
	阿德拉尔Adrar     feud.Barony
	瓦尔塔Alouarta   feud.Barony
	阿乌格鲁特Aougrout feud.Barony
	沙尔温Charouine  feud.Barony
	克尔扎兹Kerzaz    feud.Barony
	梅达塔Metarta    feud.Barony
	雷甘Reggane     feud.Barony
	泰加扎Taghaza    feud.Barony
}

func (c *泰加扎TaghazaCounty) BAdrar阿德拉尔() feud.Barony {
	return c.阿德拉尔Adrar
}

func (c *泰加扎TaghazaCounty) BAlouarta瓦尔塔() feud.Barony {
	return c.瓦尔塔Alouarta
}

func (c *泰加扎TaghazaCounty) BAougrout阿乌格鲁特() feud.Barony {
	return c.阿乌格鲁特Aougrout
}

func (c *泰加扎TaghazaCounty) BCharouine沙尔温() feud.Barony {
	return c.沙尔温Charouine
}

func (c *泰加扎TaghazaCounty) BKerzaz克尔扎兹() feud.Barony {
	return c.克尔扎兹Kerzaz
}

func (c *泰加扎TaghazaCounty) BMetarta梅达塔() feud.Barony {
	return c.梅达塔Metarta
}

func (c *泰加扎TaghazaCounty) BReggane雷甘() feud.Barony {
	return c.雷甘Reggane
}

func (c *泰加扎TaghazaCounty) BTaghaza泰加扎() feud.Barony {
	return c.泰加扎Taghaza
}

var CTaghaza泰加扎 TaghazaCounty = &泰加扎TaghazaCounty{}

func init() {
	f := CTaghaza泰加扎.(*泰加扎TaghazaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "916",
		Title:     "taghaza",
		TitleName: "泰加扎",
		TitleCode: "c_taghaza",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德拉尔Adrar = BAdrar阿德拉尔
	f.阿德拉尔Adrar.SetParent(f)

	f.瓦尔塔Alouarta = BAlouarta瓦尔塔
	f.瓦尔塔Alouarta.SetParent(f)

	f.阿乌格鲁特Aougrout = BAougrout阿乌格鲁特
	f.阿乌格鲁特Aougrout.SetParent(f)

	f.沙尔温Charouine = BCharouine沙尔温
	f.沙尔温Charouine.SetParent(f)

	f.克尔扎兹Kerzaz = BKerzaz克尔扎兹
	f.克尔扎兹Kerzaz.SetParent(f)

	f.梅达塔Metarta = BMetarta梅达塔
	f.梅达塔Metarta.SetParent(f)

	f.雷甘Reggane = BReggane雷甘
	f.雷甘Reggane.SetParent(f)

	f.泰加扎Taghaza = BTaghaza泰加扎
	f.泰加扎Taghaza.SetParent(f)

}
