package volkovysk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VolkovyskCounty interface {
	feud.County
	BKlotsk克洛茨克() feud.Barony
	BLyakhavichy利亚哈维奇() feud.Barony
	BOuslonim乌斯洛尼姆() feud.Barony
	BRuzhany鲁扎内() feud.Barony
	BStudzieniki斯图泽尼采() feud.Barony
	BVolkovysk沃尔科维斯克() feud.Barony
	BZyrovichy日罗维奇() feud.Barony
}

type 沃尔科维斯克VolkovyskCounty struct {
	feud.BaseCounty
	克洛茨克Klotsk       feud.Barony
	利亚哈维奇Lyakhavichy feud.Barony
	乌斯洛尼姆Ouslonim    feud.Barony
	鲁扎内Ruzhany       feud.Barony
	斯图泽尼采Studzieniki feud.Barony
	沃尔科维斯克Volkovysk  feud.Barony
	日罗维奇Zyrovichy    feud.Barony
}

func (c *沃尔科维斯克VolkovyskCounty) BKlotsk克洛茨克() feud.Barony {
	return c.克洛茨克Klotsk
}

func (c *沃尔科维斯克VolkovyskCounty) BLyakhavichy利亚哈维奇() feud.Barony {
	return c.利亚哈维奇Lyakhavichy
}

func (c *沃尔科维斯克VolkovyskCounty) BOuslonim乌斯洛尼姆() feud.Barony {
	return c.乌斯洛尼姆Ouslonim
}

func (c *沃尔科维斯克VolkovyskCounty) BRuzhany鲁扎内() feud.Barony {
	return c.鲁扎内Ruzhany
}

func (c *沃尔科维斯克VolkovyskCounty) BStudzieniki斯图泽尼采() feud.Barony {
	return c.斯图泽尼采Studzieniki
}

func (c *沃尔科维斯克VolkovyskCounty) BVolkovysk沃尔科维斯克() feud.Barony {
	return c.沃尔科维斯克Volkovysk
}

func (c *沃尔科维斯克VolkovyskCounty) BZyrovichy日罗维奇() feud.Barony {
	return c.日罗维奇Zyrovichy
}

var CVolkovysk沃尔科维斯克 VolkovyskCounty = &沃尔科维斯克VolkovyskCounty{}

func init() {
	f := CVolkovysk沃尔科维斯克.(*沃尔科维斯克VolkovyskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1655",
		Title:     "volkovysk",
		TitleName: "沃尔科维斯克",
		TitleCode: "c_volkovysk",
		Baronies:  map[string]feud.Barony{},
	}

	f.克洛茨克Klotsk = BKlotsk克洛茨克
	f.克洛茨克Klotsk.SetParent(f)

	f.利亚哈维奇Lyakhavichy = BLyakhavichy利亚哈维奇
	f.利亚哈维奇Lyakhavichy.SetParent(f)

	f.乌斯洛尼姆Ouslonim = BOuslonim乌斯洛尼姆
	f.乌斯洛尼姆Ouslonim.SetParent(f)

	f.鲁扎内Ruzhany = BRuzhany鲁扎内
	f.鲁扎内Ruzhany.SetParent(f)

	f.斯图泽尼采Studzieniki = BStudzieniki斯图泽尼采
	f.斯图泽尼采Studzieniki.SetParent(f)

	f.沃尔科维斯克Volkovysk = BVolkovysk沃尔科维斯克
	f.沃尔科维斯克Volkovysk.SetParent(f)

	f.日罗维奇Zyrovichy = BZyrovichy日罗维奇
	f.日罗维奇Zyrovichy.SetParent(f)

}
