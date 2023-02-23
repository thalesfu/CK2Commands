package kolomna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KolomnaCounty interface {
	feud.County
	BKochira科奇拉() feud.Barony
	BKonchinka孔钦卡() feud.Barony
	BLaptevo拉普捷沃() feud.Barony
	BOzherelye奥热列利耶() feud.Barony
	BRostislavl罗斯季斯拉夫尔() feud.Barony
	BTula图拉() feud.Barony
	BVeniov韦尼奥夫() feud.Barony
}

type 图拉KolomnaCounty struct {
	feud.BaseCounty
	科奇拉Kochira        feud.Barony
	孔钦卡Konchinka      feud.Barony
	拉普捷沃Laptevo       feud.Barony
	奥热列利耶Ozherelye    feud.Barony
	罗斯季斯拉夫尔Rostislavl feud.Barony
	图拉Tula            feud.Barony
	韦尼奥夫Veniov        feud.Barony
}

func (c *图拉KolomnaCounty) BKochira科奇拉() feud.Barony {
	return c.科奇拉Kochira
}

func (c *图拉KolomnaCounty) BKonchinka孔钦卡() feud.Barony {
	return c.孔钦卡Konchinka
}

func (c *图拉KolomnaCounty) BLaptevo拉普捷沃() feud.Barony {
	return c.拉普捷沃Laptevo
}

func (c *图拉KolomnaCounty) BOzherelye奥热列利耶() feud.Barony {
	return c.奥热列利耶Ozherelye
}

func (c *图拉KolomnaCounty) BRostislavl罗斯季斯拉夫尔() feud.Barony {
	return c.罗斯季斯拉夫尔Rostislavl
}

func (c *图拉KolomnaCounty) BTula图拉() feud.Barony {
	return c.图拉Tula
}

func (c *图拉KolomnaCounty) BVeniov韦尼奥夫() feud.Barony {
	return c.韦尼奥夫Veniov
}

var CKolomna图拉 KolomnaCounty = &图拉KolomnaCounty{}

func init() {
	f := CKolomna图拉.(*图拉KolomnaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "578",
		Title:     "kolomna",
		TitleName: "图拉",
		TitleCode: "c_kolomna",
		Baronies:  map[string]feud.Barony{},
	}

	f.科奇拉Kochira = BKochira科奇拉
	f.科奇拉Kochira.SetParent(f)

	f.孔钦卡Konchinka = BKonchinka孔钦卡
	f.孔钦卡Konchinka.SetParent(f)

	f.拉普捷沃Laptevo = BLaptevo拉普捷沃
	f.拉普捷沃Laptevo.SetParent(f)

	f.奥热列利耶Ozherelye = BOzherelye奥热列利耶
	f.奥热列利耶Ozherelye.SetParent(f)

	f.罗斯季斯拉夫尔Rostislavl = BRostislavl罗斯季斯拉夫尔
	f.罗斯季斯拉夫尔Rostislavl.SetParent(f)

	f.图拉Tula = BTula图拉
	f.图拉Tula.SetParent(f)

	f.韦尼奥夫Veniov = BVeniov韦尼奥夫
	f.韦尼奥夫Veniov.SetParent(f)

}
