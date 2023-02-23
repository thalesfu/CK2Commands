package bukhara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BukharaCounty interface {
	feud.County
	BAyrybaba艾里巴巴() feud.Barony
	BBukhara布哈拉() feud.Barony
	BChardjul查尔朱() feud.Barony
	BGizhduvan吉日杜万() feud.Barony
	BKarkuh卡尔库赫() feud.Barony
	BKogon科贡() feud.Barony
	BVabkent恰什肯特() feud.Barony
}

type 布哈拉BukharaCounty struct {
	feud.BaseCounty
	艾里巴巴Ayrybaba  feud.Barony
	布哈拉Bukhara    feud.Barony
	查尔朱Chardjul   feud.Barony
	吉日杜万Gizhduvan feud.Barony
	卡尔库赫Karkuh    feud.Barony
	科贡Kogon       feud.Barony
	恰什肯特Vabkent   feud.Barony
}

func (c *布哈拉BukharaCounty) BAyrybaba艾里巴巴() feud.Barony {
	return c.艾里巴巴Ayrybaba
}

func (c *布哈拉BukharaCounty) BBukhara布哈拉() feud.Barony {
	return c.布哈拉Bukhara
}

func (c *布哈拉BukharaCounty) BChardjul查尔朱() feud.Barony {
	return c.查尔朱Chardjul
}

func (c *布哈拉BukharaCounty) BGizhduvan吉日杜万() feud.Barony {
	return c.吉日杜万Gizhduvan
}

func (c *布哈拉BukharaCounty) BKarkuh卡尔库赫() feud.Barony {
	return c.卡尔库赫Karkuh
}

func (c *布哈拉BukharaCounty) BKogon科贡() feud.Barony {
	return c.科贡Kogon
}

func (c *布哈拉BukharaCounty) BVabkent恰什肯特() feud.Barony {
	return c.恰什肯特Vabkent
}

var CBukhara布哈拉 BukharaCounty = &布哈拉BukharaCounty{}

func init() {
	f := CBukhara布哈拉.(*布哈拉BukharaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "628",
		Title:     "bukhara",
		TitleName: "布哈拉",
		TitleCode: "c_bukhara",
		Baronies:  map[string]feud.Barony{},
	}

	f.艾里巴巴Ayrybaba = BAyrybaba艾里巴巴
	f.艾里巴巴Ayrybaba.SetParent(f)

	f.布哈拉Bukhara = BBukhara布哈拉
	f.布哈拉Bukhara.SetParent(f)

	f.查尔朱Chardjul = BChardjul查尔朱
	f.查尔朱Chardjul.SetParent(f)

	f.吉日杜万Gizhduvan = BGizhduvan吉日杜万
	f.吉日杜万Gizhduvan.SetParent(f)

	f.卡尔库赫Karkuh = BKarkuh卡尔库赫
	f.卡尔库赫Karkuh.SetParent(f)

	f.科贡Kogon = BKogon科贡
	f.科贡Kogon.SetParent(f)

	f.恰什肯特Vabkent = BVabkent恰什肯特
	f.恰什肯特Vabkent.SetParent(f)

}
