package ngari

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ngari/coqen"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ngari/gar"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ngari/gerze"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ngari/kunlun"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ngari/rutog"
	"github.com/thalesfu/CK2Commands/earth/tibet/guge/ngari/tsakha"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NgariDuke interface {
	feud.Duke
	CCoqen措勤() coqen.CoqenCounty
	CGar噶尔() gar.GarCounty
	CGerze改则() gerze.GerzeCounty
	CKunlun昆仑() kunlun.KunlunCounty
	CRutog日托() rutog.RutogCounty
	CTsakha茶卡() tsakha.TsakhaCounty
}

type 纳里NgariDuke struct {
	feud.BaseDuke
	措勤Coqen  coqen.CoqenCounty
	噶尔Gar    gar.GarCounty
	改则Gerze  gerze.GerzeCounty
	昆仑Kunlun kunlun.KunlunCounty
	日托Rutog  rutog.RutogCounty
	茶卡Tsakha tsakha.TsakhaCounty
}

func (d *纳里NgariDuke) CCoqen措勤() coqen.CoqenCounty {
	return d.措勤Coqen
}

func (d *纳里NgariDuke) CGar噶尔() gar.GarCounty {
	return d.噶尔Gar
}

func (d *纳里NgariDuke) CGerze改则() gerze.GerzeCounty {
	return d.改则Gerze
}

func (d *纳里NgariDuke) CKunlun昆仑() kunlun.KunlunCounty {
	return d.昆仑Kunlun
}

func (d *纳里NgariDuke) CRutog日托() rutog.RutogCounty {
	return d.日托Rutog
}

func (d *纳里NgariDuke) CTsakha茶卡() tsakha.TsakhaCounty {
	return d.茶卡Tsakha
}

var DNgari纳里 NgariDuke = &纳里NgariDuke{}

func init() {
	f := DNgari纳里.(*纳里NgariDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ngari",
		TitleName: "纳里",
		TitleCode: "d_ngari",
		Counties:  map[string]feud.County{},
	}

	f.措勤Coqen = coqen.CCoqen措勤
	f.措勤Coqen.SetParent(f)

	f.噶尔Gar = gar.CGar噶尔
	f.噶尔Gar.SetParent(f)

	f.改则Gerze = gerze.CGerze改则
	f.改则Gerze.SetParent(f)

	f.昆仑Kunlun = kunlun.CKunlun昆仑
	f.昆仑Kunlun.SetParent(f)

	f.日托Rutog = rutog.CRutog日托
	f.日托Rutog.SetParent(f)

	f.茶卡Tsakha = tsakha.CTsakha茶卡
	f.茶卡Tsakha.SetParent(f)

}
