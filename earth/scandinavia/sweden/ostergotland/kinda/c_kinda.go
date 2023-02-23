package kinda

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KindaCounty interface {
	feud.County
	BHagerstad海格斯塔德() feud.Barony
	BHycklinge许克灵厄() feud.Barony
	BKisa希萨() feud.Barony
	BOppeby奥珀比() feud.Barony
	BTidersrum蒂德什鲁姆() feud.Barony
	BTjarstad蒂耶什塔德() feud.Barony
}

type 欣达KindaCounty struct {
	feud.BaseCounty
	海格斯塔德Hagerstad feud.Barony
	许克灵厄Hycklinge  feud.Barony
	希萨Kisa         feud.Barony
	奥珀比Oppeby      feud.Barony
	蒂德什鲁姆Tidersrum feud.Barony
	蒂耶什塔德Tjarstad  feud.Barony
}

func (c *欣达KindaCounty) BHagerstad海格斯塔德() feud.Barony {
	return c.海格斯塔德Hagerstad
}

func (c *欣达KindaCounty) BHycklinge许克灵厄() feud.Barony {
	return c.许克灵厄Hycklinge
}

func (c *欣达KindaCounty) BKisa希萨() feud.Barony {
	return c.希萨Kisa
}

func (c *欣达KindaCounty) BOppeby奥珀比() feud.Barony {
	return c.奥珀比Oppeby
}

func (c *欣达KindaCounty) BTidersrum蒂德什鲁姆() feud.Barony {
	return c.蒂德什鲁姆Tidersrum
}

func (c *欣达KindaCounty) BTjarstad蒂耶什塔德() feud.Barony {
	return c.蒂耶什塔德Tjarstad
}

var CKinda欣达 KindaCounty = &欣达KindaCounty{}

func init() {
	f := CKinda欣达.(*欣达KindaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1705",
		Title:     "kinda",
		TitleName: "欣达",
		TitleCode: "c_kinda",
		Baronies:  map[string]feud.Barony{},
	}

	f.海格斯塔德Hagerstad = BHagerstad海格斯塔德
	f.海格斯塔德Hagerstad.SetParent(f)

	f.许克灵厄Hycklinge = BHycklinge许克灵厄
	f.许克灵厄Hycklinge.SetParent(f)

	f.希萨Kisa = BKisa希萨
	f.希萨Kisa.SetParent(f)

	f.奥珀比Oppeby = BOppeby奥珀比
	f.奥珀比Oppeby.SetParent(f)

	f.蒂德什鲁姆Tidersrum = BTidersrum蒂德什鲁姆
	f.蒂德什鲁姆Tidersrum.SetParent(f)

	f.蒂耶什塔德Tjarstad = BTjarstad蒂耶什塔德
	f.蒂耶什塔德Tjarstad.SetParent(f)

}
