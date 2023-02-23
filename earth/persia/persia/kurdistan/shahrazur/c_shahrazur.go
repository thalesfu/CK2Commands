package shahrazur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ShahrazurCounty interface {
	feud.County
	BAmkaleh阿姆卡勒() feud.Barony
	BEskelabad埃斯克拉巴德() feud.Barony
	BKabgan卡卜甘() feud.Barony
	BKavand卡万德() feud.Barony
	BSanda桑达() feud.Barony
	BSaqqez萨盖兹() feud.Barony
	BShahrazur沙赫里祖尔() feud.Barony
}

type 沙赫拉祖尔ShahrazurCounty struct {
	feud.BaseCounty
	阿姆卡勒Amkaleh     feud.Barony
	埃斯克拉巴德Eskelabad feud.Barony
	卡卜甘Kabgan       feud.Barony
	卡万德Kavand       feud.Barony
	桑达Sanda         feud.Barony
	萨盖兹Saqqez       feud.Barony
	沙赫里祖尔Shahrazur  feud.Barony
}

func (c *沙赫拉祖尔ShahrazurCounty) BAmkaleh阿姆卡勒() feud.Barony {
	return c.阿姆卡勒Amkaleh
}

func (c *沙赫拉祖尔ShahrazurCounty) BEskelabad埃斯克拉巴德() feud.Barony {
	return c.埃斯克拉巴德Eskelabad
}

func (c *沙赫拉祖尔ShahrazurCounty) BKabgan卡卜甘() feud.Barony {
	return c.卡卜甘Kabgan
}

func (c *沙赫拉祖尔ShahrazurCounty) BKavand卡万德() feud.Barony {
	return c.卡万德Kavand
}

func (c *沙赫拉祖尔ShahrazurCounty) BSanda桑达() feud.Barony {
	return c.桑达Sanda
}

func (c *沙赫拉祖尔ShahrazurCounty) BSaqqez萨盖兹() feud.Barony {
	return c.萨盖兹Saqqez
}

func (c *沙赫拉祖尔ShahrazurCounty) BShahrazur沙赫里祖尔() feud.Barony {
	return c.沙赫里祖尔Shahrazur
}

var CShahrazur沙赫拉祖尔 ShahrazurCounty = &沙赫拉祖尔ShahrazurCounty{}

func init() {
	f := CShahrazur沙赫拉祖尔.(*沙赫拉祖尔ShahrazurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "784",
		Title:     "shahrazur",
		TitleName: "沙赫拉祖尔",
		TitleCode: "c_shahrazur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿姆卡勒Amkaleh = BAmkaleh阿姆卡勒
	f.阿姆卡勒Amkaleh.SetParent(f)

	f.埃斯克拉巴德Eskelabad = BEskelabad埃斯克拉巴德
	f.埃斯克拉巴德Eskelabad.SetParent(f)

	f.卡卜甘Kabgan = BKabgan卡卜甘
	f.卡卜甘Kabgan.SetParent(f)

	f.卡万德Kavand = BKavand卡万德
	f.卡万德Kavand.SetParent(f)

	f.桑达Sanda = BSanda桑达
	f.桑达Sanda.SetParent(f)

	f.萨盖兹Saqqez = BSaqqez萨盖兹
	f.萨盖兹Saqqez.SetParent(f)

	f.沙赫里祖尔Shahrazur = BShahrazur沙赫里祖尔
	f.沙赫里祖尔Shahrazur.SetParent(f)

}
