package maymana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MaymanaCounty interface {
	feud.County
	BAlmar阿尔玛勒() feud.Barony
	BBasin贝欣() feud.Barony
	BBilcheragh比尔切拉格() feud.Barony
	BDarzab杜尔札卜() feud.Barony
	BDihekhala迪厄哈拉() feud.Barony
	BGurziwan古尔齐万() feud.Barony
	BMaymana梅马内() feud.Barony
}

type 梅马内MaymanaCounty struct {
	feud.BaseCounty
	阿尔玛勒Almar       feud.Barony
	贝欣Basin         feud.Barony
	比尔切拉格Bilcheragh feud.Barony
	杜尔札卜Darzab      feud.Barony
	迪厄哈拉Dihekhala   feud.Barony
	古尔齐万Gurziwan    feud.Barony
	梅马内Maymana      feud.Barony
}

func (c *梅马内MaymanaCounty) BAlmar阿尔玛勒() feud.Barony {
	return c.阿尔玛勒Almar
}

func (c *梅马内MaymanaCounty) BBasin贝欣() feud.Barony {
	return c.贝欣Basin
}

func (c *梅马内MaymanaCounty) BBilcheragh比尔切拉格() feud.Barony {
	return c.比尔切拉格Bilcheragh
}

func (c *梅马内MaymanaCounty) BDarzab杜尔札卜() feud.Barony {
	return c.杜尔札卜Darzab
}

func (c *梅马内MaymanaCounty) BDihekhala迪厄哈拉() feud.Barony {
	return c.迪厄哈拉Dihekhala
}

func (c *梅马内MaymanaCounty) BGurziwan古尔齐万() feud.Barony {
	return c.古尔齐万Gurziwan
}

func (c *梅马内MaymanaCounty) BMaymana梅马内() feud.Barony {
	return c.梅马内Maymana
}

var CMaymana梅马内 MaymanaCounty = &梅马内MaymanaCounty{}

func init() {
	f := CMaymana梅马内.(*梅马内MaymanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1377",
		Title:     "maymana",
		TitleName: "梅马内",
		TitleCode: "c_maymana",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔玛勒Almar = BAlmar阿尔玛勒
	f.阿尔玛勒Almar.SetParent(f)

	f.贝欣Basin = BBasin贝欣
	f.贝欣Basin.SetParent(f)

	f.比尔切拉格Bilcheragh = BBilcheragh比尔切拉格
	f.比尔切拉格Bilcheragh.SetParent(f)

	f.杜尔札卜Darzab = BDarzab杜尔札卜
	f.杜尔札卜Darzab.SetParent(f)

	f.迪厄哈拉Dihekhala = BDihekhala迪厄哈拉
	f.迪厄哈拉Dihekhala.SetParent(f)

	f.古尔齐万Gurziwan = BGurziwan古尔齐万
	f.古尔齐万Gurziwan.SetParent(f)

	f.梅马内Maymana = BMaymana梅马内
	f.梅马内Maymana.SetParent(f)

}
