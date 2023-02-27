package navarra

import (
	"github.com/thalesfu/CK2Commands/earth/spain/navarra/navarra"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NavarraKingdom interface {
    feud.Kingdom
    DNavarra纳瓦拉() 	navarra.NavarraDuke
}

type 纳瓦拉NavarraKingdom struct {
	feud.BaseKingdom
	纳瓦拉Navarra 	navarra.NavarraDuke
}

func (k *纳瓦拉NavarraKingdom) DNavarra纳瓦拉() navarra.NavarraDuke {
	return k.纳瓦拉Navarra
}
    
var KNavarra纳瓦拉 NavarraKingdom = &纳瓦拉NavarraKingdom{}

func init() {
	f := KNavarra纳瓦拉.(*纳瓦拉NavarraKingdom)
	f.BaseKingdom = feud.BaseKingdom{
		Title:     "navarra",
		TitleName: "纳瓦拉",
		TitleCode: "k_navarra",
		Dukes:  map[string]feud.Duke{},
	}

	f.纳瓦拉Navarra = navarra.DNavarra纳瓦拉
	f.纳瓦拉Navarra.SetParent(f)
	
}
