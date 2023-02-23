package pisa

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PisaCounty interface {
	feud.County
	BCanefro卡内夫罗() feud.Barony
	BGiglio吉廖() feud.Barony
	BLivorno利沃诺() feud.Barony
	BPisa比萨() feud.Barony
	BPortoferraio费拉约港() feud.Barony
	BSanminiato圣米尼亚托() feud.Barony
	BVicopisano维科皮萨诺() feud.Barony
	BVolterra沃尔泰拉() feud.Barony
}

type 比萨PisaCounty struct {
	feud.BaseCounty
	卡内夫罗Canefro      feud.Barony
	吉廖Giglio         feud.Barony
	利沃诺Livorno       feud.Barony
	比萨Pisa           feud.Barony
	费拉约港Portoferraio feud.Barony
	圣米尼亚托Sanminiato  feud.Barony
	维科皮萨诺Vicopisano  feud.Barony
	沃尔泰拉Volterra     feud.Barony
}

func (c *比萨PisaCounty) BCanefro卡内夫罗() feud.Barony {
	return c.卡内夫罗Canefro
}

func (c *比萨PisaCounty) BGiglio吉廖() feud.Barony {
	return c.吉廖Giglio
}

func (c *比萨PisaCounty) BLivorno利沃诺() feud.Barony {
	return c.利沃诺Livorno
}

func (c *比萨PisaCounty) BPisa比萨() feud.Barony {
	return c.比萨Pisa
}

func (c *比萨PisaCounty) BPortoferraio费拉约港() feud.Barony {
	return c.费拉约港Portoferraio
}

func (c *比萨PisaCounty) BSanminiato圣米尼亚托() feud.Barony {
	return c.圣米尼亚托Sanminiato
}

func (c *比萨PisaCounty) BVicopisano维科皮萨诺() feud.Barony {
	return c.维科皮萨诺Vicopisano
}

func (c *比萨PisaCounty) BVolterra沃尔泰拉() feud.Barony {
	return c.沃尔泰拉Volterra
}

var CPisa比萨 PisaCounty = &比萨PisaCounty{}

func init() {
	f := CPisa比萨.(*比萨PisaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "327",
		Title:     "pisa",
		TitleName: "比萨",
		TitleCode: "c_pisa",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡内夫罗Canefro = BCanefro卡内夫罗
	f.卡内夫罗Canefro.SetParent(f)

	f.吉廖Giglio = BGiglio吉廖
	f.吉廖Giglio.SetParent(f)

	f.利沃诺Livorno = BLivorno利沃诺
	f.利沃诺Livorno.SetParent(f)

	f.比萨Pisa = BPisa比萨
	f.比萨Pisa.SetParent(f)

	f.费拉约港Portoferraio = BPortoferraio费拉约港
	f.费拉约港Portoferraio.SetParent(f)

	f.圣米尼亚托Sanminiato = BSanminiato圣米尼亚托
	f.圣米尼亚托Sanminiato.SetParent(f)

	f.维科皮萨诺Vicopisano = BVicopisano维科皮萨诺
	f.维科皮萨诺Vicopisano.SetParent(f)

	f.沃尔泰拉Volterra = BVolterra沃尔泰拉
	f.沃尔泰拉Volterra.SetParent(f)

}
