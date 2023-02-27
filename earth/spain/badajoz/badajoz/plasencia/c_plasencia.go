package plasencia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PlasenciaCounty interface {
    feud.County
    BHervas埃尔瓦斯() 	feud.Barony
    BJaraiz哈赖斯() 	feud.Barony
    BJarandilla哈兰迪利亚() 	feud.Barony
    BLazarza拉萨尔萨() 	feud.Barony
    BMontehermoso蒙特埃莫索() 	feud.Barony
    BPlasencia普拉森西亚() 	feud.Barony
    BTalayuela塔拉尤埃拉() 	feud.Barony
    BVentadelmoral本塔德尔莫拉尔() 	feud.Barony
}

type 普拉森西亚PlasenciaCounty struct {
	feud.BaseCounty
	埃尔瓦斯Hervas 	feud.Barony
	哈赖斯Jaraiz 	feud.Barony
	哈兰迪利亚Jarandilla 	feud.Barony
	拉萨尔萨Lazarza 	feud.Barony
	蒙特埃莫索Montehermoso 	feud.Barony
	普拉森西亚Plasencia 	feud.Barony
	塔拉尤埃拉Talayuela 	feud.Barony
	本塔德尔莫拉尔Ventadelmoral 	feud.Barony
}

func (c *普拉森西亚PlasenciaCounty) BHervas埃尔瓦斯() feud.Barony {
	return c.埃尔瓦斯Hervas
}
    
func (c *普拉森西亚PlasenciaCounty) BJaraiz哈赖斯() feud.Barony {
	return c.哈赖斯Jaraiz
}
    
func (c *普拉森西亚PlasenciaCounty) BJarandilla哈兰迪利亚() feud.Barony {
	return c.哈兰迪利亚Jarandilla
}
    
func (c *普拉森西亚PlasenciaCounty) BLazarza拉萨尔萨() feud.Barony {
	return c.拉萨尔萨Lazarza
}
    
func (c *普拉森西亚PlasenciaCounty) BMontehermoso蒙特埃莫索() feud.Barony {
	return c.蒙特埃莫索Montehermoso
}
    
func (c *普拉森西亚PlasenciaCounty) BPlasencia普拉森西亚() feud.Barony {
	return c.普拉森西亚Plasencia
}
    
func (c *普拉森西亚PlasenciaCounty) BTalayuela塔拉尤埃拉() feud.Barony {
	return c.塔拉尤埃拉Talayuela
}
    
func (c *普拉森西亚PlasenciaCounty) BVentadelmoral本塔德尔莫拉尔() feud.Barony {
	return c.本塔德尔莫拉尔Ventadelmoral
}
    
var CPlasencia普拉森西亚 PlasenciaCounty = &普拉森西亚PlasenciaCounty{}

func init() {
	f := CPlasencia普拉森西亚.(*普拉森西亚PlasenciaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "194",
		Title:     "plasencia",
		TitleName: "普拉森西亚",
		TitleCode: "c_plasencia",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃尔瓦斯Hervas = BHervas埃尔瓦斯
	f.埃尔瓦斯Hervas.SetParent(f)
	
	f.哈赖斯Jaraiz = BJaraiz哈赖斯
	f.哈赖斯Jaraiz.SetParent(f)
	
	f.哈兰迪利亚Jarandilla = BJarandilla哈兰迪利亚
	f.哈兰迪利亚Jarandilla.SetParent(f)
	
	f.拉萨尔萨Lazarza = BLazarza拉萨尔萨
	f.拉萨尔萨Lazarza.SetParent(f)
	
	f.蒙特埃莫索Montehermoso = BMontehermoso蒙特埃莫索
	f.蒙特埃莫索Montehermoso.SetParent(f)
	
	f.普拉森西亚Plasencia = BPlasencia普拉森西亚
	f.普拉森西亚Plasencia.SetParent(f)
	
	f.塔拉尤埃拉Talayuela = BTalayuela塔拉尤埃拉
	f.塔拉尤埃拉Talayuela.SetParent(f)
	
	f.本塔德尔莫拉尔Ventadelmoral = BVentadelmoral本塔德尔莫拉尔
	f.本塔德尔莫拉尔Ventadelmoral.SetParent(f)
	
}
