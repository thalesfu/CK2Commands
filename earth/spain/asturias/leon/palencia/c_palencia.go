package palencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PalenciaCounty interface {
    feud.County
    BMonte_el_viejo埃尔别霍山() 	feud.Barony
    BOlmeda_palencia奥尔梅达() 	feud.Barony
    BPalencia帕伦西亚() 	feud.Barony
    BParedes帕雷德斯() 	feud.Barony
    BSaint_antolin圣安托林() 	feud.Barony
    BSan_salvador圣萨尔瓦多() 	feud.Barony
    BVaccaei瓦凯伊() 	feud.Barony
}

type 帕伦西亚PalenciaCounty struct {
	feud.BaseCounty
	埃尔别霍山Monte_el_viejo 	feud.Barony
	奥尔梅达Olmeda_palencia 	feud.Barony
	帕伦西亚Palencia 	feud.Barony
	帕雷德斯Paredes 	feud.Barony
	圣安托林Saint_antolin 	feud.Barony
	圣萨尔瓦多San_salvador 	feud.Barony
	瓦凯伊Vaccaei 	feud.Barony
}

func (c *帕伦西亚PalenciaCounty) BMonte_el_viejo埃尔别霍山() feud.Barony {
	return c.埃尔别霍山Monte_el_viejo
}
    
func (c *帕伦西亚PalenciaCounty) BOlmeda_palencia奥尔梅达() feud.Barony {
	return c.奥尔梅达Olmeda_palencia
}
    
func (c *帕伦西亚PalenciaCounty) BPalencia帕伦西亚() feud.Barony {
	return c.帕伦西亚Palencia
}
    
func (c *帕伦西亚PalenciaCounty) BParedes帕雷德斯() feud.Barony {
	return c.帕雷德斯Paredes
}
    
func (c *帕伦西亚PalenciaCounty) BSaint_antolin圣安托林() feud.Barony {
	return c.圣安托林Saint_antolin
}
    
func (c *帕伦西亚PalenciaCounty) BSan_salvador圣萨尔瓦多() feud.Barony {
	return c.圣萨尔瓦多San_salvador
}
    
func (c *帕伦西亚PalenciaCounty) BVaccaei瓦凯伊() feud.Barony {
	return c.瓦凯伊Vaccaei
}
    
var CPalencia帕伦西亚 PalenciaCounty = &帕伦西亚PalenciaCounty{}

func init() {
	f := CPalencia帕伦西亚.(*帕伦西亚PalenciaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1975",
		Title:     "palencia",
		TitleName: "帕伦西亚",
		TitleCode: "c_palencia",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃尔别霍山Monte_el_viejo = BMonte_el_viejo埃尔别霍山
	f.埃尔别霍山Monte_el_viejo.SetParent(f)
	
	f.奥尔梅达Olmeda_palencia = BOlmeda_palencia奥尔梅达
	f.奥尔梅达Olmeda_palencia.SetParent(f)
	
	f.帕伦西亚Palencia = BPalencia帕伦西亚
	f.帕伦西亚Palencia.SetParent(f)
	
	f.帕雷德斯Paredes = BParedes帕雷德斯
	f.帕雷德斯Paredes.SetParent(f)
	
	f.圣安托林Saint_antolin = BSaint_antolin圣安托林
	f.圣安托林Saint_antolin.SetParent(f)
	
	f.圣萨尔瓦多San_salvador = BSan_salvador圣萨尔瓦多
	f.圣萨尔瓦多San_salvador.SetParent(f)
	
	f.瓦凯伊Vaccaei = BVaccaei瓦凯伊
	f.瓦凯伊Vaccaei.SetParent(f)
	
}
