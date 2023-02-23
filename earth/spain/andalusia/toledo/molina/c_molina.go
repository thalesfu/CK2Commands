package molina

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MolinaCounty interface {
	feud.County
	BCabanillasdelcampo卡瓦尼利亚斯德尔坎波() feud.Barony
	BElcasar埃尔卡萨尔() feud.Barony
	BElpedregal埃尔佩德雷加尔() feud.Barony
	BHinojosa伊诺霍萨() feud.Barony
	BMaranchon马兰琼() feud.Barony
	BMolina莫利纳() feud.Barony
	BOlmeda奥尔梅达() feud.Barony
	BPinilla皮尼利亚() feud.Barony
}

type 莫利纳MolinaCounty struct {
	feud.BaseCounty
	卡瓦尼利亚斯德尔坎波Cabanillasdelcampo feud.Barony
	埃尔卡萨尔Elcasar                 feud.Barony
	埃尔佩德雷加尔Elpedregal            feud.Barony
	伊诺霍萨Hinojosa                 feud.Barony
	马兰琼Maranchon                 feud.Barony
	莫利纳Molina                    feud.Barony
	奥尔梅达Olmeda                   feud.Barony
	皮尼利亚Pinilla                  feud.Barony
}

func (c *莫利纳MolinaCounty) BCabanillasdelcampo卡瓦尼利亚斯德尔坎波() feud.Barony {
	return c.卡瓦尼利亚斯德尔坎波Cabanillasdelcampo
}

func (c *莫利纳MolinaCounty) BElcasar埃尔卡萨尔() feud.Barony {
	return c.埃尔卡萨尔Elcasar
}

func (c *莫利纳MolinaCounty) BElpedregal埃尔佩德雷加尔() feud.Barony {
	return c.埃尔佩德雷加尔Elpedregal
}

func (c *莫利纳MolinaCounty) BHinojosa伊诺霍萨() feud.Barony {
	return c.伊诺霍萨Hinojosa
}

func (c *莫利纳MolinaCounty) BMaranchon马兰琼() feud.Barony {
	return c.马兰琼Maranchon
}

func (c *莫利纳MolinaCounty) BMolina莫利纳() feud.Barony {
	return c.莫利纳Molina
}

func (c *莫利纳MolinaCounty) BOlmeda奥尔梅达() feud.Barony {
	return c.奥尔梅达Olmeda
}

func (c *莫利纳MolinaCounty) BPinilla皮尼利亚() feud.Barony {
	return c.皮尼利亚Pinilla
}

var CMolina莫利纳 MolinaCounty = &莫利纳MolinaCounty{}

func init() {
	f := CMolina莫利纳.(*莫利纳MolinaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "176",
		Title:     "molina",
		TitleName: "莫利纳",
		TitleCode: "c_molina",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡瓦尼利亚斯德尔坎波Cabanillasdelcampo = BCabanillasdelcampo卡瓦尼利亚斯德尔坎波
	f.卡瓦尼利亚斯德尔坎波Cabanillasdelcampo.SetParent(f)

	f.埃尔卡萨尔Elcasar = BElcasar埃尔卡萨尔
	f.埃尔卡萨尔Elcasar.SetParent(f)

	f.埃尔佩德雷加尔Elpedregal = BElpedregal埃尔佩德雷加尔
	f.埃尔佩德雷加尔Elpedregal.SetParent(f)

	f.伊诺霍萨Hinojosa = BHinojosa伊诺霍萨
	f.伊诺霍萨Hinojosa.SetParent(f)

	f.马兰琼Maranchon = BMaranchon马兰琼
	f.马兰琼Maranchon.SetParent(f)

	f.莫利纳Molina = BMolina莫利纳
	f.莫利纳Molina.SetParent(f)

	f.奥尔梅达Olmeda = BOlmeda奥尔梅达
	f.奥尔梅达Olmeda.SetParent(f)

	f.皮尼利亚Pinilla = BPinilla皮尼利亚
	f.皮尼利亚Pinilla.SetParent(f)

}
