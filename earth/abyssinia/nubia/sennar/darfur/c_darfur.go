package darfur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DarfurCounty interface {
	feud.County
	BKalkal卡尔卡尔() feud.Barony
	BMasalat马萨拉特() feud.Barony
	BUri乌里() feud.Barony
}

type 达尔富尔DarfurCounty struct {
	feud.BaseCounty
	卡尔卡尔Kalkal  feud.Barony
	马萨拉特Masalat feud.Barony
	乌里Uri       feud.Barony
}

func (c *达尔富尔DarfurCounty) BKalkal卡尔卡尔() feud.Barony {
	return c.卡尔卡尔Kalkal
}

func (c *达尔富尔DarfurCounty) BMasalat马萨拉特() feud.Barony {
	return c.马萨拉特Masalat
}

func (c *达尔富尔DarfurCounty) BUri乌里() feud.Barony {
	return c.乌里Uri
}

var CDarfur达尔富尔 DarfurCounty = &达尔富尔DarfurCounty{}

func init() {
	f := CDarfur达尔富尔.(*达尔富尔DarfurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1772",
		Title:     "darfur",
		TitleName: "达尔富尔",
		TitleCode: "c_darfur",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡尔卡尔Kalkal = BKalkal卡尔卡尔
	f.卡尔卡尔Kalkal.SetParent(f)

	f.马萨拉特Masalat = BMasalat马萨拉特
	f.马萨拉特Masalat.SetParent(f)

	f.乌里Uri = BUri乌里
	f.乌里Uri.SetParent(f)

}
