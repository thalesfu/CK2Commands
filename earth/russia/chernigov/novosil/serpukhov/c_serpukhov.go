package serpukhov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SerpukhovCounty interface {
	feud.County
	BBogucharovo博古恰罗沃() feud.Barony
	BDugna杜格纳() feud.Barony
	BKolyupanovo科柳帕诺沃() feud.Barony
	BNenashevo涅纳舍沃() feud.Barony
	BNovogurovsky诺沃古罗夫斯基() feud.Barony
	BSerpukhov谢尔普霍夫() feud.Barony
	BZaokskiy扎奥克斯基() feud.Barony
}

type 谢尔普霍夫SerpukhovCounty struct {
	feud.BaseCounty
	博古恰罗沃Bogucharovo    feud.Barony
	杜格纳Dugna            feud.Barony
	科柳帕诺沃Kolyupanovo    feud.Barony
	涅纳舍沃Nenashevo       feud.Barony
	诺沃古罗夫斯基Novogurovsky feud.Barony
	谢尔普霍夫Serpukhov      feud.Barony
	扎奥克斯基Zaokskiy       feud.Barony
}

func (c *谢尔普霍夫SerpukhovCounty) BBogucharovo博古恰罗沃() feud.Barony {
	return c.博古恰罗沃Bogucharovo
}

func (c *谢尔普霍夫SerpukhovCounty) BDugna杜格纳() feud.Barony {
	return c.杜格纳Dugna
}

func (c *谢尔普霍夫SerpukhovCounty) BKolyupanovo科柳帕诺沃() feud.Barony {
	return c.科柳帕诺沃Kolyupanovo
}

func (c *谢尔普霍夫SerpukhovCounty) BNenashevo涅纳舍沃() feud.Barony {
	return c.涅纳舍沃Nenashevo
}

func (c *谢尔普霍夫SerpukhovCounty) BNovogurovsky诺沃古罗夫斯基() feud.Barony {
	return c.诺沃古罗夫斯基Novogurovsky
}

func (c *谢尔普霍夫SerpukhovCounty) BSerpukhov谢尔普霍夫() feud.Barony {
	return c.谢尔普霍夫Serpukhov
}

func (c *谢尔普霍夫SerpukhovCounty) BZaokskiy扎奥克斯基() feud.Barony {
	return c.扎奥克斯基Zaokskiy
}

var CSerpukhov谢尔普霍夫 SerpukhovCounty = &谢尔普霍夫SerpukhovCounty{}

func init() {
	f := CSerpukhov谢尔普霍夫.(*谢尔普霍夫SerpukhovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1679",
		Title:     "serpukhov",
		TitleName: "谢尔普霍夫",
		TitleCode: "c_serpukhov",
		Baronies:  map[string]feud.Barony{},
	}

	f.博古恰罗沃Bogucharovo = BBogucharovo博古恰罗沃
	f.博古恰罗沃Bogucharovo.SetParent(f)

	f.杜格纳Dugna = BDugna杜格纳
	f.杜格纳Dugna.SetParent(f)

	f.科柳帕诺沃Kolyupanovo = BKolyupanovo科柳帕诺沃
	f.科柳帕诺沃Kolyupanovo.SetParent(f)

	f.涅纳舍沃Nenashevo = BNenashevo涅纳舍沃
	f.涅纳舍沃Nenashevo.SetParent(f)

	f.诺沃古罗夫斯基Novogurovsky = BNovogurovsky诺沃古罗夫斯基
	f.诺沃古罗夫斯基Novogurovsky.SetParent(f)

	f.谢尔普霍夫Serpukhov = BSerpukhov谢尔普霍夫
	f.谢尔普霍夫Serpukhov.SetParent(f)

	f.扎奥克斯基Zaokskiy = BZaokskiy扎奥克斯基
	f.扎奥克斯基Zaokskiy.SetParent(f)

}
