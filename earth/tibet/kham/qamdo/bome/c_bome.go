package bome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BomeCounty interface {
	feud.County
	BAlamdo阿兰多() feud.Barony
	BBome波密() feud.Barony
	BCizula次足拉() feud.Barony
	BNaha那哈() feud.Barony
	BYiong易翁() feud.Barony
	BYupu玉普() feud.Barony
	BZhamo扎木() feud.Barony
}

type 波密BomeCounty struct {
	feud.BaseCounty
	阿兰多Alamdo feud.Barony
	波密Bome    feud.Barony
	次足拉Cizula feud.Barony
	那哈Naha    feud.Barony
	易翁Yiong   feud.Barony
	玉普Yupu    feud.Barony
	扎木Zhamo   feud.Barony
}

func (c *波密BomeCounty) BAlamdo阿兰多() feud.Barony {
	return c.阿兰多Alamdo
}

func (c *波密BomeCounty) BBome波密() feud.Barony {
	return c.波密Bome
}

func (c *波密BomeCounty) BCizula次足拉() feud.Barony {
	return c.次足拉Cizula
}

func (c *波密BomeCounty) BNaha那哈() feud.Barony {
	return c.那哈Naha
}

func (c *波密BomeCounty) BYiong易翁() feud.Barony {
	return c.易翁Yiong
}

func (c *波密BomeCounty) BYupu玉普() feud.Barony {
	return c.玉普Yupu
}

func (c *波密BomeCounty) BZhamo扎木() feud.Barony {
	return c.扎木Zhamo
}

var CBome波密 BomeCounty = &波密BomeCounty{}

func init() {
	f := CBome波密.(*波密BomeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1553",
		Title:     "bome",
		TitleName: "波密",
		TitleCode: "c_bome",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿兰多Alamdo = BAlamdo阿兰多
	f.阿兰多Alamdo.SetParent(f)

	f.波密Bome = BBome波密
	f.波密Bome.SetParent(f)

	f.次足拉Cizula = BCizula次足拉
	f.次足拉Cizula.SetParent(f)

	f.那哈Naha = BNaha那哈
	f.那哈Naha.SetParent(f)

	f.易翁Yiong = BYiong易翁
	f.易翁Yiong.SetParent(f)

	f.玉普Yupu = BYupu玉普
	f.玉普Yupu.SetParent(f)

	f.扎木Zhamo = BZhamo扎木
	f.扎木Zhamo.SetParent(f)

}
