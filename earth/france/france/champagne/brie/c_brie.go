package brie

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BrieCounty interface {
	feud.County
	BChateauthierry蒂耶里堡() feud.Barony
	BMeaux莫城() feud.Barony
	BMontmirail蒙米拉伊() feud.Barony
	BProvins普罗万() feud.Barony
	BSezanne塞扎讷() feud.Barony
	BVertus韦尔蒂() feud.Barony
}

type 布里BrieCounty struct {
	feud.BaseCounty
	蒂耶里堡Chateauthierry feud.Barony
	莫城Meaux            feud.Barony
	蒙米拉伊Montmirail     feud.Barony
	普罗万Provins         feud.Barony
	塞扎讷Sezanne         feud.Barony
	韦尔蒂Vertus          feud.Barony
}

func (c *布里BrieCounty) BChateauthierry蒂耶里堡() feud.Barony {
	return c.蒂耶里堡Chateauthierry
}

func (c *布里BrieCounty) BMeaux莫城() feud.Barony {
	return c.莫城Meaux
}

func (c *布里BrieCounty) BMontmirail蒙米拉伊() feud.Barony {
	return c.蒙米拉伊Montmirail
}

func (c *布里BrieCounty) BProvins普罗万() feud.Barony {
	return c.普罗万Provins
}

func (c *布里BrieCounty) BSezanne塞扎讷() feud.Barony {
	return c.塞扎讷Sezanne
}

func (c *布里BrieCounty) BVertus韦尔蒂() feud.Barony {
	return c.韦尔蒂Vertus
}

var CBrie布里 BrieCounty = &布里BrieCounty{}

func init() {
	f := CBrie布里.(*布里BrieCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1964",
		Title:     "brie",
		TitleName: "布里",
		TitleCode: "c_brie",
		Baronies:  map[string]feud.Barony{},
	}

	f.蒂耶里堡Chateauthierry = BChateauthierry蒂耶里堡
	f.蒂耶里堡Chateauthierry.SetParent(f)

	f.莫城Meaux = BMeaux莫城
	f.莫城Meaux.SetParent(f)

	f.蒙米拉伊Montmirail = BMontmirail蒙米拉伊
	f.蒙米拉伊Montmirail.SetParent(f)

	f.普罗万Provins = BProvins普罗万
	f.普罗万Provins.SetParent(f)

	f.塞扎讷Sezanne = BSezanne塞扎讷
	f.塞扎讷Sezanne.SetParent(f)

	f.韦尔蒂Vertus = BVertus韦尔蒂
	f.韦尔蒂Vertus.SetParent(f)

}
