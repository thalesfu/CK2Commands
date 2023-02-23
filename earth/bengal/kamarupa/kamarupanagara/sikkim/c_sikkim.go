package sikkim

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SikkimCounty interface {
	feud.County
	BGangdoz甘托克() feud.Barony
	BGyalshing加欣() feud.Barony
	BMangan曼甘() feud.Barony
	BNamchi纳姆吉() feud.Barony
	BRalang拉浪() feud.Barony
	BRangpo朗格波() feud.Barony
	BRumtheg隆德() feud.Barony
}

type 锡金SikkimCounty struct {
	feud.BaseCounty
	甘托克Gangdoz  feud.Barony
	加欣Gyalshing feud.Barony
	曼甘Mangan    feud.Barony
	纳姆吉Namchi   feud.Barony
	拉浪Ralang    feud.Barony
	朗格波Rangpo   feud.Barony
	隆德Rumtheg   feud.Barony
}

func (c *锡金SikkimCounty) BGangdoz甘托克() feud.Barony {
	return c.甘托克Gangdoz
}

func (c *锡金SikkimCounty) BGyalshing加欣() feud.Barony {
	return c.加欣Gyalshing
}

func (c *锡金SikkimCounty) BMangan曼甘() feud.Barony {
	return c.曼甘Mangan
}

func (c *锡金SikkimCounty) BNamchi纳姆吉() feud.Barony {
	return c.纳姆吉Namchi
}

func (c *锡金SikkimCounty) BRalang拉浪() feud.Barony {
	return c.拉浪Ralang
}

func (c *锡金SikkimCounty) BRangpo朗格波() feud.Barony {
	return c.朗格波Rangpo
}

func (c *锡金SikkimCounty) BRumtheg隆德() feud.Barony {
	return c.隆德Rumtheg
}

var CSikkim锡金 SikkimCounty = &锡金SikkimCounty{}

func init() {
	f := CSikkim锡金.(*锡金SikkimCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1480",
		Title:     "sikkim",
		TitleName: "锡金",
		TitleCode: "c_sikkim",
		Baronies:  map[string]feud.Barony{},
	}

	f.甘托克Gangdoz = BGangdoz甘托克
	f.甘托克Gangdoz.SetParent(f)

	f.加欣Gyalshing = BGyalshing加欣
	f.加欣Gyalshing.SetParent(f)

	f.曼甘Mangan = BMangan曼甘
	f.曼甘Mangan.SetParent(f)

	f.纳姆吉Namchi = BNamchi纳姆吉
	f.纳姆吉Namchi.SetParent(f)

	f.拉浪Ralang = BRalang拉浪
	f.拉浪Ralang.SetParent(f)

	f.朗格波Rangpo = BRangpo朗格波
	f.朗格波Rangpo.SetParent(f)

	f.隆德Rumtheg = BRumtheg隆德
	f.隆德Rumtheg.SetParent(f)

}
