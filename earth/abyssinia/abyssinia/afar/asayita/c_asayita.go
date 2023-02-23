package asayita

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AsayitaCounty interface {
	feud.County
	BAfambo阿凡博() feud.Barony
	BAsayita阿塞塔() feud.Barony
	BDedai德代() feud.Barony
	BJewaha吉瓦哈() feud.Barony
	BLofefle洛菲弗勒() feud.Barony
	BSerdo瑟尔多() feud.Barony
	BTendaho滕达霍() feud.Barony
}

type 阿塞塔AsayitaCounty struct {
	feud.BaseCounty
	阿凡博Afambo   feud.Barony
	阿塞塔Asayita  feud.Barony
	德代Dedai     feud.Barony
	吉瓦哈Jewaha   feud.Barony
	洛菲弗勒Lofefle feud.Barony
	瑟尔多Serdo    feud.Barony
	滕达霍Tendaho  feud.Barony
}

func (c *阿塞塔AsayitaCounty) BAfambo阿凡博() feud.Barony {
	return c.阿凡博Afambo
}

func (c *阿塞塔AsayitaCounty) BAsayita阿塞塔() feud.Barony {
	return c.阿塞塔Asayita
}

func (c *阿塞塔AsayitaCounty) BDedai德代() feud.Barony {
	return c.德代Dedai
}

func (c *阿塞塔AsayitaCounty) BJewaha吉瓦哈() feud.Barony {
	return c.吉瓦哈Jewaha
}

func (c *阿塞塔AsayitaCounty) BLofefle洛菲弗勒() feud.Barony {
	return c.洛菲弗勒Lofefle
}

func (c *阿塞塔AsayitaCounty) BSerdo瑟尔多() feud.Barony {
	return c.瑟尔多Serdo
}

func (c *阿塞塔AsayitaCounty) BTendaho滕达霍() feud.Barony {
	return c.滕达霍Tendaho
}

var CAsayita阿塞塔 AsayitaCounty = &阿塞塔AsayitaCounty{}

func init() {
	f := CAsayita阿塞塔.(*阿塞塔AsayitaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1379",
		Title:     "asayita",
		TitleName: "阿塞塔",
		TitleCode: "c_asayita",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿凡博Afambo = BAfambo阿凡博
	f.阿凡博Afambo.SetParent(f)

	f.阿塞塔Asayita = BAsayita阿塞塔
	f.阿塞塔Asayita.SetParent(f)

	f.德代Dedai = BDedai德代
	f.德代Dedai.SetParent(f)

	f.吉瓦哈Jewaha = BJewaha吉瓦哈
	f.吉瓦哈Jewaha.SetParent(f)

	f.洛菲弗勒Lofefle = BLofefle洛菲弗勒
	f.洛菲弗勒Lofefle.SetParent(f)

	f.瑟尔多Serdo = BSerdo瑟尔多
	f.瑟尔多Serdo.SetParent(f)

	f.滕达霍Tendaho = BTendaho滕达霍
	f.滕达霍Tendaho.SetParent(f)

}
