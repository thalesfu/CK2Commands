package mecca

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MeccaCounty interface {
	feud.County
	BJmumum扎姆穆() feud.Barony
	BMecca麦加() feud.Barony
	BQarn卡鲁纳() feud.Barony
	BTaif塔伊夫() feud.Barony
	BTurubah图赖拜() feud.Barony
}

type 麦加MeccaCounty struct {
	feud.BaseCounty
	扎姆穆Jmumum  feud.Barony
	麦加Mecca    feud.Barony
	卡鲁纳Qarn    feud.Barony
	塔伊夫Taif    feud.Barony
	图赖拜Turubah feud.Barony
}

func (c *麦加MeccaCounty) BJmumum扎姆穆() feud.Barony {
	return c.扎姆穆Jmumum
}

func (c *麦加MeccaCounty) BMecca麦加() feud.Barony {
	return c.麦加Mecca
}

func (c *麦加MeccaCounty) BQarn卡鲁纳() feud.Barony {
	return c.卡鲁纳Qarn
}

func (c *麦加MeccaCounty) BTaif塔伊夫() feud.Barony {
	return c.塔伊夫Taif
}

func (c *麦加MeccaCounty) BTurubah图赖拜() feud.Barony {
	return c.图赖拜Turubah
}

var CMecca麦加 MeccaCounty = &麦加MeccaCounty{}

func init() {
	f := CMecca麦加.(*麦加MeccaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "719",
		Title:     "mecca",
		TitleName: "麦加",
		TitleCode: "c_mecca",
		Baronies:  map[string]feud.Barony{},
	}

	f.扎姆穆Jmumum = BJmumum扎姆穆
	f.扎姆穆Jmumum.SetParent(f)

	f.麦加Mecca = BMecca麦加
	f.麦加Mecca.SetParent(f)

	f.卡鲁纳Qarn = BQarn卡鲁纳
	f.卡鲁纳Qarn.SetParent(f)

	f.塔伊夫Taif = BTaif塔伊夫
	f.塔伊夫Taif.SetParent(f)

	f.图赖拜Turubah = BTurubah图赖拜
	f.图赖拜Turubah.SetParent(f)

}
