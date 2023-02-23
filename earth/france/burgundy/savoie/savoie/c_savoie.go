package savoie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SavoieCounty interface {
	feud.County
	BChambery尚贝里() feud.Barony
	BCluses克吕斯() feud.Barony
	BConflens孔福朗() feud.Barony
	BFaucigny福西尼() feud.Barony
	BTarentaise塔朗泰斯() feud.Barony
	BThonon托农() feud.Barony
}

type 萨伏依SavoieCounty struct {
	feud.BaseCounty
	尚贝里Chambery    feud.Barony
	克吕斯Cluses      feud.Barony
	孔福朗Conflens    feud.Barony
	福西尼Faucigny    feud.Barony
	塔朗泰斯Tarentaise feud.Barony
	托农Thonon       feud.Barony
}

func (c *萨伏依SavoieCounty) BChambery尚贝里() feud.Barony {
	return c.尚贝里Chambery
}

func (c *萨伏依SavoieCounty) BCluses克吕斯() feud.Barony {
	return c.克吕斯Cluses
}

func (c *萨伏依SavoieCounty) BConflens孔福朗() feud.Barony {
	return c.孔福朗Conflens
}

func (c *萨伏依SavoieCounty) BFaucigny福西尼() feud.Barony {
	return c.福西尼Faucigny
}

func (c *萨伏依SavoieCounty) BTarentaise塔朗泰斯() feud.Barony {
	return c.塔朗泰斯Tarentaise
}

func (c *萨伏依SavoieCounty) BThonon托农() feud.Barony {
	return c.托农Thonon
}

var CSavoie萨伏依 SavoieCounty = &萨伏依SavoieCounty{}

func init() {
	f := CSavoie萨伏依.(*萨伏依SavoieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "237",
		Title:     "savoie",
		TitleName: "萨伏依",
		TitleCode: "c_savoie",
		Baronies:  map[string]feud.Barony{},
	}

	f.尚贝里Chambery = BChambery尚贝里
	f.尚贝里Chambery.SetParent(f)

	f.克吕斯Cluses = BCluses克吕斯
	f.克吕斯Cluses.SetParent(f)

	f.孔福朗Conflens = BConflens孔福朗
	f.孔福朗Conflens.SetParent(f)

	f.福西尼Faucigny = BFaucigny福西尼
	f.福西尼Faucigny.SetParent(f)

	f.塔朗泰斯Tarentaise = BTarentaise塔朗泰斯
	f.塔朗泰斯Tarentaise.SetParent(f)

	f.托农Thonon = BThonon托农
	f.托农Thonon.SetParent(f)

}
