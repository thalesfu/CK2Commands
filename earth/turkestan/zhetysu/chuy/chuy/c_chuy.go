package chuy

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChuyCounty interface {
	feud.County
	BAlmatu阿拉木图() feud.Barony
	BBalasagun裴罗将军城() feud.Barony
	BBishkek比什凯克() feud.Barony
	BKeru克鲁() feud.Barony
	BKorday科尔代() feud.Barony
	BSuyab碎叶() feud.Barony
	BTokmok托克马克() feud.Barony
}

type 裴罗将军城ChuyCounty struct {
	feud.BaseCounty
	阿拉木图Almatu     feud.Barony
	裴罗将军城Balasagun feud.Barony
	比什凯克Bishkek    feud.Barony
	克鲁Keru         feud.Barony
	科尔代Korday      feud.Barony
	碎叶Suyab        feud.Barony
	托克马克Tokmok     feud.Barony
}

func (c *裴罗将军城ChuyCounty) BAlmatu阿拉木图() feud.Barony {
	return c.阿拉木图Almatu
}

func (c *裴罗将军城ChuyCounty) BBalasagun裴罗将军城() feud.Barony {
	return c.裴罗将军城Balasagun
}

func (c *裴罗将军城ChuyCounty) BBishkek比什凯克() feud.Barony {
	return c.比什凯克Bishkek
}

func (c *裴罗将军城ChuyCounty) BKeru克鲁() feud.Barony {
	return c.克鲁Keru
}

func (c *裴罗将军城ChuyCounty) BKorday科尔代() feud.Barony {
	return c.科尔代Korday
}

func (c *裴罗将军城ChuyCounty) BSuyab碎叶() feud.Barony {
	return c.碎叶Suyab
}

func (c *裴罗将军城ChuyCounty) BTokmok托克马克() feud.Barony {
	return c.托克马克Tokmok
}

var CChuy裴罗将军城 ChuyCounty = &裴罗将军城ChuyCounty{}

func init() {
	f := CChuy裴罗将军城.(*裴罗将军城ChuyCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1424",
		Title:     "chuy",
		TitleName: "裴罗将军城",
		TitleCode: "c_chuy",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿拉木图Almatu = BAlmatu阿拉木图
	f.阿拉木图Almatu.SetParent(f)

	f.裴罗将军城Balasagun = BBalasagun裴罗将军城
	f.裴罗将军城Balasagun.SetParent(f)

	f.比什凯克Bishkek = BBishkek比什凯克
	f.比什凯克Bishkek.SetParent(f)

	f.克鲁Keru = BKeru克鲁
	f.克鲁Keru.SetParent(f)

	f.科尔代Korday = BKorday科尔代
	f.科尔代Korday.SetParent(f)

	f.碎叶Suyab = BSuyab碎叶
	f.碎叶Suyab.SetParent(f)

	f.托克马克Tokmok = BTokmok托克马克
	f.托克马克Tokmok.SetParent(f)

}
