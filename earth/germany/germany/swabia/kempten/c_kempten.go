package kempten

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KemptenCounty interface {
	feud.County
	BAugsburg奥格斯堡() feud.Barony
	BDonauworth多瑙沃特() feud.Barony
	BFriedberg弗里德贝格() feud.Barony
	BHochstadt赫希施塔特() feud.Barony
	BKaufbeuren考夫博伊伦() feud.Barony
	BKempten肯普滕() feud.Barony
	BRain赖恩() feud.Barony
	BZumarshausen楚斯马斯豪森() feud.Barony
}

type 奥格斯堡KemptenCounty struct {
	feud.BaseCounty
	奥格斯堡Augsburg       feud.Barony
	多瑙沃特Donauworth     feud.Barony
	弗里德贝格Friedberg     feud.Barony
	赫希施塔特Hochstadt     feud.Barony
	考夫博伊伦Kaufbeuren    feud.Barony
	肯普滕Kempten         feud.Barony
	赖恩Rain             feud.Barony
	楚斯马斯豪森Zumarshausen feud.Barony
}

func (c *奥格斯堡KemptenCounty) BAugsburg奥格斯堡() feud.Barony {
	return c.奥格斯堡Augsburg
}

func (c *奥格斯堡KemptenCounty) BDonauworth多瑙沃特() feud.Barony {
	return c.多瑙沃特Donauworth
}

func (c *奥格斯堡KemptenCounty) BFriedberg弗里德贝格() feud.Barony {
	return c.弗里德贝格Friedberg
}

func (c *奥格斯堡KemptenCounty) BHochstadt赫希施塔特() feud.Barony {
	return c.赫希施塔特Hochstadt
}

func (c *奥格斯堡KemptenCounty) BKaufbeuren考夫博伊伦() feud.Barony {
	return c.考夫博伊伦Kaufbeuren
}

func (c *奥格斯堡KemptenCounty) BKempten肯普滕() feud.Barony {
	return c.肯普滕Kempten
}

func (c *奥格斯堡KemptenCounty) BRain赖恩() feud.Barony {
	return c.赖恩Rain
}

func (c *奥格斯堡KemptenCounty) BZumarshausen楚斯马斯豪森() feud.Barony {
	return c.楚斯马斯豪森Zumarshausen
}

var CKempten奥格斯堡 KemptenCounty = &奥格斯堡KemptenCounty{}

func init() {
	f := CKempten奥格斯堡.(*奥格斯堡KemptenCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "315",
		Title:     "kempten",
		TitleName: "奥格斯堡",
		TitleCode: "c_kempten",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥格斯堡Augsburg = BAugsburg奥格斯堡
	f.奥格斯堡Augsburg.SetParent(f)

	f.多瑙沃特Donauworth = BDonauworth多瑙沃特
	f.多瑙沃特Donauworth.SetParent(f)

	f.弗里德贝格Friedberg = BFriedberg弗里德贝格
	f.弗里德贝格Friedberg.SetParent(f)

	f.赫希施塔特Hochstadt = BHochstadt赫希施塔特
	f.赫希施塔特Hochstadt.SetParent(f)

	f.考夫博伊伦Kaufbeuren = BKaufbeuren考夫博伊伦
	f.考夫博伊伦Kaufbeuren.SetParent(f)

	f.肯普滕Kempten = BKempten肯普滕
	f.肯普滕Kempten.SetParent(f)

	f.赖恩Rain = BRain赖恩
	f.赖恩Rain.SetParent(f)

	f.楚斯马斯豪森Zumarshausen = BZumarshausen楚斯马斯豪森
	f.楚斯马斯豪森Zumarshausen.SetParent(f)

}
