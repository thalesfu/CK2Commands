package navarra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NavarraCounty interface {
	feud.County
	BCarcastillo卡尔卡斯蒂略() feud.Barony
	BEstella埃斯特利亚() feud.Barony
	BLeyre莱雷() feud.Barony
	BOlite奥利特() feud.Barony
	BPamplona潘普洛纳() feud.Barony
	BSanguesa桑圭萨() feud.Barony
	BTafalla塔法利亚() feud.Barony
	BTudela图德拉() feud.Barony
}

type 纳瓦拉NavarraCounty struct {
	feud.BaseCounty
	卡尔卡斯蒂略Carcastillo feud.Barony
	埃斯特利亚Estella      feud.Barony
	莱雷Leyre           feud.Barony
	奥利特Olite          feud.Barony
	潘普洛纳Pamplona      feud.Barony
	桑圭萨Sanguesa       feud.Barony
	塔法利亚Tafalla       feud.Barony
	图德拉Tudela         feud.Barony
}

func (c *纳瓦拉NavarraCounty) BCarcastillo卡尔卡斯蒂略() feud.Barony {
	return c.卡尔卡斯蒂略Carcastillo
}

func (c *纳瓦拉NavarraCounty) BEstella埃斯特利亚() feud.Barony {
	return c.埃斯特利亚Estella
}

func (c *纳瓦拉NavarraCounty) BLeyre莱雷() feud.Barony {
	return c.莱雷Leyre
}

func (c *纳瓦拉NavarraCounty) BOlite奥利特() feud.Barony {
	return c.奥利特Olite
}

func (c *纳瓦拉NavarraCounty) BPamplona潘普洛纳() feud.Barony {
	return c.潘普洛纳Pamplona
}

func (c *纳瓦拉NavarraCounty) BSanguesa桑圭萨() feud.Barony {
	return c.桑圭萨Sanguesa
}

func (c *纳瓦拉NavarraCounty) BTafalla塔法利亚() feud.Barony {
	return c.塔法利亚Tafalla
}

func (c *纳瓦拉NavarraCounty) BTudela图德拉() feud.Barony {
	return c.图德拉Tudela
}

var CNavarra纳瓦拉 NavarraCounty = &纳瓦拉NavarraCounty{}

func init() {
	f := CNavarra纳瓦拉.(*纳瓦拉NavarraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "152",
		Title:     "navarra",
		TitleName: "纳瓦拉",
		TitleCode: "c_navarra",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡尔卡斯蒂略Carcastillo = BCarcastillo卡尔卡斯蒂略
	f.卡尔卡斯蒂略Carcastillo.SetParent(f)

	f.埃斯特利亚Estella = BEstella埃斯特利亚
	f.埃斯特利亚Estella.SetParent(f)

	f.莱雷Leyre = BLeyre莱雷
	f.莱雷Leyre.SetParent(f)

	f.奥利特Olite = BOlite奥利特
	f.奥利特Olite.SetParent(f)

	f.潘普洛纳Pamplona = BPamplona潘普洛纳
	f.潘普洛纳Pamplona.SetParent(f)

	f.桑圭萨Sanguesa = BSanguesa桑圭萨
	f.桑圭萨Sanguesa.SetParent(f)

	f.塔法利亚Tafalla = BTafalla塔法利亚
	f.塔法利亚Tafalla.SetParent(f)

	f.图德拉Tudela = BTudela图德拉
	f.图德拉Tudela.SetParent(f)

}
