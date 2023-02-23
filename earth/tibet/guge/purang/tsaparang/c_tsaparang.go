package tsaparang

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TsaparangCounty interface {
	feud.County
	BBoyi博依() feud.Barony
	BGartok噶大克() feud.Barony
	BMangga芒嘎() feud.Barony
	BMangnang芒囊() feud.Barony
	BTholing托林() feud.Barony
	BTholinggompa托林贡巴() feud.Barony
	BTsaparang札布让() feud.Barony
}

type 札布让TsaparangCounty struct {
	feud.BaseCounty
	博依Boyi           feud.Barony
	噶大克Gartok        feud.Barony
	芒嘎Mangga         feud.Barony
	芒囊Mangnang       feud.Barony
	托林Tholing        feud.Barony
	托林贡巴Tholinggompa feud.Barony
	札布让Tsaparang     feud.Barony
}

func (c *札布让TsaparangCounty) BBoyi博依() feud.Barony {
	return c.博依Boyi
}

func (c *札布让TsaparangCounty) BGartok噶大克() feud.Barony {
	return c.噶大克Gartok
}

func (c *札布让TsaparangCounty) BMangga芒嘎() feud.Barony {
	return c.芒嘎Mangga
}

func (c *札布让TsaparangCounty) BMangnang芒囊() feud.Barony {
	return c.芒囊Mangnang
}

func (c *札布让TsaparangCounty) BTholing托林() feud.Barony {
	return c.托林Tholing
}

func (c *札布让TsaparangCounty) BTholinggompa托林贡巴() feud.Barony {
	return c.托林贡巴Tholinggompa
}

func (c *札布让TsaparangCounty) BTsaparang札布让() feud.Barony {
	return c.札布让Tsaparang
}

var CTsaparang札布让 TsaparangCounty = &札布让TsaparangCounty{}

func init() {
	f := CTsaparang札布让.(*札布让TsaparangCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1486",
		Title:     "tsaparang",
		TitleName: "札布让",
		TitleCode: "c_tsaparang",
		Baronies:  map[string]feud.Barony{},
	}

	f.博依Boyi = BBoyi博依
	f.博依Boyi.SetParent(f)

	f.噶大克Gartok = BGartok噶大克
	f.噶大克Gartok.SetParent(f)

	f.芒嘎Mangga = BMangga芒嘎
	f.芒嘎Mangga.SetParent(f)

	f.芒囊Mangnang = BMangnang芒囊
	f.芒囊Mangnang.SetParent(f)

	f.托林Tholing = BTholing托林
	f.托林Tholing.SetParent(f)

	f.托林贡巴Tholinggompa = BTholinggompa托林贡巴
	f.托林贡巴Tholinggompa.SetParent(f)

	f.札布让Tsaparang = BTsaparang札布让
	f.札布让Tsaparang.SetParent(f)

}
