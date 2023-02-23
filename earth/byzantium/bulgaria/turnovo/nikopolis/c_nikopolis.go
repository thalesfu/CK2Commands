package nikopolis

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NikopolisCounty interface {
	feud.County
	BBelene贝莱内() feud.Barony
	BDolnidabknik下德布尼克() feud.Barony
	BIskar伊斯克尔() feud.Barony
	BKnezha克内扎() feud.Barony
	BNikopolis尼科波利斯() feud.Barony
	BOescus埃斯库斯() feud.Barony
	BPleven普莱文() feud.Barony
	BPordim波尔迪姆() feud.Barony
}

type 尼科波利斯NikopolisCounty struct {
	feud.BaseCounty
	贝莱内Belene         feud.Barony
	下德布尼克Dolnidabknik feud.Barony
	伊斯克尔Iskar         feud.Barony
	克内扎Knezha         feud.Barony
	尼科波利斯Nikopolis    feud.Barony
	埃斯库斯Oescus        feud.Barony
	普莱文Pleven         feud.Barony
	波尔迪姆Pordim        feud.Barony
}

func (c *尼科波利斯NikopolisCounty) BBelene贝莱内() feud.Barony {
	return c.贝莱内Belene
}

func (c *尼科波利斯NikopolisCounty) BDolnidabknik下德布尼克() feud.Barony {
	return c.下德布尼克Dolnidabknik
}

func (c *尼科波利斯NikopolisCounty) BIskar伊斯克尔() feud.Barony {
	return c.伊斯克尔Iskar
}

func (c *尼科波利斯NikopolisCounty) BKnezha克内扎() feud.Barony {
	return c.克内扎Knezha
}

func (c *尼科波利斯NikopolisCounty) BNikopolis尼科波利斯() feud.Barony {
	return c.尼科波利斯Nikopolis
}

func (c *尼科波利斯NikopolisCounty) BOescus埃斯库斯() feud.Barony {
	return c.埃斯库斯Oescus
}

func (c *尼科波利斯NikopolisCounty) BPleven普莱文() feud.Barony {
	return c.普莱文Pleven
}

func (c *尼科波利斯NikopolisCounty) BPordim波尔迪姆() feud.Barony {
	return c.波尔迪姆Pordim
}

var CNikopolis尼科波利斯 NikopolisCounty = &尼科波利斯NikopolisCounty{}

func init() {
	f := CNikopolis尼科波利斯.(*尼科波利斯NikopolisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "507",
		Title:     "nikopolis",
		TitleName: "尼科波利斯",
		TitleCode: "c_nikopolis",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝莱内Belene = BBelene贝莱内
	f.贝莱内Belene.SetParent(f)

	f.下德布尼克Dolnidabknik = BDolnidabknik下德布尼克
	f.下德布尼克Dolnidabknik.SetParent(f)

	f.伊斯克尔Iskar = BIskar伊斯克尔
	f.伊斯克尔Iskar.SetParent(f)

	f.克内扎Knezha = BKnezha克内扎
	f.克内扎Knezha.SetParent(f)

	f.尼科波利斯Nikopolis = BNikopolis尼科波利斯
	f.尼科波利斯Nikopolis.SetParent(f)

	f.埃斯库斯Oescus = BOescus埃斯库斯
	f.埃斯库斯Oescus.SetParent(f)

	f.普莱文Pleven = BPleven普莱文
	f.普莱文Pleven.SetParent(f)

	f.波尔迪姆Pordim = BPordim波尔迪姆
	f.波尔迪姆Pordim.SetParent(f)

}
