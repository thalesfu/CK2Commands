package aswan

import (
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/aswan/alqusair"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/aswan/aswan"
	"github.com/thalesfu/CK2Commands/earth/arabia/egypt/aswan/quena"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AswanDuke interface {
    feud.Duke
    CAlqusair古赛尔() 	alqusair.AlqusairCounty
    CAswan阿斯旺() 	aswan.AswanCounty
    CQuena基纳() 	quena.QuenaCounty
}

type 阿斯旺AswanDuke struct {
	feud.BaseDuke
	古赛尔Alqusair 	alqusair.AlqusairCounty
	阿斯旺Aswan 	aswan.AswanCounty
	基纳Quena 	quena.QuenaCounty
}

func (d *阿斯旺AswanDuke) CAlqusair古赛尔() alqusair.AlqusairCounty {
	return d.古赛尔Alqusair
}
    
func (d *阿斯旺AswanDuke) CAswan阿斯旺() aswan.AswanCounty {
	return d.阿斯旺Aswan
}
    
func (d *阿斯旺AswanDuke) CQuena基纳() quena.QuenaCounty {
	return d.基纳Quena
}
    
var DAswan阿斯旺 AswanDuke = &阿斯旺AswanDuke{}

func init() {
	f := DAswan阿斯旺.(*阿斯旺AswanDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "aswan",
		TitleName: "阿斯旺",
		TitleCode: "d_aswan",
		Counties:  map[string]feud.County{},
	}

	f.古赛尔Alqusair = alqusair.CAlqusair古赛尔
	f.古赛尔Alqusair.SetParent(f)
	
	f.阿斯旺Aswan = aswan.CAswan阿斯旺
	f.阿斯旺Aswan.SetParent(f)
	
	f.基纳Quena = quena.CQuena基纳
	f.基纳Quena.SetParent(f)
	
}
