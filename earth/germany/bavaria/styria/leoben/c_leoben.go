package leoben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeobenCounty interface {
	feud.County
	BEppenstein埃彭施泰因() feud.Barony
	BGallenstein加仑施泰因() feud.Barony
	BJudenburg尤登堡() feud.Barony
	BLeoben莱奥本() feud.Barony
	BNeuberg诺伊贝格() feud.Barony
	BRottenmann罗滕曼() feud.Barony
}

type 埃彭施泰因LeobenCounty struct {
	feud.BaseCounty
	埃彭施泰因Eppenstein  feud.Barony
	加仑施泰因Gallenstein feud.Barony
	尤登堡Judenburg     feud.Barony
	莱奥本Leoben        feud.Barony
	诺伊贝格Neuberg      feud.Barony
	罗滕曼Rottenmann    feud.Barony
}

func (c *埃彭施泰因LeobenCounty) BEppenstein埃彭施泰因() feud.Barony {
	return c.埃彭施泰因Eppenstein
}

func (c *埃彭施泰因LeobenCounty) BGallenstein加仑施泰因() feud.Barony {
	return c.加仑施泰因Gallenstein
}

func (c *埃彭施泰因LeobenCounty) BJudenburg尤登堡() feud.Barony {
	return c.尤登堡Judenburg
}

func (c *埃彭施泰因LeobenCounty) BLeoben莱奥本() feud.Barony {
	return c.莱奥本Leoben
}

func (c *埃彭施泰因LeobenCounty) BNeuberg诺伊贝格() feud.Barony {
	return c.诺伊贝格Neuberg
}

func (c *埃彭施泰因LeobenCounty) BRottenmann罗滕曼() feud.Barony {
	return c.罗滕曼Rottenmann
}

var CLeoben埃彭施泰因 LeobenCounty = &埃彭施泰因LeobenCounty{}

func init() {
	f := CLeoben埃彭施泰因.(*埃彭施泰因LeobenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1692",
		Title:     "leoben",
		TitleName: "埃彭施泰因",
		TitleCode: "c_leoben",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃彭施泰因Eppenstein = BEppenstein埃彭施泰因
	f.埃彭施泰因Eppenstein.SetParent(f)

	f.加仑施泰因Gallenstein = BGallenstein加仑施泰因
	f.加仑施泰因Gallenstein.SetParent(f)

	f.尤登堡Judenburg = BJudenburg尤登堡
	f.尤登堡Judenburg.SetParent(f)

	f.莱奥本Leoben = BLeoben莱奥本
	f.莱奥本Leoben.SetParent(f)

	f.诺伊贝格Neuberg = BNeuberg诺伊贝格
	f.诺伊贝格Neuberg.SetParent(f)

	f.罗滕曼Rottenmann = BRottenmann罗滕曼
	f.罗滕曼Rottenmann.SetParent(f)

}
