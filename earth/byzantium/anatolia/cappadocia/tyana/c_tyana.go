package tyana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TyanaCounty interface {
	feud.County
	BAnatoliaheraklea赫拉克勒亚() feud.Barony
	BCybistra基比斯特拉() feud.Barony
	BFaustinopolis福斯蒂诺波利斯() feud.Barony
	BGamar伽马() feud.Barony
	BLoulon卢隆() feud.Barony
	BTomarza托马尔扎() feud.Barony
	BTyana堤亚纳() feud.Barony
}

type 堤亚纳TyanaCounty struct {
	feud.BaseCounty
	赫拉克勒亚Anatoliaheraklea feud.Barony
	基比斯特拉Cybistra         feud.Barony
	福斯蒂诺波利斯Faustinopolis  feud.Barony
	伽马Gamar               feud.Barony
	卢隆Loulon              feud.Barony
	托马尔扎Tomarza           feud.Barony
	堤亚纳Tyana              feud.Barony
}

func (c *堤亚纳TyanaCounty) BAnatoliaheraklea赫拉克勒亚() feud.Barony {
	return c.赫拉克勒亚Anatoliaheraklea
}

func (c *堤亚纳TyanaCounty) BCybistra基比斯特拉() feud.Barony {
	return c.基比斯特拉Cybistra
}

func (c *堤亚纳TyanaCounty) BFaustinopolis福斯蒂诺波利斯() feud.Barony {
	return c.福斯蒂诺波利斯Faustinopolis
}

func (c *堤亚纳TyanaCounty) BGamar伽马() feud.Barony {
	return c.伽马Gamar
}

func (c *堤亚纳TyanaCounty) BLoulon卢隆() feud.Barony {
	return c.卢隆Loulon
}

func (c *堤亚纳TyanaCounty) BTomarza托马尔扎() feud.Barony {
	return c.托马尔扎Tomarza
}

func (c *堤亚纳TyanaCounty) BTyana堤亚纳() feud.Barony {
	return c.堤亚纳Tyana
}

var CTyana堤亚纳 TyanaCounty = &堤亚纳TyanaCounty{}

func init() {
	f := CTyana堤亚纳.(*堤亚纳TyanaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "760",
		Title:     "tyana",
		TitleName: "堤亚纳",
		TitleCode: "c_tyana",
		Baronies:  map[string]feud.Barony{},
	}

	f.赫拉克勒亚Anatoliaheraklea = BAnatoliaheraklea赫拉克勒亚
	f.赫拉克勒亚Anatoliaheraklea.SetParent(f)

	f.基比斯特拉Cybistra = BCybistra基比斯特拉
	f.基比斯特拉Cybistra.SetParent(f)

	f.福斯蒂诺波利斯Faustinopolis = BFaustinopolis福斯蒂诺波利斯
	f.福斯蒂诺波利斯Faustinopolis.SetParent(f)

	f.伽马Gamar = BGamar伽马
	f.伽马Gamar.SetParent(f)

	f.卢隆Loulon = BLoulon卢隆
	f.卢隆Loulon.SetParent(f)

	f.托马尔扎Tomarza = BTomarza托马尔扎
	f.托马尔扎Tomarza.SetParent(f)

	f.堤亚纳Tyana = BTyana堤亚纳
	f.堤亚纳Tyana.SetParent(f)

}
