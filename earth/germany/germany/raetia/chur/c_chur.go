package chur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChurCounty interface {
	feud.County
	BChur库尔() feud.Barony
	BChurwalden库尔瓦尔登() feud.Barony
	BDavos达沃斯() feud.Barony
	BGlurns格卢恩斯() feud.Barony
	BIllanz伊兰茨() feud.Barony
	BMaienfeld迈恩费尔德() feud.Barony
	BThusis图西斯() feud.Barony
	BZuoz楚奥茨() feud.Barony
}

type 库尔ChurCounty struct {
	feud.BaseCounty
	库尔Chur          feud.Barony
	库尔瓦尔登Churwalden feud.Barony
	达沃斯Davos        feud.Barony
	格卢恩斯Glurns      feud.Barony
	伊兰茨Illanz       feud.Barony
	迈恩费尔德Maienfeld  feud.Barony
	图西斯Thusis       feud.Barony
	楚奥茨Zuoz         feud.Barony
}

func (c *库尔ChurCounty) BChur库尔() feud.Barony {
	return c.库尔Chur
}

func (c *库尔ChurCounty) BChurwalden库尔瓦尔登() feud.Barony {
	return c.库尔瓦尔登Churwalden
}

func (c *库尔ChurCounty) BDavos达沃斯() feud.Barony {
	return c.达沃斯Davos
}

func (c *库尔ChurCounty) BGlurns格卢恩斯() feud.Barony {
	return c.格卢恩斯Glurns
}

func (c *库尔ChurCounty) BIllanz伊兰茨() feud.Barony {
	return c.伊兰茨Illanz
}

func (c *库尔ChurCounty) BMaienfeld迈恩费尔德() feud.Barony {
	return c.迈恩费尔德Maienfeld
}

func (c *库尔ChurCounty) BThusis图西斯() feud.Barony {
	return c.图西斯Thusis
}

func (c *库尔ChurCounty) BZuoz楚奥茨() feud.Barony {
	return c.楚奥茨Zuoz
}

var CChur库尔 ChurCounty = &库尔ChurCounty{}

func init() {
	f := CChur库尔.(*库尔ChurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "247",
		Title:     "chur",
		TitleName: "库尔",
		TitleCode: "c_chur",
		Baronies:  map[string]feud.Barony{},
	}

	f.库尔Chur = BChur库尔
	f.库尔Chur.SetParent(f)

	f.库尔瓦尔登Churwalden = BChurwalden库尔瓦尔登
	f.库尔瓦尔登Churwalden.SetParent(f)

	f.达沃斯Davos = BDavos达沃斯
	f.达沃斯Davos.SetParent(f)

	f.格卢恩斯Glurns = BGlurns格卢恩斯
	f.格卢恩斯Glurns.SetParent(f)

	f.伊兰茨Illanz = BIllanz伊兰茨
	f.伊兰茨Illanz.SetParent(f)

	f.迈恩费尔德Maienfeld = BMaienfeld迈恩费尔德
	f.迈恩费尔德Maienfeld.SetParent(f)

	f.图西斯Thusis = BThusis图西斯
	f.图西斯Thusis.SetParent(f)

	f.楚奥茨Zuoz = BZuoz楚奥茨
	f.楚奥茨Zuoz.SetParent(f)

}
