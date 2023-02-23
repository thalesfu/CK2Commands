package saluzzo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SaluzzoCounty interface {
	feud.County
	BBusca布斯卡() feud.Barony
	BCaraglio卡拉廖() feud.Barony
	BCuneo库内奥() feud.Barony
	BPaesana帕埃萨纳() feud.Barony
	BPinerolo皮内罗洛() feud.Barony
	BSaluzzo萨卢佐() feud.Barony
	BSavigliano萨维利亚诺() feud.Barony
	BVerzuolo韦尔佐洛() feud.Barony
}

type 萨卢佐SaluzzoCounty struct {
	feud.BaseCounty
	布斯卡Busca        feud.Barony
	卡拉廖Caraglio     feud.Barony
	库内奥Cuneo        feud.Barony
	帕埃萨纳Paesana     feud.Barony
	皮内罗洛Pinerolo    feud.Barony
	萨卢佐Saluzzo      feud.Barony
	萨维利亚诺Savigliano feud.Barony
	韦尔佐洛Verzuolo    feud.Barony
}

func (c *萨卢佐SaluzzoCounty) BBusca布斯卡() feud.Barony {
	return c.布斯卡Busca
}

func (c *萨卢佐SaluzzoCounty) BCaraglio卡拉廖() feud.Barony {
	return c.卡拉廖Caraglio
}

func (c *萨卢佐SaluzzoCounty) BCuneo库内奥() feud.Barony {
	return c.库内奥Cuneo
}

func (c *萨卢佐SaluzzoCounty) BPaesana帕埃萨纳() feud.Barony {
	return c.帕埃萨纳Paesana
}

func (c *萨卢佐SaluzzoCounty) BPinerolo皮内罗洛() feud.Barony {
	return c.皮内罗洛Pinerolo
}

func (c *萨卢佐SaluzzoCounty) BSaluzzo萨卢佐() feud.Barony {
	return c.萨卢佐Saluzzo
}

func (c *萨卢佐SaluzzoCounty) BSavigliano萨维利亚诺() feud.Barony {
	return c.萨维利亚诺Savigliano
}

func (c *萨卢佐SaluzzoCounty) BVerzuolo韦尔佐洛() feud.Barony {
	return c.韦尔佐洛Verzuolo
}

var CSaluzzo萨卢佐 SaluzzoCounty = &萨卢佐SaluzzoCounty{}

func init() {
	f := CSaluzzo萨卢佐.(*萨卢佐SaluzzoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "231",
		Title:     "saluzzo",
		TitleName: "萨卢佐",
		TitleCode: "c_saluzzo",
		Baronies:  map[string]feud.Barony{},
	}

	f.布斯卡Busca = BBusca布斯卡
	f.布斯卡Busca.SetParent(f)

	f.卡拉廖Caraglio = BCaraglio卡拉廖
	f.卡拉廖Caraglio.SetParent(f)

	f.库内奥Cuneo = BCuneo库内奥
	f.库内奥Cuneo.SetParent(f)

	f.帕埃萨纳Paesana = BPaesana帕埃萨纳
	f.帕埃萨纳Paesana.SetParent(f)

	f.皮内罗洛Pinerolo = BPinerolo皮内罗洛
	f.皮内罗洛Pinerolo.SetParent(f)

	f.萨卢佐Saluzzo = BSaluzzo萨卢佐
	f.萨卢佐Saluzzo.SetParent(f)

	f.萨维利亚诺Savigliano = BSavigliano萨维利亚诺
	f.萨维利亚诺Savigliano.SetParent(f)

	f.韦尔佐洛Verzuolo = BVerzuolo韦尔佐洛
	f.韦尔佐洛Verzuolo.SetParent(f)

}
