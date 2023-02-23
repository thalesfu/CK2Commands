package tigris

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TigrisCounty interface {
	feud.County
	BBismaya比斯玛亚() feud.Barony
	BIshan伊斯罕() feud.Barony
	BMajaralkabir马贾阿卡比尔() feud.Barony
	BNuffar努法尔() feud.Barony
	BQalatsjergat喀拉特斯吉格特() feud.Barony
	BSamawah萨马瓦() feud.Barony
	BTellelhiba特尔赫巴() feud.Barony
}

type 苏萨TigrisCounty struct {
	feud.BaseCounty
	比斯玛亚Bismaya         feud.Barony
	伊斯罕Ishan            feud.Barony
	马贾阿卡比尔Majaralkabir  feud.Barony
	努法尔Nuffar           feud.Barony
	喀拉特斯吉格特Qalatsjergat feud.Barony
	萨马瓦Samawah          feud.Barony
	特尔赫巴Tellelhiba      feud.Barony
}

func (c *苏萨TigrisCounty) BBismaya比斯玛亚() feud.Barony {
	return c.比斯玛亚Bismaya
}

func (c *苏萨TigrisCounty) BIshan伊斯罕() feud.Barony {
	return c.伊斯罕Ishan
}

func (c *苏萨TigrisCounty) BMajaralkabir马贾阿卡比尔() feud.Barony {
	return c.马贾阿卡比尔Majaralkabir
}

func (c *苏萨TigrisCounty) BNuffar努法尔() feud.Barony {
	return c.努法尔Nuffar
}

func (c *苏萨TigrisCounty) BQalatsjergat喀拉特斯吉格特() feud.Barony {
	return c.喀拉特斯吉格特Qalatsjergat
}

func (c *苏萨TigrisCounty) BSamawah萨马瓦() feud.Barony {
	return c.萨马瓦Samawah
}

func (c *苏萨TigrisCounty) BTellelhiba特尔赫巴() feud.Barony {
	return c.特尔赫巴Tellelhiba
}

var CTigris苏萨 TigrisCounty = &苏萨TigrisCounty{}

func init() {
	f := CTigris苏萨.(*苏萨TigrisCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "656",
		Title:     "tigris",
		TitleName: "苏萨",
		TitleCode: "c_tigris",
		Baronies:  map[string]feud.Barony{},
	}

	f.比斯玛亚Bismaya = BBismaya比斯玛亚
	f.比斯玛亚Bismaya.SetParent(f)

	f.伊斯罕Ishan = BIshan伊斯罕
	f.伊斯罕Ishan.SetParent(f)

	f.马贾阿卡比尔Majaralkabir = BMajaralkabir马贾阿卡比尔
	f.马贾阿卡比尔Majaralkabir.SetParent(f)

	f.努法尔Nuffar = BNuffar努法尔
	f.努法尔Nuffar.SetParent(f)

	f.喀拉特斯吉格特Qalatsjergat = BQalatsjergat喀拉特斯吉格特
	f.喀拉特斯吉格特Qalatsjergat.SetParent(f)

	f.萨马瓦Samawah = BSamawah萨马瓦
	f.萨马瓦Samawah.SetParent(f)

	f.特尔赫巴Tellelhiba = BTellelhiba特尔赫巴
	f.特尔赫巴Tellelhiba.SetParent(f)

}
