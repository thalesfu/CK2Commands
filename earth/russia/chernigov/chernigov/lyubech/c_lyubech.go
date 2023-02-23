package lyubech

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LyubechCounty interface {
	feud.County
	BCierachouka切拉霍卡() feud.Barony
	BDobrouch多布鲁什() feud.Barony
	BDobryanka多布良卡() feud.Barony
	BLysyye雷瑟耶() feud.Barony
	BLyubech柳别奇() feud.Barony
	BPerepys佩列皮斯() feud.Barony
	BSkytok斯克托克() feud.Barony
}

type 柳别奇LyubechCounty struct {
	feud.BaseCounty
	切拉霍卡Cierachouka feud.Barony
	多布鲁什Dobrouch    feud.Barony
	多布良卡Dobryanka   feud.Barony
	雷瑟耶Lysyye       feud.Barony
	柳别奇Lyubech      feud.Barony
	佩列皮斯Perepys     feud.Barony
	斯克托克Skytok      feud.Barony
}

func (c *柳别奇LyubechCounty) BCierachouka切拉霍卡() feud.Barony {
	return c.切拉霍卡Cierachouka
}

func (c *柳别奇LyubechCounty) BDobrouch多布鲁什() feud.Barony {
	return c.多布鲁什Dobrouch
}

func (c *柳别奇LyubechCounty) BDobryanka多布良卡() feud.Barony {
	return c.多布良卡Dobryanka
}

func (c *柳别奇LyubechCounty) BLysyye雷瑟耶() feud.Barony {
	return c.雷瑟耶Lysyye
}

func (c *柳别奇LyubechCounty) BLyubech柳别奇() feud.Barony {
	return c.柳别奇Lyubech
}

func (c *柳别奇LyubechCounty) BPerepys佩列皮斯() feud.Barony {
	return c.佩列皮斯Perepys
}

func (c *柳别奇LyubechCounty) BSkytok斯克托克() feud.Barony {
	return c.斯克托克Skytok
}

var CLyubech柳别奇 LyubechCounty = &柳别奇LyubechCounty{}

func init() {
	f := CLyubech柳别奇.(*柳别奇LyubechCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "553",
		Title:     "lyubech",
		TitleName: "柳别奇",
		TitleCode: "c_lyubech",
		Baronies:  map[string]feud.Barony{},
	}

	f.切拉霍卡Cierachouka = BCierachouka切拉霍卡
	f.切拉霍卡Cierachouka.SetParent(f)

	f.多布鲁什Dobrouch = BDobrouch多布鲁什
	f.多布鲁什Dobrouch.SetParent(f)

	f.多布良卡Dobryanka = BDobryanka多布良卡
	f.多布良卡Dobryanka.SetParent(f)

	f.雷瑟耶Lysyye = BLysyye雷瑟耶
	f.雷瑟耶Lysyye.SetParent(f)

	f.柳别奇Lyubech = BLyubech柳别奇
	f.柳别奇Lyubech.SetParent(f)

	f.佩列皮斯Perepys = BPerepys佩列皮斯
	f.佩列皮斯Perepys.SetParent(f)

	f.斯克托克Skytok = BSkytok斯克托克
	f.斯克托克Skytok.SetParent(f)

}
