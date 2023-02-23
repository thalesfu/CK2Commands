package slutsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SlutskCounty interface {
	feud.County
	BGomel戈梅利() feud.Barony
	BKazlovichy科兹洛维奇() feud.Barony
	BLoyew洛耶夫() feud.Barony
	BMazyr莫济里() feud.Barony
	BRechytsa列奇察() feud.Barony
	BSlutsk斯卢茨克() feud.Barony
	BZhytkavichy日特科维奇() feud.Barony
}

type 斯卢茨克SlutskCounty struct {
	feud.BaseCounty
	戈梅利Gomel         feud.Barony
	科兹洛维奇Kazlovichy  feud.Barony
	洛耶夫Loyew         feud.Barony
	莫济里Mazyr         feud.Barony
	列奇察Rechytsa      feud.Barony
	斯卢茨克Slutsk       feud.Barony
	日特科维奇Zhytkavichy feud.Barony
}

func (c *斯卢茨克SlutskCounty) BGomel戈梅利() feud.Barony {
	return c.戈梅利Gomel
}

func (c *斯卢茨克SlutskCounty) BKazlovichy科兹洛维奇() feud.Barony {
	return c.科兹洛维奇Kazlovichy
}

func (c *斯卢茨克SlutskCounty) BLoyew洛耶夫() feud.Barony {
	return c.洛耶夫Loyew
}

func (c *斯卢茨克SlutskCounty) BMazyr莫济里() feud.Barony {
	return c.莫济里Mazyr
}

func (c *斯卢茨克SlutskCounty) BRechytsa列奇察() feud.Barony {
	return c.列奇察Rechytsa
}

func (c *斯卢茨克SlutskCounty) BSlutsk斯卢茨克() feud.Barony {
	return c.斯卢茨克Slutsk
}

func (c *斯卢茨克SlutskCounty) BZhytkavichy日特科维奇() feud.Barony {
	return c.日特科维奇Zhytkavichy
}

var CSlutsk斯卢茨克 SlutskCounty = &斯卢茨克SlutskCounty{}

func init() {
	f := CSlutsk斯卢茨克.(*斯卢茨克SlutskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1653",
		Title:     "slutsk",
		TitleName: "斯卢茨克",
		TitleCode: "c_slutsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.戈梅利Gomel = BGomel戈梅利
	f.戈梅利Gomel.SetParent(f)

	f.科兹洛维奇Kazlovichy = BKazlovichy科兹洛维奇
	f.科兹洛维奇Kazlovichy.SetParent(f)

	f.洛耶夫Loyew = BLoyew洛耶夫
	f.洛耶夫Loyew.SetParent(f)

	f.莫济里Mazyr = BMazyr莫济里
	f.莫济里Mazyr.SetParent(f)

	f.列奇察Rechytsa = BRechytsa列奇察
	f.列奇察Rechytsa.SetParent(f)

	f.斯卢茨克Slutsk = BSlutsk斯卢茨克
	f.斯卢茨克Slutsk.SetParent(f)

	f.日特科维奇Zhytkavichy = BZhytkavichy日特科维奇
	f.日特科维奇Zhytkavichy.SetParent(f)

}
