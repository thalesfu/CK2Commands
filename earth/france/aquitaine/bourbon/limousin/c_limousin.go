package limousin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LimousinCounty interface {
	feud.County
	BChalus沙吕() feud.Barony
	BComborn孔博恩() feud.Barony
	BLimoges利摩日() feud.Barony
	BRochechouart罗什舒瓦尔() feud.Barony
	BStleonard圣莱昂纳尔() feud.Barony
	BThiviers蒂维耶() feud.Barony
	BTurenne蒂雷纳() feud.Barony
	BVentadour旺塔杜尔() feud.Barony
}

type 利穆赞LimousinCounty struct {
	feud.BaseCounty
	沙吕Chalus          feud.Barony
	孔博恩Comborn        feud.Barony
	利摩日Limoges        feud.Barony
	罗什舒瓦尔Rochechouart feud.Barony
	圣莱昂纳尔Stleonard    feud.Barony
	蒂维耶Thiviers       feud.Barony
	蒂雷纳Turenne        feud.Barony
	旺塔杜尔Ventadour     feud.Barony
}

func (c *利穆赞LimousinCounty) BChalus沙吕() feud.Barony {
	return c.沙吕Chalus
}

func (c *利穆赞LimousinCounty) BComborn孔博恩() feud.Barony {
	return c.孔博恩Comborn
}

func (c *利穆赞LimousinCounty) BLimoges利摩日() feud.Barony {
	return c.利摩日Limoges
}

func (c *利穆赞LimousinCounty) BRochechouart罗什舒瓦尔() feud.Barony {
	return c.罗什舒瓦尔Rochechouart
}

func (c *利穆赞LimousinCounty) BStleonard圣莱昂纳尔() feud.Barony {
	return c.圣莱昂纳尔Stleonard
}

func (c *利穆赞LimousinCounty) BThiviers蒂维耶() feud.Barony {
	return c.蒂维耶Thiviers
}

func (c *利穆赞LimousinCounty) BTurenne蒂雷纳() feud.Barony {
	return c.蒂雷纳Turenne
}

func (c *利穆赞LimousinCounty) BVentadour旺塔杜尔() feud.Barony {
	return c.旺塔杜尔Ventadour
}

var CLimousin利穆赞 LimousinCounty = &利穆赞LimousinCounty{}

func init() {
	f := CLimousin利穆赞.(*利穆赞LimousinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "147",
		Title:     "limousin",
		TitleName: "利穆赞",
		TitleCode: "c_limousin",
		Baronies:  map[string]feud.Barony{},
	}

	f.沙吕Chalus = BChalus沙吕
	f.沙吕Chalus.SetParent(f)

	f.孔博恩Comborn = BComborn孔博恩
	f.孔博恩Comborn.SetParent(f)

	f.利摩日Limoges = BLimoges利摩日
	f.利摩日Limoges.SetParent(f)

	f.罗什舒瓦尔Rochechouart = BRochechouart罗什舒瓦尔
	f.罗什舒瓦尔Rochechouart.SetParent(f)

	f.圣莱昂纳尔Stleonard = BStleonard圣莱昂纳尔
	f.圣莱昂纳尔Stleonard.SetParent(f)

	f.蒂维耶Thiviers = BThiviers蒂维耶
	f.蒂维耶Thiviers.SetParent(f)

	f.蒂雷纳Turenne = BTurenne蒂雷纳
	f.蒂雷纳Turenne.SetParent(f)

	f.旺塔杜尔Ventadour = BVentadour旺塔杜尔
	f.旺塔杜尔Ventadour.SetParent(f)

}
