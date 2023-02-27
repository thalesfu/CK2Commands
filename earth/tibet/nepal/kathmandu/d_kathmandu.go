package kathmandu

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/kathmandu/janakpur"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/kathmandu/kathmandu"
	"github.com/thalesfu/CK2Commands/earth/tibet/nepal/kathmandu/limbuwan"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type KathmanduDuke interface {
    feud.Duke
    CJanakpur阇那迦补罗() 	janakpur.JanakpurCounty
    CKathmandu加德满都() 	kathmandu.KathmanduCounty
    CLimbuwan林布旺() 	limbuwan.LimbuwanCounty
}

type 加德满都KathmanduDuke struct {
	feud.BaseDuke
	阇那迦补罗Janakpur 	janakpur.JanakpurCounty
	加德满都Kathmandu 	kathmandu.KathmanduCounty
	林布旺Limbuwan 	limbuwan.LimbuwanCounty
}

func (d *加德满都KathmanduDuke) CJanakpur阇那迦补罗() janakpur.JanakpurCounty {
	return d.阇那迦补罗Janakpur
}
    
func (d *加德满都KathmanduDuke) CKathmandu加德满都() kathmandu.KathmanduCounty {
	return d.加德满都Kathmandu
}
    
func (d *加德满都KathmanduDuke) CLimbuwan林布旺() limbuwan.LimbuwanCounty {
	return d.林布旺Limbuwan
}
    
var DKathmandu加德满都 KathmanduDuke = &加德满都KathmanduDuke{}

func init() {
	f := DKathmandu加德满都.(*加德满都KathmanduDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "kathmandu",
		TitleName: "加德满都",
		TitleCode: "d_kathmandu",
		Counties:  map[string]feud.County{},
	}

	f.阇那迦补罗Janakpur = janakpur.CJanakpur阇那迦补罗
	f.阇那迦补罗Janakpur.SetParent(f)
	
	f.加德满都Kathmandu = kathmandu.CKathmandu加德满都
	f.加德满都Kathmandu.SetParent(f)
	
	f.林布旺Limbuwan = limbuwan.CLimbuwan林布旺
	f.林布旺Limbuwan.SetParent(f)
	
}
