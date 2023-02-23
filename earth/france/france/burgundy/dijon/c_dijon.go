package dijon

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DijonCounty interface {
	feud.County
	BAutun奥坦() feud.Barony
	BAvallon阿瓦隆() feud.Barony
	BBeaune博讷() feud.Barony
	BCiteaux西托() feud.Barony
	BDijon第戎() feud.Barony
	BNoyers努瓦耶() feud.Barony
	BSemur瑟米尔() feud.Barony
	BVezelay韦泽莱() feud.Barony
}

type 第戎DijonCounty struct {
	feud.BaseCounty
	奥坦Autun    feud.Barony
	阿瓦隆Avallon feud.Barony
	博讷Beaune   feud.Barony
	西托Citeaux  feud.Barony
	第戎Dijon    feud.Barony
	努瓦耶Noyers  feud.Barony
	瑟米尔Semur   feud.Barony
	韦泽莱Vezelay feud.Barony
}

func (c *第戎DijonCounty) BAutun奥坦() feud.Barony {
	return c.奥坦Autun
}

func (c *第戎DijonCounty) BAvallon阿瓦隆() feud.Barony {
	return c.阿瓦隆Avallon
}

func (c *第戎DijonCounty) BBeaune博讷() feud.Barony {
	return c.博讷Beaune
}

func (c *第戎DijonCounty) BCiteaux西托() feud.Barony {
	return c.西托Citeaux
}

func (c *第戎DijonCounty) BDijon第戎() feud.Barony {
	return c.第戎Dijon
}

func (c *第戎DijonCounty) BNoyers努瓦耶() feud.Barony {
	return c.努瓦耶Noyers
}

func (c *第戎DijonCounty) BSemur瑟米尔() feud.Barony {
	return c.瑟米尔Semur
}

func (c *第戎DijonCounty) BVezelay韦泽莱() feud.Barony {
	return c.韦泽莱Vezelay
}

var CDijon第戎 DijonCounty = &第戎DijonCounty{}

func init() {
	f := CDijon第戎.(*第戎DijonCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "136",
		Title:     "dijon",
		TitleName: "第戎",
		TitleCode: "c_dijon",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥坦Autun = BAutun奥坦
	f.奥坦Autun.SetParent(f)

	f.阿瓦隆Avallon = BAvallon阿瓦隆
	f.阿瓦隆Avallon.SetParent(f)

	f.博讷Beaune = BBeaune博讷
	f.博讷Beaune.SetParent(f)

	f.西托Citeaux = BCiteaux西托
	f.西托Citeaux.SetParent(f)

	f.第戎Dijon = BDijon第戎
	f.第戎Dijon.SetParent(f)

	f.努瓦耶Noyers = BNoyers努瓦耶
	f.努瓦耶Noyers.SetParent(f)

	f.瑟米尔Semur = BSemur瑟米尔
	f.瑟米尔Semur.SetParent(f)

	f.韦泽莱Vezelay = BVezelay韦泽莱
	f.韦泽莱Vezelay.SetParent(f)

}
