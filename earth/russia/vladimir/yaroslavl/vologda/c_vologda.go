package vologda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VologdaCounty interface {
	feud.County
	BBolshayamurga大穆尔加() feud.Barony
	BDvinitsa德维尼察() feud.Barony
	BKadnikovskaya卡德尼科夫斯卡亚() feud.Barony
	BMotyn马蒂恩() feud.Barony
	BSokol索科尔() feud.Barony
	BStaroye斯塔罗耶() feud.Barony
	BUste乌斯季耶() feud.Barony
	BVologda沃洛格达() feud.Barony
}

type 沃洛格达VologdaCounty struct {
	feud.BaseCounty
	大穆尔加Bolshayamurga     feud.Barony
	德维尼察Dvinitsa          feud.Barony
	卡德尼科夫斯卡亚Kadnikovskaya feud.Barony
	马蒂恩Motyn              feud.Barony
	索科尔Sokol              feud.Barony
	斯塔罗耶Staroye           feud.Barony
	乌斯季耶Uste              feud.Barony
	沃洛格达Vologda           feud.Barony
}

func (c *沃洛格达VologdaCounty) BBolshayamurga大穆尔加() feud.Barony {
	return c.大穆尔加Bolshayamurga
}

func (c *沃洛格达VologdaCounty) BDvinitsa德维尼察() feud.Barony {
	return c.德维尼察Dvinitsa
}

func (c *沃洛格达VologdaCounty) BKadnikovskaya卡德尼科夫斯卡亚() feud.Barony {
	return c.卡德尼科夫斯卡亚Kadnikovskaya
}

func (c *沃洛格达VologdaCounty) BMotyn马蒂恩() feud.Barony {
	return c.马蒂恩Motyn
}

func (c *沃洛格达VologdaCounty) BSokol索科尔() feud.Barony {
	return c.索科尔Sokol
}

func (c *沃洛格达VologdaCounty) BStaroye斯塔罗耶() feud.Barony {
	return c.斯塔罗耶Staroye
}

func (c *沃洛格达VologdaCounty) BUste乌斯季耶() feud.Barony {
	return c.乌斯季耶Uste
}

func (c *沃洛格达VologdaCounty) BVologda沃洛格达() feud.Barony {
	return c.沃洛格达Vologda
}

var CVologda沃洛格达 VologdaCounty = &沃洛格达VologdaCounty{}

func init() {
	f := CVologda沃洛格达.(*沃洛格达VologdaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "406",
		Title:     "vologda",
		TitleName: "沃洛格达",
		TitleCode: "c_vologda",
		Baronies:  map[string]feud.Barony{},
	}

	f.大穆尔加Bolshayamurga = BBolshayamurga大穆尔加
	f.大穆尔加Bolshayamurga.SetParent(f)

	f.德维尼察Dvinitsa = BDvinitsa德维尼察
	f.德维尼察Dvinitsa.SetParent(f)

	f.卡德尼科夫斯卡亚Kadnikovskaya = BKadnikovskaya卡德尼科夫斯卡亚
	f.卡德尼科夫斯卡亚Kadnikovskaya.SetParent(f)

	f.马蒂恩Motyn = BMotyn马蒂恩
	f.马蒂恩Motyn.SetParent(f)

	f.索科尔Sokol = BSokol索科尔
	f.索科尔Sokol.SetParent(f)

	f.斯塔罗耶Staroye = BStaroye斯塔罗耶
	f.斯塔罗耶Staroye.SetParent(f)

	f.乌斯季耶Uste = BUste乌斯季耶
	f.乌斯季耶Uste.SetParent(f)

	f.沃洛格达Vologda = BVologda沃洛格达
	f.沃洛格达Vologda.SetParent(f)

}
