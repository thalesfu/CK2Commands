package manan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MananCounty interface {
	feud.County
	BArkolo阿尔科洛() feud.Barony
	BFurtu富尔图() feud.Barony
	BGaroumele加鲁梅勒() feud.Barony
	BGoyawa戈亚瓦() feud.Barony
	BManan马南() feud.Barony
	BNankassa南卡萨() feud.Barony
	BTchoukoutoli楚库托利() feud.Barony
}

type 马南MananCounty struct {
	feud.BaseCounty
	阿尔科洛Arkolo       feud.Barony
	富尔图Furtu         feud.Barony
	加鲁梅勒Garoumele    feud.Barony
	戈亚瓦Goyawa        feud.Barony
	马南Manan          feud.Barony
	南卡萨Nankassa      feud.Barony
	楚库托利Tchoukoutoli feud.Barony
}

func (c *马南MananCounty) BArkolo阿尔科洛() feud.Barony {
	return c.阿尔科洛Arkolo
}

func (c *马南MananCounty) BFurtu富尔图() feud.Barony {
	return c.富尔图Furtu
}

func (c *马南MananCounty) BGaroumele加鲁梅勒() feud.Barony {
	return c.加鲁梅勒Garoumele
}

func (c *马南MananCounty) BGoyawa戈亚瓦() feud.Barony {
	return c.戈亚瓦Goyawa
}

func (c *马南MananCounty) BManan马南() feud.Barony {
	return c.马南Manan
}

func (c *马南MananCounty) BNankassa南卡萨() feud.Barony {
	return c.南卡萨Nankassa
}

func (c *马南MananCounty) BTchoukoutoli楚库托利() feud.Barony {
	return c.楚库托利Tchoukoutoli
}

var CManan马南 MananCounty = &马南MananCounty{}

func init() {
	f := CManan马南.(*马南MananCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1736",
		Title:     "manan",
		TitleName: "马南",
		TitleCode: "c_manan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿尔科洛Arkolo = BArkolo阿尔科洛
	f.阿尔科洛Arkolo.SetParent(f)

	f.富尔图Furtu = BFurtu富尔图
	f.富尔图Furtu.SetParent(f)

	f.加鲁梅勒Garoumele = BGaroumele加鲁梅勒
	f.加鲁梅勒Garoumele.SetParent(f)

	f.戈亚瓦Goyawa = BGoyawa戈亚瓦
	f.戈亚瓦Goyawa.SetParent(f)

	f.马南Manan = BManan马南
	f.马南Manan.SetParent(f)

	f.南卡萨Nankassa = BNankassa南卡萨
	f.南卡萨Nankassa.SetParent(f)

	f.楚库托利Tchoukoutoli = BTchoukoutoli楚库托利
	f.楚库托利Tchoukoutoli.SetParent(f)

}
