package rimini

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RiminiCounty interface {
	feud.County
	BBellaria贝拉里亚() feud.Barony
	BGabicce加比切() feud.Barony
	BGatteo加泰奥() feud.Barony
	BPennabilli彭纳比利() feud.Barony
	BPesaro佩萨罗() feud.Barony
	BRimini里米尼() feud.Barony
}

type 里米尼RiminiCounty struct {
	feud.BaseCounty
	贝拉里亚Bellaria   feud.Barony
	加比切Gabicce     feud.Barony
	加泰奥Gatteo      feud.Barony
	彭纳比利Pennabilli feud.Barony
	佩萨罗Pesaro      feud.Barony
	里米尼Rimini      feud.Barony
}

func (c *里米尼RiminiCounty) BBellaria贝拉里亚() feud.Barony {
	return c.贝拉里亚Bellaria
}

func (c *里米尼RiminiCounty) BGabicce加比切() feud.Barony {
	return c.加比切Gabicce
}

func (c *里米尼RiminiCounty) BGatteo加泰奥() feud.Barony {
	return c.加泰奥Gatteo
}

func (c *里米尼RiminiCounty) BPennabilli彭纳比利() feud.Barony {
	return c.彭纳比利Pennabilli
}

func (c *里米尼RiminiCounty) BPesaro佩萨罗() feud.Barony {
	return c.佩萨罗Pesaro
}

func (c *里米尼RiminiCounty) BRimini里米尼() feud.Barony {
	return c.里米尼Rimini
}

var CRimini里米尼 RiminiCounty = &里米尼RiminiCounty{}

func init() {
	f := CRimini里米尼.(*里米尼RiminiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1711",
		Title:     "rimini",
		TitleName: "里米尼",
		TitleCode: "c_rimini",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝拉里亚Bellaria = BBellaria贝拉里亚
	f.贝拉里亚Bellaria.SetParent(f)

	f.加比切Gabicce = BGabicce加比切
	f.加比切Gabicce.SetParent(f)

	f.加泰奥Gatteo = BGatteo加泰奥
	f.加泰奥Gatteo.SetParent(f)

	f.彭纳比利Pennabilli = BPennabilli彭纳比利
	f.彭纳比利Pennabilli.SetParent(f)

	f.佩萨罗Pesaro = BPesaro佩萨罗
	f.佩萨罗Pesaro.SetParent(f)

	f.里米尼Rimini = BRimini里米尼
	f.里米尼Rimini.SetParent(f)

}
