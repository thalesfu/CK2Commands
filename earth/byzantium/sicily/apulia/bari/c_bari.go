package bari

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BariCounty interface {
	feud.County
	BAndria安德里亚() feud.Barony
	BBari巴里() feud.Barony
	BBitonto比通托() feud.Barony
	BConversano孔韦尔萨诺() feud.Barony
	BGiovinazzo焦维纳佐() feud.Barony
	BMolietta莫列塔() feud.Barony
	BPolignano波利尼亚诺() feud.Barony
	BRuvo鲁沃() feud.Barony
}

type 巴里BariCounty struct {
	feud.BaseCounty
	安德里亚Andria      feud.Barony
	巴里Bari          feud.Barony
	比通托Bitonto      feud.Barony
	孔韦尔萨诺Conversano feud.Barony
	焦维纳佐Giovinazzo  feud.Barony
	莫列塔Molietta     feud.Barony
	波利尼亚诺Polignano  feud.Barony
	鲁沃Ruvo          feud.Barony
}

func (c *巴里BariCounty) BAndria安德里亚() feud.Barony {
	return c.安德里亚Andria
}

func (c *巴里BariCounty) BBari巴里() feud.Barony {
	return c.巴里Bari
}

func (c *巴里BariCounty) BBitonto比通托() feud.Barony {
	return c.比通托Bitonto
}

func (c *巴里BariCounty) BConversano孔韦尔萨诺() feud.Barony {
	return c.孔韦尔萨诺Conversano
}

func (c *巴里BariCounty) BGiovinazzo焦维纳佐() feud.Barony {
	return c.焦维纳佐Giovinazzo
}

func (c *巴里BariCounty) BMolietta莫列塔() feud.Barony {
	return c.莫列塔Molietta
}

func (c *巴里BariCounty) BPolignano波利尼亚诺() feud.Barony {
	return c.波利尼亚诺Polignano
}

func (c *巴里BariCounty) BRuvo鲁沃() feud.Barony {
	return c.鲁沃Ruvo
}

var CBari巴里 BariCounty = &巴里BariCounty{}

func init() {
	f := CBari巴里.(*巴里BariCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "346",
		Title:     "bari",
		TitleName: "巴里",
		TitleCode: "c_bari",
		Baronies:  map[string]feud.Barony{},
	}

	f.安德里亚Andria = BAndria安德里亚
	f.安德里亚Andria.SetParent(f)

	f.巴里Bari = BBari巴里
	f.巴里Bari.SetParent(f)

	f.比通托Bitonto = BBitonto比通托
	f.比通托Bitonto.SetParent(f)

	f.孔韦尔萨诺Conversano = BConversano孔韦尔萨诺
	f.孔韦尔萨诺Conversano.SetParent(f)

	f.焦维纳佐Giovinazzo = BGiovinazzo焦维纳佐
	f.焦维纳佐Giovinazzo.SetParent(f)

	f.莫列塔Molietta = BMolietta莫列塔
	f.莫列塔Molietta.SetParent(f)

	f.波利尼亚诺Polignano = BPolignano波利尼亚诺
	f.波利尼亚诺Polignano.SetParent(f)

	f.鲁沃Ruvo = BRuvo鲁沃
	f.鲁沃Ruvo.SetParent(f)

}
