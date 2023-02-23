package seres

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SeresCounty interface {
	feud.County
	BBisaltia比萨尔提亚() feud.Barony
	BChristoupolis赫里斯图波利斯() feud.Barony
	BDrama兹拉马() feud.Barony
	BMelnik梅尔尼克() feud.Barony
	BPhilippi腓立比() feud.Barony
	BSeres塞雷斯() feud.Barony
	BSiderokastron锡季罗卡斯特龙() feud.Barony
}

type 塞雷斯SeresCounty struct {
	feud.BaseCounty
	比萨尔提亚Bisaltia        feud.Barony
	赫里斯图波利斯Christoupolis feud.Barony
	兹拉马Drama             feud.Barony
	梅尔尼克Melnik           feud.Barony
	腓立比Philippi          feud.Barony
	塞雷斯Seres             feud.Barony
	锡季罗卡斯特龙Siderokastron feud.Barony
}

func (c *塞雷斯SeresCounty) BBisaltia比萨尔提亚() feud.Barony {
	return c.比萨尔提亚Bisaltia
}

func (c *塞雷斯SeresCounty) BChristoupolis赫里斯图波利斯() feud.Barony {
	return c.赫里斯图波利斯Christoupolis
}

func (c *塞雷斯SeresCounty) BDrama兹拉马() feud.Barony {
	return c.兹拉马Drama
}

func (c *塞雷斯SeresCounty) BMelnik梅尔尼克() feud.Barony {
	return c.梅尔尼克Melnik
}

func (c *塞雷斯SeresCounty) BPhilippi腓立比() feud.Barony {
	return c.腓立比Philippi
}

func (c *塞雷斯SeresCounty) BSeres塞雷斯() feud.Barony {
	return c.塞雷斯Seres
}

func (c *塞雷斯SeresCounty) BSiderokastron锡季罗卡斯特龙() feud.Barony {
	return c.锡季罗卡斯特龙Siderokastron
}

var CSeres塞雷斯 SeresCounty = &塞雷斯SeresCounty{}

func init() {
	f := CSeres塞雷斯.(*塞雷斯SeresCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1882",
		Title:     "seres",
		TitleName: "塞雷斯",
		TitleCode: "c_seres",
		Baronies:  map[string]feud.Barony{},
	}

	f.比萨尔提亚Bisaltia = BBisaltia比萨尔提亚
	f.比萨尔提亚Bisaltia.SetParent(f)

	f.赫里斯图波利斯Christoupolis = BChristoupolis赫里斯图波利斯
	f.赫里斯图波利斯Christoupolis.SetParent(f)

	f.兹拉马Drama = BDrama兹拉马
	f.兹拉马Drama.SetParent(f)

	f.梅尔尼克Melnik = BMelnik梅尔尼克
	f.梅尔尼克Melnik.SetParent(f)

	f.腓立比Philippi = BPhilippi腓立比
	f.腓立比Philippi.SetParent(f)

	f.塞雷斯Seres = BSeres塞雷斯
	f.塞雷斯Seres.SetParent(f)

	f.锡季罗卡斯特龙Siderokastron = BSiderokastron锡季罗卡斯特龙
	f.锡季罗卡斯特龙Siderokastron.SetParent(f)

}
