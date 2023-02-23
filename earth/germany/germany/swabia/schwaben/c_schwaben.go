package schwaben

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SchwabenCounty interface {
	feud.County
	BFriedrichshafen腓特烈港() feud.Barony
	BHeiligenberg海利根贝格() feud.Barony
	BHohenberg霍恩贝格() feud.Barony
	BKonstanz康斯坦茨() feud.Barony
	BLindau林道() feud.Barony
	BTubingen蒂宾根() feud.Barony
	BUberlingen于伯林根() feud.Barony
	BWangen旺根() feud.Barony
}

type 施瓦本SchwabenCounty struct {
	feud.BaseCounty
	腓特烈港Friedrichshafen feud.Barony
	海利根贝格Heiligenberg   feud.Barony
	霍恩贝格Hohenberg       feud.Barony
	康斯坦茨Konstanz        feud.Barony
	林道Lindau            feud.Barony
	蒂宾根Tubingen         feud.Barony
	于伯林根Uberlingen      feud.Barony
	旺根Wangen            feud.Barony
}

func (c *施瓦本SchwabenCounty) BFriedrichshafen腓特烈港() feud.Barony {
	return c.腓特烈港Friedrichshafen
}

func (c *施瓦本SchwabenCounty) BHeiligenberg海利根贝格() feud.Barony {
	return c.海利根贝格Heiligenberg
}

func (c *施瓦本SchwabenCounty) BHohenberg霍恩贝格() feud.Barony {
	return c.霍恩贝格Hohenberg
}

func (c *施瓦本SchwabenCounty) BKonstanz康斯坦茨() feud.Barony {
	return c.康斯坦茨Konstanz
}

func (c *施瓦本SchwabenCounty) BLindau林道() feud.Barony {
	return c.林道Lindau
}

func (c *施瓦本SchwabenCounty) BTubingen蒂宾根() feud.Barony {
	return c.蒂宾根Tubingen
}

func (c *施瓦本SchwabenCounty) BUberlingen于伯林根() feud.Barony {
	return c.于伯林根Uberlingen
}

func (c *施瓦本SchwabenCounty) BWangen旺根() feud.Barony {
	return c.旺根Wangen
}

var CSchwaben施瓦本 SchwabenCounty = &施瓦本SchwabenCounty{}

func init() {
	f := CSchwaben施瓦本.(*施瓦本SchwabenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "249",
		Title:     "schwaben",
		TitleName: "施瓦本",
		TitleCode: "c_schwaben",
		Baronies:  map[string]feud.Barony{},
	}

	f.腓特烈港Friedrichshafen = BFriedrichshafen腓特烈港
	f.腓特烈港Friedrichshafen.SetParent(f)

	f.海利根贝格Heiligenberg = BHeiligenberg海利根贝格
	f.海利根贝格Heiligenberg.SetParent(f)

	f.霍恩贝格Hohenberg = BHohenberg霍恩贝格
	f.霍恩贝格Hohenberg.SetParent(f)

	f.康斯坦茨Konstanz = BKonstanz康斯坦茨
	f.康斯坦茨Konstanz.SetParent(f)

	f.林道Lindau = BLindau林道
	f.林道Lindau.SetParent(f)

	f.蒂宾根Tubingen = BTubingen蒂宾根
	f.蒂宾根Tubingen.SetParent(f)

	f.于伯林根Uberlingen = BUberlingen于伯林根
	f.于伯林根Uberlingen.SetParent(f)

	f.旺根Wangen = BWangen旺根
	f.旺根Wangen.SetParent(f)

}
