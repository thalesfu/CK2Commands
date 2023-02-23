package hesse

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HesseCounty interface {
	feud.County
	BBlankenstein布兰肯施泰因() feud.Barony
	BGladenbach格拉登巴赫() feud.Barony
	BHesse黑森() feud.Barony
	BKirchhain基希海恩() feud.Barony
	BMarburg马堡() feud.Barony
	BMerenberg梅伦贝格() feud.Barony
	BWetzlar韦茨拉尔() feud.Barony
}

type 马堡HesseCounty struct {
	feud.BaseCounty
	布兰肯施泰因Blankenstein feud.Barony
	格拉登巴赫Gladenbach    feud.Barony
	黑森Hesse            feud.Barony
	基希海恩Kirchhain      feud.Barony
	马堡Marburg          feud.Barony
	梅伦贝格Merenberg      feud.Barony
	韦茨拉尔Wetzlar        feud.Barony
}

func (c *马堡HesseCounty) BBlankenstein布兰肯施泰因() feud.Barony {
	return c.布兰肯施泰因Blankenstein
}

func (c *马堡HesseCounty) BGladenbach格拉登巴赫() feud.Barony {
	return c.格拉登巴赫Gladenbach
}

func (c *马堡HesseCounty) BHesse黑森() feud.Barony {
	return c.黑森Hesse
}

func (c *马堡HesseCounty) BKirchhain基希海恩() feud.Barony {
	return c.基希海恩Kirchhain
}

func (c *马堡HesseCounty) BMarburg马堡() feud.Barony {
	return c.马堡Marburg
}

func (c *马堡HesseCounty) BMerenberg梅伦贝格() feud.Barony {
	return c.梅伦贝格Merenberg
}

func (c *马堡HesseCounty) BWetzlar韦茨拉尔() feud.Barony {
	return c.韦茨拉尔Wetzlar
}

var CHesse马堡 HesseCounty = &马堡HesseCounty{}

func init() {
	f := CHesse马堡.(*马堡HesseCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1987",
		Title:     "hesse",
		TitleName: "马堡",
		TitleCode: "c_hesse",
		Baronies:  map[string]feud.Barony{},
	}

	f.布兰肯施泰因Blankenstein = BBlankenstein布兰肯施泰因
	f.布兰肯施泰因Blankenstein.SetParent(f)

	f.格拉登巴赫Gladenbach = BGladenbach格拉登巴赫
	f.格拉登巴赫Gladenbach.SetParent(f)

	f.黑森Hesse = BHesse黑森
	f.黑森Hesse.SetParent(f)

	f.基希海恩Kirchhain = BKirchhain基希海恩
	f.基希海恩Kirchhain.SetParent(f)

	f.马堡Marburg = BMarburg马堡
	f.马堡Marburg.SetParent(f)

	f.梅伦贝格Merenberg = BMerenberg梅伦贝格
	f.梅伦贝格Merenberg.SetParent(f)

	f.韦茨拉尔Wetzlar = BWetzlar韦茨拉尔
	f.韦茨拉尔Wetzlar.SetParent(f)

}
