package onogost

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OnogostCounty interface {
	feud.County
	BBudimlja布迪姆利亚() feud.Barony
	BDesnik德尔斯尼克() feud.Barony
	BDestinik德斯蒂尼克() feud.Barony
	BGradina格拉迪纳() feud.Barony
	BOnogost奥诺戈什特() feud.Barony
	BRibnica里布尼察() feud.Barony
}

type 奥诺戈什特OnogostCounty struct {
	feud.BaseCounty
	布迪姆利亚Budimlja feud.Barony
	德尔斯尼克Desnik   feud.Barony
	德斯蒂尼克Destinik feud.Barony
	格拉迪纳Gradina   feud.Barony
	奥诺戈什特Onogost  feud.Barony
	里布尼察Ribnica   feud.Barony
}

func (c *奥诺戈什特OnogostCounty) BBudimlja布迪姆利亚() feud.Barony {
	return c.布迪姆利亚Budimlja
}

func (c *奥诺戈什特OnogostCounty) BDesnik德尔斯尼克() feud.Barony {
	return c.德尔斯尼克Desnik
}

func (c *奥诺戈什特OnogostCounty) BDestinik德斯蒂尼克() feud.Barony {
	return c.德斯蒂尼克Destinik
}

func (c *奥诺戈什特OnogostCounty) BGradina格拉迪纳() feud.Barony {
	return c.格拉迪纳Gradina
}

func (c *奥诺戈什特OnogostCounty) BOnogost奥诺戈什特() feud.Barony {
	return c.奥诺戈什特Onogost
}

func (c *奥诺戈什特OnogostCounty) BRibnica里布尼察() feud.Barony {
	return c.里布尼察Ribnica
}

var COnogost奥诺戈什特 OnogostCounty = &奥诺戈什特OnogostCounty{}

func init() {
	f := COnogost奥诺戈什特.(*奥诺戈什特OnogostCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1973",
		Title:     "onogost",
		TitleName: "奥诺戈什特",
		TitleCode: "c_onogost",
		Baronies:  map[string]feud.Barony{},
	}

	f.布迪姆利亚Budimlja = BBudimlja布迪姆利亚
	f.布迪姆利亚Budimlja.SetParent(f)

	f.德尔斯尼克Desnik = BDesnik德尔斯尼克
	f.德尔斯尼克Desnik.SetParent(f)

	f.德斯蒂尼克Destinik = BDestinik德斯蒂尼克
	f.德斯蒂尼克Destinik.SetParent(f)

	f.格拉迪纳Gradina = BGradina格拉迪纳
	f.格拉迪纳Gradina.SetParent(f)

	f.奥诺戈什特Onogost = BOnogost奥诺戈什特
	f.奥诺戈什特Onogost.SetParent(f)

	f.里布尼察Ribnica = BRibnica里布尼察
	f.里布尼察Ribnica.SetParent(f)

}
