package galicia

import (
	"github.com/thalesfu/CK2Commands/earth/spain/spanish_galicia/galicia/coruna"
	"github.com/thalesfu/CK2Commands/earth/spain/spanish_galicia/galicia/lugo"
	"github.com/thalesfu/CK2Commands/earth/spain/spanish_galicia/galicia/santiago"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GaliciaDuke interface {
    feud.Duke
    CCoruna科鲁尼亚() 	coruna.CorunaCounty
    CLugo卢戈() 	lugo.LugoCounty
    CSantiago圣地亚哥() 	santiago.SantiagoCounty
}

type 加利西亚GaliciaDuke struct {
	feud.BaseDuke
	科鲁尼亚Coruna 	coruna.CorunaCounty
	卢戈Lugo 	lugo.LugoCounty
	圣地亚哥Santiago 	santiago.SantiagoCounty
}

func (d *加利西亚GaliciaDuke) CCoruna科鲁尼亚() coruna.CorunaCounty {
	return d.科鲁尼亚Coruna
}
    
func (d *加利西亚GaliciaDuke) CLugo卢戈() lugo.LugoCounty {
	return d.卢戈Lugo
}
    
func (d *加利西亚GaliciaDuke) CSantiago圣地亚哥() santiago.SantiagoCounty {
	return d.圣地亚哥Santiago
}
    
var DGalicia加利西亚 GaliciaDuke = &加利西亚GaliciaDuke{}

func init() {
	f := DGalicia加利西亚.(*加利西亚GaliciaDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "galicia",
		TitleName: "加利西亚",
		TitleCode: "d_galicia",
		Counties:  map[string]feud.County{},
	}

	f.科鲁尼亚Coruna = coruna.CCoruna科鲁尼亚
	f.科鲁尼亚Coruna.SetParent(f)
	
	f.卢戈Lugo = lugo.CLugo卢戈
	f.卢戈Lugo.SetParent(f)
	
	f.圣地亚哥Santiago = santiago.CSantiago圣地亚哥
	f.圣地亚哥Santiago.SetParent(f)
	
}
