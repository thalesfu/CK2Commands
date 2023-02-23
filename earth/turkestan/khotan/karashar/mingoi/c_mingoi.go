package mingoi

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MingoiCounty interface {
	feud.County
	BMingoi明屋() feud.Barony
	BPingpo平坡() feud.Barony
	BQigexing七个星() feud.Barony
	BShikshin锡克沁() feud.Barony
	BShorshuq舒尔楚克() feud.Barony
	BTaibaiya太白垭() feud.Barony
	BYushe玉舍() feud.Barony
}

type 明屋MingoiCounty struct {
	feud.BaseCounty
	明屋Mingoi     feud.Barony
	平坡Pingpo     feud.Barony
	七个星Qigexing  feud.Barony
	锡克沁Shikshin  feud.Barony
	舒尔楚克Shorshuq feud.Barony
	太白垭Taibaiya  feud.Barony
	玉舍Yushe      feud.Barony
}

func (c *明屋MingoiCounty) BMingoi明屋() feud.Barony {
	return c.明屋Mingoi
}

func (c *明屋MingoiCounty) BPingpo平坡() feud.Barony {
	return c.平坡Pingpo
}

func (c *明屋MingoiCounty) BQigexing七个星() feud.Barony {
	return c.七个星Qigexing
}

func (c *明屋MingoiCounty) BShikshin锡克沁() feud.Barony {
	return c.锡克沁Shikshin
}

func (c *明屋MingoiCounty) BShorshuq舒尔楚克() feud.Barony {
	return c.舒尔楚克Shorshuq
}

func (c *明屋MingoiCounty) BTaibaiya太白垭() feud.Barony {
	return c.太白垭Taibaiya
}

func (c *明屋MingoiCounty) BYushe玉舍() feud.Barony {
	return c.玉舍Yushe
}

var CMingoi明屋 MingoiCounty = &明屋MingoiCounty{}

func init() {
	f := CMingoi明屋.(*明屋MingoiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1521",
		Title:     "mingoi",
		TitleName: "明屋",
		TitleCode: "c_mingoi",
		Baronies:  map[string]feud.Barony{},
	}

	f.明屋Mingoi = BMingoi明屋
	f.明屋Mingoi.SetParent(f)

	f.平坡Pingpo = BPingpo平坡
	f.平坡Pingpo.SetParent(f)

	f.七个星Qigexing = BQigexing七个星
	f.七个星Qigexing.SetParent(f)

	f.锡克沁Shikshin = BShikshin锡克沁
	f.锡克沁Shikshin.SetParent(f)

	f.舒尔楚克Shorshuq = BShorshuq舒尔楚克
	f.舒尔楚克Shorshuq.SetParent(f)

	f.太白垭Taibaiya = BTaibaiya太白垭
	f.太白垭Taibaiya.SetParent(f)

	f.玉舍Yushe = BYushe玉舍
	f.玉舍Yushe.SetParent(f)

}
