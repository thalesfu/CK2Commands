package imeretia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ImeretiaCounty interface {
	feud.County
	BBagrati巴格拉季() feud.Barony
	BGeguti戈古提() feud.Barony
	BGhelati盖拉蒂() feud.Barony
	BKhoni霍尼() feud.Barony
	BKutaisi库塔伊西() feud.Barony
	BLentekhi伦泰希() feud.Barony
	BMotsameta莫察梅塔() feud.Barony
	BTsikhegoji齐赫戈吉() feud.Barony
}

type 伊梅列季亚ImeretiaCounty struct {
	feud.BaseCounty
	巴格拉季Bagrati    feud.Barony
	戈古提Geguti      feud.Barony
	盖拉蒂Ghelati     feud.Barony
	霍尼Khoni        feud.Barony
	库塔伊西Kutaisi    feud.Barony
	伦泰希Lentekhi    feud.Barony
	莫察梅塔Motsameta  feud.Barony
	齐赫戈吉Tsikhegoji feud.Barony
}

func (c *伊梅列季亚ImeretiaCounty) BBagrati巴格拉季() feud.Barony {
	return c.巴格拉季Bagrati
}

func (c *伊梅列季亚ImeretiaCounty) BGeguti戈古提() feud.Barony {
	return c.戈古提Geguti
}

func (c *伊梅列季亚ImeretiaCounty) BGhelati盖拉蒂() feud.Barony {
	return c.盖拉蒂Ghelati
}

func (c *伊梅列季亚ImeretiaCounty) BKhoni霍尼() feud.Barony {
	return c.霍尼Khoni
}

func (c *伊梅列季亚ImeretiaCounty) BKutaisi库塔伊西() feud.Barony {
	return c.库塔伊西Kutaisi
}

func (c *伊梅列季亚ImeretiaCounty) BLentekhi伦泰希() feud.Barony {
	return c.伦泰希Lentekhi
}

func (c *伊梅列季亚ImeretiaCounty) BMotsameta莫察梅塔() feud.Barony {
	return c.莫察梅塔Motsameta
}

func (c *伊梅列季亚ImeretiaCounty) BTsikhegoji齐赫戈吉() feud.Barony {
	return c.齐赫戈吉Tsikhegoji
}

var CImeretia伊梅列季亚 ImeretiaCounty = &伊梅列季亚ImeretiaCounty{}

func init() {
	f := CImeretia伊梅列季亚.(*伊梅列季亚ImeretiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "601",
		Title:     "imeretia",
		TitleName: "伊梅列季亚",
		TitleCode: "c_imeretia",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴格拉季Bagrati = BBagrati巴格拉季
	f.巴格拉季Bagrati.SetParent(f)

	f.戈古提Geguti = BGeguti戈古提
	f.戈古提Geguti.SetParent(f)

	f.盖拉蒂Ghelati = BGhelati盖拉蒂
	f.盖拉蒂Ghelati.SetParent(f)

	f.霍尼Khoni = BKhoni霍尼
	f.霍尼Khoni.SetParent(f)

	f.库塔伊西Kutaisi = BKutaisi库塔伊西
	f.库塔伊西Kutaisi.SetParent(f)

	f.伦泰希Lentekhi = BLentekhi伦泰希
	f.伦泰希Lentekhi.SetParent(f)

	f.莫察梅塔Motsameta = BMotsameta莫察梅塔
	f.莫察梅塔Motsameta.SetParent(f)

	f.齐赫戈吉Tsikhegoji = BTsikhegoji齐赫戈吉
	f.齐赫戈吉Tsikhegoji.SetParent(f)

}
