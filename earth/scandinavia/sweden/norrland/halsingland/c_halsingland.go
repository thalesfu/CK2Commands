package halsingland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HalsinglandCounty interface {
	feud.County
	BAlir阿利尔() feud.Barony
	BForsa佛尔萨() feud.Barony
	BHog赫格() feud.Barony
	BJattendal耶滕达尔() feud.Barony
	BNordanstig努丹斯蒂格() feud.Barony
	BNorrala诺腊拉() feud.Barony
	BSundhed松德海德() feud.Barony
	BTasta托斯塔() feud.Barony
}

type 海尔辛兰HalsinglandCounty struct {
	feud.BaseCounty
	阿利尔Alir         feud.Barony
	佛尔萨Forsa        feud.Barony
	赫格Hog           feud.Barony
	耶滕达尔Jattendal   feud.Barony
	努丹斯蒂格Nordanstig feud.Barony
	诺腊拉Norrala      feud.Barony
	松德海德Sundhed     feud.Barony
	托斯塔Tasta        feud.Barony
}

func (c *海尔辛兰HalsinglandCounty) BAlir阿利尔() feud.Barony {
	return c.阿利尔Alir
}

func (c *海尔辛兰HalsinglandCounty) BForsa佛尔萨() feud.Barony {
	return c.佛尔萨Forsa
}

func (c *海尔辛兰HalsinglandCounty) BHog赫格() feud.Barony {
	return c.赫格Hog
}

func (c *海尔辛兰HalsinglandCounty) BJattendal耶滕达尔() feud.Barony {
	return c.耶滕达尔Jattendal
}

func (c *海尔辛兰HalsinglandCounty) BNordanstig努丹斯蒂格() feud.Barony {
	return c.努丹斯蒂格Nordanstig
}

func (c *海尔辛兰HalsinglandCounty) BNorrala诺腊拉() feud.Barony {
	return c.诺腊拉Norrala
}

func (c *海尔辛兰HalsinglandCounty) BSundhed松德海德() feud.Barony {
	return c.松德海德Sundhed
}

func (c *海尔辛兰HalsinglandCounty) BTasta托斯塔() feud.Barony {
	return c.托斯塔Tasta
}

var CHalsingland海尔辛兰 HalsinglandCounty = &海尔辛兰HalsinglandCounty{}

func init() {
	f := CHalsingland海尔辛兰.(*海尔辛兰HalsinglandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "285",
		Title:     "halsingland",
		TitleName: "海尔辛兰",
		TitleCode: "c_halsingland",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿利尔Alir = BAlir阿利尔
	f.阿利尔Alir.SetParent(f)

	f.佛尔萨Forsa = BForsa佛尔萨
	f.佛尔萨Forsa.SetParent(f)

	f.赫格Hog = BHog赫格
	f.赫格Hog.SetParent(f)

	f.耶滕达尔Jattendal = BJattendal耶滕达尔
	f.耶滕达尔Jattendal.SetParent(f)

	f.努丹斯蒂格Nordanstig = BNordanstig努丹斯蒂格
	f.努丹斯蒂格Nordanstig.SetParent(f)

	f.诺腊拉Norrala = BNorrala诺腊拉
	f.诺腊拉Norrala.SetParent(f)

	f.松德海德Sundhed = BSundhed松德海德
	f.松德海德Sundhed.SetParent(f)

	f.托斯塔Tasta = BTasta托斯塔
	f.托斯塔Tasta.SetParent(f)

}
