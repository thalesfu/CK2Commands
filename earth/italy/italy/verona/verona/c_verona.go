package verona

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VeronaCounty interface {
	feud.County
	BArzignano阿尔齐尼亚诺() feud.Barony
	BBarbarano巴尔巴拉诺() feud.Barony
	BLonigo洛尼戈() feud.Barony
	BMontecchio蒙特基奥() feud.Barony
	BSanmartino圣马蒂诺() feud.Barony
	BSchio斯基奥() feud.Barony
	BVerona维罗纳() feud.Barony
	BVicenza维琴察() feud.Barony
}

type 维罗纳VeronaCounty struct {
	feud.BaseCounty
	阿尔齐尼亚诺Arzignano feud.Barony
	巴尔巴拉诺Barbarano  feud.Barony
	洛尼戈Lonigo       feud.Barony
	蒙特基奥Montecchio  feud.Barony
	圣马蒂诺Sanmartino  feud.Barony
	斯基奥Schio        feud.Barony
	维罗纳Verona       feud.Barony
	维琴察Vicenza      feud.Barony
}

func (c *维罗纳VeronaCounty) BArzignano阿尔齐尼亚诺() feud.Barony {
	return c.阿尔齐尼亚诺Arzignano
}

func (c *维罗纳VeronaCounty) BBarbarano巴尔巴拉诺() feud.Barony {
	return c.巴尔巴拉诺Barbarano
}

func (c *维罗纳VeronaCounty) BLonigo洛尼戈() feud.Barony {
	return c.洛尼戈Lonigo
}

func (c *维罗纳VeronaCounty) BMontecchio蒙特基奥() feud.Barony {
	return c.蒙特基奥Montecchio
}

func (c *维罗纳VeronaCounty) BSanmartino圣马蒂诺() feud.Barony {
	return c.圣马蒂诺Sanmartino
}

func (c *维罗纳VeronaCounty) BSchio斯基奥() feud.Barony {
	return c.斯基奥Schio
}

func (c *维罗纳VeronaCounty) BVerona维罗纳() feud.Barony {
	return c.维罗纳Verona
}

func (c *维罗纳VeronaCounty) BVicenza维琴察() feud.Barony {
	return c.维琴察Vicenza
}

var CVerona维罗纳 VeronaCounty = &维罗纳VeronaCounty{}

func init() {
	f := CVerona维罗纳.(*维罗纳VeronaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "319",
		Title:     "verona",
		TitleName: "维罗纳",
		TitleCode: "c_verona",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔齐尼亚诺Arzignano = BArzignano阿尔齐尼亚诺
	f.阿尔齐尼亚诺Arzignano.SetParent(f)

	f.巴尔巴拉诺Barbarano = BBarbarano巴尔巴拉诺
	f.巴尔巴拉诺Barbarano.SetParent(f)

	f.洛尼戈Lonigo = BLonigo洛尼戈
	f.洛尼戈Lonigo.SetParent(f)

	f.蒙特基奥Montecchio = BMontecchio蒙特基奥
	f.蒙特基奥Montecchio.SetParent(f)

	f.圣马蒂诺Sanmartino = BSanmartino圣马蒂诺
	f.圣马蒂诺Sanmartino.SetParent(f)

	f.斯基奥Schio = BSchio斯基奥
	f.斯基奥Schio.SetParent(f)

	f.维罗纳Verona = BVerona维罗纳
	f.维罗纳Verona.SetParent(f)

	f.维琴察Vicenza = BVicenza维琴察
	f.维琴察Vicenza.SetParent(f)

}
