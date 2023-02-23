package tsakha

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TsakhaCounty interface {
	feud.County
	BGacha嘎查() feud.Barony
	BNyer聂尔() feud.Barony
	BOmbulung屋木龙() feud.Barony
	BPerotse别若则() feud.Barony
	BRungra绒热() feud.Barony
	BSharma夏玛() feud.Barony
	BTsakha茶卡() feud.Barony
}

type 茶卡TsakhaCounty struct {
	feud.BaseCounty
	嘎查Gacha     feud.Barony
	聂尔Nyer      feud.Barony
	屋木龙Ombulung feud.Barony
	别若则Perotse  feud.Barony
	绒热Rungra    feud.Barony
	夏玛Sharma    feud.Barony
	茶卡Tsakha    feud.Barony
}

func (c *茶卡TsakhaCounty) BGacha嘎查() feud.Barony {
	return c.嘎查Gacha
}

func (c *茶卡TsakhaCounty) BNyer聂尔() feud.Barony {
	return c.聂尔Nyer
}

func (c *茶卡TsakhaCounty) BOmbulung屋木龙() feud.Barony {
	return c.屋木龙Ombulung
}

func (c *茶卡TsakhaCounty) BPerotse别若则() feud.Barony {
	return c.别若则Perotse
}

func (c *茶卡TsakhaCounty) BRungra绒热() feud.Barony {
	return c.绒热Rungra
}

func (c *茶卡TsakhaCounty) BSharma夏玛() feud.Barony {
	return c.夏玛Sharma
}

func (c *茶卡TsakhaCounty) BTsakha茶卡() feud.Barony {
	return c.茶卡Tsakha
}

var CTsakha茶卡 TsakhaCounty = &茶卡TsakhaCounty{}

func init() {
	f := CTsakha茶卡.(*茶卡TsakhaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1565",
		Title:     "tsakha",
		TitleName: "茶卡",
		TitleCode: "c_tsakha",
		Baronies:  map[string]feud.Barony{},
	}

	f.嘎查Gacha = BGacha嘎查
	f.嘎查Gacha.SetParent(f)

	f.聂尔Nyer = BNyer聂尔
	f.聂尔Nyer.SetParent(f)

	f.屋木龙Ombulung = BOmbulung屋木龙
	f.屋木龙Ombulung.SetParent(f)

	f.别若则Perotse = BPerotse别若则
	f.别若则Perotse.SetParent(f)

	f.绒热Rungra = BRungra绒热
	f.绒热Rungra.SetParent(f)

	f.夏玛Sharma = BSharma夏玛
	f.夏玛Sharma.SetParent(f)

	f.茶卡Tsakha = BTsakha茶卡
	f.茶卡Tsakha.SetParent(f)

}
