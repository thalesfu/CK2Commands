package timbuktu

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TimbuktuCounty interface {
	feud.County
	BDire迪雷() feud.Barony
	BGundam贡达姆() feud.Barony
	BKabara卡巴拉() feud.Barony
	BSafanqu萨方克() feud.Barony
	BSankore桑科雷() feud.Barony
	BTendirma滕迪尔马() feud.Barony
	BTimbuktu廷巴克图() feud.Barony
	BTirakka蒂拉卡() feud.Barony
}

type 廷巴克图TimbuktuCounty struct {
	feud.BaseCounty
	迪雷Dire       feud.Barony
	贡达姆Gundam    feud.Barony
	卡巴拉Kabara    feud.Barony
	萨方克Safanqu   feud.Barony
	桑科雷Sankore   feud.Barony
	滕迪尔马Tendirma feud.Barony
	廷巴克图Timbuktu feud.Barony
	蒂拉卡Tirakka   feud.Barony
}

func (c *廷巴克图TimbuktuCounty) BDire迪雷() feud.Barony {
	return c.迪雷Dire
}

func (c *廷巴克图TimbuktuCounty) BGundam贡达姆() feud.Barony {
	return c.贡达姆Gundam
}

func (c *廷巴克图TimbuktuCounty) BKabara卡巴拉() feud.Barony {
	return c.卡巴拉Kabara
}

func (c *廷巴克图TimbuktuCounty) BSafanqu萨方克() feud.Barony {
	return c.萨方克Safanqu
}

func (c *廷巴克图TimbuktuCounty) BSankore桑科雷() feud.Barony {
	return c.桑科雷Sankore
}

func (c *廷巴克图TimbuktuCounty) BTendirma滕迪尔马() feud.Barony {
	return c.滕迪尔马Tendirma
}

func (c *廷巴克图TimbuktuCounty) BTimbuktu廷巴克图() feud.Barony {
	return c.廷巴克图Timbuktu
}

func (c *廷巴克图TimbuktuCounty) BTirakka蒂拉卡() feud.Barony {
	return c.蒂拉卡Tirakka
}

var CTimbuktu廷巴克图 TimbuktuCounty = &廷巴克图TimbuktuCounty{}

func init() {
	f := CTimbuktu廷巴克图.(*廷巴克图TimbuktuCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "911",
		Title:     "timbuktu",
		TitleName: "廷巴克图",
		TitleCode: "c_timbuktu",
		Baronies:  map[string]feud.Barony{},
	}

	f.迪雷Dire = BDire迪雷
	f.迪雷Dire.SetParent(f)

	f.贡达姆Gundam = BGundam贡达姆
	f.贡达姆Gundam.SetParent(f)

	f.卡巴拉Kabara = BKabara卡巴拉
	f.卡巴拉Kabara.SetParent(f)

	f.萨方克Safanqu = BSafanqu萨方克
	f.萨方克Safanqu.SetParent(f)

	f.桑科雷Sankore = BSankore桑科雷
	f.桑科雷Sankore.SetParent(f)

	f.滕迪尔马Tendirma = BTendirma滕迪尔马
	f.滕迪尔马Tendirma.SetParent(f)

	f.廷巴克图Timbuktu = BTimbuktu廷巴克图
	f.廷巴克图Timbuktu.SetParent(f)

	f.蒂拉卡Tirakka = BTirakka蒂拉卡
	f.蒂拉卡Tirakka.SetParent(f)

}
