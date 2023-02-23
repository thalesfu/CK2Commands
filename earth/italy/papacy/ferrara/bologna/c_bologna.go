package bologna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BolognaCounty interface {
	feud.County
	BBagnacavallo巴尼亚卡瓦洛() feud.Barony
	BBentivoglio本蒂沃利奥() feud.Barony
	BBologna博洛尼亚() feud.Barony
	BBudno布德诺() feud.Barony
	BCastelguelfo圭尔福堡() feud.Barony
	BFaenza法恩扎() feud.Barony
	BForli弗利() feud.Barony
	BImola伊莫拉() feud.Barony
}

type 博洛尼亚BolognaCounty struct {
	feud.BaseCounty
	巴尼亚卡瓦洛Bagnacavallo feud.Barony
	本蒂沃利奥Bentivoglio   feud.Barony
	博洛尼亚Bologna        feud.Barony
	布德诺Budno           feud.Barony
	圭尔福堡Castelguelfo   feud.Barony
	法恩扎Faenza          feud.Barony
	弗利Forli            feud.Barony
	伊莫拉Imola           feud.Barony
}

func (c *博洛尼亚BolognaCounty) BBagnacavallo巴尼亚卡瓦洛() feud.Barony {
	return c.巴尼亚卡瓦洛Bagnacavallo
}

func (c *博洛尼亚BolognaCounty) BBentivoglio本蒂沃利奥() feud.Barony {
	return c.本蒂沃利奥Bentivoglio
}

func (c *博洛尼亚BolognaCounty) BBologna博洛尼亚() feud.Barony {
	return c.博洛尼亚Bologna
}

func (c *博洛尼亚BolognaCounty) BBudno布德诺() feud.Barony {
	return c.布德诺Budno
}

func (c *博洛尼亚BolognaCounty) BCastelguelfo圭尔福堡() feud.Barony {
	return c.圭尔福堡Castelguelfo
}

func (c *博洛尼亚BolognaCounty) BFaenza法恩扎() feud.Barony {
	return c.法恩扎Faenza
}

func (c *博洛尼亚BolognaCounty) BForli弗利() feud.Barony {
	return c.弗利Forli
}

func (c *博洛尼亚BolognaCounty) BImola伊莫拉() feud.Barony {
	return c.伊莫拉Imola
}

var CBologna博洛尼亚 BolognaCounty = &博洛尼亚BolognaCounty{}

func init() {
	f := CBologna博洛尼亚.(*博洛尼亚BolognaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "352",
		Title:     "bologna",
		TitleName: "博洛尼亚",
		TitleCode: "c_bologna",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尼亚卡瓦洛Bagnacavallo = BBagnacavallo巴尼亚卡瓦洛
	f.巴尼亚卡瓦洛Bagnacavallo.SetParent(f)

	f.本蒂沃利奥Bentivoglio = BBentivoglio本蒂沃利奥
	f.本蒂沃利奥Bentivoglio.SetParent(f)

	f.博洛尼亚Bologna = BBologna博洛尼亚
	f.博洛尼亚Bologna.SetParent(f)

	f.布德诺Budno = BBudno布德诺
	f.布德诺Budno.SetParent(f)

	f.圭尔福堡Castelguelfo = BCastelguelfo圭尔福堡
	f.圭尔福堡Castelguelfo.SetParent(f)

	f.法恩扎Faenza = BFaenza法恩扎
	f.法恩扎Faenza.SetParent(f)

	f.弗利Forli = BForli弗利
	f.弗利Forli.SetParent(f)

	f.伊莫拉Imola = BImola伊莫拉
	f.伊莫拉Imola.SetParent(f)

}
