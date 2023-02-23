package leix

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeixCounty interface {
	feud.County
	BDunamase敦纳梅斯() feud.Barony
	BLeix莱伊什() feud.Barony
	BLoigis洛伊吉什() feud.Barony
	BMountrath芒特拉斯() feud.Barony
	BTimahoe蒂马胡() feud.Barony
}

type 莱伊什LeixCounty struct {
	feud.BaseCounty
	敦纳梅斯Dunamase  feud.Barony
	莱伊什Leix       feud.Barony
	洛伊吉什Loigis    feud.Barony
	芒特拉斯Mountrath feud.Barony
	蒂马胡Timahoe    feud.Barony
}

func (c *莱伊什LeixCounty) BDunamase敦纳梅斯() feud.Barony {
	return c.敦纳梅斯Dunamase
}

func (c *莱伊什LeixCounty) BLeix莱伊什() feud.Barony {
	return c.莱伊什Leix
}

func (c *莱伊什LeixCounty) BLoigis洛伊吉什() feud.Barony {
	return c.洛伊吉什Loigis
}

func (c *莱伊什LeixCounty) BMountrath芒特拉斯() feud.Barony {
	return c.芒特拉斯Mountrath
}

func (c *莱伊什LeixCounty) BTimahoe蒂马胡() feud.Barony {
	return c.蒂马胡Timahoe
}

var CLeix莱伊什 LeixCounty = &莱伊什LeixCounty{}

func init() {
	f := CLeix莱伊什.(*莱伊什LeixCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1953",
		Title:     "leix",
		TitleName: "莱伊什",
		TitleCode: "c_leix",
		Baronies:  map[string]feud.Barony{},
	}

	f.敦纳梅斯Dunamase = BDunamase敦纳梅斯
	f.敦纳梅斯Dunamase.SetParent(f)

	f.莱伊什Leix = BLeix莱伊什
	f.莱伊什Leix.SetParent(f)

	f.洛伊吉什Loigis = BLoigis洛伊吉什
	f.洛伊吉什Loigis.SetParent(f)

	f.芒特拉斯Mountrath = BMountrath芒特拉斯
	f.芒特拉斯Mountrath.SetParent(f)

	f.蒂马胡Timahoe = BTimahoe蒂马胡
	f.蒂马胡Timahoe.SetParent(f)

}
