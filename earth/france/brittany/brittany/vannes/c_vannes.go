package vannes

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VannesCounty interface {
	feud.County
	BAuray欧赖() feud.Barony
	BHennebont埃讷邦() feud.Barony
	BJosselin若瑟兰() feud.Barony
	BLocmine洛克米内() feud.Barony
	BPloermel普洛埃梅勒() feud.Barony
	BPontivy蓬蒂维() feud.Barony
	BVannes瓦讷() feud.Barony
}

type 瓦讷VannesCounty struct {
	feud.BaseCounty
	欧赖Auray       feud.Barony
	埃讷邦Hennebont  feud.Barony
	若瑟兰Josselin   feud.Barony
	洛克米内Locmine   feud.Barony
	普洛埃梅勒Ploermel feud.Barony
	蓬蒂维Pontivy    feud.Barony
	瓦讷Vannes      feud.Barony
}

func (c *瓦讷VannesCounty) BAuray欧赖() feud.Barony {
	return c.欧赖Auray
}

func (c *瓦讷VannesCounty) BHennebont埃讷邦() feud.Barony {
	return c.埃讷邦Hennebont
}

func (c *瓦讷VannesCounty) BJosselin若瑟兰() feud.Barony {
	return c.若瑟兰Josselin
}

func (c *瓦讷VannesCounty) BLocmine洛克米内() feud.Barony {
	return c.洛克米内Locmine
}

func (c *瓦讷VannesCounty) BPloermel普洛埃梅勒() feud.Barony {
	return c.普洛埃梅勒Ploermel
}

func (c *瓦讷VannesCounty) BPontivy蓬蒂维() feud.Barony {
	return c.蓬蒂维Pontivy
}

func (c *瓦讷VannesCounty) BVannes瓦讷() feud.Barony {
	return c.瓦讷Vannes
}

var CVannes瓦讷 VannesCounty = &瓦讷VannesCounty{}

func init() {
	f := CVannes瓦讷.(*瓦讷VannesCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "105",
		Title:     "vannes",
		TitleName: "瓦讷",
		TitleCode: "c_vannes",
		Baronies:  map[string]feud.Barony{},
	}

	f.欧赖Auray = BAuray欧赖
	f.欧赖Auray.SetParent(f)

	f.埃讷邦Hennebont = BHennebont埃讷邦
	f.埃讷邦Hennebont.SetParent(f)

	f.若瑟兰Josselin = BJosselin若瑟兰
	f.若瑟兰Josselin.SetParent(f)

	f.洛克米内Locmine = BLocmine洛克米内
	f.洛克米内Locmine.SetParent(f)

	f.普洛埃梅勒Ploermel = BPloermel普洛埃梅勒
	f.普洛埃梅勒Ploermel.SetParent(f)

	f.蓬蒂维Pontivy = BPontivy蓬蒂维
	f.蓬蒂维Pontivy.SetParent(f)

	f.瓦讷Vannes = BVannes瓦讷
	f.瓦讷Vannes.SetParent(f)

}
