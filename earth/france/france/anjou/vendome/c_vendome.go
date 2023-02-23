package vendome

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VendomeCounty interface {
	feud.County
	BChateaurenault雷诺堡() feud.Barony
	BLavardin拉瓦尔丹() feud.Barony
	BMontoire蒙图瓦尔() feud.Barony
	BOucques乌克() feud.Barony
	BRomilly罗米伊() feud.Barony
	BStavit圣阿维() feud.Barony
	BVendome旺多姆() feud.Barony
}

type 旺多姆VendomeCounty struct {
	feud.BaseCounty
	雷诺堡Chateaurenault feud.Barony
	拉瓦尔丹Lavardin      feud.Barony
	蒙图瓦尔Montoire      feud.Barony
	乌克Oucques         feud.Barony
	罗米伊Romilly        feud.Barony
	圣阿维Stavit         feud.Barony
	旺多姆Vendome        feud.Barony
}

func (c *旺多姆VendomeCounty) BChateaurenault雷诺堡() feud.Barony {
	return c.雷诺堡Chateaurenault
}

func (c *旺多姆VendomeCounty) BLavardin拉瓦尔丹() feud.Barony {
	return c.拉瓦尔丹Lavardin
}

func (c *旺多姆VendomeCounty) BMontoire蒙图瓦尔() feud.Barony {
	return c.蒙图瓦尔Montoire
}

func (c *旺多姆VendomeCounty) BOucques乌克() feud.Barony {
	return c.乌克Oucques
}

func (c *旺多姆VendomeCounty) BRomilly罗米伊() feud.Barony {
	return c.罗米伊Romilly
}

func (c *旺多姆VendomeCounty) BStavit圣阿维() feud.Barony {
	return c.圣阿维Stavit
}

func (c *旺多姆VendomeCounty) BVendome旺多姆() feud.Barony {
	return c.旺多姆Vendome
}

var CVendome旺多姆 VendomeCounty = &旺多姆VendomeCounty{}

func init() {
	f := CVendome旺多姆.(*旺多姆VendomeCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "109",
		Title:     "vendome",
		TitleName: "旺多姆",
		TitleCode: "c_vendome",
		Baronies:  map[string]feud.Barony{},
	}

	f.雷诺堡Chateaurenault = BChateaurenault雷诺堡
	f.雷诺堡Chateaurenault.SetParent(f)

	f.拉瓦尔丹Lavardin = BLavardin拉瓦尔丹
	f.拉瓦尔丹Lavardin.SetParent(f)

	f.蒙图瓦尔Montoire = BMontoire蒙图瓦尔
	f.蒙图瓦尔Montoire.SetParent(f)

	f.乌克Oucques = BOucques乌克
	f.乌克Oucques.SetParent(f)

	f.罗米伊Romilly = BRomilly罗米伊
	f.罗米伊Romilly.SetParent(f)

	f.圣阿维Stavit = BStavit圣阿维
	f.圣阿维Stavit.SetParent(f)

	f.旺多姆Vendome = BVendome旺多姆
	f.旺多姆Vendome.SetParent(f)

}
