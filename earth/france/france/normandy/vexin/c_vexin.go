package vexin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VexinCounty interface {
	feud.County
	BAbbayedemortemer莫特梅修道院() feud.Barony
	BAndely昂德利() feud.Barony
	BHarcourt阿尔库尔() feud.Barony
	BIvry伊夫里() feud.Barony
	BLarocheguyon拉罗什_居永() feud.Barony
	BMantes芒特() feud.Barony
	BPontoise蓬图瓦兹() feud.Barony
}

type 韦克桑VexinCounty struct {
	feud.BaseCounty
	莫特梅修道院Abbayedemortemer feud.Barony
	昂德利Andely              feud.Barony
	阿尔库尔Harcourt           feud.Barony
	伊夫里Ivry                feud.Barony
	拉罗什_居永Larocheguyon     feud.Barony
	芒特Mantes               feud.Barony
	蓬图瓦兹Pontoise           feud.Barony
}

func (c *韦克桑VexinCounty) BAbbayedemortemer莫特梅修道院() feud.Barony {
	return c.莫特梅修道院Abbayedemortemer
}

func (c *韦克桑VexinCounty) BAndely昂德利() feud.Barony {
	return c.昂德利Andely
}

func (c *韦克桑VexinCounty) BHarcourt阿尔库尔() feud.Barony {
	return c.阿尔库尔Harcourt
}

func (c *韦克桑VexinCounty) BIvry伊夫里() feud.Barony {
	return c.伊夫里Ivry
}

func (c *韦克桑VexinCounty) BLarocheguyon拉罗什_居永() feud.Barony {
	return c.拉罗什_居永Larocheguyon
}

func (c *韦克桑VexinCounty) BMantes芒特() feud.Barony {
	return c.芒特Mantes
}

func (c *韦克桑VexinCounty) BPontoise蓬图瓦兹() feud.Barony {
	return c.蓬图瓦兹Pontoise
}

var CVexin韦克桑 VexinCounty = &韦克桑VexinCounty{}

func init() {
	f := CVexin韦克桑.(*韦克桑VexinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "98",
		Title:     "vexin",
		TitleName: "韦克桑",
		TitleCode: "c_vexin",
		Baronies:  map[string]feud.Barony{},
	}

	f.莫特梅修道院Abbayedemortemer = BAbbayedemortemer莫特梅修道院
	f.莫特梅修道院Abbayedemortemer.SetParent(f)

	f.昂德利Andely = BAndely昂德利
	f.昂德利Andely.SetParent(f)

	f.阿尔库尔Harcourt = BHarcourt阿尔库尔
	f.阿尔库尔Harcourt.SetParent(f)

	f.伊夫里Ivry = BIvry伊夫里
	f.伊夫里Ivry.SetParent(f)

	f.拉罗什_居永Larocheguyon = BLarocheguyon拉罗什_居永
	f.拉罗什_居永Larocheguyon.SetParent(f)

	f.芒特Mantes = BMantes芒特
	f.芒特Mantes.SetParent(f)

	f.蓬图瓦兹Pontoise = BPontoise蓬图瓦兹
	f.蓬图瓦兹Pontoise.SetParent(f)

}
