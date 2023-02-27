package turkestan

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/turkestan/aral"
	"github.com/thalesfu/CK2Commands/earth/turkestan/turkestan/turkestan/turkestan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type TurkestanDuke interface {
    feud.Duke
    CAral库坦布拉克() 	aral.AralCounty
    CTurkestan阿拉尔() 	turkestan.TurkestanCounty
}

type 阿拉尔TurkestanDuke struct {
	feud.BaseDuke
	库坦布拉克Aral 	aral.AralCounty
	阿拉尔Turkestan 	turkestan.TurkestanCounty
}

func (d *阿拉尔TurkestanDuke) CAral库坦布拉克() aral.AralCounty {
	return d.库坦布拉克Aral
}
    
func (d *阿拉尔TurkestanDuke) CTurkestan阿拉尔() turkestan.TurkestanCounty {
	return d.阿拉尔Turkestan
}
    
var DTurkestan阿拉尔 TurkestanDuke = &阿拉尔TurkestanDuke{}

func init() {
	f := DTurkestan阿拉尔.(*阿拉尔TurkestanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "turkestan",
		TitleName: "阿拉尔",
		TitleCode: "d_turkestan",
		Counties:  map[string]feud.County{},
	}

	f.库坦布拉克Aral = aral.CAral库坦布拉克
	f.库坦布拉克Aral.SetParent(f)
	
	f.阿拉尔Turkestan = turkestan.CTurkestan阿拉尔
	f.阿拉尔Turkestan.SetParent(f)
	
}
