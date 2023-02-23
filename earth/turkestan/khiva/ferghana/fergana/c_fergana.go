package fergana

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type FerganaCounty interface {
	feud.County
	BAndijan俺的干() feud.Barony
	BAval阿瓦尔() feud.Barony
	BMarginan麻耳亦囊() feud.Barony
	BNadaq纳达克() feud.Barony
	BOsh我失() feud.Barony
	BRishton里什顿() feud.Barony
	BUzkand讹迹邗() feud.Barony
}

type 费尔干纳FerganaCounty struct {
	feud.BaseCounty
	俺的干Andijan   feud.Barony
	阿瓦尔Aval      feud.Barony
	麻耳亦囊Marginan feud.Barony
	纳达克Nadaq     feud.Barony
	我失Osh        feud.Barony
	里什顿Rishton   feud.Barony
	讹迹邗Uzkand    feud.Barony
}

func (c *费尔干纳FerganaCounty) BAndijan俺的干() feud.Barony {
	return c.俺的干Andijan
}

func (c *费尔干纳FerganaCounty) BAval阿瓦尔() feud.Barony {
	return c.阿瓦尔Aval
}

func (c *费尔干纳FerganaCounty) BMarginan麻耳亦囊() feud.Barony {
	return c.麻耳亦囊Marginan
}

func (c *费尔干纳FerganaCounty) BNadaq纳达克() feud.Barony {
	return c.纳达克Nadaq
}

func (c *费尔干纳FerganaCounty) BOsh我失() feud.Barony {
	return c.我失Osh
}

func (c *费尔干纳FerganaCounty) BRishton里什顿() feud.Barony {
	return c.里什顿Rishton
}

func (c *费尔干纳FerganaCounty) BUzkand讹迹邗() feud.Barony {
	return c.讹迹邗Uzkand
}

var CFergana费尔干纳 FerganaCounty = &费尔干纳FerganaCounty{}

func init() {
	f := CFergana费尔干纳.(*费尔干纳FerganaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1423",
		Title:     "fergana",
		TitleName: "费尔干纳",
		TitleCode: "c_fergana",
		Baronies:  map[string]feud.Barony{},
	}

	f.俺的干Andijan = BAndijan俺的干
	f.俺的干Andijan.SetParent(f)

	f.阿瓦尔Aval = BAval阿瓦尔
	f.阿瓦尔Aval.SetParent(f)

	f.麻耳亦囊Marginan = BMarginan麻耳亦囊
	f.麻耳亦囊Marginan.SetParent(f)

	f.纳达克Nadaq = BNadaq纳达克
	f.纳达克Nadaq.SetParent(f)

	f.我失Osh = BOsh我失
	f.我失Osh.SetParent(f)

	f.里什顿Rishton = BRishton里什顿
	f.里什顿Rishton.SetParent(f)

	f.讹迹邗Uzkand = BUzkand讹迹邗
	f.讹迹邗Uzkand.SetParent(f)

}
