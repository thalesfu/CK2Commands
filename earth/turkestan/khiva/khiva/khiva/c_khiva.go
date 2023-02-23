package khiva

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KhivaCounty interface {
	feud.County
	BAtajab阿塔亚布() feud.Barony
	BDarvaza达瓦扎() feud.Barony
	BKath卡斯() feud.Barony
	BKhiva希瓦() feud.Barony
	BKizketken基克肯() feud.Barony
	BKuskupir库斯库皮尔() feud.Barony
	BSumanay苏曼纳亚() feud.Barony
	BTahta塔赫塔() feud.Barony
}

type 希瓦KhivaCounty struct {
	feud.BaseCounty
	阿塔亚布Atajab    feud.Barony
	达瓦扎Darvaza    feud.Barony
	卡斯Kath        feud.Barony
	希瓦Khiva       feud.Barony
	基克肯Kizketken  feud.Barony
	库斯库皮尔Kuskupir feud.Barony
	苏曼纳亚Sumanay   feud.Barony
	塔赫塔Tahta      feud.Barony
}

func (c *希瓦KhivaCounty) BAtajab阿塔亚布() feud.Barony {
	return c.阿塔亚布Atajab
}

func (c *希瓦KhivaCounty) BDarvaza达瓦扎() feud.Barony {
	return c.达瓦扎Darvaza
}

func (c *希瓦KhivaCounty) BKath卡斯() feud.Barony {
	return c.卡斯Kath
}

func (c *希瓦KhivaCounty) BKhiva希瓦() feud.Barony {
	return c.希瓦Khiva
}

func (c *希瓦KhivaCounty) BKizketken基克肯() feud.Barony {
	return c.基克肯Kizketken
}

func (c *希瓦KhivaCounty) BKuskupir库斯库皮尔() feud.Barony {
	return c.库斯库皮尔Kuskupir
}

func (c *希瓦KhivaCounty) BSumanay苏曼纳亚() feud.Barony {
	return c.苏曼纳亚Sumanay
}

func (c *希瓦KhivaCounty) BTahta塔赫塔() feud.Barony {
	return c.塔赫塔Tahta
}

var CKhiva希瓦 KhivaCounty = &希瓦KhivaCounty{}

func init() {
	f := CKhiva希瓦.(*希瓦KhivaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "626",
		Title:     "khiva",
		TitleName: "希瓦",
		TitleCode: "c_khiva",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿塔亚布Atajab = BAtajab阿塔亚布
	f.阿塔亚布Atajab.SetParent(f)

	f.达瓦扎Darvaza = BDarvaza达瓦扎
	f.达瓦扎Darvaza.SetParent(f)

	f.卡斯Kath = BKath卡斯
	f.卡斯Kath.SetParent(f)

	f.希瓦Khiva = BKhiva希瓦
	f.希瓦Khiva.SetParent(f)

	f.基克肯Kizketken = BKizketken基克肯
	f.基克肯Kizketken.SetParent(f)

	f.库斯库皮尔Kuskupir = BKuskupir库斯库皮尔
	f.库斯库皮尔Kuskupir.SetParent(f)

	f.苏曼纳亚Sumanay = BSumanay苏曼纳亚
	f.苏曼纳亚Sumanay.SetParent(f)

	f.塔赫塔Tahta = BTahta塔赫塔
	f.塔赫塔Tahta.SetParent(f)

}
