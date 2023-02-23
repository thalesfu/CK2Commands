package bayuda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BayudaCounty interface {
	feud.County
	BAlmeragh梅拉格() feud.Barony
	BAtongour阿通古尔() feud.Barony
	BBayuda拜尤达() feud.Barony
	BElkhandaq汉代格() feud.Barony
	BJeroko杰罗科() feud.Barony
}

type 拜尤达BayudaCounty struct {
	feud.BaseCounty
	梅拉格Almeragh  feud.Barony
	阿通古尔Atongour feud.Barony
	拜尤达Bayuda    feud.Barony
	汉代格Elkhandaq feud.Barony
	杰罗科Jeroko    feud.Barony
}

func (c *拜尤达BayudaCounty) BAlmeragh梅拉格() feud.Barony {
	return c.梅拉格Almeragh
}

func (c *拜尤达BayudaCounty) BAtongour阿通古尔() feud.Barony {
	return c.阿通古尔Atongour
}

func (c *拜尤达BayudaCounty) BBayuda拜尤达() feud.Barony {
	return c.拜尤达Bayuda
}

func (c *拜尤达BayudaCounty) BElkhandaq汉代格() feud.Barony {
	return c.汉代格Elkhandaq
}

func (c *拜尤达BayudaCounty) BJeroko杰罗科() feud.Barony {
	return c.杰罗科Jeroko
}

var CBayuda拜尤达 BayudaCounty = &拜尤达BayudaCounty{}

func init() {
	f := CBayuda拜尤达.(*拜尤达BayudaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1994",
		Title:     "bayuda",
		TitleName: "拜尤达",
		TitleCode: "c_bayuda",
		Baronies:  map[string]feud.Barony{},
	}

	f.梅拉格Almeragh = BAlmeragh梅拉格
	f.梅拉格Almeragh.SetParent(f)

	f.阿通古尔Atongour = BAtongour阿通古尔
	f.阿通古尔Atongour.SetParent(f)

	f.拜尤达Bayuda = BBayuda拜尤达
	f.拜尤达Bayuda.SetParent(f)

	f.汉代格Elkhandaq = BElkhandaq汉代格
	f.汉代格Elkhandaq.SetParent(f)

	f.杰罗科Jeroko = BJeroko杰罗科
	f.杰罗科Jeroko.SetParent(f)

}
