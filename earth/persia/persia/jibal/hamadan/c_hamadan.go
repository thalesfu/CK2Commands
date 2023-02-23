package hamadan

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HamadanCounty interface {
	feud.County
	BAsadabad阿萨达巴德() feud.Barony
	BEcbatana埃克巴坦那() feud.Barony
	BGanjnameh甘吉纳梅() feud.Barony
	BHamadan哈马丹() feud.Barony
	BKabudrahang卡布德拉汉格() feud.Barony
	BMalayer马拉耶尔() feud.Barony
	BNahavand纳哈万德() feud.Barony
	BRoudavar鲁德阿瓦尔() feud.Barony
}

type 哈马丹HamadanCounty struct {
	feud.BaseCounty
	阿萨达巴德Asadabad     feud.Barony
	埃克巴坦那Ecbatana     feud.Barony
	甘吉纳梅Ganjnameh     feud.Barony
	哈马丹Hamadan        feud.Barony
	卡布德拉汉格Kabudrahang feud.Barony
	马拉耶尔Malayer       feud.Barony
	纳哈万德Nahavand      feud.Barony
	鲁德阿瓦尔Roudavar     feud.Barony
}

func (c *哈马丹HamadanCounty) BAsadabad阿萨达巴德() feud.Barony {
	return c.阿萨达巴德Asadabad
}

func (c *哈马丹HamadanCounty) BEcbatana埃克巴坦那() feud.Barony {
	return c.埃克巴坦那Ecbatana
}

func (c *哈马丹HamadanCounty) BGanjnameh甘吉纳梅() feud.Barony {
	return c.甘吉纳梅Ganjnameh
}

func (c *哈马丹HamadanCounty) BHamadan哈马丹() feud.Barony {
	return c.哈马丹Hamadan
}

func (c *哈马丹HamadanCounty) BKabudrahang卡布德拉汉格() feud.Barony {
	return c.卡布德拉汉格Kabudrahang
}

func (c *哈马丹HamadanCounty) BMalayer马拉耶尔() feud.Barony {
	return c.马拉耶尔Malayer
}

func (c *哈马丹HamadanCounty) BNahavand纳哈万德() feud.Barony {
	return c.纳哈万德Nahavand
}

func (c *哈马丹HamadanCounty) BRoudavar鲁德阿瓦尔() feud.Barony {
	return c.鲁德阿瓦尔Roudavar
}

var CHamadan哈马丹 HamadanCounty = &哈马丹HamadanCounty{}

func init() {
	f := CHamadan哈马丹.(*哈马丹HamadanCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "658",
		Title:     "hamadan",
		TitleName: "哈马丹",
		TitleCode: "c_hamadan",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿萨达巴德Asadabad = BAsadabad阿萨达巴德
	f.阿萨达巴德Asadabad.SetParent(f)

	f.埃克巴坦那Ecbatana = BEcbatana埃克巴坦那
	f.埃克巴坦那Ecbatana.SetParent(f)

	f.甘吉纳梅Ganjnameh = BGanjnameh甘吉纳梅
	f.甘吉纳梅Ganjnameh.SetParent(f)

	f.哈马丹Hamadan = BHamadan哈马丹
	f.哈马丹Hamadan.SetParent(f)

	f.卡布德拉汉格Kabudrahang = BKabudrahang卡布德拉汉格
	f.卡布德拉汉格Kabudrahang.SetParent(f)

	f.马拉耶尔Malayer = BMalayer马拉耶尔
	f.马拉耶尔Malayer.SetParent(f)

	f.纳哈万德Nahavand = BNahavand纳哈万德
	f.纳哈万德Nahavand.SetParent(f)

	f.鲁德阿瓦尔Roudavar = BRoudavar鲁德阿瓦尔
	f.鲁德阿瓦尔Roudavar.SetParent(f)

}
