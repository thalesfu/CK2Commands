package nangqen

import (
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nangqen/lingtsang"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nangqen/nangqen"
	"github.com/thalesfu/CK2Commands/earth/tibet/xixia/nangqen/zadoi"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NangqenDuke interface {
    feud.Duke
    CLingtsang灵藏() 	lingtsang.LingtsangCounty
    CNangqen囊谦() 	nangqen.NangqenCounty
    CZadoi杂多() 	zadoi.ZadoiCounty
}

type 囊谦NangqenDuke struct {
	feud.BaseDuke
	灵藏Lingtsang 	lingtsang.LingtsangCounty
	囊谦Nangqen 	nangqen.NangqenCounty
	杂多Zadoi 	zadoi.ZadoiCounty
}

func (d *囊谦NangqenDuke) CLingtsang灵藏() lingtsang.LingtsangCounty {
	return d.灵藏Lingtsang
}
    
func (d *囊谦NangqenDuke) CNangqen囊谦() nangqen.NangqenCounty {
	return d.囊谦Nangqen
}
    
func (d *囊谦NangqenDuke) CZadoi杂多() zadoi.ZadoiCounty {
	return d.杂多Zadoi
}
    
var DNangqen囊谦 NangqenDuke = &囊谦NangqenDuke{}

func init() {
	f := DNangqen囊谦.(*囊谦NangqenDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "nangqen",
		TitleName: "囊谦",
		TitleCode: "d_nangqen",
		Counties:  map[string]feud.County{},
	}

	f.灵藏Lingtsang = lingtsang.CLingtsang灵藏
	f.灵藏Lingtsang.SetParent(f)
	
	f.囊谦Nangqen = nangqen.CNangqen囊谦
	f.囊谦Nangqen.SetParent(f)
	
	f.杂多Zadoi = zadoi.CZadoi杂多
	f.杂多Zadoi.SetParent(f)
	
}
