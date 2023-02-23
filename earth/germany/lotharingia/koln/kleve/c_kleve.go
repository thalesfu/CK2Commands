package kleve

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KleveCounty interface {
	feud.County
	BAnholt安霍尔特() feud.Barony
	BEmmerich埃默里希() feud.Barony
	BGoch戈赫() feud.Barony
	BIsselburg伊瑟尔堡() feud.Barony
	BKleve克莱沃() feud.Barony
	BRees雷斯() feud.Barony
	BWesel韦塞尔() feud.Barony
}

type 克莱沃KleveCounty struct {
	feud.BaseCounty
	安霍尔特Anholt    feud.Barony
	埃默里希Emmerich  feud.Barony
	戈赫Goch        feud.Barony
	伊瑟尔堡Isselburg feud.Barony
	克莱沃Kleve      feud.Barony
	雷斯Rees        feud.Barony
	韦塞尔Wesel      feud.Barony
}

func (c *克莱沃KleveCounty) BAnholt安霍尔特() feud.Barony {
	return c.安霍尔特Anholt
}

func (c *克莱沃KleveCounty) BEmmerich埃默里希() feud.Barony {
	return c.埃默里希Emmerich
}

func (c *克莱沃KleveCounty) BGoch戈赫() feud.Barony {
	return c.戈赫Goch
}

func (c *克莱沃KleveCounty) BIsselburg伊瑟尔堡() feud.Barony {
	return c.伊瑟尔堡Isselburg
}

func (c *克莱沃KleveCounty) BKleve克莱沃() feud.Barony {
	return c.克莱沃Kleve
}

func (c *克莱沃KleveCounty) BRees雷斯() feud.Barony {
	return c.雷斯Rees
}

func (c *克莱沃KleveCounty) BWesel韦塞尔() feud.Barony {
	return c.韦塞尔Wesel
}

var CKleve克莱沃 KleveCounty = &克莱沃KleveCounty{}

func init() {
	f := CKleve克莱沃.(*克莱沃KleveCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "89",
		Title:     "kleve",
		TitleName: "克莱沃",
		TitleCode: "c_kleve",
		Baronies:  map[string]feud.Barony{},
	}

	f.安霍尔特Anholt = BAnholt安霍尔特
	f.安霍尔特Anholt.SetParent(f)

	f.埃默里希Emmerich = BEmmerich埃默里希
	f.埃默里希Emmerich.SetParent(f)

	f.戈赫Goch = BGoch戈赫
	f.戈赫Goch.SetParent(f)

	f.伊瑟尔堡Isselburg = BIsselburg伊瑟尔堡
	f.伊瑟尔堡Isselburg.SetParent(f)

	f.克莱沃Kleve = BKleve克莱沃
	f.克莱沃Kleve.SetParent(f)

	f.雷斯Rees = BRees雷斯
	f.雷斯Rees.SetParent(f)

	f.韦塞尔Wesel = BWesel韦塞尔
	f.韦塞尔Wesel.SetParent(f)

}
