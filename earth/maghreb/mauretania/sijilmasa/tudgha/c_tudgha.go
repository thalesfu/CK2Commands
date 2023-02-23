package tudgha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TudghaCounty interface {
	feud.County
	BAlnif艾勒尼夫() feud.Barony
	BOuaourhrout瓦乌尔鲁() feud.Barony
	BTarzout塔尔祖特() feud.Barony
	BTudgha图德盖() feud.Barony
	BZbair兹拜尔() feud.Barony
}

type 图德盖TudghaCounty struct {
	feud.BaseCounty
	艾勒尼夫Alnif       feud.Barony
	瓦乌尔鲁Ouaourhrout feud.Barony
	塔尔祖特Tarzout     feud.Barony
	图德盖Tudgha       feud.Barony
	兹拜尔Zbair        feud.Barony
}

func (c *图德盖TudghaCounty) BAlnif艾勒尼夫() feud.Barony {
	return c.艾勒尼夫Alnif
}

func (c *图德盖TudghaCounty) BOuaourhrout瓦乌尔鲁() feud.Barony {
	return c.瓦乌尔鲁Ouaourhrout
}

func (c *图德盖TudghaCounty) BTarzout塔尔祖特() feud.Barony {
	return c.塔尔祖特Tarzout
}

func (c *图德盖TudghaCounty) BTudgha图德盖() feud.Barony {
	return c.图德盖Tudgha
}

func (c *图德盖TudghaCounty) BZbair兹拜尔() feud.Barony {
	return c.兹拜尔Zbair
}

var CTudgha图德盖 TudghaCounty = &图德盖TudghaCounty{}

func init() {
	f := CTudgha图德盖.(*图德盖TudghaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1775",
		Title:     "tudgha",
		TitleName: "图德盖",
		TitleCode: "c_tudgha",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾勒尼夫Alnif = BAlnif艾勒尼夫
	f.艾勒尼夫Alnif.SetParent(f)

	f.瓦乌尔鲁Ouaourhrout = BOuaourhrout瓦乌尔鲁
	f.瓦乌尔鲁Ouaourhrout.SetParent(f)

	f.塔尔祖特Tarzout = BTarzout塔尔祖特
	f.塔尔祖特Tarzout.SetParent(f)

	f.图德盖Tudgha = BTudgha图德盖
	f.图德盖Tudgha.SetParent(f)

	f.兹拜尔Zbair = BZbair兹拜尔
	f.兹拜尔Zbair.SetParent(f)

}
