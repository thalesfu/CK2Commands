package ulster

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/ulster/oriel"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/ulster/tyrconnell"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/ulster/tyrone"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/ulster/ulster"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type UlsterDuke interface {
    feud.Duke
    COriel奥里尔() 	oriel.OrielCounty
    CTyrconnell蒂尔康奈() 	tyrconnell.TyrconnellCounty
    CTyrone蒂龙() 	tyrone.TyroneCounty
    CUlster阿尔斯特() 	ulster.UlsterCounty
}

type 阿尔斯特UlsterDuke struct {
	feud.BaseDuke
	奥里尔Oriel 	oriel.OrielCounty
	蒂尔康奈Tyrconnell 	tyrconnell.TyrconnellCounty
	蒂龙Tyrone 	tyrone.TyroneCounty
	阿尔斯特Ulster 	ulster.UlsterCounty
}

func (d *阿尔斯特UlsterDuke) COriel奥里尔() oriel.OrielCounty {
	return d.奥里尔Oriel
}
    
func (d *阿尔斯特UlsterDuke) CTyrconnell蒂尔康奈() tyrconnell.TyrconnellCounty {
	return d.蒂尔康奈Tyrconnell
}
    
func (d *阿尔斯特UlsterDuke) CTyrone蒂龙() tyrone.TyroneCounty {
	return d.蒂龙Tyrone
}
    
func (d *阿尔斯特UlsterDuke) CUlster阿尔斯特() ulster.UlsterCounty {
	return d.阿尔斯特Ulster
}
    
var DUlster阿尔斯特 UlsterDuke = &阿尔斯特UlsterDuke{}

func init() {
	f := DUlster阿尔斯特.(*阿尔斯特UlsterDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ulster",
		TitleName: "阿尔斯特",
		TitleCode: "d_ulster",
		Counties:  map[string]feud.County{},
	}

	f.奥里尔Oriel = oriel.COriel奥里尔
	f.奥里尔Oriel.SetParent(f)
	
	f.蒂尔康奈Tyrconnell = tyrconnell.CTyrconnell蒂尔康奈
	f.蒂尔康奈Tyrconnell.SetParent(f)
	
	f.蒂龙Tyrone = tyrone.CTyrone蒂龙
	f.蒂龙Tyrone.SetParent(f)
	
	f.阿尔斯特Ulster = ulster.CUlster阿尔斯特
	f.阿尔斯特Ulster.SetParent(f)
	
}
