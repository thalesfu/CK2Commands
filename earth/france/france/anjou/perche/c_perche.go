package perche

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PercheCounty interface {
	feud.County
	BBelleme贝莱姆() feud.Barony
	BLongny隆尼() feud.Barony
	BLuigny吕伊尼() feud.Barony
	BMortagne莫尔塔涅() feud.Barony
	BSenonches瑟农什() feud.Barony
	BVerneuil韦尔讷伊() feud.Barony
}

type 佩尔什PercheCounty struct {
	feud.BaseCounty
	贝莱姆Belleme   feud.Barony
	隆尼Longny     feud.Barony
	吕伊尼Luigny    feud.Barony
	莫尔塔涅Mortagne feud.Barony
	瑟农什Senonches feud.Barony
	韦尔讷伊Verneuil feud.Barony
}

func (c *佩尔什PercheCounty) BBelleme贝莱姆() feud.Barony {
	return c.贝莱姆Belleme
}

func (c *佩尔什PercheCounty) BLongny隆尼() feud.Barony {
	return c.隆尼Longny
}

func (c *佩尔什PercheCounty) BLuigny吕伊尼() feud.Barony {
	return c.吕伊尼Luigny
}

func (c *佩尔什PercheCounty) BMortagne莫尔塔涅() feud.Barony {
	return c.莫尔塔涅Mortagne
}

func (c *佩尔什PercheCounty) BSenonches瑟农什() feud.Barony {
	return c.瑟农什Senonches
}

func (c *佩尔什PercheCounty) BVerneuil韦尔讷伊() feud.Barony {
	return c.韦尔讷伊Verneuil
}

var CPerche佩尔什 PercheCounty = &佩尔什PercheCounty{}

func init() {
	f := CPerche佩尔什.(*佩尔什PercheCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1958",
		Title:     "perche",
		TitleName: "佩尔什",
		TitleCode: "c_perche",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝莱姆Belleme = BBelleme贝莱姆
	f.贝莱姆Belleme.SetParent(f)

	f.隆尼Longny = BLongny隆尼
	f.隆尼Longny.SetParent(f)

	f.吕伊尼Luigny = BLuigny吕伊尼
	f.吕伊尼Luigny.SetParent(f)

	f.莫尔塔涅Mortagne = BMortagne莫尔塔涅
	f.莫尔塔涅Mortagne.SetParent(f)

	f.瑟农什Senonches = BSenonches瑟农什
	f.瑟农什Senonches.SetParent(f)

	f.韦尔讷伊Verneuil = BVerneuil韦尔讷伊
	f.韦尔讷伊Verneuil.SetParent(f)

}
