package lhasa

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/lhasa/kunggar"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/lhasa/lhasa"
	"github.com/thalesfu/CK2Commands/earth/tibet/yarlung/lhasa/lhunzhub"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LhasaDuke interface {
    feud.Duke
    CKunggar工卡() 	kunggar.KunggarCounty
    CLhasa逻些() 	lhasa.LhasaCounty
    CLhunzhub林周() 	lhunzhub.LhunzhubCounty
}

type 逻些LhasaDuke struct {
	feud.BaseDuke
	工卡Kunggar 	kunggar.KunggarCounty
	逻些Lhasa 	lhasa.LhasaCounty
	林周Lhunzhub 	lhunzhub.LhunzhubCounty
}

func (d *逻些LhasaDuke) CKunggar工卡() kunggar.KunggarCounty {
	return d.工卡Kunggar
}
    
func (d *逻些LhasaDuke) CLhasa逻些() lhasa.LhasaCounty {
	return d.逻些Lhasa
}
    
func (d *逻些LhasaDuke) CLhunzhub林周() lhunzhub.LhunzhubCounty {
	return d.林周Lhunzhub
}
    
var DLhasa逻些 LhasaDuke = &逻些LhasaDuke{}

func init() {
	f := DLhasa逻些.(*逻些LhasaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lhasa",
		TitleName: "逻些",
		TitleCode: "d_lhasa",
		Counties:  map[string]feud.County{},
	}

	f.工卡Kunggar = kunggar.CKunggar工卡
	f.工卡Kunggar.SetParent(f)
	
	f.逻些Lhasa = lhasa.CLhasa逻些
	f.逻些Lhasa.SetParent(f)
	
	f.林周Lhunzhub = lhunzhub.CLhunzhub林周
	f.林周Lhunzhub.SetParent(f)
	
}
