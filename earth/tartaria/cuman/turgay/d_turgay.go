package turgay

import (
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/turgay/boqtybay"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/turgay/irgiz"
	"github.com/thalesfu/CK2Commands/earth/tartaria/cuman/turgay/turgay"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurgayDuke interface {
    feud.Duke
    CBoqtybay博克特拜() 	boqtybay.BoqtybayCounty
    CIrgiz伊尔吉兹() 	irgiz.IrgizCounty
    CTurgay图尔盖() 	turgay.TurgayCounty
}

type 图尔盖TurgayDuke struct {
	feud.BaseDuke
	博克特拜Boqtybay 	boqtybay.BoqtybayCounty
	伊尔吉兹Irgiz 	irgiz.IrgizCounty
	图尔盖Turgay 	turgay.TurgayCounty
}

func (d *图尔盖TurgayDuke) CBoqtybay博克特拜() boqtybay.BoqtybayCounty {
	return d.博克特拜Boqtybay
}
    
func (d *图尔盖TurgayDuke) CIrgiz伊尔吉兹() irgiz.IrgizCounty {
	return d.伊尔吉兹Irgiz
}
    
func (d *图尔盖TurgayDuke) CTurgay图尔盖() turgay.TurgayCounty {
	return d.图尔盖Turgay
}
    
var DTurgay图尔盖 TurgayDuke = &图尔盖TurgayDuke{}

func init() {
	f := DTurgay图尔盖.(*图尔盖TurgayDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "turgay",
		TitleName: "图尔盖",
		TitleCode: "d_turgay",
		Counties:  map[string]feud.County{},
	}

	f.博克特拜Boqtybay = boqtybay.CBoqtybay博克特拜
	f.博克特拜Boqtybay.SetParent(f)
	
	f.伊尔吉兹Irgiz = irgiz.CIrgiz伊尔吉兹
	f.伊尔吉兹Irgiz.SetParent(f)
	
	f.图尔盖Turgay = turgay.CTurgay图尔盖
	f.图尔盖Turgay.SetParent(f)
	
}
