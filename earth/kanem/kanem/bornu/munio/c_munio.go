package munio

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MunioCounty interface {
	feud.County
	BAmalget阿马尔盖特() feud.Barony
	BDiaka迪亚卡() feud.Barony
	BGumsa古姆萨() feud.Barony
	BMashina马希纳() feud.Barony
	BNguru恩古鲁() feud.Barony
	BTsaskiya扎斯基亚() feud.Barony
}

type 穆尼奥MunioCounty struct {
	feud.BaseCounty
	阿马尔盖特Amalget feud.Barony
	迪亚卡Diaka     feud.Barony
	古姆萨Gumsa     feud.Barony
	马希纳Mashina   feud.Barony
	恩古鲁Nguru     feud.Barony
	扎斯基亚Tsaskiya feud.Barony
}

func (c *穆尼奥MunioCounty) BAmalget阿马尔盖特() feud.Barony {
	return c.阿马尔盖特Amalget
}

func (c *穆尼奥MunioCounty) BDiaka迪亚卡() feud.Barony {
	return c.迪亚卡Diaka
}

func (c *穆尼奥MunioCounty) BGumsa古姆萨() feud.Barony {
	return c.古姆萨Gumsa
}

func (c *穆尼奥MunioCounty) BMashina马希纳() feud.Barony {
	return c.马希纳Mashina
}

func (c *穆尼奥MunioCounty) BNguru恩古鲁() feud.Barony {
	return c.恩古鲁Nguru
}

func (c *穆尼奥MunioCounty) BTsaskiya扎斯基亚() feud.Barony {
	return c.扎斯基亚Tsaskiya
}

var CMunio穆尼奥 MunioCounty = &穆尼奥MunioCounty{}

func init() {
	f := CMunio穆尼奥.(*穆尼奥MunioCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1744",
		Title:     "munio",
		TitleName: "穆尼奥",
		TitleCode: "c_munio",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿马尔盖特Amalget = BAmalget阿马尔盖特
	f.阿马尔盖特Amalget.SetParent(f)

	f.迪亚卡Diaka = BDiaka迪亚卡
	f.迪亚卡Diaka.SetParent(f)

	f.古姆萨Gumsa = BGumsa古姆萨
	f.古姆萨Gumsa.SetParent(f)

	f.马希纳Mashina = BMashina马希纳
	f.马希纳Mashina.SetParent(f)

	f.恩古鲁Nguru = BNguru恩古鲁
	f.恩古鲁Nguru.SetParent(f)

	f.扎斯基亚Tsaskiya = BTsaskiya扎斯基亚
	f.扎斯基亚Tsaskiya.SetParent(f)

}
