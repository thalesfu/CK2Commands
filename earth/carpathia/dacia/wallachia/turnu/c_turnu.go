package turnu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurnuCounty interface {
	feud.County
	BBucuresti布加勒斯特() feud.Barony
	BComana科马纳() feud.Barony
	BGiurgiu久尔久() feud.Barony
	BGlavacioc格拉瓦乔克() feud.Barony
	BRosiorriidevede罗肖里德韦代() feud.Barony
	BTurnu图尔努() feud.Barony
	BZimnicea济姆尼恰() feud.Barony
}

type 图尔努TurnuCounty struct {
	feud.BaseCounty
	布加勒斯特Bucuresti        feud.Barony
	科马纳Comana             feud.Barony
	久尔久Giurgiu            feud.Barony
	格拉瓦乔克Glavacioc        feud.Barony
	罗肖里德韦代Rosiorriidevede feud.Barony
	图尔努Turnu              feud.Barony
	济姆尼恰Zimnicea          feud.Barony
}

func (c *图尔努TurnuCounty) BBucuresti布加勒斯特() feud.Barony {
	return c.布加勒斯特Bucuresti
}

func (c *图尔努TurnuCounty) BComana科马纳() feud.Barony {
	return c.科马纳Comana
}

func (c *图尔努TurnuCounty) BGiurgiu久尔久() feud.Barony {
	return c.久尔久Giurgiu
}

func (c *图尔努TurnuCounty) BGlavacioc格拉瓦乔克() feud.Barony {
	return c.格拉瓦乔克Glavacioc
}

func (c *图尔努TurnuCounty) BRosiorriidevede罗肖里德韦代() feud.Barony {
	return c.罗肖里德韦代Rosiorriidevede
}

func (c *图尔努TurnuCounty) BTurnu图尔努() feud.Barony {
	return c.图尔努Turnu
}

func (c *图尔努TurnuCounty) BZimnicea济姆尼恰() feud.Barony {
	return c.济姆尼恰Zimnicea
}

var CTurnu图尔努 TurnuCounty = &图尔努TurnuCounty{}

func init() {
	f := CTurnu图尔努.(*图尔努TurnuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "514",
		Title:     "turnu",
		TitleName: "图尔努",
		TitleCode: "c_turnu",
		Baronies:  map[string]feud.Barony{},
	}

	f.布加勒斯特Bucuresti = BBucuresti布加勒斯特
	f.布加勒斯特Bucuresti.SetParent(f)

	f.科马纳Comana = BComana科马纳
	f.科马纳Comana.SetParent(f)

	f.久尔久Giurgiu = BGiurgiu久尔久
	f.久尔久Giurgiu.SetParent(f)

	f.格拉瓦乔克Glavacioc = BGlavacioc格拉瓦乔克
	f.格拉瓦乔克Glavacioc.SetParent(f)

	f.罗肖里德韦代Rosiorriidevede = BRosiorriidevede罗肖里德韦代
	f.罗肖里德韦代Rosiorriidevede.SetParent(f)

	f.图尔努Turnu = BTurnu图尔努
	f.图尔努Turnu.SetParent(f)

	f.济姆尼恰Zimnicea = BZimnicea济姆尼恰
	f.济姆尼恰Zimnicea.SetParent(f)

}
