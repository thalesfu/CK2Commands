package cebta

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CebtaCounty interface {
	feud.County
	BAulef奥莱法() feud.Barony
	BBabtaza巴卜塔宰() feud.Barony
	BCueta休达() feud.Barony
	BMartil迈尔提勒() feud.Barony
	BMdiq迈迪格() feud.Barony
	BTarguist塔尔吉斯特() feud.Barony
	BTetouan得土安() feud.Barony
}

type 休达CebtaCounty struct {
	feud.BaseCounty
	奥莱法Aulef      feud.Barony
	巴卜塔宰Babtaza   feud.Barony
	休达Cueta       feud.Barony
	迈尔提勒Martil    feud.Barony
	迈迪格Mdiq       feud.Barony
	塔尔吉斯特Targuist feud.Barony
	得土安Tetouan    feud.Barony
}

func (c *休达CebtaCounty) BAulef奥莱法() feud.Barony {
	return c.奥莱法Aulef
}

func (c *休达CebtaCounty) BBabtaza巴卜塔宰() feud.Barony {
	return c.巴卜塔宰Babtaza
}

func (c *休达CebtaCounty) BCueta休达() feud.Barony {
	return c.休达Cueta
}

func (c *休达CebtaCounty) BMartil迈尔提勒() feud.Barony {
	return c.迈尔提勒Martil
}

func (c *休达CebtaCounty) BMdiq迈迪格() feud.Barony {
	return c.迈迪格Mdiq
}

func (c *休达CebtaCounty) BTarguist塔尔吉斯特() feud.Barony {
	return c.塔尔吉斯特Targuist
}

func (c *休达CebtaCounty) BTetouan得土安() feud.Barony {
	return c.得土安Tetouan
}

var CCebta休达 CebtaCounty = &休达CebtaCounty{}

func init() {
	f := CCebta休达.(*休达CebtaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "839",
		Title:     "cebta",
		TitleName: "休达",
		TitleCode: "c_cebta",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥莱法Aulef = BAulef奥莱法
	f.奥莱法Aulef.SetParent(f)

	f.巴卜塔宰Babtaza = BBabtaza巴卜塔宰
	f.巴卜塔宰Babtaza.SetParent(f)

	f.休达Cueta = BCueta休达
	f.休达Cueta.SetParent(f)

	f.迈尔提勒Martil = BMartil迈尔提勒
	f.迈尔提勒Martil.SetParent(f)

	f.迈迪格Mdiq = BMdiq迈迪格
	f.迈迪格Mdiq.SetParent(f)

	f.塔尔吉斯特Targuist = BTarguist塔尔吉斯特
	f.塔尔吉斯特Targuist.SetParent(f)

	f.得土安Tetouan = BTetouan得土安
	f.得土安Tetouan.SetParent(f)

}
