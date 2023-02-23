package leh

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LehCounty interface {
	feud.County
	BChumathang楚姆唐() feud.Barony
	BGya盖亚() feud.Barony
	BKeylong吉隆() feud.Barony
	BLeh列城() feud.Barony
	BShey雪依() feud.Barony
	BThiksey提克西() feud.Barony
}

type 列城LehCounty struct {
	feud.BaseCounty
	楚姆唐Chumathang feud.Barony
	盖亚Gya         feud.Barony
	吉隆Keylong     feud.Barony
	列城Leh         feud.Barony
	雪依Shey        feud.Barony
	提克西Thiksey    feud.Barony
}

func (c *列城LehCounty) BChumathang楚姆唐() feud.Barony {
	return c.楚姆唐Chumathang
}

func (c *列城LehCounty) BGya盖亚() feud.Barony {
	return c.盖亚Gya
}

func (c *列城LehCounty) BKeylong吉隆() feud.Barony {
	return c.吉隆Keylong
}

func (c *列城LehCounty) BLeh列城() feud.Barony {
	return c.列城Leh
}

func (c *列城LehCounty) BShey雪依() feud.Barony {
	return c.雪依Shey
}

func (c *列城LehCounty) BThiksey提克西() feud.Barony {
	return c.提克西Thiksey
}

var CLeh列城 LehCounty = &列城LehCounty{}

func init() {
	f := CLeh列城.(*列城LehCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1468",
		Title:     "leh",
		TitleName: "列城",
		TitleCode: "c_leh",
		Baronies:  map[string]feud.Barony{},
	}

	f.楚姆唐Chumathang = BChumathang楚姆唐
	f.楚姆唐Chumathang.SetParent(f)

	f.盖亚Gya = BGya盖亚
	f.盖亚Gya.SetParent(f)

	f.吉隆Keylong = BKeylong吉隆
	f.吉隆Keylong.SetParent(f)

	f.列城Leh = BLeh列城
	f.列城Leh.SetParent(f)

	f.雪依Shey = BShey雪依
	f.雪依Shey.SetParent(f)

	f.提克西Thiksey = BThiksey提克西
	f.提克西Thiksey.SetParent(f)

}
