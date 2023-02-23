package quena

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type QuenaCounty interface {
	feud.County
	BBilbeis比勒拜斯() feud.Barony
	BKhalaf哈拉夫() feud.Barony
	BLuxor卢克索() feud.Barony
	BQift吉夫特() feud.Barony
	BQuena基纳() feud.Barony
	BQus古斯() feud.Barony
	BShiblanjah希卜兰杰() feud.Barony
}

type 基纳QuenaCounty struct {
	feud.BaseCounty
	比勒拜斯Bilbeis    feud.Barony
	哈拉夫Khalaf      feud.Barony
	卢克索Luxor       feud.Barony
	吉夫特Qift        feud.Barony
	基纳Quena        feud.Barony
	古斯Qus          feud.Barony
	希卜兰杰Shiblanjah feud.Barony
}

func (c *基纳QuenaCounty) BBilbeis比勒拜斯() feud.Barony {
	return c.比勒拜斯Bilbeis
}

func (c *基纳QuenaCounty) BKhalaf哈拉夫() feud.Barony {
	return c.哈拉夫Khalaf
}

func (c *基纳QuenaCounty) BLuxor卢克索() feud.Barony {
	return c.卢克索Luxor
}

func (c *基纳QuenaCounty) BQift吉夫特() feud.Barony {
	return c.吉夫特Qift
}

func (c *基纳QuenaCounty) BQuena基纳() feud.Barony {
	return c.基纳Quena
}

func (c *基纳QuenaCounty) BQus古斯() feud.Barony {
	return c.古斯Qus
}

func (c *基纳QuenaCounty) BShiblanjah希卜兰杰() feud.Barony {
	return c.希卜兰杰Shiblanjah
}

var CQuena基纳 QuenaCounty = &基纳QuenaCounty{}

func init() {
	f := CQuena基纳.(*基纳QuenaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "791",
		Title:     "quena",
		TitleName: "基纳",
		TitleCode: "c_quena",
		Baronies:  map[string]feud.Barony{},
	}

	f.比勒拜斯Bilbeis = BBilbeis比勒拜斯
	f.比勒拜斯Bilbeis.SetParent(f)

	f.哈拉夫Khalaf = BKhalaf哈拉夫
	f.哈拉夫Khalaf.SetParent(f)

	f.卢克索Luxor = BLuxor卢克索
	f.卢克索Luxor.SetParent(f)

	f.吉夫特Qift = BQift吉夫特
	f.吉夫特Qift.SetParent(f)

	f.基纳Quena = BQuena基纳
	f.基纳Quena.SetParent(f)

	f.古斯Qus = BQus古斯
	f.古斯Qus.SetParent(f)

	f.希卜兰杰Shiblanjah = BShiblanjah希卜兰杰
	f.希卜兰杰Shiblanjah.SetParent(f)

}
