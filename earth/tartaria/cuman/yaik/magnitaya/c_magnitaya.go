package magnitaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MagnitayaCounty interface {
	feud.County
	BBalkany巴尔卡内() feud.Barony
	BBikkulovo比库洛沃() feud.Barony
	BGusevo古谢沃() feud.Barony
	BKyubas丘巴斯() feud.Barony
	BMagnitaya马格尼特纳亚() feud.Barony
	BTirmen季尔缅() feud.Barony
	BUlyandy乌良德() feud.Barony
}

type 马格尼特纳亚MagnitayaCounty struct {
	feud.BaseCounty
	巴尔卡内Balkany     feud.Barony
	比库洛沃Bikkulovo   feud.Barony
	古谢沃Gusevo       feud.Barony
	丘巴斯Kyubas       feud.Barony
	马格尼特纳亚Magnitaya feud.Barony
	季尔缅Tirmen       feud.Barony
	乌良德Ulyandy      feud.Barony
}

func (c *马格尼特纳亚MagnitayaCounty) BBalkany巴尔卡内() feud.Barony {
	return c.巴尔卡内Balkany
}

func (c *马格尼特纳亚MagnitayaCounty) BBikkulovo比库洛沃() feud.Barony {
	return c.比库洛沃Bikkulovo
}

func (c *马格尼特纳亚MagnitayaCounty) BGusevo古谢沃() feud.Barony {
	return c.古谢沃Gusevo
}

func (c *马格尼特纳亚MagnitayaCounty) BKyubas丘巴斯() feud.Barony {
	return c.丘巴斯Kyubas
}

func (c *马格尼特纳亚MagnitayaCounty) BMagnitaya马格尼特纳亚() feud.Barony {
	return c.马格尼特纳亚Magnitaya
}

func (c *马格尼特纳亚MagnitayaCounty) BTirmen季尔缅() feud.Barony {
	return c.季尔缅Tirmen
}

func (c *马格尼特纳亚MagnitayaCounty) BUlyandy乌良德() feud.Barony {
	return c.乌良德Ulyandy
}

var CMagnitaya马格尼特纳亚 MagnitayaCounty = &马格尼特纳亚MagnitayaCounty{}

func init() {
	f := CMagnitaya马格尼特纳亚.(*马格尼特纳亚MagnitayaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1853",
		Title:     "magnitaya",
		TitleName: "马格尼特纳亚",
		TitleCode: "c_magnitaya",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴尔卡内Balkany = BBalkany巴尔卡内
	f.巴尔卡内Balkany.SetParent(f)

	f.比库洛沃Bikkulovo = BBikkulovo比库洛沃
	f.比库洛沃Bikkulovo.SetParent(f)

	f.古谢沃Gusevo = BGusevo古谢沃
	f.古谢沃Gusevo.SetParent(f)

	f.丘巴斯Kyubas = BKyubas丘巴斯
	f.丘巴斯Kyubas.SetParent(f)

	f.马格尼特纳亚Magnitaya = BMagnitaya马格尼特纳亚
	f.马格尼特纳亚Magnitaya.SetParent(f)

	f.季尔缅Tirmen = BTirmen季尔缅
	f.季尔缅Tirmen.SetParent(f)

	f.乌良德Ulyandy = BUlyandy乌良德
	f.乌良德Ulyandy.SetParent(f)

}
