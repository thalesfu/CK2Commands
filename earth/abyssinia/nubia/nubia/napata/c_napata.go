package napata

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NapataCounty interface {
	feud.County
	BGhazali加扎利() feud.Barony
	BKanisah凯尼塞() feud.Barony
	BKarima凯里迈() feud.Barony
	BKeheilli科赫伊利() feud.Barony
	BKorti库尔提() feud.Barony
	BMarawi麦罗维() feud.Barony
	BNapata纳帕塔() feud.Barony
}

type 纳帕塔NapataCounty struct {
	feud.BaseCounty
	加扎利Ghazali   feud.Barony
	凯尼塞Kanisah   feud.Barony
	凯里迈Karima    feud.Barony
	科赫伊利Keheilli feud.Barony
	库尔提Korti     feud.Barony
	麦罗维Marawi    feud.Barony
	纳帕塔Napata    feud.Barony
}

func (c *纳帕塔NapataCounty) BGhazali加扎利() feud.Barony {
	return c.加扎利Ghazali
}

func (c *纳帕塔NapataCounty) BKanisah凯尼塞() feud.Barony {
	return c.凯尼塞Kanisah
}

func (c *纳帕塔NapataCounty) BKarima凯里迈() feud.Barony {
	return c.凯里迈Karima
}

func (c *纳帕塔NapataCounty) BKeheilli科赫伊利() feud.Barony {
	return c.科赫伊利Keheilli
}

func (c *纳帕塔NapataCounty) BKorti库尔提() feud.Barony {
	return c.库尔提Korti
}

func (c *纳帕塔NapataCounty) BMarawi麦罗维() feud.Barony {
	return c.麦罗维Marawi
}

func (c *纳帕塔NapataCounty) BNapata纳帕塔() feud.Barony {
	return c.纳帕塔Napata
}

var CNapata纳帕塔 NapataCounty = &纳帕塔NapataCounty{}

func init() {
	f := CNapata纳帕塔.(*纳帕塔NapataCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1326",
		Title:     "napata",
		TitleName: "纳帕塔",
		TitleCode: "c_napata",
		Baronies:  map[string]feud.Barony{},
	}

	f.加扎利Ghazali = BGhazali加扎利
	f.加扎利Ghazali.SetParent(f)

	f.凯尼塞Kanisah = BKanisah凯尼塞
	f.凯尼塞Kanisah.SetParent(f)

	f.凯里迈Karima = BKarima凯里迈
	f.凯里迈Karima.SetParent(f)

	f.科赫伊利Keheilli = BKeheilli科赫伊利
	f.科赫伊利Keheilli.SetParent(f)

	f.库尔提Korti = BKorti库尔提
	f.库尔提Korti.SetParent(f)

	f.麦罗维Marawi = BMarawi麦罗维
	f.麦罗维Marawi.SetParent(f)

	f.纳帕塔Napata = BNapata纳帕塔
	f.纳帕塔Napata.SetParent(f)

}
