package teviotdale

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TeviotdaleCounty interface {
	feud.County
	BEdnam埃德纳姆() feud.Barony
	BJedburgh杰德堡() feud.Barony
	BKelso凯尔索() feud.Barony
	BMaxwell马克斯韦尔() feud.Barony
	BMelrose梅尔罗斯() feud.Barony
	BPeebles安嫩代尔() feud.Barony
	BRoxburgh罗克斯堡() feud.Barony
	BSelkirk塞尔扣克() feud.Barony
}

type 蒂维厄特代尔TeviotdaleCounty struct {
	feud.BaseCounty
	埃德纳姆Ednam    feud.Barony
	杰德堡Jedburgh  feud.Barony
	凯尔索Kelso     feud.Barony
	马克斯韦尔Maxwell feud.Barony
	梅尔罗斯Melrose  feud.Barony
	安嫩代尔Peebles  feud.Barony
	罗克斯堡Roxburgh feud.Barony
	塞尔扣克Selkirk  feud.Barony
}

func (c *蒂维厄特代尔TeviotdaleCounty) BEdnam埃德纳姆() feud.Barony {
	return c.埃德纳姆Ednam
}

func (c *蒂维厄特代尔TeviotdaleCounty) BJedburgh杰德堡() feud.Barony {
	return c.杰德堡Jedburgh
}

func (c *蒂维厄特代尔TeviotdaleCounty) BKelso凯尔索() feud.Barony {
	return c.凯尔索Kelso
}

func (c *蒂维厄特代尔TeviotdaleCounty) BMaxwell马克斯韦尔() feud.Barony {
	return c.马克斯韦尔Maxwell
}

func (c *蒂维厄特代尔TeviotdaleCounty) BMelrose梅尔罗斯() feud.Barony {
	return c.梅尔罗斯Melrose
}

func (c *蒂维厄特代尔TeviotdaleCounty) BPeebles安嫩代尔() feud.Barony {
	return c.安嫩代尔Peebles
}

func (c *蒂维厄特代尔TeviotdaleCounty) BRoxburgh罗克斯堡() feud.Barony {
	return c.罗克斯堡Roxburgh
}

func (c *蒂维厄特代尔TeviotdaleCounty) BSelkirk塞尔扣克() feud.Barony {
	return c.塞尔扣克Selkirk
}

var CTeviotdale蒂维厄特代尔 TeviotdaleCounty = &蒂维厄特代尔TeviotdaleCounty{}

func init() {
	f := CTeviotdale蒂维厄特代尔.(*蒂维厄特代尔TeviotdaleCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "38",
		Title:     "teviotdale",
		TitleName: "蒂维厄特代尔",
		TitleCode: "c_teviotdale",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃德纳姆Ednam = BEdnam埃德纳姆
	f.埃德纳姆Ednam.SetParent(f)

	f.杰德堡Jedburgh = BJedburgh杰德堡
	f.杰德堡Jedburgh.SetParent(f)

	f.凯尔索Kelso = BKelso凯尔索
	f.凯尔索Kelso.SetParent(f)

	f.马克斯韦尔Maxwell = BMaxwell马克斯韦尔
	f.马克斯韦尔Maxwell.SetParent(f)

	f.梅尔罗斯Melrose = BMelrose梅尔罗斯
	f.梅尔罗斯Melrose.SetParent(f)

	f.安嫩代尔Peebles = BPeebles安嫩代尔
	f.安嫩代尔Peebles.SetParent(f)

	f.罗克斯堡Roxburgh = BRoxburgh罗克斯堡
	f.罗克斯堡Roxburgh.SetParent(f)

	f.塞尔扣克Selkirk = BSelkirk塞尔扣克
	f.塞尔扣克Selkirk.SetParent(f)

}
