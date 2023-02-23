package lakomelza

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LakomelzaCounty interface {
	feud.County
	BCherkwa车尔科瓦() feud.Barony
	BDessye德西() feud.Barony
	BHabru哈布茹() feud.Barony
	BKembolcha孔博勒查() feud.Barony
	BMetene麼忒讷() feud.Barony
	BTigaja提嘎加() feud.Barony
	BZiya兹雅() feud.Barony
}

type 拉科美尔扎LakomelzaCounty struct {
	feud.BaseCounty
	车尔科瓦Cherkwa   feud.Barony
	德西Dessye      feud.Barony
	哈布茹Habru      feud.Barony
	孔博勒查Kembolcha feud.Barony
	麼忒讷Metene     feud.Barony
	提嘎加Tigaja     feud.Barony
	兹雅Ziya        feud.Barony
}

func (c *拉科美尔扎LakomelzaCounty) BCherkwa车尔科瓦() feud.Barony {
	return c.车尔科瓦Cherkwa
}

func (c *拉科美尔扎LakomelzaCounty) BDessye德西() feud.Barony {
	return c.德西Dessye
}

func (c *拉科美尔扎LakomelzaCounty) BHabru哈布茹() feud.Barony {
	return c.哈布茹Habru
}

func (c *拉科美尔扎LakomelzaCounty) BKembolcha孔博勒查() feud.Barony {
	return c.孔博勒查Kembolcha
}

func (c *拉科美尔扎LakomelzaCounty) BMetene麼忒讷() feud.Barony {
	return c.麼忒讷Metene
}

func (c *拉科美尔扎LakomelzaCounty) BTigaja提嘎加() feud.Barony {
	return c.提嘎加Tigaja
}

func (c *拉科美尔扎LakomelzaCounty) BZiya兹雅() feud.Barony {
	return c.兹雅Ziya
}

var CLakomelza拉科美尔扎 LakomelzaCounty = &拉科美尔扎LakomelzaCounty{}

func init() {
	f := CLakomelza拉科美尔扎.(*拉科美尔扎LakomelzaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1436",
		Title:     "lakomelza",
		TitleName: "拉科美尔扎",
		TitleCode: "c_lakomelza",
		Baronies:  map[string]feud.Barony{},
	}

	f.车尔科瓦Cherkwa = BCherkwa车尔科瓦
	f.车尔科瓦Cherkwa.SetParent(f)

	f.德西Dessye = BDessye德西
	f.德西Dessye.SetParent(f)

	f.哈布茹Habru = BHabru哈布茹
	f.哈布茹Habru.SetParent(f)

	f.孔博勒查Kembolcha = BKembolcha孔博勒查
	f.孔博勒查Kembolcha.SetParent(f)

	f.麼忒讷Metene = BMetene麼忒讷
	f.麼忒讷Metene.SetParent(f)

	f.提嘎加Tigaja = BTigaja提嘎加
	f.提嘎加Tigaja.SetParent(f)

	f.兹雅Ziya = BZiya兹雅
	f.兹雅Ziya.SetParent(f)

}
