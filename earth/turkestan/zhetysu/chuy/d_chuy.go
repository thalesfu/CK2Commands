package chuy

import (
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/chuy/barskhan"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/chuy/chuy"
	"github.com/thalesfu/CK2Commands/earth/turkestan/zhetysu/chuy/talas"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ChuyDuke interface {
    feud.Duke
    CBarskhan拔塞干() 	barskhan.BarskhanCounty
    CChuy裴罗将军城() 	chuy.ChuyCounty
    CTalas怛罗斯() 	talas.TalasCounty
}

type 垂河ChuyDuke struct {
	feud.BaseDuke
	拔塞干Barskhan 	barskhan.BarskhanCounty
	裴罗将军城Chuy 	chuy.ChuyCounty
	怛罗斯Talas 	talas.TalasCounty
}

func (d *垂河ChuyDuke) CBarskhan拔塞干() barskhan.BarskhanCounty {
	return d.拔塞干Barskhan
}
    
func (d *垂河ChuyDuke) CChuy裴罗将军城() chuy.ChuyCounty {
	return d.裴罗将军城Chuy
}
    
func (d *垂河ChuyDuke) CTalas怛罗斯() talas.TalasCounty {
	return d.怛罗斯Talas
}
    
var DChuy垂河 ChuyDuke = &垂河ChuyDuke{}

func init() {
	f := DChuy垂河.(*垂河ChuyDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "chuy",
		TitleName: "垂河",
		TitleCode: "d_chuy",
		Counties:  map[string]feud.County{},
	}

	f.拔塞干Barskhan = barskhan.CBarskhan拔塞干
	f.拔塞干Barskhan.SetParent(f)
	
	f.裴罗将军城Chuy = chuy.CChuy裴罗将军城
	f.裴罗将军城Chuy.SetParent(f)
	
	f.怛罗斯Talas = talas.CTalas怛罗斯
	f.怛罗斯Talas.SetParent(f)
	
}
