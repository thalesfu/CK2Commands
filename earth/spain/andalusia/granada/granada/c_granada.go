package granada

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GranadaCounty interface {
	feud.County
	BAntequara安特克拉() feud.Barony
	BElvira埃尔薇拉() feud.Barony
	BGranada格拉纳达() feud.Barony
	BGuadix瓜迪克斯() feud.Barony
	BHuelma韦尔马() feud.Barony
	BIznajar伊斯纳哈尔() feud.Barony
	BMoclin莫克林() feud.Barony
}

type 格拉纳达GranadaCounty struct {
	feud.BaseCounty
	安特克拉Antequara feud.Barony
	埃尔薇拉Elvira    feud.Barony
	格拉纳达Granada   feud.Barony
	瓜迪克斯Guadix    feud.Barony
	韦尔马Huelma     feud.Barony
	伊斯纳哈尔Iznajar  feud.Barony
	莫克林Moclin     feud.Barony
}

func (c *格拉纳达GranadaCounty) BAntequara安特克拉() feud.Barony {
	return c.安特克拉Antequara
}

func (c *格拉纳达GranadaCounty) BElvira埃尔薇拉() feud.Barony {
	return c.埃尔薇拉Elvira
}

func (c *格拉纳达GranadaCounty) BGranada格拉纳达() feud.Barony {
	return c.格拉纳达Granada
}

func (c *格拉纳达GranadaCounty) BGuadix瓜迪克斯() feud.Barony {
	return c.瓜迪克斯Guadix
}

func (c *格拉纳达GranadaCounty) BHuelma韦尔马() feud.Barony {
	return c.韦尔马Huelma
}

func (c *格拉纳达GranadaCounty) BIznajar伊斯纳哈尔() feud.Barony {
	return c.伊斯纳哈尔Iznajar
}

func (c *格拉纳达GranadaCounty) BMoclin莫克林() feud.Barony {
	return c.莫克林Moclin
}

var CGranada格拉纳达 GranadaCounty = &格拉纳达GranadaCounty{}

func init() {
	f := CGranada格拉纳达.(*格拉纳达GranadaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "180",
		Title:     "granada",
		TitleName: "格拉纳达",
		TitleCode: "c_granada",
		Baronies:  map[string]feud.Barony{},
	}

	f.安特克拉Antequara = BAntequara安特克拉
	f.安特克拉Antequara.SetParent(f)

	f.埃尔薇拉Elvira = BElvira埃尔薇拉
	f.埃尔薇拉Elvira.SetParent(f)

	f.格拉纳达Granada = BGranada格拉纳达
	f.格拉纳达Granada.SetParent(f)

	f.瓜迪克斯Guadix = BGuadix瓜迪克斯
	f.瓜迪克斯Guadix.SetParent(f)

	f.韦尔马Huelma = BHuelma韦尔马
	f.韦尔马Huelma.SetParent(f)

	f.伊斯纳哈尔Iznajar = BIznajar伊斯纳哈尔
	f.伊斯纳哈尔Iznajar.SetParent(f)

	f.莫克林Moclin = BMoclin莫克林
	f.莫克林Moclin.SetParent(f)

}
