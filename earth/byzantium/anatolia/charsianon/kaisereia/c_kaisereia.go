package kaisereia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KaisereiaCounty interface {
	feud.County
	BDobada多巴达() feud.Barony
	BKaisereia凯撒利亚() feud.Barony
	BMasaka马萨卡() feud.Barony
	BMisti米斯蒂() feud.Barony
	BSariz萨勒兹() feud.Barony
	BTalas塔拉斯() feud.Barony
	BVenessa维尼萨() feud.Barony
	BZoropassos佐罗帕索斯() feud.Barony
}

type 凯撒利亚KaisereiaCounty struct {
	feud.BaseCounty
	多巴达Dobada       feud.Barony
	凯撒利亚Kaisereia   feud.Barony
	马萨卡Masaka       feud.Barony
	米斯蒂Misti        feud.Barony
	萨勒兹Sariz        feud.Barony
	塔拉斯Talas        feud.Barony
	维尼萨Venessa      feud.Barony
	佐罗帕索斯Zoropassos feud.Barony
}

func (c *凯撒利亚KaisereiaCounty) BDobada多巴达() feud.Barony {
	return c.多巴达Dobada
}

func (c *凯撒利亚KaisereiaCounty) BKaisereia凯撒利亚() feud.Barony {
	return c.凯撒利亚Kaisereia
}

func (c *凯撒利亚KaisereiaCounty) BMasaka马萨卡() feud.Barony {
	return c.马萨卡Masaka
}

func (c *凯撒利亚KaisereiaCounty) BMisti米斯蒂() feud.Barony {
	return c.米斯蒂Misti
}

func (c *凯撒利亚KaisereiaCounty) BSariz萨勒兹() feud.Barony {
	return c.萨勒兹Sariz
}

func (c *凯撒利亚KaisereiaCounty) BTalas塔拉斯() feud.Barony {
	return c.塔拉斯Talas
}

func (c *凯撒利亚KaisereiaCounty) BVenessa维尼萨() feud.Barony {
	return c.维尼萨Venessa
}

func (c *凯撒利亚KaisereiaCounty) BZoropassos佐罗帕索斯() feud.Barony {
	return c.佐罗帕索斯Zoropassos
}

var CKaisereia凯撒利亚 KaisereiaCounty = &凯撒利亚KaisereiaCounty{}

func init() {
	f := CKaisereia凯撒利亚.(*凯撒利亚KaisereiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "737",
		Title:     "kaisereia",
		TitleName: "凯撒利亚",
		TitleCode: "c_kaisereia",
		Baronies:  map[string]feud.Barony{},
	}

	f.多巴达Dobada = BDobada多巴达
	f.多巴达Dobada.SetParent(f)

	f.凯撒利亚Kaisereia = BKaisereia凯撒利亚
	f.凯撒利亚Kaisereia.SetParent(f)

	f.马萨卡Masaka = BMasaka马萨卡
	f.马萨卡Masaka.SetParent(f)

	f.米斯蒂Misti = BMisti米斯蒂
	f.米斯蒂Misti.SetParent(f)

	f.萨勒兹Sariz = BSariz萨勒兹
	f.萨勒兹Sariz.SetParent(f)

	f.塔拉斯Talas = BTalas塔拉斯
	f.塔拉斯Talas.SetParent(f)

	f.维尼萨Venessa = BVenessa维尼萨
	f.维尼萨Venessa.SetParent(f)

	f.佐罗帕索斯Zoropassos = BZoropassos佐罗帕索斯
	f.佐罗帕索斯Zoropassos.SetParent(f)

}
