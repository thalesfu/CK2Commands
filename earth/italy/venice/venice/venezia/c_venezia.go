package venezia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VeneziaCounty interface {
	feud.County
	BFusina富西纳() feud.Barony
	BJesolo耶索洛() feud.Barony
	BLido利多() feud.Barony
	BMurano穆拉诺() feud.Barony
	BPallestrina帕莱斯特里纳() feud.Barony
	BRialto里亚尔托() feud.Barony
	BTorcello托尔切洛() feud.Barony
	BVenezia威尼斯() feud.Barony
}

type 威尼斯VeneziaCounty struct {
	feud.BaseCounty
	富西纳Fusina         feud.Barony
	耶索洛Jesolo         feud.Barony
	利多Lido            feud.Barony
	穆拉诺Murano         feud.Barony
	帕莱斯特里纳Pallestrina feud.Barony
	里亚尔托Rialto        feud.Barony
	托尔切洛Torcello      feud.Barony
	威尼斯Venezia        feud.Barony
}

func (c *威尼斯VeneziaCounty) BFusina富西纳() feud.Barony {
	return c.富西纳Fusina
}

func (c *威尼斯VeneziaCounty) BJesolo耶索洛() feud.Barony {
	return c.耶索洛Jesolo
}

func (c *威尼斯VeneziaCounty) BLido利多() feud.Barony {
	return c.利多Lido
}

func (c *威尼斯VeneziaCounty) BMurano穆拉诺() feud.Barony {
	return c.穆拉诺Murano
}

func (c *威尼斯VeneziaCounty) BPallestrina帕莱斯特里纳() feud.Barony {
	return c.帕莱斯特里纳Pallestrina
}

func (c *威尼斯VeneziaCounty) BRialto里亚尔托() feud.Barony {
	return c.里亚尔托Rialto
}

func (c *威尼斯VeneziaCounty) BTorcello托尔切洛() feud.Barony {
	return c.托尔切洛Torcello
}

func (c *威尼斯VeneziaCounty) BVenezia威尼斯() feud.Barony {
	return c.威尼斯Venezia
}

var CVenezia威尼斯 VeneziaCounty = &威尼斯VeneziaCounty{}

func init() {
	f := CVenezia威尼斯.(*威尼斯VeneziaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "356",
		Title:     "venezia",
		TitleName: "威尼斯",
		TitleCode: "c_venezia",
		Baronies:  map[string]feud.Barony{},
	}

	f.富西纳Fusina = BFusina富西纳
	f.富西纳Fusina.SetParent(f)

	f.耶索洛Jesolo = BJesolo耶索洛
	f.耶索洛Jesolo.SetParent(f)

	f.利多Lido = BLido利多
	f.利多Lido.SetParent(f)

	f.穆拉诺Murano = BMurano穆拉诺
	f.穆拉诺Murano.SetParent(f)

	f.帕莱斯特里纳Pallestrina = BPallestrina帕莱斯特里纳
	f.帕莱斯特里纳Pallestrina.SetParent(f)

	f.里亚尔托Rialto = BRialto里亚尔托
	f.里亚尔托Rialto.SetParent(f)

	f.托尔切洛Torcello = BTorcello托尔切洛
	f.托尔切洛Torcello.SetParent(f)

	f.威尼斯Venezia = BVenezia威尼斯
	f.威尼斯Venezia.SetParent(f)

}
