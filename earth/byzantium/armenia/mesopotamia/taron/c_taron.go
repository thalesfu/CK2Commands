package taron

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TaronCounty interface {
	feud.County
	BAkhlat阿赫拉特() feud.Barony
	BArakelots阿拉克洛特() feud.Barony
	BArarati亚拉拉特() feud.Barony
	BArtchesh阿彻什() feud.Barony
	BGlak格拉卡() feud.Barony
	BManzikert曼齐刻尔特() feud.Barony
	BMush慕斯() feud.Barony
	BSason萨松() feud.Barony
}

type 塔龙TaronCounty struct {
	feud.BaseCounty
	阿赫拉特Akhlat     feud.Barony
	阿拉克洛特Arakelots feud.Barony
	亚拉拉特Ararati    feud.Barony
	阿彻什Artchesh    feud.Barony
	格拉卡Glak        feud.Barony
	曼齐刻尔特Manzikert feud.Barony
	慕斯Mush         feud.Barony
	萨松Sason        feud.Barony
}

func (c *塔龙TaronCounty) BAkhlat阿赫拉特() feud.Barony {
	return c.阿赫拉特Akhlat
}

func (c *塔龙TaronCounty) BArakelots阿拉克洛特() feud.Barony {
	return c.阿拉克洛特Arakelots
}

func (c *塔龙TaronCounty) BArarati亚拉拉特() feud.Barony {
	return c.亚拉拉特Ararati
}

func (c *塔龙TaronCounty) BArtchesh阿彻什() feud.Barony {
	return c.阿彻什Artchesh
}

func (c *塔龙TaronCounty) BGlak格拉卡() feud.Barony {
	return c.格拉卡Glak
}

func (c *塔龙TaronCounty) BManzikert曼齐刻尔特() feud.Barony {
	return c.曼齐刻尔特Manzikert
}

func (c *塔龙TaronCounty) BMush慕斯() feud.Barony {
	return c.慕斯Mush
}

func (c *塔龙TaronCounty) BSason萨松() feud.Barony {
	return c.萨松Sason
}

var CTaron塔龙 TaronCounty = &塔龙TaronCounty{}

func init() {
	f := CTaron塔龙.(*塔龙TaronCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "701",
		Title:     "taron",
		TitleName: "塔龙",
		TitleCode: "c_taron",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿赫拉特Akhlat = BAkhlat阿赫拉特
	f.阿赫拉特Akhlat.SetParent(f)

	f.阿拉克洛特Arakelots = BArakelots阿拉克洛特
	f.阿拉克洛特Arakelots.SetParent(f)

	f.亚拉拉特Ararati = BArarati亚拉拉特
	f.亚拉拉特Ararati.SetParent(f)

	f.阿彻什Artchesh = BArtchesh阿彻什
	f.阿彻什Artchesh.SetParent(f)

	f.格拉卡Glak = BGlak格拉卡
	f.格拉卡Glak.SetParent(f)

	f.曼齐刻尔特Manzikert = BManzikert曼齐刻尔特
	f.曼齐刻尔特Manzikert.SetParent(f)

	f.慕斯Mush = BMush慕斯
	f.慕斯Mush.SetParent(f)

	f.萨松Sason = BSason萨松
	f.萨松Sason.SetParent(f)

}
