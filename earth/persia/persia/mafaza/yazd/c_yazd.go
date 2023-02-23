package yazd

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type YazdCounty interface {
	feud.County
	BArdakan阿尔达坎() feud.Barony
	BMeybod梅博德() feud.Barony
	BNir尼尔() feud.Barony
	BTaft塔夫特() feud.Barony
	BYazd亚兹德() feud.Barony
	BZarch扎尔奇() feud.Barony
}

type 亚兹德YazdCounty struct {
	feud.BaseCounty
	阿尔达坎Ardakan feud.Barony
	梅博德Meybod   feud.Barony
	尼尔Nir       feud.Barony
	塔夫特Taft     feud.Barony
	亚兹德Yazd     feud.Barony
	扎尔奇Zarch    feud.Barony
}

func (c *亚兹德YazdCounty) BArdakan阿尔达坎() feud.Barony {
	return c.阿尔达坎Ardakan
}

func (c *亚兹德YazdCounty) BMeybod梅博德() feud.Barony {
	return c.梅博德Meybod
}

func (c *亚兹德YazdCounty) BNir尼尔() feud.Barony {
	return c.尼尔Nir
}

func (c *亚兹德YazdCounty) BTaft塔夫特() feud.Barony {
	return c.塔夫特Taft
}

func (c *亚兹德YazdCounty) BYazd亚兹德() feud.Barony {
	return c.亚兹德Yazd
}

func (c *亚兹德YazdCounty) BZarch扎尔奇() feud.Barony {
	return c.扎尔奇Zarch
}

var CYazd亚兹德 YazdCounty = &亚兹德YazdCounty{}

func init() {
	f := CYazd亚兹德.(*亚兹德YazdCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "638",
		Title:     "yazd",
		TitleName: "亚兹德",
		TitleCode: "c_yazd",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔达坎Ardakan = BArdakan阿尔达坎
	f.阿尔达坎Ardakan.SetParent(f)

	f.梅博德Meybod = BMeybod梅博德
	f.梅博德Meybod.SetParent(f)

	f.尼尔Nir = BNir尼尔
	f.尼尔Nir.SetParent(f)

	f.塔夫特Taft = BTaft塔夫特
	f.塔夫特Taft.SetParent(f)

	f.亚兹德Yazd = BYazd亚兹德
	f.亚兹德Yazd.SetParent(f)

	f.扎尔奇Zarch = BZarch扎尔奇
	f.扎尔奇Zarch.SetParent(f)

}
