package lahore

import (
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/lahore/bhera"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/lahore/dipalpur"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/lahore/lahur"
	"github.com/thalesfu/CK2Commands/earth/rajastan/punjab/lahore/shorkot"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LahoreDuke interface {
    feud.Duke
    CBhera吠罗() 	bhera.BheraCounty
    CDipalpur提波罗补罗() 	dipalpur.DipalpurCounty
    CLahur拉合尔() 	lahur.LahurCounty
    CShorkot绍科特() 	shorkot.ShorkotCounty
}

type 拉合尔LahoreDuke struct {
	feud.BaseDuke
	吠罗Bhera 	bhera.BheraCounty
	提波罗补罗Dipalpur 	dipalpur.DipalpurCounty
	拉合尔Lahur 	lahur.LahurCounty
	绍科特Shorkot 	shorkot.ShorkotCounty
}

func (d *拉合尔LahoreDuke) CBhera吠罗() bhera.BheraCounty {
	return d.吠罗Bhera
}
    
func (d *拉合尔LahoreDuke) CDipalpur提波罗补罗() dipalpur.DipalpurCounty {
	return d.提波罗补罗Dipalpur
}
    
func (d *拉合尔LahoreDuke) CLahur拉合尔() lahur.LahurCounty {
	return d.拉合尔Lahur
}
    
func (d *拉合尔LahoreDuke) CShorkot绍科特() shorkot.ShorkotCounty {
	return d.绍科特Shorkot
}
    
var DLahore拉合尔 LahoreDuke = &拉合尔LahoreDuke{}

func init() {
	f := DLahore拉合尔.(*拉合尔LahoreDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lahore",
		TitleName: "拉合尔",
		TitleCode: "d_lahore",
		Counties:  map[string]feud.County{},
	}

	f.吠罗Bhera = bhera.CBhera吠罗
	f.吠罗Bhera.SetParent(f)
	
	f.提波罗补罗Dipalpur = dipalpur.CDipalpur提波罗补罗
	f.提波罗补罗Dipalpur.SetParent(f)
	
	f.拉合尔Lahur = lahur.CLahur拉合尔
	f.拉合尔Lahur.SetParent(f)
	
	f.绍科特Shorkot = shorkot.CShorkot绍科特
	f.绍科特Shorkot.SetParent(f)
	
}
