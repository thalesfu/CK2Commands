package rajrappa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RajrappaCounty interface {
	feud.County
	BDhaurpur豆富() feud.Barony
	BDigapahandi提迦波汉提() feud.Barony
	BKajgaon迦阇伽罗摩() feud.Barony
	BKanmukal乾牟伽() feud.Barony
	BKenduli艺那头梨() feud.Barony
	BMahobkanth摩呼乾() feud.Barony
	BRajrappa罗奢罗波() feud.Barony
}

type 罗奢罗波RajrappaCounty struct {
	feud.BaseCounty
	豆富Dhaurpur       feud.Barony
	提迦波汉提Digapahandi feud.Barony
	迦阇伽罗摩Kajgaon     feud.Barony
	乾牟伽Kanmukal      feud.Barony
	艺那头梨Kenduli      feud.Barony
	摩呼乾Mahobkanth    feud.Barony
	罗奢罗波Rajrappa     feud.Barony
}

func (c *罗奢罗波RajrappaCounty) BDhaurpur豆富() feud.Barony {
	return c.豆富Dhaurpur
}

func (c *罗奢罗波RajrappaCounty) BDigapahandi提迦波汉提() feud.Barony {
	return c.提迦波汉提Digapahandi
}

func (c *罗奢罗波RajrappaCounty) BKajgaon迦阇伽罗摩() feud.Barony {
	return c.迦阇伽罗摩Kajgaon
}

func (c *罗奢罗波RajrappaCounty) BKanmukal乾牟伽() feud.Barony {
	return c.乾牟伽Kanmukal
}

func (c *罗奢罗波RajrappaCounty) BKenduli艺那头梨() feud.Barony {
	return c.艺那头梨Kenduli
}

func (c *罗奢罗波RajrappaCounty) BMahobkanth摩呼乾() feud.Barony {
	return c.摩呼乾Mahobkanth
}

func (c *罗奢罗波RajrappaCounty) BRajrappa罗奢罗波() feud.Barony {
	return c.罗奢罗波Rajrappa
}

var CRajrappa罗奢罗波 RajrappaCounty = &罗奢罗波RajrappaCounty{}

func init() {
	f := CRajrappa罗奢罗波.(*罗奢罗波RajrappaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1237",
		Title:     "rajrappa",
		TitleName: "罗奢罗波",
		TitleCode: "c_rajrappa",
		Baronies:  map[string]feud.Barony{},
	}

	f.豆富Dhaurpur = BDhaurpur豆富
	f.豆富Dhaurpur.SetParent(f)

	f.提迦波汉提Digapahandi = BDigapahandi提迦波汉提
	f.提迦波汉提Digapahandi.SetParent(f)

	f.迦阇伽罗摩Kajgaon = BKajgaon迦阇伽罗摩
	f.迦阇伽罗摩Kajgaon.SetParent(f)

	f.乾牟伽Kanmukal = BKanmukal乾牟伽
	f.乾牟伽Kanmukal.SetParent(f)

	f.艺那头梨Kenduli = BKenduli艺那头梨
	f.艺那头梨Kenduli.SetParent(f)

	f.摩呼乾Mahobkanth = BMahobkanth摩呼乾
	f.摩呼乾Mahobkanth.SetParent(f)

	f.罗奢罗波Rajrappa = BRajrappa罗奢罗波
	f.罗奢罗波Rajrappa.SetParent(f)

}
