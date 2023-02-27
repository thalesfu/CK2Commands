package lancaster

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/lancaster/chester"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/lancaster/derby"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/lancaster/lancaster"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LancasterDuke interface {
    feud.Duke
    CChester切斯特() 	chester.ChesterCounty
    CDerby德比() 	derby.DerbyCounty
    CLancaster兰开斯特() 	lancaster.LancasterCounty
}

type 兰开斯特LancasterDuke struct {
	feud.BaseDuke
	切斯特Chester 	chester.ChesterCounty
	德比Derby 	derby.DerbyCounty
	兰开斯特Lancaster 	lancaster.LancasterCounty
}

func (d *兰开斯特LancasterDuke) CChester切斯特() chester.ChesterCounty {
	return d.切斯特Chester
}
    
func (d *兰开斯特LancasterDuke) CDerby德比() derby.DerbyCounty {
	return d.德比Derby
}
    
func (d *兰开斯特LancasterDuke) CLancaster兰开斯特() lancaster.LancasterCounty {
	return d.兰开斯特Lancaster
}
    
var DLancaster兰开斯特 LancasterDuke = &兰开斯特LancasterDuke{}

func init() {
	f := DLancaster兰开斯特.(*兰开斯特LancasterDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "lancaster",
		TitleName: "兰开斯特",
		TitleCode: "d_lancaster",
		Counties:  map[string]feud.County{},
	}

	f.切斯特Chester = chester.CChester切斯特
	f.切斯特Chester.SetParent(f)
	
	f.德比Derby = derby.CDerby德比
	f.德比Derby.SetParent(f)
	
	f.兰开斯特Lancaster = lancaster.CLancaster兰开斯特
	f.兰开斯特Lancaster.SetParent(f)
	
}
