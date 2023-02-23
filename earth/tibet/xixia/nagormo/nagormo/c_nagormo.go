package nagormo

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NagormoCounty interface {
	feud.County
	BBaoku宝库() feud.Barony
	BDabuxun达布逊() feud.Barony
	BDagele大格勒() feud.Barony
	BGolmud郭勒木德() feud.Barony
	BQarham察尔汗() feud.Barony
	BQingshui清水() feud.Barony
	BXinle新乐() feud.Barony
}

type 格尔木NagormoCounty struct {
	feud.BaseCounty
	宝库Baoku    feud.Barony
	达布逊Dabuxun feud.Barony
	大格勒Dagele  feud.Barony
	郭勒木德Golmud feud.Barony
	察尔汗Qarham  feud.Barony
	清水Qingshui feud.Barony
	新乐Xinle    feud.Barony
}

func (c *格尔木NagormoCounty) BBaoku宝库() feud.Barony {
	return c.宝库Baoku
}

func (c *格尔木NagormoCounty) BDabuxun达布逊() feud.Barony {
	return c.达布逊Dabuxun
}

func (c *格尔木NagormoCounty) BDagele大格勒() feud.Barony {
	return c.大格勒Dagele
}

func (c *格尔木NagormoCounty) BGolmud郭勒木德() feud.Barony {
	return c.郭勒木德Golmud
}

func (c *格尔木NagormoCounty) BQarham察尔汗() feud.Barony {
	return c.察尔汗Qarham
}

func (c *格尔木NagormoCounty) BQingshui清水() feud.Barony {
	return c.清水Qingshui
}

func (c *格尔木NagormoCounty) BXinle新乐() feud.Barony {
	return c.新乐Xinle
}

var CNagormo格尔木 NagormoCounty = &格尔木NagormoCounty{}

func init() {
	f := CNagormo格尔木.(*格尔木NagormoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1501",
		Title:     "nagormo",
		TitleName: "格尔木",
		TitleCode: "c_nagormo",
		Baronies:  map[string]feud.Barony{},
	}

	f.宝库Baoku = BBaoku宝库
	f.宝库Baoku.SetParent(f)

	f.达布逊Dabuxun = BDabuxun达布逊
	f.达布逊Dabuxun.SetParent(f)

	f.大格勒Dagele = BDagele大格勒
	f.大格勒Dagele.SetParent(f)

	f.郭勒木德Golmud = BGolmud郭勒木德
	f.郭勒木德Golmud.SetParent(f)

	f.察尔汗Qarham = BQarham察尔汗
	f.察尔汗Qarham.SetParent(f)

	f.清水Qingshui = BQingshui清水
	f.清水Qingshui.SetParent(f)

	f.新乐Xinle = BXinle新乐
	f.新乐Xinle.SetParent(f)

}
