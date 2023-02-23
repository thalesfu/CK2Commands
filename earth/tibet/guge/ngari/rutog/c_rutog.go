package rutog

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RutogCounty interface {
	feud.County
	BDerok德汝() feud.Barony
	BDomar多玛() feud.Barony
	BGobak果巴卡() feud.Barony
	BRutog日托() feud.Barony
	BSumshi松西() feud.Barony
	BTsapuk扎普() feud.Barony
}

type 日托RutogCounty struct {
	feud.BaseCounty
	德汝Derok  feud.Barony
	多玛Domar  feud.Barony
	果巴卡Gobak feud.Barony
	日托Rutog  feud.Barony
	松西Sumshi feud.Barony
	扎普Tsapuk feud.Barony
}

func (c *日托RutogCounty) BDerok德汝() feud.Barony {
	return c.德汝Derok
}

func (c *日托RutogCounty) BDomar多玛() feud.Barony {
	return c.多玛Domar
}

func (c *日托RutogCounty) BGobak果巴卡() feud.Barony {
	return c.果巴卡Gobak
}

func (c *日托RutogCounty) BRutog日托() feud.Barony {
	return c.日托Rutog
}

func (c *日托RutogCounty) BSumshi松西() feud.Barony {
	return c.松西Sumshi
}

func (c *日托RutogCounty) BTsapuk扎普() feud.Barony {
	return c.扎普Tsapuk
}

var CRutog日托 RutogCounty = &日托RutogCounty{}

func init() {
	f := CRutog日托.(*日托RutogCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1485",
		Title:     "rutog",
		TitleName: "日托",
		TitleCode: "c_rutog",
		Baronies:  map[string]feud.Barony{},
	}

	f.德汝Derok = BDerok德汝
	f.德汝Derok.SetParent(f)

	f.多玛Domar = BDomar多玛
	f.多玛Domar.SetParent(f)

	f.果巴卡Gobak = BGobak果巴卡
	f.果巴卡Gobak.SetParent(f)

	f.日托Rutog = BRutog日托
	f.日托Rutog.SetParent(f)

	f.松西Sumshi = BSumshi松西
	f.松西Sumshi.SetParent(f)

	f.扎普Tsapuk = BTsapuk扎普
	f.扎普Tsapuk.SetParent(f)

}
