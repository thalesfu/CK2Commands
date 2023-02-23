package medelike

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MedelikeCounty interface {
	feud.County
	BGaming加明() feud.Barony
	BLilienfeld利林费尔德() feud.Barony
	BMelk梅尔克() feud.Barony
	BPochlarn珀希拉恩() feud.Barony
	BStpolten圣珀尔滕() feud.Barony
	BWaidhofen魏德霍芬() feud.Barony
	BYbbs伊布斯() feud.Barony
	BZelking策尔京() feud.Barony
}

type 梅德利克MedelikeCounty struct {
	feud.BaseCounty
	加明Gaming        feud.Barony
	利林费尔德Lilienfeld feud.Barony
	梅尔克Melk         feud.Barony
	珀希拉恩Pochlarn    feud.Barony
	圣珀尔滕Stpolten    feud.Barony
	魏德霍芬Waidhofen   feud.Barony
	伊布斯Ybbs         feud.Barony
	策尔京Zelking      feud.Barony
}

func (c *梅德利克MedelikeCounty) BGaming加明() feud.Barony {
	return c.加明Gaming
}

func (c *梅德利克MedelikeCounty) BLilienfeld利林费尔德() feud.Barony {
	return c.利林费尔德Lilienfeld
}

func (c *梅德利克MedelikeCounty) BMelk梅尔克() feud.Barony {
	return c.梅尔克Melk
}

func (c *梅德利克MedelikeCounty) BPochlarn珀希拉恩() feud.Barony {
	return c.珀希拉恩Pochlarn
}

func (c *梅德利克MedelikeCounty) BStpolten圣珀尔滕() feud.Barony {
	return c.圣珀尔滕Stpolten
}

func (c *梅德利克MedelikeCounty) BWaidhofen魏德霍芬() feud.Barony {
	return c.魏德霍芬Waidhofen
}

func (c *梅德利克MedelikeCounty) BYbbs伊布斯() feud.Barony {
	return c.伊布斯Ybbs
}

func (c *梅德利克MedelikeCounty) BZelking策尔京() feud.Barony {
	return c.策尔京Zelking
}

var CMedelike梅德利克 MedelikeCounty = &梅德利克MedelikeCounty{}

func init() {
	f := CMedelike梅德利克.(*梅德利克MedelikeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1699",
		Title:     "medelike",
		TitleName: "梅德利克",
		TitleCode: "c_medelike",
		Baronies:  map[string]feud.Barony{},
	}

	f.加明Gaming = BGaming加明
	f.加明Gaming.SetParent(f)

	f.利林费尔德Lilienfeld = BLilienfeld利林费尔德
	f.利林费尔德Lilienfeld.SetParent(f)

	f.梅尔克Melk = BMelk梅尔克
	f.梅尔克Melk.SetParent(f)

	f.珀希拉恩Pochlarn = BPochlarn珀希拉恩
	f.珀希拉恩Pochlarn.SetParent(f)

	f.圣珀尔滕Stpolten = BStpolten圣珀尔滕
	f.圣珀尔滕Stpolten.SetParent(f)

	f.魏德霍芬Waidhofen = BWaidhofen魏德霍芬
	f.魏德霍芬Waidhofen.SetParent(f)

	f.伊布斯Ybbs = BYbbs伊布斯
	f.伊布斯Ybbs.SetParent(f)

	f.策尔京Zelking = BZelking策尔京
	f.策尔京Zelking.SetParent(f)

}
