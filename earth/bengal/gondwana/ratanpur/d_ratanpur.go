package ratanpur

import (
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana/ratanpur/bandhugadha"
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana/ratanpur/ratanpur"
	"github.com/thalesfu/CK2Commands/earth/bengal/gondwana/ratanpur/tummana"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RatanpurDuke interface {
    feud.Duke
    CBandhugadha畔度伽荼() 	bandhugadha.BandhugadhaCounty
    CRatanpur剌怛那补罗() 	ratanpur.RatanpurCounty
    CTummana都摩那() 	tummana.TummanaCounty
}

type 剌怛那补罗RatanpurDuke struct {
	feud.BaseDuke
	畔度伽荼Bandhugadha 	bandhugadha.BandhugadhaCounty
	剌怛那补罗Ratanpur 	ratanpur.RatanpurCounty
	都摩那Tummana 	tummana.TummanaCounty
}

func (d *剌怛那补罗RatanpurDuke) CBandhugadha畔度伽荼() bandhugadha.BandhugadhaCounty {
	return d.畔度伽荼Bandhugadha
}
    
func (d *剌怛那补罗RatanpurDuke) CRatanpur剌怛那补罗() ratanpur.RatanpurCounty {
	return d.剌怛那补罗Ratanpur
}
    
func (d *剌怛那补罗RatanpurDuke) CTummana都摩那() tummana.TummanaCounty {
	return d.都摩那Tummana
}
    
var DRatanpur剌怛那补罗 RatanpurDuke = &剌怛那补罗RatanpurDuke{}

func init() {
	f := DRatanpur剌怛那补罗.(*剌怛那补罗RatanpurDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "ratanpur",
		TitleName: "剌怛那补罗",
		TitleCode: "d_ratanpur",
		Counties:  map[string]feud.County{},
	}

	f.畔度伽荼Bandhugadha = bandhugadha.CBandhugadha畔度伽荼
	f.畔度伽荼Bandhugadha.SetParent(f)
	
	f.剌怛那补罗Ratanpur = ratanpur.CRatanpur剌怛那补罗
	f.剌怛那补罗Ratanpur.SetParent(f)
	
	f.都摩那Tummana = tummana.CTummana都摩那
	f.都摩那Tummana.SetParent(f)
	
}
