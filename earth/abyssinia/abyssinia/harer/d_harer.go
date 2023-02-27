package harer

import (
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/harer/berbera"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/harer/busaso"
	"github.com/thalesfu/CK2Commands/earth/abyssinia/abyssinia/harer/zeila"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type HarerDuke interface {
    feud.Duke
    CBerbera拨拔力() 	berbera.BerberaCounty
    CBusaso博萨索() 	busaso.BusasoCounty
    CZeila泽拉() 	zeila.ZeilaCounty
}

type 拨拔力HarerDuke struct {
	feud.BaseDuke
	拨拔力Berbera 	berbera.BerberaCounty
	博萨索Busaso 	busaso.BusasoCounty
	泽拉Zeila 	zeila.ZeilaCounty
}

func (d *拨拔力HarerDuke) CBerbera拨拔力() berbera.BerberaCounty {
	return d.拨拔力Berbera
}
    
func (d *拨拔力HarerDuke) CBusaso博萨索() busaso.BusasoCounty {
	return d.博萨索Busaso
}
    
func (d *拨拔力HarerDuke) CZeila泽拉() zeila.ZeilaCounty {
	return d.泽拉Zeila
}
    
var DHarer拨拔力 HarerDuke = &拨拔力HarerDuke{}

func init() {
	f := DHarer拨拔力.(*拨拔力HarerDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "harer",
		TitleName: "拨拔力",
		TitleCode: "d_harer",
		Counties:  map[string]feud.County{},
	}

	f.拨拔力Berbera = berbera.CBerbera拨拔力
	f.拨拔力Berbera.SetParent(f)
	
	f.博萨索Busaso = busaso.CBusaso博萨索
	f.博萨索Busaso.SetParent(f)
	
	f.泽拉Zeila = zeila.CZeila泽拉
	f.泽拉Zeila.SetParent(f)
	
}
