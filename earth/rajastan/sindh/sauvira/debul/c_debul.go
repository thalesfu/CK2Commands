package debul

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DebulCounty interface {
	feud.County
	BBanbhore斑波尔() feud.Barony
	BDebul提部罗() feud.Barony
	BKolachi拘罗支() feud.Barony
	BLandhi兰蒂希() feud.Barony
	BMakli摩柯尼() feud.Barony
	BShahbandar舍般陀() feud.Barony
	BThatta塔达() feud.Barony
}

type 提部罗DebulCounty struct {
	feud.BaseCounty
	斑波尔Banbhore   feud.Barony
	提部罗Debul      feud.Barony
	拘罗支Kolachi    feud.Barony
	兰蒂希Landhi     feud.Barony
	摩柯尼Makli      feud.Barony
	舍般陀Shahbandar feud.Barony
	塔达Thatta      feud.Barony
}

func (c *提部罗DebulCounty) BBanbhore斑波尔() feud.Barony {
	return c.斑波尔Banbhore
}

func (c *提部罗DebulCounty) BDebul提部罗() feud.Barony {
	return c.提部罗Debul
}

func (c *提部罗DebulCounty) BKolachi拘罗支() feud.Barony {
	return c.拘罗支Kolachi
}

func (c *提部罗DebulCounty) BLandhi兰蒂希() feud.Barony {
	return c.兰蒂希Landhi
}

func (c *提部罗DebulCounty) BMakli摩柯尼() feud.Barony {
	return c.摩柯尼Makli
}

func (c *提部罗DebulCounty) BShahbandar舍般陀() feud.Barony {
	return c.舍般陀Shahbandar
}

func (c *提部罗DebulCounty) BThatta塔达() feud.Barony {
	return c.塔达Thatta
}

var CDebul提部罗 DebulCounty = &提部罗DebulCounty{}

func init() {
	f := CDebul提部罗.(*提部罗DebulCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1297",
		Title:     "debul",
		TitleName: "提部罗",
		TitleCode: "c_debul",
		Baronies:  map[string]feud.Barony{},
	}

	f.斑波尔Banbhore = BBanbhore斑波尔
	f.斑波尔Banbhore.SetParent(f)

	f.提部罗Debul = BDebul提部罗
	f.提部罗Debul.SetParent(f)

	f.拘罗支Kolachi = BKolachi拘罗支
	f.拘罗支Kolachi.SetParent(f)

	f.兰蒂希Landhi = BLandhi兰蒂希
	f.兰蒂希Landhi.SetParent(f)

	f.摩柯尼Makli = BMakli摩柯尼
	f.摩柯尼Makli.SetParent(f)

	f.舍般陀Shahbandar = BShahbandar舍般陀
	f.舍般陀Shahbandar.SetParent(f)

	f.塔达Thatta = BThatta塔达
	f.塔达Thatta.SetParent(f)

}
