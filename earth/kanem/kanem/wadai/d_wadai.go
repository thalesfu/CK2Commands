package wadai

import (
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/wadai/ennedi"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/wadai/fitri"
	"github.com/thalesfu/CK2Commands/earth/kanem/kanem/wadai/wadai"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type WadaiDuke interface {
    feud.Duke
    CEnnedi恩内迪() 	ennedi.EnnediCounty
    CFitri菲特里() 	fitri.FitriCounty
    CWadai瓦达伊() 	wadai.WadaiCounty
}

type 瓦达伊WadaiDuke struct {
	feud.BaseDuke
	恩内迪Ennedi 	ennedi.EnnediCounty
	菲特里Fitri 	fitri.FitriCounty
	瓦达伊Wadai 	wadai.WadaiCounty
}

func (d *瓦达伊WadaiDuke) CEnnedi恩内迪() ennedi.EnnediCounty {
	return d.恩内迪Ennedi
}
    
func (d *瓦达伊WadaiDuke) CFitri菲特里() fitri.FitriCounty {
	return d.菲特里Fitri
}
    
func (d *瓦达伊WadaiDuke) CWadai瓦达伊() wadai.WadaiCounty {
	return d.瓦达伊Wadai
}
    
var DWadai瓦达伊 WadaiDuke = &瓦达伊WadaiDuke{}

func init() {
	f := DWadai瓦达伊.(*瓦达伊WadaiDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "wadai",
		TitleName: "瓦达伊",
		TitleCode: "d_wadai",
		Counties:  map[string]feud.County{},
	}

	f.恩内迪Ennedi = ennedi.CEnnedi恩内迪
	f.恩内迪Ennedi.SetParent(f)
	
	f.菲特里Fitri = fitri.CFitri菲特里
	f.菲特里Fitri.SetParent(f)
	
	f.瓦达伊Wadai = wadai.CWadai瓦达伊
	f.瓦达伊Wadai.SetParent(f)
	
}
