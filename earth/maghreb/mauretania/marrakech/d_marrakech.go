package marrakech

import (
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/marrakech/canarias"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/marrakech/marrakech"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/marrakech/massat"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/marrakech/tadla"
	"github.com/thalesfu/CK2Commands/earth/maghreb/mauretania/marrakech/tinmallal"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MarrakechDuke interface {
    feud.Duke
    CCanarias加那利() 	canarias.CanariasCounty
    CMarrakech马拉喀什() 	marrakech.MarrakechCounty
    CMassat安法() 	massat.MassatCounty
    CTadla塔德莱() 	tadla.TadlaCounty
    CTinmallal廷马拉尔() 	tinmallal.TinmallalCounty
}

type 马拉喀什MarrakechDuke struct {
	feud.BaseDuke
	加那利Canarias 	canarias.CanariasCounty
	马拉喀什Marrakech 	marrakech.MarrakechCounty
	安法Massat 	massat.MassatCounty
	塔德莱Tadla 	tadla.TadlaCounty
	廷马拉尔Tinmallal 	tinmallal.TinmallalCounty
}

func (d *马拉喀什MarrakechDuke) CCanarias加那利() canarias.CanariasCounty {
	return d.加那利Canarias
}
    
func (d *马拉喀什MarrakechDuke) CMarrakech马拉喀什() marrakech.MarrakechCounty {
	return d.马拉喀什Marrakech
}
    
func (d *马拉喀什MarrakechDuke) CMassat安法() massat.MassatCounty {
	return d.安法Massat
}
    
func (d *马拉喀什MarrakechDuke) CTadla塔德莱() tadla.TadlaCounty {
	return d.塔德莱Tadla
}
    
func (d *马拉喀什MarrakechDuke) CTinmallal廷马拉尔() tinmallal.TinmallalCounty {
	return d.廷马拉尔Tinmallal
}
    
var DMarrakech马拉喀什 MarrakechDuke = &马拉喀什MarrakechDuke{}

func init() {
	f := DMarrakech马拉喀什.(*马拉喀什MarrakechDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "marrakech",
		TitleName: "马拉喀什",
		TitleCode: "d_marrakech",
		Counties:  map[string]feud.County{},
	}

	f.加那利Canarias = canarias.CCanarias加那利
	f.加那利Canarias.SetParent(f)
	
	f.马拉喀什Marrakech = marrakech.CMarrakech马拉喀什
	f.马拉喀什Marrakech.SetParent(f)
	
	f.安法Massat = massat.CMassat安法
	f.安法Massat.SetParent(f)
	
	f.塔德莱Tadla = tadla.CTadla塔德莱
	f.塔德莱Tadla.SetParent(f)
	
	f.廷马拉尔Tinmallal = tinmallal.CTinmallal廷马拉尔
	f.廷马拉尔Tinmallal.SetParent(f)
	
}
