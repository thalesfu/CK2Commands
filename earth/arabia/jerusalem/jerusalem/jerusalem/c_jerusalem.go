package jerusalem

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type JerusalemCounty interface {
	feud.County
	BBeitnuba贝特努巴() feud.Barony
	BJericho杰里科() feud.Barony
	BJerusalem耶路撒冷() feud.Barony
	BMirabel米拉贝尔() feud.Barony
	BMontgisard蒙吉萨() feud.Barony
	BNablus纳布卢斯() feud.Barony
	BRammala拉马拉() feud.Barony
	BSaintsamuel圣撒慕尔() feud.Barony
}

type 耶路撒冷JerusalemCounty struct {
	feud.BaseCounty
	贝特努巴Beitnuba    feud.Barony
	杰里科Jericho      feud.Barony
	耶路撒冷Jerusalem   feud.Barony
	米拉贝尔Mirabel     feud.Barony
	蒙吉萨Montgisard   feud.Barony
	纳布卢斯Nablus      feud.Barony
	拉马拉Rammala      feud.Barony
	圣撒慕尔Saintsamuel feud.Barony
}

func (c *耶路撒冷JerusalemCounty) BBeitnuba贝特努巴() feud.Barony {
	return c.贝特努巴Beitnuba
}

func (c *耶路撒冷JerusalemCounty) BJericho杰里科() feud.Barony {
	return c.杰里科Jericho
}

func (c *耶路撒冷JerusalemCounty) BJerusalem耶路撒冷() feud.Barony {
	return c.耶路撒冷Jerusalem
}

func (c *耶路撒冷JerusalemCounty) BMirabel米拉贝尔() feud.Barony {
	return c.米拉贝尔Mirabel
}

func (c *耶路撒冷JerusalemCounty) BMontgisard蒙吉萨() feud.Barony {
	return c.蒙吉萨Montgisard
}

func (c *耶路撒冷JerusalemCounty) BNablus纳布卢斯() feud.Barony {
	return c.纳布卢斯Nablus
}

func (c *耶路撒冷JerusalemCounty) BRammala拉马拉() feud.Barony {
	return c.拉马拉Rammala
}

func (c *耶路撒冷JerusalemCounty) BSaintsamuel圣撒慕尔() feud.Barony {
	return c.圣撒慕尔Saintsamuel
}

var CJerusalem耶路撒冷 JerusalemCounty = &耶路撒冷JerusalemCounty{}

func init() {
	f := CJerusalem耶路撒冷.(*耶路撒冷JerusalemCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "774",
		Title:     "jerusalem",
		TitleName: "耶路撒冷",
		TitleCode: "c_jerusalem",
		Baronies:  map[string]feud.Barony{},
	}

	f.贝特努巴Beitnuba = BBeitnuba贝特努巴
	f.贝特努巴Beitnuba.SetParent(f)

	f.杰里科Jericho = BJericho杰里科
	f.杰里科Jericho.SetParent(f)

	f.耶路撒冷Jerusalem = BJerusalem耶路撒冷
	f.耶路撒冷Jerusalem.SetParent(f)

	f.米拉贝尔Mirabel = BMirabel米拉贝尔
	f.米拉贝尔Mirabel.SetParent(f)

	f.蒙吉萨Montgisard = BMontgisard蒙吉萨
	f.蒙吉萨Montgisard.SetParent(f)

	f.纳布卢斯Nablus = BNablus纳布卢斯
	f.纳布卢斯Nablus.SetParent(f)

	f.拉马拉Rammala = BRammala拉马拉
	f.拉马拉Rammala.SetParent(f)

	f.圣撒慕尔Saintsamuel = BSaintsamuel圣撒慕尔
	f.圣撒慕尔Saintsamuel.SetParent(f)

}
