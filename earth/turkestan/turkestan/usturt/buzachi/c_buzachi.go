package buzachi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BuzachiCounty interface {
	feud.County
	BBuzachi布扎奇() feud.Barony
	BDurneva杜尔涅瓦() feud.Barony
	BKarakishu卡拉基丘() feud.Barony
	BTyuleniy秋列尼() feud.Barony
}

type 布扎奇BuzachiCounty struct {
	feud.BaseCounty
	布扎奇Buzachi    feud.Barony
	杜尔涅瓦Durneva   feud.Barony
	卡拉基丘Karakishu feud.Barony
	秋列尼Tyuleniy   feud.Barony
}

func (c *布扎奇BuzachiCounty) BBuzachi布扎奇() feud.Barony {
	return c.布扎奇Buzachi
}

func (c *布扎奇BuzachiCounty) BDurneva杜尔涅瓦() feud.Barony {
	return c.杜尔涅瓦Durneva
}

func (c *布扎奇BuzachiCounty) BKarakishu卡拉基丘() feud.Barony {
	return c.卡拉基丘Karakishu
}

func (c *布扎奇BuzachiCounty) BTyuleniy秋列尼() feud.Barony {
	return c.秋列尼Tyuleniy
}

var CBuzachi布扎奇 BuzachiCounty = &布扎奇BuzachiCounty{}

func init() {
	f := CBuzachi布扎奇.(*布扎奇BuzachiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1787",
		Title:     "buzachi",
		TitleName: "布扎奇",
		TitleCode: "c_buzachi",
		Baronies:  map[string]feud.Barony{},
	}

	f.布扎奇Buzachi = BBuzachi布扎奇
	f.布扎奇Buzachi.SetParent(f)

	f.杜尔涅瓦Durneva = BDurneva杜尔涅瓦
	f.杜尔涅瓦Durneva.SetParent(f)

	f.卡拉基丘Karakishu = BKarakishu卡拉基丘
	f.卡拉基丘Karakishu.SetParent(f)

	f.秋列尼Tyuleniy = BTyuleniy秋列尼
	f.秋列尼Tyuleniy.SetParent(f)

}
