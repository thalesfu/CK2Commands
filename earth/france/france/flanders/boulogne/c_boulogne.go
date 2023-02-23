package boulogne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BoulogneCounty interface {
	feud.County
	BArdres阿德尔() feud.Barony
	BBoulogne布洛涅() feud.Barony
	BFauquembergues福康贝格() feud.Barony
	BHesdin埃丹() feud.Barony
	BMontreuil蒙特勒伊() feud.Barony
	BSaintpol圣波勒() feud.Barony
	BTernouis泰努瓦斯() feud.Barony
	BTherouanne泰鲁阿讷() feud.Barony
}

type 布洛涅BoulogneCounty struct {
	feud.BaseCounty
	阿德尔Ardres          feud.Barony
	布洛涅Boulogne        feud.Barony
	福康贝格Fauquembergues feud.Barony
	埃丹Hesdin           feud.Barony
	蒙特勒伊Montreuil      feud.Barony
	圣波勒Saintpol        feud.Barony
	泰努瓦斯Ternouis       feud.Barony
	泰鲁阿讷Therouanne     feud.Barony
}

func (c *布洛涅BoulogneCounty) BArdres阿德尔() feud.Barony {
	return c.阿德尔Ardres
}

func (c *布洛涅BoulogneCounty) BBoulogne布洛涅() feud.Barony {
	return c.布洛涅Boulogne
}

func (c *布洛涅BoulogneCounty) BFauquembergues福康贝格() feud.Barony {
	return c.福康贝格Fauquembergues
}

func (c *布洛涅BoulogneCounty) BHesdin埃丹() feud.Barony {
	return c.埃丹Hesdin
}

func (c *布洛涅BoulogneCounty) BMontreuil蒙特勒伊() feud.Barony {
	return c.蒙特勒伊Montreuil
}

func (c *布洛涅BoulogneCounty) BSaintpol圣波勒() feud.Barony {
	return c.圣波勒Saintpol
}

func (c *布洛涅BoulogneCounty) BTernouis泰努瓦斯() feud.Barony {
	return c.泰努瓦斯Ternouis
}

func (c *布洛涅BoulogneCounty) BTherouanne泰鲁阿讷() feud.Barony {
	return c.泰鲁阿讷Therouanne
}

var CBoulogne布洛涅 BoulogneCounty = &布洛涅BoulogneCounty{}

func init() {
	f := CBoulogne布洛涅.(*布洛涅BoulogneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "75",
		Title:     "boulogne",
		TitleName: "布洛涅",
		TitleCode: "c_boulogne",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德尔Ardres = BArdres阿德尔
	f.阿德尔Ardres.SetParent(f)

	f.布洛涅Boulogne = BBoulogne布洛涅
	f.布洛涅Boulogne.SetParent(f)

	f.福康贝格Fauquembergues = BFauquembergues福康贝格
	f.福康贝格Fauquembergues.SetParent(f)

	f.埃丹Hesdin = BHesdin埃丹
	f.埃丹Hesdin.SetParent(f)

	f.蒙特勒伊Montreuil = BMontreuil蒙特勒伊
	f.蒙特勒伊Montreuil.SetParent(f)

	f.圣波勒Saintpol = BSaintpol圣波勒
	f.圣波勒Saintpol.SetParent(f)

	f.泰努瓦斯Ternouis = BTernouis泰努瓦斯
	f.泰努瓦斯Ternouis.SetParent(f)

	f.泰鲁阿讷Therouanne = BTherouanne泰鲁阿讷
	f.泰鲁阿讷Therouanne.SetParent(f)

}
