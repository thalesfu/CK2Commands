package leinster

import (
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/leinster/leinster"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/leinster/leix"
	"github.com/thalesfu/CK2Commands/earth/britannia/ireland/leinster/ossory"
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type LeinsterDuke interface {
    feud.Duke
    CLeinster伦斯特() 	leinster.LeinsterCounty
    CLeix莱伊什() 	leix.LeixCounty
    COssory奥索里() 	ossory.OssoryCounty
}

type 伦斯特LeinsterDuke struct {
	feud.BaseDuke
	伦斯特Leinster 	leinster.LeinsterCounty
	莱伊什Leix 	leix.LeixCounty
	奥索里Ossory 	ossory.OssoryCounty
}

func (d *伦斯特LeinsterDuke) CLeinster伦斯特() leinster.LeinsterCounty {
	return d.伦斯特Leinster
}
    
func (d *伦斯特LeinsterDuke) CLeix莱伊什() leix.LeixCounty {
	return d.莱伊什Leix
}
    
func (d *伦斯特LeinsterDuke) COssory奥索里() ossory.OssoryCounty {
	return d.奥索里Ossory
}
    
var DLeinster伦斯特 LeinsterDuke = &伦斯特LeinsterDuke{}

func init() {
	f := DLeinster伦斯特.(*伦斯特LeinsterDuke)
	f.BaseDuke = feud.BaseDuke{
		Title:     "leinster",
		TitleName: "伦斯特",
		TitleCode: "d_leinster",
		Counties:  map[string]feud.County{},
	}

	f.伦斯特Leinster = leinster.CLeinster伦斯特
	f.伦斯特Leinster.SetParent(f)
	
	f.莱伊什Leix = leix.CLeix莱伊什
	f.莱伊什Leix.SetParent(f)
	
	f.奥索里Ossory = ossory.COssory奥索里
	f.奥索里Ossory.SetParent(f)
	
}
