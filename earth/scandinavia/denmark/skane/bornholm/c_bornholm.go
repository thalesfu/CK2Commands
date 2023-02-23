package bornholm

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BornholmCounty interface {
	feud.County
	BAakirkeby奥基克比() feud.Barony
	BGudhjem古兹耶姆() feud.Barony
	BHammershus哈默斯胡斯() feud.Barony
	BHasle海斯勒() feud.Barony
	BKnudsker克罗斯齐亚() feud.Barony
	BNexo内克瑟() feud.Barony
	BRonne伦讷() feud.Barony
	BSvaneke斯瓦讷克() feud.Barony
}

type 博恩霍尔姆BornholmCounty struct {
	feud.BaseCounty
	奥基克比Aakirkeby   feud.Barony
	古兹耶姆Gudhjem     feud.Barony
	哈默斯胡斯Hammershus feud.Barony
	海斯勒Hasle        feud.Barony
	克罗斯齐亚Knudsker   feud.Barony
	内克瑟Nexo         feud.Barony
	伦讷Ronne         feud.Barony
	斯瓦讷克Svaneke     feud.Barony
}

func (c *博恩霍尔姆BornholmCounty) BAakirkeby奥基克比() feud.Barony {
	return c.奥基克比Aakirkeby
}

func (c *博恩霍尔姆BornholmCounty) BGudhjem古兹耶姆() feud.Barony {
	return c.古兹耶姆Gudhjem
}

func (c *博恩霍尔姆BornholmCounty) BHammershus哈默斯胡斯() feud.Barony {
	return c.哈默斯胡斯Hammershus
}

func (c *博恩霍尔姆BornholmCounty) BHasle海斯勒() feud.Barony {
	return c.海斯勒Hasle
}

func (c *博恩霍尔姆BornholmCounty) BKnudsker克罗斯齐亚() feud.Barony {
	return c.克罗斯齐亚Knudsker
}

func (c *博恩霍尔姆BornholmCounty) BNexo内克瑟() feud.Barony {
	return c.内克瑟Nexo
}

func (c *博恩霍尔姆BornholmCounty) BRonne伦讷() feud.Barony {
	return c.伦讷Ronne
}

func (c *博恩霍尔姆BornholmCounty) BSvaneke斯瓦讷克() feud.Barony {
	return c.斯瓦讷克Svaneke
}

var CBornholm博恩霍尔姆 BornholmCounty = &博恩霍尔姆BornholmCounty{}

func init() {
	f := CBornholm博恩霍尔姆.(*博恩霍尔姆BornholmCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "305",
		Title:     "bornholm",
		TitleName: "博恩霍尔姆",
		TitleCode: "c_bornholm",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥基克比Aakirkeby = BAakirkeby奥基克比
	f.奥基克比Aakirkeby.SetParent(f)

	f.古兹耶姆Gudhjem = BGudhjem古兹耶姆
	f.古兹耶姆Gudhjem.SetParent(f)

	f.哈默斯胡斯Hammershus = BHammershus哈默斯胡斯
	f.哈默斯胡斯Hammershus.SetParent(f)

	f.海斯勒Hasle = BHasle海斯勒
	f.海斯勒Hasle.SetParent(f)

	f.克罗斯齐亚Knudsker = BKnudsker克罗斯齐亚
	f.克罗斯齐亚Knudsker.SetParent(f)

	f.内克瑟Nexo = BNexo内克瑟
	f.内克瑟Nexo.SetParent(f)

	f.伦讷Ronne = BRonne伦讷
	f.伦讷Ronne.SetParent(f)

	f.斯瓦讷克Svaneke = BSvaneke斯瓦讷克
	f.斯瓦讷克Svaneke.SetParent(f)

}
