package gorodez

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GorodezCounty interface {
	feud.County
	BBor博尔() feud.Barony
	BGorodez戈罗杰茨() feud.Barony
	BKitezh基捷日() feud.Barony
	BPuchishche普奇谢() feud.Barony
	BSpasskoye斯帕斯科耶() feud.Barony
	BTarasikha塔拉西哈() feud.Barony
}

type 戈罗杰茨GorodezCounty struct {
	feud.BaseCounty
	博尔Bor          feud.Barony
	戈罗杰茨Gorodez    feud.Barony
	基捷日Kitezh      feud.Barony
	普奇谢Puchishche  feud.Barony
	斯帕斯科耶Spasskoye feud.Barony
	塔拉西哈Tarasikha  feud.Barony
}

func (c *戈罗杰茨GorodezCounty) BBor博尔() feud.Barony {
	return c.博尔Bor
}

func (c *戈罗杰茨GorodezCounty) BGorodez戈罗杰茨() feud.Barony {
	return c.戈罗杰茨Gorodez
}

func (c *戈罗杰茨GorodezCounty) BKitezh基捷日() feud.Barony {
	return c.基捷日Kitezh
}

func (c *戈罗杰茨GorodezCounty) BPuchishche普奇谢() feud.Barony {
	return c.普奇谢Puchishche
}

func (c *戈罗杰茨GorodezCounty) BSpasskoye斯帕斯科耶() feud.Barony {
	return c.斯帕斯科耶Spasskoye
}

func (c *戈罗杰茨GorodezCounty) BTarasikha塔拉西哈() feud.Barony {
	return c.塔拉西哈Tarasikha
}

var CGorodez戈罗杰茨 GorodezCounty = &戈罗杰茨GorodezCounty{}

func init() {
	f := CGorodez戈罗杰茨.(*戈罗杰茨GorodezCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "585",
		Title:     "gorodez",
		TitleName: "戈罗杰茨",
		TitleCode: "c_gorodez",
		Baronies:  map[string]feud.Barony{},
	}

	f.博尔Bor = BBor博尔
	f.博尔Bor.SetParent(f)

	f.戈罗杰茨Gorodez = BGorodez戈罗杰茨
	f.戈罗杰茨Gorodez.SetParent(f)

	f.基捷日Kitezh = BKitezh基捷日
	f.基捷日Kitezh.SetParent(f)

	f.普奇谢Puchishche = BPuchishche普奇谢
	f.普奇谢Puchishche.SetParent(f)

	f.斯帕斯科耶Spasskoye = BSpasskoye斯帕斯科耶
	f.斯帕斯科耶Spasskoye.SetParent(f)

	f.塔拉西哈Tarasikha = BTarasikha塔拉西哈
	f.塔拉西哈Tarasikha.SetParent(f)

}
