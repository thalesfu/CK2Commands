package ishim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type IshimCounty interface {
	feud.County
	BEsil叶西尔() feud.Barony
	BEsut埃苏特() feud.Barony
	BIshim伊希姆() feud.Barony
	BMamlyut马姆柳特() feud.Barony
	BMamshim马姆希姆() feud.Barony
	BUlsut乌尔苏特() feud.Barony
}

type 伊希姆IshimCounty struct {
	feud.BaseCounty
	叶西尔Esil     feud.Barony
	埃苏特Esut     feud.Barony
	伊希姆Ishim    feud.Barony
	马姆柳特Mamlyut feud.Barony
	马姆希姆Mamshim feud.Barony
	乌尔苏特Ulsut   feud.Barony
}

func (c *伊希姆IshimCounty) BEsil叶西尔() feud.Barony {
	return c.叶西尔Esil
}

func (c *伊希姆IshimCounty) BEsut埃苏特() feud.Barony {
	return c.埃苏特Esut
}

func (c *伊希姆IshimCounty) BIshim伊希姆() feud.Barony {
	return c.伊希姆Ishim
}

func (c *伊希姆IshimCounty) BMamlyut马姆柳特() feud.Barony {
	return c.马姆柳特Mamlyut
}

func (c *伊希姆IshimCounty) BMamshim马姆希姆() feud.Barony {
	return c.马姆希姆Mamshim
}

func (c *伊希姆IshimCounty) BUlsut乌尔苏特() feud.Barony {
	return c.乌尔苏特Ulsut
}

var CIshim伊希姆 IshimCounty = &伊希姆IshimCounty{}

func init() {
	f := CIshim伊希姆.(*伊希姆IshimCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1223",
		Title:     "ishim",
		TitleName: "伊希姆",
		TitleCode: "c_ishim",
		Baronies:  map[string]feud.Barony{},
	}

	f.叶西尔Esil = BEsil叶西尔
	f.叶西尔Esil.SetParent(f)

	f.埃苏特Esut = BEsut埃苏特
	f.埃苏特Esut.SetParent(f)

	f.伊希姆Ishim = BIshim伊希姆
	f.伊希姆Ishim.SetParent(f)

	f.马姆柳特Mamlyut = BMamlyut马姆柳特
	f.马姆柳特Mamlyut.SetParent(f)

	f.马姆希姆Mamshim = BMamshim马姆希姆
	f.马姆希姆Mamshim.SetParent(f)

	f.乌尔苏特Ulsut = BUlsut乌尔苏特
	f.乌尔苏特Ulsut.SetParent(f)

}
