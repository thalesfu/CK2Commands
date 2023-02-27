package auvergne

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AuvergneCounty interface {
    feud.County
    BBrioude布里乌德() 	feud.Barony
    BClermont克莱蒙() 	feud.Barony
    BDomeyrat多梅拉() 	feud.Barony
    BMontferrand蒙特费朗() 	feud.Barony
    BMozac莫扎克() 	feud.Barony
    BMurol米罗勒() 	feud.Barony
    BTournoel图尔纳尔() 	feud.Barony
}

type 奥弗涅AuvergneCounty struct {
	feud.BaseCounty
	布里乌德Brioude 	feud.Barony
	克莱蒙Clermont 	feud.Barony
	多梅拉Domeyrat 	feud.Barony
	蒙特费朗Montferrand 	feud.Barony
	莫扎克Mozac 	feud.Barony
	米罗勒Murol 	feud.Barony
	图尔纳尔Tournoel 	feud.Barony
}

func (c *奥弗涅AuvergneCounty) BBrioude布里乌德() feud.Barony {
	return c.布里乌德Brioude
}
    
func (c *奥弗涅AuvergneCounty) BClermont克莱蒙() feud.Barony {
	return c.克莱蒙Clermont
}
    
func (c *奥弗涅AuvergneCounty) BDomeyrat多梅拉() feud.Barony {
	return c.多梅拉Domeyrat
}
    
func (c *奥弗涅AuvergneCounty) BMontferrand蒙特费朗() feud.Barony {
	return c.蒙特费朗Montferrand
}
    
func (c *奥弗涅AuvergneCounty) BMozac莫扎克() feud.Barony {
	return c.莫扎克Mozac
}
    
func (c *奥弗涅AuvergneCounty) BMurol米罗勒() feud.Barony {
	return c.米罗勒Murol
}
    
func (c *奥弗涅AuvergneCounty) BTournoel图尔纳尔() feud.Barony {
	return c.图尔纳尔Tournoel
}
    
var CAuvergne奥弗涅 AuvergneCounty = &奥弗涅AuvergneCounty{}

func init() {
	f := CAuvergne奥弗涅.(*奥弗涅AuvergneCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "217",
		Title:     "auvergne",
		TitleName: "奥弗涅",
		TitleCode: "c_auvergne",
		Baronies:  map[string]feud.Barony{},
	}

	f.布里乌德Brioude = BBrioude布里乌德
	f.布里乌德Brioude.SetParent(f)
	
	f.克莱蒙Clermont = BClermont克莱蒙
	f.克莱蒙Clermont.SetParent(f)
	
	f.多梅拉Domeyrat = BDomeyrat多梅拉
	f.多梅拉Domeyrat.SetParent(f)
	
	f.蒙特费朗Montferrand = BMontferrand蒙特费朗
	f.蒙特费朗Montferrand.SetParent(f)
	
	f.莫扎克Mozac = BMozac莫扎克
	f.莫扎克Mozac.SetParent(f)
	
	f.米罗勒Murol = BMurol米罗勒
	f.米罗勒Murol.SetParent(f)
	
	f.图尔纳尔Tournoel = BTournoel图尔纳尔
	f.图尔纳尔Tournoel.SetParent(f)
	
}
