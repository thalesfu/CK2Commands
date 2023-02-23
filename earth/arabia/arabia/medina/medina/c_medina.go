package medina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedinaCounty interface {
	feud.County
	BFarsh法拉斯() feud.Barony
	BKuba库巴() feud.Barony
	BMedina麦地那() feud.Barony
	BMilha米哈() feud.Barony
}

type 麦地那MedinaCounty struct {
	feud.BaseCounty
	法拉斯Farsh  feud.Barony
	库巴Kuba    feud.Barony
	麦地那Medina feud.Barony
	米哈Milha   feud.Barony
}

func (c *麦地那MedinaCounty) BFarsh法拉斯() feud.Barony {
	return c.法拉斯Farsh
}

func (c *麦地那MedinaCounty) BKuba库巴() feud.Barony {
	return c.库巴Kuba
}

func (c *麦地那MedinaCounty) BMedina麦地那() feud.Barony {
	return c.麦地那Medina
}

func (c *麦地那MedinaCounty) BMilha米哈() feud.Barony {
	return c.米哈Milha
}

var CMedina麦地那 MedinaCounty = &麦地那MedinaCounty{}

func init() {
	f := CMedina麦地那.(*麦地那MedinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "718",
		Title:     "medina",
		TitleName: "麦地那",
		TitleCode: "c_medina",
		Baronies:  map[string]feud.Barony{},
	}

	f.法拉斯Farsh = BFarsh法拉斯
	f.法拉斯Farsh.SetParent(f)

	f.库巴Kuba = BKuba库巴
	f.库巴Kuba.SetParent(f)

	f.麦地那Medina = BMedina麦地那
	f.麦地那Medina.SetParent(f)

	f.米哈Milha = BMilha米哈
	f.米哈Milha.SetParent(f)

}
