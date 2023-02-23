package tadmekka

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TadmekkaCounty interface {
	feud.County
	BDjila吉拉() feud.Barony
	BDongossela东戈塞拉() feud.Barony
	BEskako埃斯卡科() feud.Barony
	BEssako艾萨科() feud.Barony
	BKidako基达科() feud.Barony
	BKidal基达勒() feud.Barony
	BTadmekka塔得迈卡() feud.Barony
}

type 塔得迈卡TadmekkaCounty struct {
	feud.BaseCounty
	吉拉Djila        feud.Barony
	东戈塞拉Dongossela feud.Barony
	埃斯卡科Eskako     feud.Barony
	艾萨科Essako      feud.Barony
	基达科Kidako      feud.Barony
	基达勒Kidal       feud.Barony
	塔得迈卡Tadmekka   feud.Barony
}

func (c *塔得迈卡TadmekkaCounty) BDjila吉拉() feud.Barony {
	return c.吉拉Djila
}

func (c *塔得迈卡TadmekkaCounty) BDongossela东戈塞拉() feud.Barony {
	return c.东戈塞拉Dongossela
}

func (c *塔得迈卡TadmekkaCounty) BEskako埃斯卡科() feud.Barony {
	return c.埃斯卡科Eskako
}

func (c *塔得迈卡TadmekkaCounty) BEssako艾萨科() feud.Barony {
	return c.艾萨科Essako
}

func (c *塔得迈卡TadmekkaCounty) BKidako基达科() feud.Barony {
	return c.基达科Kidako
}

func (c *塔得迈卡TadmekkaCounty) BKidal基达勒() feud.Barony {
	return c.基达勒Kidal
}

func (c *塔得迈卡TadmekkaCounty) BTadmekka塔得迈卡() feud.Barony {
	return c.塔得迈卡Tadmekka
}

var CTadmekka塔得迈卡 TadmekkaCounty = &塔得迈卡TadmekkaCounty{}

func init() {
	f := CTadmekka塔得迈卡.(*塔得迈卡TadmekkaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "922",
		Title:     "tadmekka",
		TitleName: "塔得迈卡",
		TitleCode: "c_tadmekka",
		Baronies:  map[string]feud.Barony{},
	}

	f.吉拉Djila = BDjila吉拉
	f.吉拉Djila.SetParent(f)

	f.东戈塞拉Dongossela = BDongossela东戈塞拉
	f.东戈塞拉Dongossela.SetParent(f)

	f.埃斯卡科Eskako = BEskako埃斯卡科
	f.埃斯卡科Eskako.SetParent(f)

	f.艾萨科Essako = BEssako艾萨科
	f.艾萨科Essako.SetParent(f)

	f.基达科Kidako = BKidako基达科
	f.基达科Kidako.SetParent(f)

	f.基达勒Kidal = BKidal基达勒
	f.基达勒Kidal.SetParent(f)

	f.塔得迈卡Tadmekka = BTadmekka塔得迈卡
	f.塔得迈卡Tadmekka.SetParent(f)

}
