package vukovar

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VukovarCounty interface {
	feud.County
	BCibalae基巴莱() feud.Barony
	BCornacum科尔纳库姆() feud.Barony
	BIlko伊尔科() feud.Barony
	BPosga波斯加() feud.Barony
	BPosgawar波斯加瓦尔() feud.Barony
	BValka瓦尔卡() feud.Barony
	BVukovar武科瓦尔() feud.Barony
}

type 武科瓦尔VukovarCounty struct {
	feud.BaseCounty
	基巴莱Cibalae    feud.Barony
	科尔纳库姆Cornacum feud.Barony
	伊尔科Ilko       feud.Barony
	波斯加Posga      feud.Barony
	波斯加瓦尔Posgawar feud.Barony
	瓦尔卡Valka      feud.Barony
	武科瓦尔Vukovar   feud.Barony
}

func (c *武科瓦尔VukovarCounty) BCibalae基巴莱() feud.Barony {
	return c.基巴莱Cibalae
}

func (c *武科瓦尔VukovarCounty) BCornacum科尔纳库姆() feud.Barony {
	return c.科尔纳库姆Cornacum
}

func (c *武科瓦尔VukovarCounty) BIlko伊尔科() feud.Barony {
	return c.伊尔科Ilko
}

func (c *武科瓦尔VukovarCounty) BPosga波斯加() feud.Barony {
	return c.波斯加Posga
}

func (c *武科瓦尔VukovarCounty) BPosgawar波斯加瓦尔() feud.Barony {
	return c.波斯加瓦尔Posgawar
}

func (c *武科瓦尔VukovarCounty) BValka瓦尔卡() feud.Barony {
	return c.瓦尔卡Valka
}

func (c *武科瓦尔VukovarCounty) BVukovar武科瓦尔() feud.Barony {
	return c.武科瓦尔Vukovar
}

var CVukovar武科瓦尔 VukovarCounty = &武科瓦尔VukovarCounty{}

func init() {
	f := CVukovar武科瓦尔.(*武科瓦尔VukovarCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1969",
		Title:     "vukovar",
		TitleName: "武科瓦尔",
		TitleCode: "c_vukovar",
		Baronies:  map[string]feud.Barony{},
	}

	f.基巴莱Cibalae = BCibalae基巴莱
	f.基巴莱Cibalae.SetParent(f)

	f.科尔纳库姆Cornacum = BCornacum科尔纳库姆
	f.科尔纳库姆Cornacum.SetParent(f)

	f.伊尔科Ilko = BIlko伊尔科
	f.伊尔科Ilko.SetParent(f)

	f.波斯加Posga = BPosga波斯加
	f.波斯加Posga.SetParent(f)

	f.波斯加瓦尔Posgawar = BPosgawar波斯加瓦尔
	f.波斯加瓦尔Posgawar.SetParent(f)

	f.瓦尔卡Valka = BValka瓦尔卡
	f.瓦尔卡Valka.SetParent(f)

	f.武科瓦尔Vukovar = BVukovar武科瓦尔
	f.武科瓦尔Vukovar.SetParent(f)

}
