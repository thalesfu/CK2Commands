package minden

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MindenCounty interface {
	feud.County
	BAuburg奥堡() feud.Barony
	BDiepholz迪普霍尔茨() feud.Barony
	BDrebber德雷伯() feud.Barony
	BHockeleve霍克莱韦() feud.Barony
	BMinden明登() feud.Barony
	BPetershagen彼得斯哈根() feud.Barony
	BWagenfeld瓦根费尔德() feud.Barony
}

type 明登MindenCounty struct {
	feud.BaseCounty
	奥堡Auburg         feud.Barony
	迪普霍尔茨Diepholz    feud.Barony
	德雷伯Drebber       feud.Barony
	霍克莱韦Hockeleve    feud.Barony
	明登Minden         feud.Barony
	彼得斯哈根Petershagen feud.Barony
	瓦根费尔德Wagenfeld   feud.Barony
}

func (c *明登MindenCounty) BAuburg奥堡() feud.Barony {
	return c.奥堡Auburg
}

func (c *明登MindenCounty) BDiepholz迪普霍尔茨() feud.Barony {
	return c.迪普霍尔茨Diepholz
}

func (c *明登MindenCounty) BDrebber德雷伯() feud.Barony {
	return c.德雷伯Drebber
}

func (c *明登MindenCounty) BHockeleve霍克莱韦() feud.Barony {
	return c.霍克莱韦Hockeleve
}

func (c *明登MindenCounty) BMinden明登() feud.Barony {
	return c.明登Minden
}

func (c *明登MindenCounty) BPetershagen彼得斯哈根() feud.Barony {
	return c.彼得斯哈根Petershagen
}

func (c *明登MindenCounty) BWagenfeld瓦根费尔德() feud.Barony {
	return c.瓦根费尔德Wagenfeld
}

var CMinden明登 MindenCounty = &明登MindenCounty{}

func init() {
	f := CMinden明登.(*明登MindenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1979",
		Title:     "minden",
		TitleName: "明登",
		TitleCode: "c_minden",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥堡Auburg = BAuburg奥堡
	f.奥堡Auburg.SetParent(f)

	f.迪普霍尔茨Diepholz = BDiepholz迪普霍尔茨
	f.迪普霍尔茨Diepholz.SetParent(f)

	f.德雷伯Drebber = BDrebber德雷伯
	f.德雷伯Drebber.SetParent(f)

	f.霍克莱韦Hockeleve = BHockeleve霍克莱韦
	f.霍克莱韦Hockeleve.SetParent(f)

	f.明登Minden = BMinden明登
	f.明登Minden.SetParent(f)

	f.彼得斯哈根Petershagen = BPetershagen彼得斯哈根
	f.彼得斯哈根Petershagen.SetParent(f)

	f.瓦根费尔德Wagenfeld = BWagenfeld瓦根费尔德
	f.瓦根费尔德Wagenfeld.SetParent(f)

}
