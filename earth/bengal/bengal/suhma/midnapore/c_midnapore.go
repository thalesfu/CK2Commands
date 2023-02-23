package midnapore

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MidnaporeCounty interface {
	feud.County
	BDantan檀檀() feud.Barony
	BHilji醯梨恃() feud.Barony
	BKashipur迦尸城() feud.Barony
	BMidnapore迷地尼城() feud.Barony
	BMoghalmari牟迦罗摩利() feud.Barony
	BNimon尼牟() feud.Barony
	BParodi波卢提() feud.Barony
}

type 迷地尼城MidnaporeCounty struct {
	feud.BaseCounty
	檀檀Dantan        feud.Barony
	醯梨恃Hilji        feud.Barony
	迦尸城Kashipur     feud.Barony
	迷地尼城Midnapore   feud.Barony
	牟迦罗摩利Moghalmari feud.Barony
	尼牟Nimon         feud.Barony
	波卢提Parodi       feud.Barony
}

func (c *迷地尼城MidnaporeCounty) BDantan檀檀() feud.Barony {
	return c.檀檀Dantan
}

func (c *迷地尼城MidnaporeCounty) BHilji醯梨恃() feud.Barony {
	return c.醯梨恃Hilji
}

func (c *迷地尼城MidnaporeCounty) BKashipur迦尸城() feud.Barony {
	return c.迦尸城Kashipur
}

func (c *迷地尼城MidnaporeCounty) BMidnapore迷地尼城() feud.Barony {
	return c.迷地尼城Midnapore
}

func (c *迷地尼城MidnaporeCounty) BMoghalmari牟迦罗摩利() feud.Barony {
	return c.牟迦罗摩利Moghalmari
}

func (c *迷地尼城MidnaporeCounty) BNimon尼牟() feud.Barony {
	return c.尼牟Nimon
}

func (c *迷地尼城MidnaporeCounty) BParodi波卢提() feud.Barony {
	return c.波卢提Parodi
}

var CMidnapore迷地尼城 MidnaporeCounty = &迷地尼城MidnaporeCounty{}

func init() {
	f := CMidnapore迷地尼城.(*迷地尼城MidnaporeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1232",
		Title:     "midnapore",
		TitleName: "迷地尼城",
		TitleCode: "c_midnapore",
		Baronies:  map[string]feud.Barony{},
	}

	f.檀檀Dantan = BDantan檀檀
	f.檀檀Dantan.SetParent(f)

	f.醯梨恃Hilji = BHilji醯梨恃
	f.醯梨恃Hilji.SetParent(f)

	f.迦尸城Kashipur = BKashipur迦尸城
	f.迦尸城Kashipur.SetParent(f)

	f.迷地尼城Midnapore = BMidnapore迷地尼城
	f.迷地尼城Midnapore.SetParent(f)

	f.牟迦罗摩利Moghalmari = BMoghalmari牟迦罗摩利
	f.牟迦罗摩利Moghalmari.SetParent(f)

	f.尼牟Nimon = BNimon尼牟
	f.尼牟Nimon.SetParent(f)

	f.波卢提Parodi = BParodi波卢提
	f.波卢提Parodi.SetParent(f)

}
