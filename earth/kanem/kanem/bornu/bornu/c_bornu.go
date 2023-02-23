package bornu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BornuCounty interface {
	feud.County
	BFika菲卡() feud.Barony
	BKukawa库卡瓦() feud.Barony
	BMarzouga梅尔祖加() feud.Barony
	BNgazargamu恩加扎尔加姆() feud.Barony
	BSao萨奥() feud.Barony
	BYao亚奥() feud.Barony
	BZamtam赞坦() feud.Barony
}

type 博尔努BornuCounty struct {
	feud.BaseCounty
	菲卡Fika           feud.Barony
	库卡瓦Kukawa        feud.Barony
	梅尔祖加Marzouga     feud.Barony
	恩加扎尔加姆Ngazargamu feud.Barony
	萨奥Sao            feud.Barony
	亚奥Yao            feud.Barony
	赞坦Zamtam         feud.Barony
}

func (c *博尔努BornuCounty) BFika菲卡() feud.Barony {
	return c.菲卡Fika
}

func (c *博尔努BornuCounty) BKukawa库卡瓦() feud.Barony {
	return c.库卡瓦Kukawa
}

func (c *博尔努BornuCounty) BMarzouga梅尔祖加() feud.Barony {
	return c.梅尔祖加Marzouga
}

func (c *博尔努BornuCounty) BNgazargamu恩加扎尔加姆() feud.Barony {
	return c.恩加扎尔加姆Ngazargamu
}

func (c *博尔努BornuCounty) BSao萨奥() feud.Barony {
	return c.萨奥Sao
}

func (c *博尔努BornuCounty) BYao亚奥() feud.Barony {
	return c.亚奥Yao
}

func (c *博尔努BornuCounty) BZamtam赞坦() feud.Barony {
	return c.赞坦Zamtam
}

var CBornu博尔努 BornuCounty = &博尔努BornuCounty{}

func init() {
	f := CBornu博尔努.(*博尔努BornuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1743",
		Title:     "bornu",
		TitleName: "博尔努",
		TitleCode: "c_bornu",
		Baronies:  map[string]feud.Barony{},
	}

	f.菲卡Fika = BFika菲卡
	f.菲卡Fika.SetParent(f)

	f.库卡瓦Kukawa = BKukawa库卡瓦
	f.库卡瓦Kukawa.SetParent(f)

	f.梅尔祖加Marzouga = BMarzouga梅尔祖加
	f.梅尔祖加Marzouga.SetParent(f)

	f.恩加扎尔加姆Ngazargamu = BNgazargamu恩加扎尔加姆
	f.恩加扎尔加姆Ngazargamu.SetParent(f)

	f.萨奥Sao = BSao萨奥
	f.萨奥Sao.SetParent(f)

	f.亚奥Yao = BYao亚奥
	f.亚奥Yao.SetParent(f)

	f.赞坦Zamtam = BZamtam赞坦
	f.赞坦Zamtam.SetParent(f)

}
