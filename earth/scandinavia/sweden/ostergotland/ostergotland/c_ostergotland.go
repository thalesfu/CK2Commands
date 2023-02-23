package ostergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OstergotlandCounty interface {
	feud.County
	BLinkoping林雪平() feud.Barony
	BNorrkoping北雪平() feud.Barony
	BRingstaholm灵斯塔霍尔姆() feud.Barony
	BSkanninge申宁厄() feud.Barony
	BSoderkoping南雪平() feud.Barony
	BStegeborg斯特格堡() feud.Barony
	BUttersberg乌特什贝里() feud.Barony
	BVadstena瓦斯泰纳() feud.Barony
	BVreta弗雷塔() feud.Barony
}

type 东约特兰OstergotlandCounty struct {
	feud.BaseCounty
	林雪平Linkoping      feud.Barony
	北雪平Norrkoping     feud.Barony
	灵斯塔霍尔姆Ringstaholm feud.Barony
	申宁厄Skanninge      feud.Barony
	南雪平Soderkoping    feud.Barony
	斯特格堡Stegeborg     feud.Barony
	乌特什贝里Uttersberg   feud.Barony
	瓦斯泰纳Vadstena      feud.Barony
	弗雷塔Vreta          feud.Barony
}

func (c *东约特兰OstergotlandCounty) BLinkoping林雪平() feud.Barony {
	return c.林雪平Linkoping
}

func (c *东约特兰OstergotlandCounty) BNorrkoping北雪平() feud.Barony {
	return c.北雪平Norrkoping
}

func (c *东约特兰OstergotlandCounty) BRingstaholm灵斯塔霍尔姆() feud.Barony {
	return c.灵斯塔霍尔姆Ringstaholm
}

func (c *东约特兰OstergotlandCounty) BSkanninge申宁厄() feud.Barony {
	return c.申宁厄Skanninge
}

func (c *东约特兰OstergotlandCounty) BSoderkoping南雪平() feud.Barony {
	return c.南雪平Soderkoping
}

func (c *东约特兰OstergotlandCounty) BStegeborg斯特格堡() feud.Barony {
	return c.斯特格堡Stegeborg
}

func (c *东约特兰OstergotlandCounty) BUttersberg乌特什贝里() feud.Barony {
	return c.乌特什贝里Uttersberg
}

func (c *东约特兰OstergotlandCounty) BVadstena瓦斯泰纳() feud.Barony {
	return c.瓦斯泰纳Vadstena
}

func (c *东约特兰OstergotlandCounty) BVreta弗雷塔() feud.Barony {
	return c.弗雷塔Vreta
}

var COstergotland东约特兰 OstergotlandCounty = &东约特兰OstergotlandCounty{}

func init() {
	f := COstergotland东约特兰.(*东约特兰OstergotlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "293",
		Title:     "ostergotland",
		TitleName: "东约特兰",
		TitleCode: "c_ostergotland",
		Baronies:  map[string]feud.Barony{},
	}

	f.林雪平Linkoping = BLinkoping林雪平
	f.林雪平Linkoping.SetParent(f)

	f.北雪平Norrkoping = BNorrkoping北雪平
	f.北雪平Norrkoping.SetParent(f)

	f.灵斯塔霍尔姆Ringstaholm = BRingstaholm灵斯塔霍尔姆
	f.灵斯塔霍尔姆Ringstaholm.SetParent(f)

	f.申宁厄Skanninge = BSkanninge申宁厄
	f.申宁厄Skanninge.SetParent(f)

	f.南雪平Soderkoping = BSoderkoping南雪平
	f.南雪平Soderkoping.SetParent(f)

	f.斯特格堡Stegeborg = BStegeborg斯特格堡
	f.斯特格堡Stegeborg.SetParent(f)

	f.乌特什贝里Uttersberg = BUttersberg乌特什贝里
	f.乌特什贝里Uttersberg.SetParent(f)

	f.瓦斯泰纳Vadstena = BVadstena瓦斯泰纳
	f.瓦斯泰纳Vadstena.SetParent(f)

	f.弗雷塔Vreta = BVreta弗雷塔
	f.弗雷塔Vreta.SetParent(f)

}
