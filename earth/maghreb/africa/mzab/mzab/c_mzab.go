package mzab

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MzabCounty interface {
	feud.County
	BBermane伯马吉() feud.Barony
	BGhardaia盖尔达耶() feud.Barony
	BMzab姆扎卜() feud.Barony
	BNoumerat努迈拉特() feud.Barony
	BOuargla乌尔格拉() feud.Barony
	BSebseb塞卜塞卜() feud.Barony
	BZelfana宰勒法纳() feud.Barony
}

type 盖尔达耶MzabCounty struct {
	feud.BaseCounty
	伯马吉Bermane   feud.Barony
	盖尔达耶Ghardaia feud.Barony
	姆扎卜Mzab      feud.Barony
	努迈拉特Noumerat feud.Barony
	乌尔格拉Ouargla  feud.Barony
	塞卜塞卜Sebseb   feud.Barony
	宰勒法纳Zelfana  feud.Barony
}

func (c *盖尔达耶MzabCounty) BBermane伯马吉() feud.Barony {
	return c.伯马吉Bermane
}

func (c *盖尔达耶MzabCounty) BGhardaia盖尔达耶() feud.Barony {
	return c.盖尔达耶Ghardaia
}

func (c *盖尔达耶MzabCounty) BMzab姆扎卜() feud.Barony {
	return c.姆扎卜Mzab
}

func (c *盖尔达耶MzabCounty) BNoumerat努迈拉特() feud.Barony {
	return c.努迈拉特Noumerat
}

func (c *盖尔达耶MzabCounty) BOuargla乌尔格拉() feud.Barony {
	return c.乌尔格拉Ouargla
}

func (c *盖尔达耶MzabCounty) BSebseb塞卜塞卜() feud.Barony {
	return c.塞卜塞卜Sebseb
}

func (c *盖尔达耶MzabCounty) BZelfana宰勒法纳() feud.Barony {
	return c.宰勒法纳Zelfana
}

var CMzab盖尔达耶 MzabCounty = &盖尔达耶MzabCounty{}

func init() {
	f := CMzab盖尔达耶.(*盖尔达耶MzabCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "829",
		Title:     "mzab",
		TitleName: "盖尔达耶",
		TitleCode: "c_mzab",
		Baronies:  map[string]feud.Barony{},
	}

	f.伯马吉Bermane = BBermane伯马吉
	f.伯马吉Bermane.SetParent(f)

	f.盖尔达耶Ghardaia = BGhardaia盖尔达耶
	f.盖尔达耶Ghardaia.SetParent(f)

	f.姆扎卜Mzab = BMzab姆扎卜
	f.姆扎卜Mzab.SetParent(f)

	f.努迈拉特Noumerat = BNoumerat努迈拉特
	f.努迈拉特Noumerat.SetParent(f)

	f.乌尔格拉Ouargla = BOuargla乌尔格拉
	f.乌尔格拉Ouargla.SetParent(f)

	f.塞卜塞卜Sebseb = BSebseb塞卜塞卜
	f.塞卜塞卜Sebseb.SetParent(f)

	f.宰勒法纳Zelfana = BZelfana宰勒法纳
	f.宰勒法纳Zelfana.SetParent(f)

}
