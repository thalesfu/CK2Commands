package azov

import (
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/azov/azov"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/azov/kuban"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/azov/sarpa"
	"github.com/thalesfu/CK2Commands/earth/pontic_steppe/alania/azov/tana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AzovDuke interface {
    feud.Duke
    CAzov亚速() 	azov.AzovCounty
    CKuban库班() 	kuban.KubanCounty
    CSarpa萨尔帕() 	sarpa.SarpaCounty
    CTana塔纳() 	tana.TanaCounty
}

type 亚速AzovDuke struct {
	feud.BaseDuke
	亚速Azov 	azov.AzovCounty
	库班Kuban 	kuban.KubanCounty
	萨尔帕Sarpa 	sarpa.SarpaCounty
	塔纳Tana 	tana.TanaCounty
}

func (d *亚速AzovDuke) CAzov亚速() azov.AzovCounty {
	return d.亚速Azov
}
    
func (d *亚速AzovDuke) CKuban库班() kuban.KubanCounty {
	return d.库班Kuban
}
    
func (d *亚速AzovDuke) CSarpa萨尔帕() sarpa.SarpaCounty {
	return d.萨尔帕Sarpa
}
    
func (d *亚速AzovDuke) CTana塔纳() tana.TanaCounty {
	return d.塔纳Tana
}
    
var DAzov亚速 AzovDuke = &亚速AzovDuke{}

func init() {
	f := DAzov亚速.(*亚速AzovDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "azov",
		TitleName: "亚速",
		TitleCode: "d_azov",
		Counties:  map[string]feud.County{},
	}

	f.亚速Azov = azov.CAzov亚速
	f.亚速Azov.SetParent(f)
	
	f.库班Kuban = kuban.CKuban库班
	f.库班Kuban.SetParent(f)
	
	f.萨尔帕Sarpa = sarpa.CSarpa萨尔帕
	f.萨尔帕Sarpa.SetParent(f)
	
	f.塔纳Tana = tana.CTana塔纳
	f.塔纳Tana.SetParent(f)
	
}
