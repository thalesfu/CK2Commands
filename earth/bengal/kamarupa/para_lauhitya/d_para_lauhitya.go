package para_lauhitya

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/para_lauhitya/dimapur"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/para_lauhitya/goalpara"
	"github.com/thalesfu/CK2Commands/earth/bengal/kamarupa/para_lauhitya/srihatta"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Para_lauhityaDuke interface {
    feud.Duke
    CDimapur提摩补罗() 	dimapur.DimapurCounty
    CGoalpara乔阿罗波罗() 	goalpara.GoalparaCounty
    CSrihatta室利诃吒() 	srihatta.SrihattaCounty
}

type 波罗劳希底耶Para_lauhityaDuke struct {
	feud.BaseDuke
	提摩补罗Dimapur 	dimapur.DimapurCounty
	乔阿罗波罗Goalpara 	goalpara.GoalparaCounty
	室利诃吒Srihatta 	srihatta.SrihattaCounty
}

func (d *波罗劳希底耶Para_lauhityaDuke) CDimapur提摩补罗() dimapur.DimapurCounty {
	return d.提摩补罗Dimapur
}
    
func (d *波罗劳希底耶Para_lauhityaDuke) CGoalpara乔阿罗波罗() goalpara.GoalparaCounty {
	return d.乔阿罗波罗Goalpara
}
    
func (d *波罗劳希底耶Para_lauhityaDuke) CSrihatta室利诃吒() srihatta.SrihattaCounty {
	return d.室利诃吒Srihatta
}
    
var DPara_lauhitya波罗劳希底耶 Para_lauhityaDuke = &波罗劳希底耶Para_lauhityaDuke{}

func init() {
	f := DPara_lauhitya波罗劳希底耶.(*波罗劳希底耶Para_lauhityaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "para_lauhitya",
		TitleName: "波罗劳希底耶",
		TitleCode: "d_para_lauhitya",
		Counties:  map[string]feud.County{},
	}

	f.提摩补罗Dimapur = dimapur.CDimapur提摩补罗
	f.提摩补罗Dimapur.SetParent(f)
	
	f.乔阿罗波罗Goalpara = goalpara.CGoalpara乔阿罗波罗
	f.乔阿罗波罗Goalpara.SetParent(f)
	
	f.室利诃吒Srihatta = srihatta.CSrihatta室利诃吒
	f.室利诃吒Srihatta.SetParent(f)
	
}
