package mainling

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MainlingCounty interface {
	feud.County
	BDengmu登木() feud.Barony
	BJindong金东() feud.Barony
	BLilong里龙() feud.Barony
	BMainling米林() feud.Barony
	BNang朗县() feud.Barony
	BNanyi南伊() feud.Barony
	BWolong卧龙() feud.Barony
}

type 米林MainlingCounty struct {
	feud.BaseCounty
	登木Dengmu   feud.Barony
	金东Jindong  feud.Barony
	里龙Lilong   feud.Barony
	米林Mainling feud.Barony
	朗县Nang     feud.Barony
	南伊Nanyi    feud.Barony
	卧龙Wolong   feud.Barony
}

func (c *米林MainlingCounty) BDengmu登木() feud.Barony {
	return c.登木Dengmu
}

func (c *米林MainlingCounty) BJindong金东() feud.Barony {
	return c.金东Jindong
}

func (c *米林MainlingCounty) BLilong里龙() feud.Barony {
	return c.里龙Lilong
}

func (c *米林MainlingCounty) BMainling米林() feud.Barony {
	return c.米林Mainling
}

func (c *米林MainlingCounty) BNang朗县() feud.Barony {
	return c.朗县Nang
}

func (c *米林MainlingCounty) BNanyi南伊() feud.Barony {
	return c.南伊Nanyi
}

func (c *米林MainlingCounty) BWolong卧龙() feud.Barony {
	return c.卧龙Wolong
}

var CMainling米林 MainlingCounty = &米林MainlingCounty{}

func init() {
	f := CMainling米林.(*米林MainlingCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1567",
		Title:     "mainling",
		TitleName: "米林",
		TitleCode: "c_mainling",
		Baronies:  map[string]feud.Barony{},
	}

	f.登木Dengmu = BDengmu登木
	f.登木Dengmu.SetParent(f)

	f.金东Jindong = BJindong金东
	f.金东Jindong.SetParent(f)

	f.里龙Lilong = BLilong里龙
	f.里龙Lilong.SetParent(f)

	f.米林Mainling = BMainling米林
	f.米林Mainling.SetParent(f)

	f.朗县Nang = BNang朗县
	f.朗县Nang.SetParent(f)

	f.南伊Nanyi = BNanyi南伊
	f.南伊Nanyi.SetParent(f)

	f.卧龙Wolong = BWolong卧龙
	f.卧龙Wolong.SetParent(f)

}
