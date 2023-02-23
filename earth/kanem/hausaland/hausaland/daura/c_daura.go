package daura

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DauraCounty interface {
	feud.County
	BDaura道拉() feud.Barony
	BIrahun伊拉洪() feud.Barony
	BLapadi拉帕迪() feud.Barony
	BMaradi马拉迪() feud.Barony
	BUper乌珀() feud.Barony
	BZinder津德尔() feud.Barony
}

type 道拉DauraCounty struct {
	feud.BaseCounty
	道拉Daura   feud.Barony
	伊拉洪Irahun feud.Barony
	拉帕迪Lapadi feud.Barony
	马拉迪Maradi feud.Barony
	乌珀Uper    feud.Barony
	津德尔Zinder feud.Barony
}

func (c *道拉DauraCounty) BDaura道拉() feud.Barony {
	return c.道拉Daura
}

func (c *道拉DauraCounty) BIrahun伊拉洪() feud.Barony {
	return c.伊拉洪Irahun
}

func (c *道拉DauraCounty) BLapadi拉帕迪() feud.Barony {
	return c.拉帕迪Lapadi
}

func (c *道拉DauraCounty) BMaradi马拉迪() feud.Barony {
	return c.马拉迪Maradi
}

func (c *道拉DauraCounty) BUper乌珀() feud.Barony {
	return c.乌珀Uper
}

func (c *道拉DauraCounty) BZinder津德尔() feud.Barony {
	return c.津德尔Zinder
}

var CDaura道拉 DauraCounty = &道拉DauraCounty{}

func init() {
	f := CDaura道拉.(*道拉DauraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1747",
		Title:     "daura",
		TitleName: "道拉",
		TitleCode: "c_daura",
		Baronies:  map[string]feud.Barony{},
	}

	f.道拉Daura = BDaura道拉
	f.道拉Daura.SetParent(f)

	f.伊拉洪Irahun = BIrahun伊拉洪
	f.伊拉洪Irahun.SetParent(f)

	f.拉帕迪Lapadi = BLapadi拉帕迪
	f.拉帕迪Lapadi.SetParent(f)

	f.马拉迪Maradi = BMaradi马拉迪
	f.马拉迪Maradi.SetParent(f)

	f.乌珀Uper = BUper乌珀
	f.乌珀Uper.SetParent(f)

	f.津德尔Zinder = BZinder津德尔
	f.津德尔Zinder.SetParent(f)

}
