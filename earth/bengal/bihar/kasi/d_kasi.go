package kasi

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/kasi/chunar"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/kasi/jaunpur"
	"github.com/thalesfu/CK2Commands/earth/bengal/bihar/kasi/varanasi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KasiDuke interface {
    feud.Duke
    CChunar枢那罗() 	chunar.ChunarCounty
    CJaunpur善补罗() 	jaunpur.JaunpurCounty
    CVaranasi波罗奈() 	varanasi.VaranasiCounty
}

type 迦尸KasiDuke struct {
	feud.BaseDuke
	枢那罗Chunar 	chunar.ChunarCounty
	善补罗Jaunpur 	jaunpur.JaunpurCounty
	波罗奈Varanasi 	varanasi.VaranasiCounty
}

func (d *迦尸KasiDuke) CChunar枢那罗() chunar.ChunarCounty {
	return d.枢那罗Chunar
}
    
func (d *迦尸KasiDuke) CJaunpur善补罗() jaunpur.JaunpurCounty {
	return d.善补罗Jaunpur
}
    
func (d *迦尸KasiDuke) CVaranasi波罗奈() varanasi.VaranasiCounty {
	return d.波罗奈Varanasi
}
    
var DKasi迦尸 KasiDuke = &迦尸KasiDuke{}

func init() {
	f := DKasi迦尸.(*迦尸KasiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kasi",
		TitleName: "迦尸",
		TitleCode: "d_kasi",
		Counties:  map[string]feud.County{},
	}

	f.枢那罗Chunar = chunar.CChunar枢那罗
	f.枢那罗Chunar.SetParent(f)
	
	f.善补罗Jaunpur = jaunpur.CJaunpur善补罗
	f.善补罗Jaunpur.SetParent(f)
	
	f.波罗奈Varanasi = varanasi.CVaranasi波罗奈
	f.波罗奈Varanasi.SetParent(f)
	
}
