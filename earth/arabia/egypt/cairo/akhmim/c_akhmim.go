package akhmim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AkhmimCounty interface {
	feud.County
	BAkhmim伊赫明() feud.Barony
	BAntinoe安提倪() feud.Barony
	BBesa贝萨() feud.Barony
	BCynopolis居诺波利斯() feud.Barony
	BHibeh希贝赫() feud.Barony
	BSharuna沙鲁纳() feud.Barony
}

type 阿赫明AkhmimCounty struct {
	feud.BaseCounty
	伊赫明Akhmim      feud.Barony
	安提倪Antinoe     feud.Barony
	贝萨Besa         feud.Barony
	居诺波利斯Cynopolis feud.Barony
	希贝赫Hibeh       feud.Barony
	沙鲁纳Sharuna     feud.Barony
}

func (c *阿赫明AkhmimCounty) BAkhmim伊赫明() feud.Barony {
	return c.伊赫明Akhmim
}

func (c *阿赫明AkhmimCounty) BAntinoe安提倪() feud.Barony {
	return c.安提倪Antinoe
}

func (c *阿赫明AkhmimCounty) BBesa贝萨() feud.Barony {
	return c.贝萨Besa
}

func (c *阿赫明AkhmimCounty) BCynopolis居诺波利斯() feud.Barony {
	return c.居诺波利斯Cynopolis
}

func (c *阿赫明AkhmimCounty) BHibeh希贝赫() feud.Barony {
	return c.希贝赫Hibeh
}

func (c *阿赫明AkhmimCounty) BSharuna沙鲁纳() feud.Barony {
	return c.沙鲁纳Sharuna
}

var CAkhmim阿赫明 AkhmimCounty = &阿赫明AkhmimCounty{}

func init() {
	f := CAkhmim阿赫明.(*阿赫明AkhmimCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "2002",
		Title:     "akhmim",
		TitleName: "阿赫明",
		TitleCode: "c_akhmim",
		Baronies:  map[string]feud.Barony{},
	}

	f.伊赫明Akhmim = BAkhmim伊赫明
	f.伊赫明Akhmim.SetParent(f)

	f.安提倪Antinoe = BAntinoe安提倪
	f.安提倪Antinoe.SetParent(f)

	f.贝萨Besa = BBesa贝萨
	f.贝萨Besa.SetParent(f)

	f.居诺波利斯Cynopolis = BCynopolis居诺波利斯
	f.居诺波利斯Cynopolis.SetParent(f)

	f.希贝赫Hibeh = BHibeh希贝赫
	f.希贝赫Hibeh.SetParent(f)

	f.沙鲁纳Sharuna = BSharuna沙鲁纳
	f.沙鲁纳Sharuna.SetParent(f)

}
