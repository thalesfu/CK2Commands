package deheubarth

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/deheubarth/ceredigion"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/deheubarth/dyfed"
	"github.com/thalesfu/CK2Commands/earth/britannia/wales/deheubarth/gower"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type DeheubarthDuke interface {
    feud.Duke
    CCeredigion凯雷迪吉昂() 	ceredigion.CeredigionCounty
    CDyfed德维得() 	dyfed.DyfedCounty
    CGower高尔() 	gower.GowerCounty
}

type 德赫巴思DeheubarthDuke struct {
	feud.BaseDuke
	凯雷迪吉昂Ceredigion 	ceredigion.CeredigionCounty
	德维得Dyfed 	dyfed.DyfedCounty
	高尔Gower 	gower.GowerCounty
}

func (d *德赫巴思DeheubarthDuke) CCeredigion凯雷迪吉昂() ceredigion.CeredigionCounty {
	return d.凯雷迪吉昂Ceredigion
}
    
func (d *德赫巴思DeheubarthDuke) CDyfed德维得() dyfed.DyfedCounty {
	return d.德维得Dyfed
}
    
func (d *德赫巴思DeheubarthDuke) CGower高尔() gower.GowerCounty {
	return d.高尔Gower
}
    
var DDeheubarth德赫巴思 DeheubarthDuke = &德赫巴思DeheubarthDuke{}

func init() {
	f := DDeheubarth德赫巴思.(*德赫巴思DeheubarthDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "deheubarth",
		TitleName: "德赫巴思",
		TitleCode: "d_deheubarth",
		Counties:  map[string]feud.County{},
	}

	f.凯雷迪吉昂Ceredigion = ceredigion.CCeredigion凯雷迪吉昂
	f.凯雷迪吉昂Ceredigion.SetParent(f)
	
	f.德维得Dyfed = dyfed.CDyfed德维得
	f.德维得Dyfed.SetParent(f)
	
	f.高尔Gower = gower.CGower高尔
	f.高尔Gower.SetParent(f)
	
}
