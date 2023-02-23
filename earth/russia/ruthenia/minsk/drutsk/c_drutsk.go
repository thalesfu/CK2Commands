package drutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DrutskCounty interface {
	feud.County
	BBabruysk博布鲁伊斯克() feud.Barony
	BBialocerkev比亚瓦采尔凯夫() feud.Barony
	BDashkovka达什科夫卡() feud.Barony
	BDrutsk德鲁茨克() feud.Barony
	BMogilev莫吉廖夫() feud.Barony
	BRahachow罗加乔夫() feud.Barony
	BZhlobin日洛宾() feud.Barony
}

type 德鲁茨克DrutskCounty struct {
	feud.BaseCounty
	博布鲁伊斯克Babruysk     feud.Barony
	比亚瓦采尔凯夫Bialocerkev feud.Barony
	达什科夫卡Dashkovka     feud.Barony
	德鲁茨克Drutsk         feud.Barony
	莫吉廖夫Mogilev        feud.Barony
	罗加乔夫Rahachow       feud.Barony
	日洛宾Zhlobin         feud.Barony
}

func (c *德鲁茨克DrutskCounty) BBabruysk博布鲁伊斯克() feud.Barony {
	return c.博布鲁伊斯克Babruysk
}

func (c *德鲁茨克DrutskCounty) BBialocerkev比亚瓦采尔凯夫() feud.Barony {
	return c.比亚瓦采尔凯夫Bialocerkev
}

func (c *德鲁茨克DrutskCounty) BDashkovka达什科夫卡() feud.Barony {
	return c.达什科夫卡Dashkovka
}

func (c *德鲁茨克DrutskCounty) BDrutsk德鲁茨克() feud.Barony {
	return c.德鲁茨克Drutsk
}

func (c *德鲁茨克DrutskCounty) BMogilev莫吉廖夫() feud.Barony {
	return c.莫吉廖夫Mogilev
}

func (c *德鲁茨克DrutskCounty) BRahachow罗加乔夫() feud.Barony {
	return c.罗加乔夫Rahachow
}

func (c *德鲁茨克DrutskCounty) BZhlobin日洛宾() feud.Barony {
	return c.日洛宾Zhlobin
}

var CDrutsk德鲁茨克 DrutskCounty = &德鲁茨克DrutskCounty{}

func init() {
	f := CDrutsk德鲁茨克.(*德鲁茨克DrutskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1659",
		Title:     "drutsk",
		TitleName: "德鲁茨克",
		TitleCode: "c_drutsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.博布鲁伊斯克Babruysk = BBabruysk博布鲁伊斯克
	f.博布鲁伊斯克Babruysk.SetParent(f)

	f.比亚瓦采尔凯夫Bialocerkev = BBialocerkev比亚瓦采尔凯夫
	f.比亚瓦采尔凯夫Bialocerkev.SetParent(f)

	f.达什科夫卡Dashkovka = BDashkovka达什科夫卡
	f.达什科夫卡Dashkovka.SetParent(f)

	f.德鲁茨克Drutsk = BDrutsk德鲁茨克
	f.德鲁茨克Drutsk.SetParent(f)

	f.莫吉廖夫Mogilev = BMogilev莫吉廖夫
	f.莫吉廖夫Mogilev.SetParent(f)

	f.罗加乔夫Rahachow = BRahachow罗加乔夫
	f.罗加乔夫Rahachow.SetParent(f)

	f.日洛宾Zhlobin = BZhlobin日洛宾
	f.日洛宾Zhlobin.SetParent(f)

}
