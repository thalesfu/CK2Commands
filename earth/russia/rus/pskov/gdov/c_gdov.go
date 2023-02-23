package gdov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GdovCounty interface {
	feud.County
	BGdov格多夫() feud.Barony
	BKolydukha科雷杜哈() feud.Barony
	BMda姆达() feud.Barony
	BPlyussa普柳萨() feud.Barony
	BStrizhki斯特里日基() feud.Barony
	BVyazki维亚兹基() feud.Barony
	BZakhodsy扎霍齐() feud.Barony
}

type 格多夫GdovCounty struct {
	feud.BaseCounty
	格多夫Gdov       feud.Barony
	科雷杜哈Kolydukha feud.Barony
	姆达Mda         feud.Barony
	普柳萨Plyussa    feud.Barony
	斯特里日基Strizhki feud.Barony
	维亚兹基Vyazki    feud.Barony
	扎霍齐Zakhodsy   feud.Barony
}

func (c *格多夫GdovCounty) BGdov格多夫() feud.Barony {
	return c.格多夫Gdov
}

func (c *格多夫GdovCounty) BKolydukha科雷杜哈() feud.Barony {
	return c.科雷杜哈Kolydukha
}

func (c *格多夫GdovCounty) BMda姆达() feud.Barony {
	return c.姆达Mda
}

func (c *格多夫GdovCounty) BPlyussa普柳萨() feud.Barony {
	return c.普柳萨Plyussa
}

func (c *格多夫GdovCounty) BStrizhki斯特里日基() feud.Barony {
	return c.斯特里日基Strizhki
}

func (c *格多夫GdovCounty) BVyazki维亚兹基() feud.Barony {
	return c.维亚兹基Vyazki
}

func (c *格多夫GdovCounty) BZakhodsy扎霍齐() feud.Barony {
	return c.扎霍齐Zakhodsy
}

var CGdov格多夫 GdovCounty = &格多夫GdovCounty{}

func init() {
	f := CGdov格多夫.(*格多夫GdovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1661",
		Title:     "gdov",
		TitleName: "格多夫",
		TitleCode: "c_gdov",
		Baronies:  map[string]feud.Barony{},
	}

	f.格多夫Gdov = BGdov格多夫
	f.格多夫Gdov.SetParent(f)

	f.科雷杜哈Kolydukha = BKolydukha科雷杜哈
	f.科雷杜哈Kolydukha.SetParent(f)

	f.姆达Mda = BMda姆达
	f.姆达Mda.SetParent(f)

	f.普柳萨Plyussa = BPlyussa普柳萨
	f.普柳萨Plyussa.SetParent(f)

	f.斯特里日基Strizhki = BStrizhki斯特里日基
	f.斯特里日基Strizhki.SetParent(f)

	f.维亚兹基Vyazki = BVyazki维亚兹基
	f.维亚兹基Vyazki.SetParent(f)

	f.扎霍齐Zakhodsy = BZakhodsy扎霍齐
	f.扎霍齐Zakhodsy.SetParent(f)

}
