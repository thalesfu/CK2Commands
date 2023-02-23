package rafha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RafhaCounty interface {
	feud.County
	BAljumaymah朱迈马() feud.Barony
	BAshshir希尔() feud.Barony
	BDuwayd杜韦德() feud.Barony
	BLawqah劳盖() feud.Barony
	BMarkuz马库兹() feud.Barony
	BRafha拉夫哈() feud.Barony
	BTimiat提米特() feud.Barony
	BUwayqilah欧韦吉莱() feud.Barony
}

type 拉夫哈RafhaCounty struct {
	feud.BaseCounty
	朱迈马Aljumaymah feud.Barony
	希尔Ashshir     feud.Barony
	杜韦德Duwayd     feud.Barony
	劳盖Lawqah      feud.Barony
	马库兹Markuz     feud.Barony
	拉夫哈Rafha      feud.Barony
	提米特Timiat     feud.Barony
	欧韦吉莱Uwayqilah feud.Barony
}

func (c *拉夫哈RafhaCounty) BAljumaymah朱迈马() feud.Barony {
	return c.朱迈马Aljumaymah
}

func (c *拉夫哈RafhaCounty) BAshshir希尔() feud.Barony {
	return c.希尔Ashshir
}

func (c *拉夫哈RafhaCounty) BDuwayd杜韦德() feud.Barony {
	return c.杜韦德Duwayd
}

func (c *拉夫哈RafhaCounty) BLawqah劳盖() feud.Barony {
	return c.劳盖Lawqah
}

func (c *拉夫哈RafhaCounty) BMarkuz马库兹() feud.Barony {
	return c.马库兹Markuz
}

func (c *拉夫哈RafhaCounty) BRafha拉夫哈() feud.Barony {
	return c.拉夫哈Rafha
}

func (c *拉夫哈RafhaCounty) BTimiat提米特() feud.Barony {
	return c.提米特Timiat
}

func (c *拉夫哈RafhaCounty) BUwayqilah欧韦吉莱() feud.Barony {
	return c.欧韦吉莱Uwayqilah
}

var CRafha拉夫哈 RafhaCounty = &拉夫哈RafhaCounty{}

func init() {
	f := CRafha拉夫哈.(*拉夫哈RafhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "865",
		Title:     "rafha",
		TitleName: "拉夫哈",
		TitleCode: "c_rafha",
		Baronies:  map[string]feud.Barony{},
	}

	f.朱迈马Aljumaymah = BAljumaymah朱迈马
	f.朱迈马Aljumaymah.SetParent(f)

	f.希尔Ashshir = BAshshir希尔
	f.希尔Ashshir.SetParent(f)

	f.杜韦德Duwayd = BDuwayd杜韦德
	f.杜韦德Duwayd.SetParent(f)

	f.劳盖Lawqah = BLawqah劳盖
	f.劳盖Lawqah.SetParent(f)

	f.马库兹Markuz = BMarkuz马库兹
	f.马库兹Markuz.SetParent(f)

	f.拉夫哈Rafha = BRafha拉夫哈
	f.拉夫哈Rafha.SetParent(f)

	f.提米特Timiat = BTimiat提米特
	f.提米特Timiat.SetParent(f)

	f.欧韦吉莱Uwayqilah = BUwayqilah欧韦吉莱
	f.欧韦吉莱Uwayqilah.SetParent(f)

}
