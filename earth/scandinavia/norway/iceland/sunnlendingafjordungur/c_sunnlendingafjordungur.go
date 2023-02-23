package sunnlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SunnlendingafjordungurCounty interface {
	feud.County
	BAlftanes奥尔塔内斯() feud.Barony
	BHlidarendi赫利达伦迪() feud.Barony
	BKjalarnes恰拉尔内斯() feud.Barony
	BOlfusa厄尔菲绍() feud.Barony
	BPingvellir辛格韦德利() feud.Barony
	BReykjavik雷克雅未克() feud.Barony
	BSkalholt斯考尔霍特() feud.Barony
}

type 南冰岛SunnlendingafjordungurCounty struct {
	feud.BaseCounty
	奥尔塔内斯Alftanes   feud.Barony
	赫利达伦迪Hlidarendi feud.Barony
	恰拉尔内斯Kjalarnes  feud.Barony
	厄尔菲绍Olfusa      feud.Barony
	辛格韦德利Pingvellir feud.Barony
	雷克雅未克Reykjavik  feud.Barony
	斯考尔霍特Skalholt   feud.Barony
}

func (c *南冰岛SunnlendingafjordungurCounty) BAlftanes奥尔塔内斯() feud.Barony {
	return c.奥尔塔内斯Alftanes
}

func (c *南冰岛SunnlendingafjordungurCounty) BHlidarendi赫利达伦迪() feud.Barony {
	return c.赫利达伦迪Hlidarendi
}

func (c *南冰岛SunnlendingafjordungurCounty) BKjalarnes恰拉尔内斯() feud.Barony {
	return c.恰拉尔内斯Kjalarnes
}

func (c *南冰岛SunnlendingafjordungurCounty) BOlfusa厄尔菲绍() feud.Barony {
	return c.厄尔菲绍Olfusa
}

func (c *南冰岛SunnlendingafjordungurCounty) BPingvellir辛格韦德利() feud.Barony {
	return c.辛格韦德利Pingvellir
}

func (c *南冰岛SunnlendingafjordungurCounty) BReykjavik雷克雅未克() feud.Barony {
	return c.雷克雅未克Reykjavik
}

func (c *南冰岛SunnlendingafjordungurCounty) BSkalholt斯考尔霍特() feud.Barony {
	return c.斯考尔霍特Skalholt
}

var CSunnlendingafjordungur南冰岛 SunnlendingafjordungurCounty = &南冰岛SunnlendingafjordungurCounty{}

func init() {
	f := CSunnlendingafjordungur南冰岛.(*南冰岛SunnlendingafjordungurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1617",
		Title:     "sunnlendingafjordungur",
		TitleName: "南冰岛",
		TitleCode: "c_sunnlendingafjordungur",
		Baronies:  map[string]feud.Barony{},
	}

	f.奥尔塔内斯Alftanes = BAlftanes奥尔塔内斯
	f.奥尔塔内斯Alftanes.SetParent(f)

	f.赫利达伦迪Hlidarendi = BHlidarendi赫利达伦迪
	f.赫利达伦迪Hlidarendi.SetParent(f)

	f.恰拉尔内斯Kjalarnes = BKjalarnes恰拉尔内斯
	f.恰拉尔内斯Kjalarnes.SetParent(f)

	f.厄尔菲绍Olfusa = BOlfusa厄尔菲绍
	f.厄尔菲绍Olfusa.SetParent(f)

	f.辛格韦德利Pingvellir = BPingvellir辛格韦德利
	f.辛格韦德利Pingvellir.SetParent(f)

	f.雷克雅未克Reykjavik = BReykjavik雷克雅未克
	f.雷克雅未克Reykjavik.SetParent(f)

	f.斯考尔霍特Skalholt = BSkalholt斯考尔霍特
	f.斯考尔霍特Skalholt.SetParent(f)

}
