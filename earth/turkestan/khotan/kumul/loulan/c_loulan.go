package loulan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LoulanCounty interface {
	feud.County
	BDonghuamen东华门() feud.Barony
	BHuxinba湖心坝() feud.Barony
	BKroran楼兰() feud.Barony
	BLoulan楼兰() feud.Barony
	BShanshan鄯善() feud.Barony
	BYingpan营盘() feud.Barony
}

type 楼兰LoulanCounty struct {
	feud.BaseCounty
	东华门Donghuamen feud.Barony
	湖心坝Huxinba    feud.Barony
	楼兰Kroran      feud.Barony
	楼兰Loulan      feud.Barony
	鄯善Shanshan    feud.Barony
	营盘Yingpan     feud.Barony
}

func (c *楼兰LoulanCounty) BDonghuamen东华门() feud.Barony {
	return c.东华门Donghuamen
}

func (c *楼兰LoulanCounty) BHuxinba湖心坝() feud.Barony {
	return c.湖心坝Huxinba
}

func (c *楼兰LoulanCounty) BKroran楼兰() feud.Barony {
	return c.楼兰Kroran
}

func (c *楼兰LoulanCounty) BLoulan楼兰() feud.Barony {
	return c.楼兰Loulan
}

func (c *楼兰LoulanCounty) BShanshan鄯善() feud.Barony {
	return c.鄯善Shanshan
}

func (c *楼兰LoulanCounty) BYingpan营盘() feud.Barony {
	return c.营盘Yingpan
}

var CLoulan楼兰 LoulanCounty = &楼兰LoulanCounty{}

func init() {
	f := CLoulan楼兰.(*楼兰LoulanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1447",
		Title:     "loulan",
		TitleName: "楼兰",
		TitleCode: "c_loulan",
		Baronies:  map[string]feud.Barony{},
	}

	f.东华门Donghuamen = BDonghuamen东华门
	f.东华门Donghuamen.SetParent(f)

	f.湖心坝Huxinba = BHuxinba湖心坝
	f.湖心坝Huxinba.SetParent(f)

	f.楼兰Kroran = BKroran楼兰
	f.楼兰Kroran.SetParent(f)

	f.楼兰Loulan = BLoulan楼兰
	f.楼兰Loulan.SetParent(f)

	f.鄯善Shanshan = BShanshan鄯善
	f.鄯善Shanshan.SetParent(f)

	f.营盘Yingpan = BYingpan营盘
	f.营盘Yingpan.SetParent(f)

}
