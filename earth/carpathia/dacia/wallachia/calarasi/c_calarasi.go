package calarasi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CalarasiCounty interface {
	feud.County
	BCalarasi克勒拉希() feud.Barony
	BCoconi科科尼() feud.Barony
	BCornatei科尔纳泰() feud.Barony
	BFetesti费泰什蒂() feud.Barony
	BFloci弗洛奇() feud.Barony
	BOltenita奥尔泰尼察() feud.Barony
	BSlobozia斯洛博齐亚() feud.Barony
}

type 克勒拉希CalarasiCounty struct {
	feud.BaseCounty
	克勒拉希Calarasi  feud.Barony
	科科尼Coconi     feud.Barony
	科尔纳泰Cornatei  feud.Barony
	费泰什蒂Fetesti   feud.Barony
	弗洛奇Floci      feud.Barony
	奥尔泰尼察Oltenita feud.Barony
	斯洛博齐亚Slobozia feud.Barony
}

func (c *克勒拉希CalarasiCounty) BCalarasi克勒拉希() feud.Barony {
	return c.克勒拉希Calarasi
}

func (c *克勒拉希CalarasiCounty) BCoconi科科尼() feud.Barony {
	return c.科科尼Coconi
}

func (c *克勒拉希CalarasiCounty) BCornatei科尔纳泰() feud.Barony {
	return c.科尔纳泰Cornatei
}

func (c *克勒拉希CalarasiCounty) BFetesti费泰什蒂() feud.Barony {
	return c.费泰什蒂Fetesti
}

func (c *克勒拉希CalarasiCounty) BFloci弗洛奇() feud.Barony {
	return c.弗洛奇Floci
}

func (c *克勒拉希CalarasiCounty) BOltenita奥尔泰尼察() feud.Barony {
	return c.奥尔泰尼察Oltenita
}

func (c *克勒拉希CalarasiCounty) BSlobozia斯洛博齐亚() feud.Barony {
	return c.斯洛博齐亚Slobozia
}

var CCalarasi克勒拉希 CalarasiCounty = &克勒拉希CalarasiCounty{}

func init() {
	f := CCalarasi克勒拉希.(*克勒拉希CalarasiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1642",
		Title:     "calarasi",
		TitleName: "克勒拉希",
		TitleCode: "c_calarasi",
		Baronies:  map[string]feud.Barony{},
	}

	f.克勒拉希Calarasi = BCalarasi克勒拉希
	f.克勒拉希Calarasi.SetParent(f)

	f.科科尼Coconi = BCoconi科科尼
	f.科科尼Coconi.SetParent(f)

	f.科尔纳泰Cornatei = BCornatei科尔纳泰
	f.科尔纳泰Cornatei.SetParent(f)

	f.费泰什蒂Fetesti = BFetesti费泰什蒂
	f.费泰什蒂Fetesti.SetParent(f)

	f.弗洛奇Floci = BFloci弗洛奇
	f.弗洛奇Floci.SetParent(f)

	f.奥尔泰尼察Oltenita = BOltenita奥尔泰尼察
	f.奥尔泰尼察Oltenita.SetParent(f)

	f.斯洛博齐亚Slobozia = BSlobozia斯洛博齐亚
	f.斯洛博齐亚Slobozia.SetParent(f)

}
