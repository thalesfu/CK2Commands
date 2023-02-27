package gloucester

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/england/gloucester/gloucester"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/gloucester/oxford"
	"github.com/thalesfu/CK2Commands/earth/britannia/england/gloucester/wiltshire"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GloucesterDuke interface {
    feud.Duke
    CGloucester格洛斯特() 	gloucester.GloucesterCounty
    COxford牛津() 	oxford.OxfordCounty
    CWiltshire威尔特郡() 	wiltshire.WiltshireCounty
}

type 赫威赛GloucesterDuke struct {
	feud.BaseDuke
	格洛斯特Gloucester 	gloucester.GloucesterCounty
	牛津Oxford 	oxford.OxfordCounty
	威尔特郡Wiltshire 	wiltshire.WiltshireCounty
}

func (d *赫威赛GloucesterDuke) CGloucester格洛斯特() gloucester.GloucesterCounty {
	return d.格洛斯特Gloucester
}
    
func (d *赫威赛GloucesterDuke) COxford牛津() oxford.OxfordCounty {
	return d.牛津Oxford
}
    
func (d *赫威赛GloucesterDuke) CWiltshire威尔特郡() wiltshire.WiltshireCounty {
	return d.威尔特郡Wiltshire
}
    
var DGloucester赫威赛 GloucesterDuke = &赫威赛GloucesterDuke{}

func init() {
	f := DGloucester赫威赛.(*赫威赛GloucesterDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "gloucester",
		TitleName: "赫威赛",
		TitleCode: "d_gloucester",
		Counties:  map[string]feud.County{},
	}

	f.格洛斯特Gloucester = gloucester.CGloucester格洛斯特
	f.格洛斯特Gloucester.SetParent(f)
	
	f.牛津Oxford = oxford.COxford牛津
	f.牛津Oxford.SetParent(f)
	
	f.威尔特郡Wiltshire = wiltshire.CWiltshire威尔特郡
	f.威尔特郡Wiltshire.SetParent(f)
	
}
