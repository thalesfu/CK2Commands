package kuznetsk

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KuznetskCounty interface {
	feud.County
	BBelogorsk别洛戈尔斯克() feud.Barony
	BKuznetsk库兹涅茨克() feud.Barony
	BNaryk纳雷克() feud.Barony
}

type 库兹涅茨克KuznetskCounty struct {
	feud.BaseCounty
	别洛戈尔斯克Belogorsk feud.Barony
	库兹涅茨克Kuznetsk   feud.Barony
	纳雷克Naryk        feud.Barony
}

func (c *库兹涅茨克KuznetskCounty) BBelogorsk别洛戈尔斯克() feud.Barony {
	return c.别洛戈尔斯克Belogorsk
}

func (c *库兹涅茨克KuznetskCounty) BKuznetsk库兹涅茨克() feud.Barony {
	return c.库兹涅茨克Kuznetsk
}

func (c *库兹涅茨克KuznetskCounty) BNaryk纳雷克() feud.Barony {
	return c.纳雷克Naryk
}

var CKuznetsk库兹涅茨克 KuznetskCounty = &库兹涅茨克KuznetskCounty{}

func init() {
	f := CKuznetsk库兹涅茨克.(*库兹涅茨克KuznetskCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1901",
		Title:     "kuznetsk",
		TitleName: "库兹涅茨克",
		TitleCode: "c_kuznetsk",
		Baronies:  map[string]feud.Barony{},
	}

	f.别洛戈尔斯克Belogorsk = BBelogorsk别洛戈尔斯克
	f.别洛戈尔斯克Belogorsk.SetParent(f)

	f.库兹涅茨克Kuznetsk = BKuznetsk库兹涅茨克
	f.库兹涅茨克Kuznetsk.SetParent(f)

	f.纳雷克Naryk = BNaryk纳雷克
	f.纳雷克Naryk.SetParent(f)

}
