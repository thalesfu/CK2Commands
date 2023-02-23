package wag

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WagCounty interface {
	feud.County
	BLashwa拉舍瓦() feud.Barony
	BLiktaba利克塔巴() feud.Barony
	BSacca萨卡() feud.Barony
	BSekota塞科塔() feud.Barony
	BTsaicha扎伊查() feud.Barony
	BZumbo宗博() feud.Barony
}

type 瓦格WagCounty struct {
	feud.BaseCounty
	拉舍瓦Lashwa   feud.Barony
	利克塔巴Liktaba feud.Barony
	萨卡Sacca     feud.Barony
	塞科塔Sekota   feud.Barony
	扎伊查Tsaicha  feud.Barony
	宗博Zumbo     feud.Barony
}

func (c *瓦格WagCounty) BLashwa拉舍瓦() feud.Barony {
	return c.拉舍瓦Lashwa
}

func (c *瓦格WagCounty) BLiktaba利克塔巴() feud.Barony {
	return c.利克塔巴Liktaba
}

func (c *瓦格WagCounty) BSacca萨卡() feud.Barony {
	return c.萨卡Sacca
}

func (c *瓦格WagCounty) BSekota塞科塔() feud.Barony {
	return c.塞科塔Sekota
}

func (c *瓦格WagCounty) BTsaicha扎伊查() feud.Barony {
	return c.扎伊查Tsaicha
}

func (c *瓦格WagCounty) BZumbo宗博() feud.Barony {
	return c.宗博Zumbo
}

var CWag瓦格 WagCounty = &瓦格WagCounty{}

func init() {
	f := CWag瓦格.(*瓦格WagCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1417",
		Title:     "wag",
		TitleName: "瓦格",
		TitleCode: "c_wag",
		Baronies:  map[string]feud.Barony{},
	}

	f.拉舍瓦Lashwa = BLashwa拉舍瓦
	f.拉舍瓦Lashwa.SetParent(f)

	f.利克塔巴Liktaba = BLiktaba利克塔巴
	f.利克塔巴Liktaba.SetParent(f)

	f.萨卡Sacca = BSacca萨卡
	f.萨卡Sacca.SetParent(f)

	f.塞科塔Sekota = BSekota塞科塔
	f.塞科塔Sekota.SetParent(f)

	f.扎伊查Tsaicha = BTsaicha扎伊查
	f.扎伊查Tsaicha.SetParent(f)

	f.宗博Zumbo = BZumbo宗博
	f.宗博Zumbo.SetParent(f)

}
