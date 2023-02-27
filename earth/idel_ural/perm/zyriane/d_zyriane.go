package zyriane

import (
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/zyriane/sedyu"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/zyriane/vychegda"
	"github.com/thalesfu/CK2Commands/earth/idel_ural/perm/zyriane/yolva"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ZyrianeDuke interface {
    feud.Duke
    CSedyu谢季尤() 	sedyu.SedyuCounty
    CVychegda维切格达() 	vychegda.VychegdaCounty
    CYolva维姆() 	yolva.YolvaCounty
}

type 济良ZyrianeDuke struct {
	feud.BaseDuke
	谢季尤Sedyu 	sedyu.SedyuCounty
	维切格达Vychegda 	vychegda.VychegdaCounty
	维姆Yolva 	yolva.YolvaCounty
}

func (d *济良ZyrianeDuke) CSedyu谢季尤() sedyu.SedyuCounty {
	return d.谢季尤Sedyu
}
    
func (d *济良ZyrianeDuke) CVychegda维切格达() vychegda.VychegdaCounty {
	return d.维切格达Vychegda
}
    
func (d *济良ZyrianeDuke) CYolva维姆() yolva.YolvaCounty {
	return d.维姆Yolva
}
    
var DZyriane济良 ZyrianeDuke = &济良ZyrianeDuke{}

func init() {
	f := DZyriane济良.(*济良ZyrianeDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "zyriane",
		TitleName: "济良",
		TitleCode: "d_zyriane",
		Counties:  map[string]feud.County{},
	}

	f.谢季尤Sedyu = sedyu.CSedyu谢季尤
	f.谢季尤Sedyu.SetParent(f)
	
	f.维切格达Vychegda = vychegda.CVychegda维切格达
	f.维切格达Vychegda.SetParent(f)
	
	f.维姆Yolva = yolva.CYolva维姆
	f.维姆Yolva.SetParent(f)
	
}
