package piombino

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PiombinoCounty interface {
	feud.County
	BCampiglia坎皮利亚() feud.Barony
	BFollonica福洛尼卡() feud.Barony
	BPiombino皮翁比诺() feud.Barony
	BPopulonia波普洛尼亚() feud.Barony
	BRadicofani拉迪科法尼() feud.Barony
	BSanvincenzo圣文琴佐() feud.Barony
	BSuvereto苏韦雷托() feud.Barony
}

type 皮翁比诺PiombinoCounty struct {
	feud.BaseCounty
	坎皮利亚Campiglia   feud.Barony
	福洛尼卡Follonica   feud.Barony
	皮翁比诺Piombino    feud.Barony
	波普洛尼亚Populonia  feud.Barony
	拉迪科法尼Radicofani feud.Barony
	圣文琴佐Sanvincenzo feud.Barony
	苏韦雷托Suvereto    feud.Barony
}

func (c *皮翁比诺PiombinoCounty) BCampiglia坎皮利亚() feud.Barony {
	return c.坎皮利亚Campiglia
}

func (c *皮翁比诺PiombinoCounty) BFollonica福洛尼卡() feud.Barony {
	return c.福洛尼卡Follonica
}

func (c *皮翁比诺PiombinoCounty) BPiombino皮翁比诺() feud.Barony {
	return c.皮翁比诺Piombino
}

func (c *皮翁比诺PiombinoCounty) BPopulonia波普洛尼亚() feud.Barony {
	return c.波普洛尼亚Populonia
}

func (c *皮翁比诺PiombinoCounty) BRadicofani拉迪科法尼() feud.Barony {
	return c.拉迪科法尼Radicofani
}

func (c *皮翁比诺PiombinoCounty) BSanvincenzo圣文琴佐() feud.Barony {
	return c.圣文琴佐Sanvincenzo
}

func (c *皮翁比诺PiombinoCounty) BSuvereto苏韦雷托() feud.Barony {
	return c.苏韦雷托Suvereto
}

var CPiombino皮翁比诺 PiombinoCounty = &皮翁比诺PiombinoCounty{}

func init() {
	f := CPiombino皮翁比诺.(*皮翁比诺PiombinoCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "331",
		Title:     "piombino",
		TitleName: "皮翁比诺",
		TitleCode: "c_piombino",
		Baronies:  map[string]feud.Barony{},
	}

	f.坎皮利亚Campiglia = BCampiglia坎皮利亚
	f.坎皮利亚Campiglia.SetParent(f)

	f.福洛尼卡Follonica = BFollonica福洛尼卡
	f.福洛尼卡Follonica.SetParent(f)

	f.皮翁比诺Piombino = BPiombino皮翁比诺
	f.皮翁比诺Piombino.SetParent(f)

	f.波普洛尼亚Populonia = BPopulonia波普洛尼亚
	f.波普洛尼亚Populonia.SetParent(f)

	f.拉迪科法尼Radicofani = BRadicofani拉迪科法尼
	f.拉迪科法尼Radicofani.SetParent(f)

	f.圣文琴佐Sanvincenzo = BSanvincenzo圣文琴佐
	f.圣文琴佐Sanvincenzo.SetParent(f)

	f.苏韦雷托Suvereto = BSuvereto苏韦雷托
	f.苏韦雷托Suvereto.SetParent(f)

}
