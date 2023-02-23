package dariya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DariyaCounty interface {
	feud.County
	BDariya德拉伊耶() feud.Barony
	BFayd费德() feud.Barony
	BZubala祖巴莱() feud.Barony
}

type 德拉伊耶DariyaCounty struct {
	feud.BaseCounty
	德拉伊耶Dariya feud.Barony
	费德Fayd     feud.Barony
	祖巴莱Zubala  feud.Barony
}

func (c *德拉伊耶DariyaCounty) BDariya德拉伊耶() feud.Barony {
	return c.德拉伊耶Dariya
}

func (c *德拉伊耶DariyaCounty) BFayd费德() feud.Barony {
	return c.费德Fayd
}

func (c *德拉伊耶DariyaCounty) BZubala祖巴莱() feud.Barony {
	return c.祖巴莱Zubala
}

var CDariya德拉伊耶 DariyaCounty = &德拉伊耶DariyaCounty{}

func init() {
	f := CDariya德拉伊耶.(*德拉伊耶DariyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1530",
		Title:     "dariya",
		TitleName: "德拉伊耶",
		TitleCode: "c_dariya",
		Baronies:  map[string]feud.Barony{},
	}

	f.德拉伊耶Dariya = BDariya德拉伊耶
	f.德拉伊耶Dariya.SetParent(f)

	f.费德Fayd = BFayd费德
	f.费德Fayd.SetParent(f)

	f.祖巴莱Zubala = BZubala祖巴莱
	f.祖巴莱Zubala.SetParent(f)

}
