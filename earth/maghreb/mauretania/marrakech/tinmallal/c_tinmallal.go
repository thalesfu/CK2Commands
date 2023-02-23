package tinmallal

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TinmallalCounty interface {
	feud.County
	BAghmat阿格马特() feud.Barony
	BNaffis纳菲斯() feud.Barony
	BTamrghoute泰姆尔古泰() feud.Barony
	BTinmallal廷马拉尔() feud.Barony
}

type 廷马拉尔TinmallalCounty struct {
	feud.BaseCounty
	阿格马特Aghmat      feud.Barony
	纳菲斯Naffis       feud.Barony
	泰姆尔古泰Tamrghoute feud.Barony
	廷马拉尔Tinmallal   feud.Barony
}

func (c *廷马拉尔TinmallalCounty) BAghmat阿格马特() feud.Barony {
	return c.阿格马特Aghmat
}

func (c *廷马拉尔TinmallalCounty) BNaffis纳菲斯() feud.Barony {
	return c.纳菲斯Naffis
}

func (c *廷马拉尔TinmallalCounty) BTamrghoute泰姆尔古泰() feud.Barony {
	return c.泰姆尔古泰Tamrghoute
}

func (c *廷马拉尔TinmallalCounty) BTinmallal廷马拉尔() feud.Barony {
	return c.廷马拉尔Tinmallal
}

var CTinmallal廷马拉尔 TinmallalCounty = &廷马拉尔TinmallalCounty{}

func init() {
	f := CTinmallal廷马拉尔.(*廷马拉尔TinmallalCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1778",
		Title:     "tinmallal",
		TitleName: "廷马拉尔",
		TitleCode: "c_tinmallal",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿格马特Aghmat = BAghmat阿格马特
	f.阿格马特Aghmat.SetParent(f)

	f.纳菲斯Naffis = BNaffis纳菲斯
	f.纳菲斯Naffis.SetParent(f)

	f.泰姆尔古泰Tamrghoute = BTamrghoute泰姆尔古泰
	f.泰姆尔古泰Tamrghoute.SetParent(f)

	f.廷马拉尔Tinmallal = BTinmallal廷马拉尔
	f.廷马拉尔Tinmallal.SetParent(f)

}
