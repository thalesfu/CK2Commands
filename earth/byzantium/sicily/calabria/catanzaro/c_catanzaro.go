package catanzaro

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CatanzaroCounty interface {
	feud.County
	BBelcastro贝尔卡斯特罗() feud.Barony
	BCatanzaro卡坦扎罗() feud.Barony
	BCrotone克罗托内() feud.Barony
	BGirifalco吉里法尔科() feud.Barony
	BKyterion基泰里翁() feud.Barony
	BSquillace斯奎拉切() feud.Barony
}

type 卡坦扎罗CatanzaroCounty struct {
	feud.BaseCounty
	贝尔卡斯特罗Belcastro feud.Barony
	卡坦扎罗Catanzaro   feud.Barony
	克罗托内Crotone     feud.Barony
	吉里法尔科Girifalco  feud.Barony
	基泰里翁Kyterion    feud.Barony
	斯奎拉切Squillace   feud.Barony
}

func (c *卡坦扎罗CatanzaroCounty) BBelcastro贝尔卡斯特罗() feud.Barony {
	return c.贝尔卡斯特罗Belcastro
}

func (c *卡坦扎罗CatanzaroCounty) BCatanzaro卡坦扎罗() feud.Barony {
	return c.卡坦扎罗Catanzaro
}

func (c *卡坦扎罗CatanzaroCounty) BCrotone克罗托内() feud.Barony {
	return c.克罗托内Crotone
}

func (c *卡坦扎罗CatanzaroCounty) BGirifalco吉里法尔科() feud.Barony {
	return c.吉里法尔科Girifalco
}

func (c *卡坦扎罗CatanzaroCounty) BKyterion基泰里翁() feud.Barony {
	return c.基泰里翁Kyterion
}

func (c *卡坦扎罗CatanzaroCounty) BSquillace斯奎拉切() feud.Barony {
	return c.斯奎拉切Squillace
}

var CCatanzaro卡坦扎罗 CatanzaroCounty = &卡坦扎罗CatanzaroCounty{}

func init() {
	f := CCatanzaro卡坦扎罗.(*卡坦扎罗CatanzaroCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1712",
		Title:     "catanzaro",
		TitleName: "卡坦扎罗",
		TitleCode: "c_catanzaro",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝尔卡斯特罗Belcastro = BBelcastro贝尔卡斯特罗
	f.贝尔卡斯特罗Belcastro.SetParent(f)

	f.卡坦扎罗Catanzaro = BCatanzaro卡坦扎罗
	f.卡坦扎罗Catanzaro.SetParent(f)

	f.克罗托内Crotone = BCrotone克罗托内
	f.克罗托内Crotone.SetParent(f)

	f.吉里法尔科Girifalco = BGirifalco吉里法尔科
	f.吉里法尔科Girifalco.SetParent(f)

	f.基泰里翁Kyterion = BKyterion基泰里翁
	f.基泰里翁Kyterion.SetParent(f)

	f.斯奎拉切Squillace = BSquillace斯奎拉切
	f.斯奎拉切Squillace.SetParent(f)

}
