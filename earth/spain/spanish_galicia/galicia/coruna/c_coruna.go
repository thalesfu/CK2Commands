package coruna

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type CorunaCounty interface {
    feud.County
    BBurela布雷拉() 	feud.Barony
    BCorunna阿科鲁尼亚() 	feud.Barony
    BFerrol费罗尔() 	feud.Barony
    BMondonedo蒙多涅多() 	feud.Barony
    BTriacastela特里亚卡斯特拉() 	feud.Barony
    BVillalba比利亚尔瓦() 	feud.Barony
    BViveiro比韦罗() 	feud.Barony
}

type 科鲁尼亚CorunaCounty struct {
	feud.BaseCounty
	布雷拉Burela 	feud.Barony
	阿科鲁尼亚Corunna 	feud.Barony
	费罗尔Ferrol 	feud.Barony
	蒙多涅多Mondonedo 	feud.Barony
	特里亚卡斯特拉Triacastela 	feud.Barony
	比利亚尔瓦Villalba 	feud.Barony
	比韦罗Viveiro 	feud.Barony
}

func (c *科鲁尼亚CorunaCounty) BBurela布雷拉() feud.Barony {
	return c.布雷拉Burela
}
    
func (c *科鲁尼亚CorunaCounty) BCorunna阿科鲁尼亚() feud.Barony {
	return c.阿科鲁尼亚Corunna
}
    
func (c *科鲁尼亚CorunaCounty) BFerrol费罗尔() feud.Barony {
	return c.费罗尔Ferrol
}
    
func (c *科鲁尼亚CorunaCounty) BMondonedo蒙多涅多() feud.Barony {
	return c.蒙多涅多Mondonedo
}
    
func (c *科鲁尼亚CorunaCounty) BTriacastela特里亚卡斯特拉() feud.Barony {
	return c.特里亚卡斯特拉Triacastela
}
    
func (c *科鲁尼亚CorunaCounty) BVillalba比利亚尔瓦() feud.Barony {
	return c.比利亚尔瓦Villalba
}
    
func (c *科鲁尼亚CorunaCounty) BViveiro比韦罗() feud.Barony {
	return c.比韦罗Viveiro
}
    
var CCoruna科鲁尼亚 CorunaCounty = &科鲁尼亚CorunaCounty{}

func init() {
	f := CCoruna科鲁尼亚.(*科鲁尼亚CorunaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "156",
		Title:     "coruna",
		TitleName: "科鲁尼亚",
		TitleCode: "c_coruna",
		Baronies:  map[string]feud.Barony{},
	}

	f.布雷拉Burela = BBurela布雷拉
	f.布雷拉Burela.SetParent(f)
	
	f.阿科鲁尼亚Corunna = BCorunna阿科鲁尼亚
	f.阿科鲁尼亚Corunna.SetParent(f)
	
	f.费罗尔Ferrol = BFerrol费罗尔
	f.费罗尔Ferrol.SetParent(f)
	
	f.蒙多涅多Mondonedo = BMondonedo蒙多涅多
	f.蒙多涅多Mondonedo.SetParent(f)
	
	f.特里亚卡斯特拉Triacastela = BTriacastela特里亚卡斯特拉
	f.特里亚卡斯特拉Triacastela.SetParent(f)
	
	f.比利亚尔瓦Villalba = BVillalba比利亚尔瓦
	f.比利亚尔瓦Villalba.SetParent(f)
	
	f.比韦罗Viveiro = BViveiro比韦罗
	f.比韦罗Viveiro.SetParent(f)
	
}
