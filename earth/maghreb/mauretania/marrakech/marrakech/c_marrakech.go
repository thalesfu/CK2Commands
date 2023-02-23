package marrakech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MarrakechCounty interface {
	feud.County
	BAsni艾斯尼() feud.Barony
	BDemnat代姆纳特() feud.Barony
	BElanakir阿巴基() feud.Barony
	BMarrakech马拉喀什() feud.Barony
	BOuarzazate瓦尔扎扎特() feud.Barony
	BSidirahhal西迪赖哈勒() feud.Barony
	BThineghir廷吉尔() feud.Barony
}

type 马拉喀什MarrakechCounty struct {
	feud.BaseCounty
	艾斯尼Asni         feud.Barony
	代姆纳特Demnat      feud.Barony
	阿巴基Elanakir     feud.Barony
	马拉喀什Marrakech   feud.Barony
	瓦尔扎扎特Ouarzazate feud.Barony
	西迪赖哈勒Sidirahhal feud.Barony
	廷吉尔Thineghir    feud.Barony
}

func (c *马拉喀什MarrakechCounty) BAsni艾斯尼() feud.Barony {
	return c.艾斯尼Asni
}

func (c *马拉喀什MarrakechCounty) BDemnat代姆纳特() feud.Barony {
	return c.代姆纳特Demnat
}

func (c *马拉喀什MarrakechCounty) BElanakir阿巴基() feud.Barony {
	return c.阿巴基Elanakir
}

func (c *马拉喀什MarrakechCounty) BMarrakech马拉喀什() feud.Barony {
	return c.马拉喀什Marrakech
}

func (c *马拉喀什MarrakechCounty) BOuarzazate瓦尔扎扎特() feud.Barony {
	return c.瓦尔扎扎特Ouarzazate
}

func (c *马拉喀什MarrakechCounty) BSidirahhal西迪赖哈勒() feud.Barony {
	return c.西迪赖哈勒Sidirahhal
}

func (c *马拉喀什MarrakechCounty) BThineghir廷吉尔() feud.Barony {
	return c.廷吉尔Thineghir
}

var CMarrakech马拉喀什 MarrakechCounty = &马拉喀什MarrakechCounty{}

func init() {
	f := CMarrakech马拉喀什.(*马拉喀什MarrakechCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "843",
		Title:     "marrakech",
		TitleName: "马拉喀什",
		TitleCode: "c_marrakech",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾斯尼Asni = BAsni艾斯尼
	f.艾斯尼Asni.SetParent(f)

	f.代姆纳特Demnat = BDemnat代姆纳特
	f.代姆纳特Demnat.SetParent(f)

	f.阿巴基Elanakir = BElanakir阿巴基
	f.阿巴基Elanakir.SetParent(f)

	f.马拉喀什Marrakech = BMarrakech马拉喀什
	f.马拉喀什Marrakech.SetParent(f)

	f.瓦尔扎扎特Ouarzazate = BOuarzazate瓦尔扎扎特
	f.瓦尔扎扎特Ouarzazate.SetParent(f)

	f.西迪赖哈勒Sidirahhal = BSidirahhal西迪赖哈勒
	f.西迪赖哈勒Sidirahhal.SetParent(f)

	f.廷吉尔Thineghir = BThineghir廷吉尔
	f.廷吉尔Thineghir.SetParent(f)

}
