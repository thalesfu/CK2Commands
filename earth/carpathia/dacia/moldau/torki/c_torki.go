package torki

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TorkiCounty interface {
	feud.County
	BBaia巴亚() feud.Barony
	BBogdana博格丹纳() feud.Barony
	BMoldovita摩尔多维察() feud.Barony
	BPutna普特纳() feud.Barony
	BRadauti勒德乌齐() feud.Barony
	BSiret锡雷特() feud.Barony
	BSuceava苏恰瓦() feud.Barony
}

type 苏恰瓦TorkiCounty struct {
	feud.BaseCounty
	巴亚Baia         feud.Barony
	博格丹纳Bogdana    feud.Barony
	摩尔多维察Moldovita feud.Barony
	普特纳Putna       feud.Barony
	勒德乌齐Radauti    feud.Barony
	锡雷特Siret       feud.Barony
	苏恰瓦Suceava     feud.Barony
}

func (c *苏恰瓦TorkiCounty) BBaia巴亚() feud.Barony {
	return c.巴亚Baia
}

func (c *苏恰瓦TorkiCounty) BBogdana博格丹纳() feud.Barony {
	return c.博格丹纳Bogdana
}

func (c *苏恰瓦TorkiCounty) BMoldovita摩尔多维察() feud.Barony {
	return c.摩尔多维察Moldovita
}

func (c *苏恰瓦TorkiCounty) BPutna普特纳() feud.Barony {
	return c.普特纳Putna
}

func (c *苏恰瓦TorkiCounty) BRadauti勒德乌齐() feud.Barony {
	return c.勒德乌齐Radauti
}

func (c *苏恰瓦TorkiCounty) BSiret锡雷特() feud.Barony {
	return c.锡雷特Siret
}

func (c *苏恰瓦TorkiCounty) BSuceava苏恰瓦() feud.Barony {
	return c.苏恰瓦Suceava
}

var CTorki苏恰瓦 TorkiCounty = &苏恰瓦TorkiCounty{}

func init() {
	f := CTorki苏恰瓦.(*苏恰瓦TorkiCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "545",
		Title:     "torki",
		TitleName: "苏恰瓦",
		TitleCode: "c_torki",
		Baronies:  map[string]feud.Barony{},
	}

	f.巴亚Baia = BBaia巴亚
	f.巴亚Baia.SetParent(f)

	f.博格丹纳Bogdana = BBogdana博格丹纳
	f.博格丹纳Bogdana.SetParent(f)

	f.摩尔多维察Moldovita = BMoldovita摩尔多维察
	f.摩尔多维察Moldovita.SetParent(f)

	f.普特纳Putna = BPutna普特纳
	f.普特纳Putna.SetParent(f)

	f.勒德乌齐Radauti = BRadauti勒德乌齐
	f.勒德乌齐Radauti.SetParent(f)

	f.锡雷特Siret = BSiret锡雷特
	f.锡雷特Siret.SetParent(f)

	f.苏恰瓦Suceava = BSuceava苏恰瓦
	f.苏恰瓦Suceava.SetParent(f)

}
