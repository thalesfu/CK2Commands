package mandavyapura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MandavyapuraCounty interface {
	feud.County
	BGhaneraov嘎内拉罗() feud.Barony
	BKorta拘罗吒() feud.Barony
	BMandavyapura曼荼毗耶补罗() feud.Barony
	BMathania摩他尼耶() feud.Barony
	BNaddula那菟罗() feud.Barony
	BOsian乌珊() feud.Barony
	BSanderaka桑啼罗迦() feud.Barony
	BSiwana斯瓦那() feud.Barony
}

type 曼荼毗耶补罗MandavyapuraCounty struct {
	feud.BaseCounty
	嘎内拉罗Ghaneraov      feud.Barony
	拘罗吒Korta           feud.Barony
	曼荼毗耶补罗Mandavyapura feud.Barony
	摩他尼耶Mathania       feud.Barony
	那菟罗Naddula         feud.Barony
	乌珊Osian            feud.Barony
	桑啼罗迦Sanderaka      feud.Barony
	斯瓦那Siwana          feud.Barony
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BGhaneraov嘎内拉罗() feud.Barony {
	return c.嘎内拉罗Ghaneraov
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BKorta拘罗吒() feud.Barony {
	return c.拘罗吒Korta
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BMandavyapura曼荼毗耶补罗() feud.Barony {
	return c.曼荼毗耶补罗Mandavyapura
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BMathania摩他尼耶() feud.Barony {
	return c.摩他尼耶Mathania
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BNaddula那菟罗() feud.Barony {
	return c.那菟罗Naddula
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BOsian乌珊() feud.Barony {
	return c.乌珊Osian
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BSanderaka桑啼罗迦() feud.Barony {
	return c.桑啼罗迦Sanderaka
}

func (c *曼荼毗耶补罗MandavyapuraCounty) BSiwana斯瓦那() feud.Barony {
	return c.斯瓦那Siwana
}

var CMandavyapura曼荼毗耶补罗 MandavyapuraCounty = &曼荼毗耶补罗MandavyapuraCounty{}

func init() {
	f := CMandavyapura曼荼毗耶补罗.(*曼荼毗耶补罗MandavyapuraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1130",
		Title:     "mandavyapura",
		TitleName: "曼荼毗耶补罗",
		TitleCode: "c_mandavyapura",
		Baronies:  map[string]feud.Barony{},
	}

	f.嘎内拉罗Ghaneraov = BGhaneraov嘎内拉罗
	f.嘎内拉罗Ghaneraov.SetParent(f)

	f.拘罗吒Korta = BKorta拘罗吒
	f.拘罗吒Korta.SetParent(f)

	f.曼荼毗耶补罗Mandavyapura = BMandavyapura曼荼毗耶补罗
	f.曼荼毗耶补罗Mandavyapura.SetParent(f)

	f.摩他尼耶Mathania = BMathania摩他尼耶
	f.摩他尼耶Mathania.SetParent(f)

	f.那菟罗Naddula = BNaddula那菟罗
	f.那菟罗Naddula.SetParent(f)

	f.乌珊Osian = BOsian乌珊
	f.乌珊Osian.SetParent(f)

	f.桑啼罗迦Sanderaka = BSanderaka桑啼罗迦
	f.桑啼罗迦Sanderaka.SetParent(f)

	f.斯瓦那Siwana = BSiwana斯瓦那
	f.斯瓦那Siwana.SetParent(f)

}
