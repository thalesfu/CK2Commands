package nordlendingafjordungur

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NordlendingafjordungurCounty interface {
	feud.County
	BAkureyri阿克雷里() feud.Barony
	BGlaumbaer格伦拜尔() feud.Barony
	BGoddalir格兹达利尔() feud.Barony
	BHolar霍拉尔() feud.Barony
	BHrisey赫里塞() feud.Barony
	BHusavik胡萨维克() feud.Barony
	BNattfaravik瑙特法拉维克() feud.Barony
}

type 北冰岛NordlendingafjordungurCounty struct {
	feud.BaseCounty
	阿克雷里Akureyri      feud.Barony
	格伦拜尔Glaumbaer     feud.Barony
	格兹达利尔Goddalir     feud.Barony
	霍拉尔Holar          feud.Barony
	赫里塞Hrisey         feud.Barony
	胡萨维克Husavik       feud.Barony
	瑙特法拉维克Nattfaravik feud.Barony
}

func (c *北冰岛NordlendingafjordungurCounty) BAkureyri阿克雷里() feud.Barony {
	return c.阿克雷里Akureyri
}

func (c *北冰岛NordlendingafjordungurCounty) BGlaumbaer格伦拜尔() feud.Barony {
	return c.格伦拜尔Glaumbaer
}

func (c *北冰岛NordlendingafjordungurCounty) BGoddalir格兹达利尔() feud.Barony {
	return c.格兹达利尔Goddalir
}

func (c *北冰岛NordlendingafjordungurCounty) BHolar霍拉尔() feud.Barony {
	return c.霍拉尔Holar
}

func (c *北冰岛NordlendingafjordungurCounty) BHrisey赫里塞() feud.Barony {
	return c.赫里塞Hrisey
}

func (c *北冰岛NordlendingafjordungurCounty) BHusavik胡萨维克() feud.Barony {
	return c.胡萨维克Husavik
}

func (c *北冰岛NordlendingafjordungurCounty) BNattfaravik瑙特法拉维克() feud.Barony {
	return c.瑙特法拉维克Nattfaravik
}

var CNordlendingafjordungur北冰岛 NordlendingafjordungurCounty = &北冰岛NordlendingafjordungurCounty{}

func init() {
	f := CNordlendingafjordungur北冰岛.(*北冰岛NordlendingafjordungurCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1616",
		Title:     "nordlendingafjordungur",
		TitleName: "北冰岛",
		TitleCode: "c_nordlendingafjordungur",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿克雷里Akureyri = BAkureyri阿克雷里
	f.阿克雷里Akureyri.SetParent(f)

	f.格伦拜尔Glaumbaer = BGlaumbaer格伦拜尔
	f.格伦拜尔Glaumbaer.SetParent(f)

	f.格兹达利尔Goddalir = BGoddalir格兹达利尔
	f.格兹达利尔Goddalir.SetParent(f)

	f.霍拉尔Holar = BHolar霍拉尔
	f.霍拉尔Holar.SetParent(f)

	f.赫里塞Hrisey = BHrisey赫里塞
	f.赫里塞Hrisey.SetParent(f)

	f.胡萨维克Husavik = BHusavik胡萨维克
	f.胡萨维克Husavik.SetParent(f)

	f.瑙特法拉维克Nattfaravik = BNattfaravik瑙特法拉维克
	f.瑙特法拉维克Nattfaravik.SetParent(f)

}
