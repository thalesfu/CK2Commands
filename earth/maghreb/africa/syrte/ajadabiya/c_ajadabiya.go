package ajadabiya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AjadabiyaCounty interface {
	feud.County
	BAjadabiya艾季达比耶() feud.Barony
	BGargaresc杰尔加雷什() feud.Barony
	BKarsah凯尔萨() feud.Barony
	BQaryatbishr喀鲁特比斯() feud.Barony
	BTagma泰格梅() feud.Barony
}

type 艾季达比耶AjadabiyaCounty struct {
	feud.BaseCounty
	艾季达比耶Ajadabiya   feud.Barony
	杰尔加雷什Gargaresc   feud.Barony
	凯尔萨Karsah        feud.Barony
	喀鲁特比斯Qaryatbishr feud.Barony
	泰格梅Tagma         feud.Barony
}

func (c *艾季达比耶AjadabiyaCounty) BAjadabiya艾季达比耶() feud.Barony {
	return c.艾季达比耶Ajadabiya
}

func (c *艾季达比耶AjadabiyaCounty) BGargaresc杰尔加雷什() feud.Barony {
	return c.杰尔加雷什Gargaresc
}

func (c *艾季达比耶AjadabiyaCounty) BKarsah凯尔萨() feud.Barony {
	return c.凯尔萨Karsah
}

func (c *艾季达比耶AjadabiyaCounty) BQaryatbishr喀鲁特比斯() feud.Barony {
	return c.喀鲁特比斯Qaryatbishr
}

func (c *艾季达比耶AjadabiyaCounty) BTagma泰格梅() feud.Barony {
	return c.泰格梅Tagma
}

var CAjadabiya艾季达比耶 AjadabiyaCounty = &艾季达比耶AjadabiyaCounty{}

func init() {
	f := CAjadabiya艾季达比耶.(*艾季达比耶AjadabiyaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1721",
		Title:     "ajadabiya",
		TitleName: "艾季达比耶",
		TitleCode: "c_ajadabiya",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾季达比耶Ajadabiya = BAjadabiya艾季达比耶
	f.艾季达比耶Ajadabiya.SetParent(f)

	f.杰尔加雷什Gargaresc = BGargaresc杰尔加雷什
	f.杰尔加雷什Gargaresc.SetParent(f)

	f.凯尔萨Karsah = BKarsah凯尔萨
	f.凯尔萨Karsah.SetParent(f)

	f.喀鲁特比斯Qaryatbishr = BQaryatbishr喀鲁特比斯
	f.喀鲁特比斯Qaryatbishr.SetParent(f)

	f.泰格梅Tagma = BTagma泰格梅
	f.泰格梅Tagma.SetParent(f)

}
