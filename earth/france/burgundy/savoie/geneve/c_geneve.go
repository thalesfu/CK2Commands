package geneve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GeneveCounty interface {
	feud.County
	BAnnecy阿讷西() feud.Barony
	BAubonne欧博讷() feud.Barony
	BGeneve日内瓦() feud.Barony
	BMeinier梅涅() feud.Barony
	BNyon尼翁() feud.Barony
	BThonex托内() feud.Barony
	BYvoire伊瓦尔() feud.Barony
}

type 日内瓦GeneveCounty struct {
	feud.BaseCounty
	阿讷西Annecy  feud.Barony
	欧博讷Aubonne feud.Barony
	日内瓦Geneve  feud.Barony
	梅涅Meinier  feud.Barony
	尼翁Nyon     feud.Barony
	托内Thonex   feud.Barony
	伊瓦尔Yvoire  feud.Barony
}

func (c *日内瓦GeneveCounty) BAnnecy阿讷西() feud.Barony {
	return c.阿讷西Annecy
}

func (c *日内瓦GeneveCounty) BAubonne欧博讷() feud.Barony {
	return c.欧博讷Aubonne
}

func (c *日内瓦GeneveCounty) BGeneve日内瓦() feud.Barony {
	return c.日内瓦Geneve
}

func (c *日内瓦GeneveCounty) BMeinier梅涅() feud.Barony {
	return c.梅涅Meinier
}

func (c *日内瓦GeneveCounty) BNyon尼翁() feud.Barony {
	return c.尼翁Nyon
}

func (c *日内瓦GeneveCounty) BThonex托内() feud.Barony {
	return c.托内Thonex
}

func (c *日内瓦GeneveCounty) BYvoire伊瓦尔() feud.Barony {
	return c.伊瓦尔Yvoire
}

var CGeneve日内瓦 GeneveCounty = &日内瓦GeneveCounty{}

func init() {
	f := CGeneve日内瓦.(*日内瓦GeneveCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "239",
		Title:     "geneve",
		TitleName: "日内瓦",
		TitleCode: "c_geneve",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿讷西Annecy = BAnnecy阿讷西
	f.阿讷西Annecy.SetParent(f)

	f.欧博讷Aubonne = BAubonne欧博讷
	f.欧博讷Aubonne.SetParent(f)

	f.日内瓦Geneve = BGeneve日内瓦
	f.日内瓦Geneve.SetParent(f)

	f.梅涅Meinier = BMeinier梅涅
	f.梅涅Meinier.SetParent(f)

	f.尼翁Nyon = BNyon尼翁
	f.尼翁Nyon.SetParent(f)

	f.托内Thonex = BThonex托内
	f.托内Thonex.SetParent(f)

	f.伊瓦尔Yvoire = BYvoire伊瓦尔
	f.伊瓦尔Yvoire.SetParent(f)

}
