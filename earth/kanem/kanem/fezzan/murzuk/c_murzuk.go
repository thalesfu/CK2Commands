package murzuk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MurzukCounty interface {
	feud.County
	BDjarma杰尔迈() feud.Barony
	BMassah马萨() feud.Barony
	BMurzuk迈尔祖格() feud.Barony
	BSebna塞卜奈() feud.Barony
	BTassawa泰萨韦() feud.Barony
	BTmessa特梅萨() feud.Barony
}

type 迈尔祖格MurzukCounty struct {
	feud.BaseCounty
	杰尔迈Djarma  feud.Barony
	马萨Massah   feud.Barony
	迈尔祖格Murzuk feud.Barony
	塞卜奈Sebna   feud.Barony
	泰萨韦Tassawa feud.Barony
	特梅萨Tmessa  feud.Barony
}

func (c *迈尔祖格MurzukCounty) BDjarma杰尔迈() feud.Barony {
	return c.杰尔迈Djarma
}

func (c *迈尔祖格MurzukCounty) BMassah马萨() feud.Barony {
	return c.马萨Massah
}

func (c *迈尔祖格MurzukCounty) BMurzuk迈尔祖格() feud.Barony {
	return c.迈尔祖格Murzuk
}

func (c *迈尔祖格MurzukCounty) BSebna塞卜奈() feud.Barony {
	return c.塞卜奈Sebna
}

func (c *迈尔祖格MurzukCounty) BTassawa泰萨韦() feud.Barony {
	return c.泰萨韦Tassawa
}

func (c *迈尔祖格MurzukCounty) BTmessa特梅萨() feud.Barony {
	return c.特梅萨Tmessa
}

var CMurzuk迈尔祖格 MurzukCounty = &迈尔祖格MurzukCounty{}

func init() {
	f := CMurzuk迈尔祖格.(*迈尔祖格MurzukCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1730",
		Title:     "murzuk",
		TitleName: "迈尔祖格",
		TitleCode: "c_murzuk",
		Baronies:  map[string]feud.Barony{},
	}

	f.杰尔迈Djarma = BDjarma杰尔迈
	f.杰尔迈Djarma.SetParent(f)

	f.马萨Massah = BMassah马萨
	f.马萨Massah.SetParent(f)

	f.迈尔祖格Murzuk = BMurzuk迈尔祖格
	f.迈尔祖格Murzuk.SetParent(f)

	f.塞卜奈Sebna = BSebna塞卜奈
	f.塞卜奈Sebna.SetParent(f)

	f.泰萨韦Tassawa = BTassawa泰萨韦
	f.泰萨韦Tassawa.SetParent(f)

	f.特梅萨Tmessa = BTmessa特梅萨
	f.特梅萨Tmessa.SetParent(f)

}
