package dwarasamudra

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DwarasamudraCounty interface {
	feud.County
	BArasikere阿尔西盖雷() feud.Barony
	BBelur贝卢尔() feud.Barony
	BDwarasamudra堕罗三慕达罗() feud.Barony
	BHassan哈桑() feud.Barony
	BJavagallu奢伐加卢() feud.Barony
	BKoravangala科罗伐那加拉() feud.Barony
	BSosavur苏萨乌尔() feud.Barony
}

type 堕罗三慕达罗DwarasamudraCounty struct {
	feud.BaseCounty
	阿尔西盖雷Arasikere     feud.Barony
	贝卢尔Belur           feud.Barony
	堕罗三慕达罗Dwarasamudra feud.Barony
	哈桑Hassan           feud.Barony
	奢伐加卢Javagallu      feud.Barony
	科罗伐那加拉Koravangala  feud.Barony
	苏萨乌尔Sosavur        feud.Barony
}

func (c *堕罗三慕达罗DwarasamudraCounty) BArasikere阿尔西盖雷() feud.Barony {
	return c.阿尔西盖雷Arasikere
}

func (c *堕罗三慕达罗DwarasamudraCounty) BBelur贝卢尔() feud.Barony {
	return c.贝卢尔Belur
}

func (c *堕罗三慕达罗DwarasamudraCounty) BDwarasamudra堕罗三慕达罗() feud.Barony {
	return c.堕罗三慕达罗Dwarasamudra
}

func (c *堕罗三慕达罗DwarasamudraCounty) BHassan哈桑() feud.Barony {
	return c.哈桑Hassan
}

func (c *堕罗三慕达罗DwarasamudraCounty) BJavagallu奢伐加卢() feud.Barony {
	return c.奢伐加卢Javagallu
}

func (c *堕罗三慕达罗DwarasamudraCounty) BKoravangala科罗伐那加拉() feud.Barony {
	return c.科罗伐那加拉Koravangala
}

func (c *堕罗三慕达罗DwarasamudraCounty) BSosavur苏萨乌尔() feud.Barony {
	return c.苏萨乌尔Sosavur
}

var CDwarasamudra堕罗三慕达罗 DwarasamudraCounty = &堕罗三慕达罗DwarasamudraCounty{}

func init() {
	f := CDwarasamudra堕罗三慕达罗.(*堕罗三慕达罗DwarasamudraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1197",
		Title:     "dwarasamudra",
		TitleName: "堕罗三慕达罗",
		TitleCode: "c_dwarasamudra",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔西盖雷Arasikere = BArasikere阿尔西盖雷
	f.阿尔西盖雷Arasikere.SetParent(f)

	f.贝卢尔Belur = BBelur贝卢尔
	f.贝卢尔Belur.SetParent(f)

	f.堕罗三慕达罗Dwarasamudra = BDwarasamudra堕罗三慕达罗
	f.堕罗三慕达罗Dwarasamudra.SetParent(f)

	f.哈桑Hassan = BHassan哈桑
	f.哈桑Hassan.SetParent(f)

	f.奢伐加卢Javagallu = BJavagallu奢伐加卢
	f.奢伐加卢Javagallu.SetParent(f)

	f.科罗伐那加拉Koravangala = BKoravangala科罗伐那加拉
	f.科罗伐那加拉Koravangala.SetParent(f)

	f.苏萨乌尔Sosavur = BSosavur苏萨乌尔
	f.苏萨乌尔Sosavur.SetParent(f)

}
