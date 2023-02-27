package munster_germany

import (
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/munster_germany/gottingen"
	"github.com/thalesfu/CK2Commands/earth/germany/saxony/munster_germany/munster"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Munster_germanyDuke interface {
    feud.Duke
    CGottingen帕德博恩() 	gottingen.GottingenCounty
    CMunster明斯特() 	munster.MunsterCounty
}

type 明斯特Munster_germanyDuke struct {
	feud.BaseDuke
	帕德博恩Gottingen 	gottingen.GottingenCounty
	明斯特Munster 	munster.MunsterCounty
}

func (d *明斯特Munster_germanyDuke) CGottingen帕德博恩() gottingen.GottingenCounty {
	return d.帕德博恩Gottingen
}
    
func (d *明斯特Munster_germanyDuke) CMunster明斯特() munster.MunsterCounty {
	return d.明斯特Munster
}
    
var DMunster_germany明斯特 Munster_germanyDuke = &明斯特Munster_germanyDuke{}

func init() {
	f := DMunster_germany明斯特.(*明斯特Munster_germanyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "munster_germany",
		TitleName: "明斯特",
		TitleCode: "d_munster_germany",
		Counties:  map[string]feud.County{},
	}

	f.帕德博恩Gottingen = gottingen.CGottingen帕德博恩
	f.帕德博恩Gottingen.SetParent(f)
	
	f.明斯特Munster = munster.CMunster明斯特
	f.明斯特Munster.SetParent(f)
	
}
