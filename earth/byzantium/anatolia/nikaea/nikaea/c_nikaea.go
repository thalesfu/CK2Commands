package nikaea

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NikaeaCounty interface {
    feud.County
    BKios基奥斯() 	feud.Barony
    BKotyaion科泰严() 	feud.Barony
    BModrene莫德里尼() 	feud.Barony
    BNikaea尼西亚() 	feud.Barony
    BOptimatum奥普提马同() 	feud.Barony
    BPalaeokastron普莱欧卡斯特朗() 	feud.Barony
    BYalova亚洛瓦() 	feud.Barony
}

type 尼西亚NikaeaCounty struct {
	feud.BaseCounty
	基奥斯Kios 	feud.Barony
	科泰严Kotyaion 	feud.Barony
	莫德里尼Modrene 	feud.Barony
	尼西亚Nikaea 	feud.Barony
	奥普提马同Optimatum 	feud.Barony
	普莱欧卡斯特朗Palaeokastron 	feud.Barony
	亚洛瓦Yalova 	feud.Barony
}

func (c *尼西亚NikaeaCounty) BKios基奥斯() feud.Barony {
	return c.基奥斯Kios
}
    
func (c *尼西亚NikaeaCounty) BKotyaion科泰严() feud.Barony {
	return c.科泰严Kotyaion
}
    
func (c *尼西亚NikaeaCounty) BModrene莫德里尼() feud.Barony {
	return c.莫德里尼Modrene
}
    
func (c *尼西亚NikaeaCounty) BNikaea尼西亚() feud.Barony {
	return c.尼西亚Nikaea
}
    
func (c *尼西亚NikaeaCounty) BOptimatum奥普提马同() feud.Barony {
	return c.奥普提马同Optimatum
}
    
func (c *尼西亚NikaeaCounty) BPalaeokastron普莱欧卡斯特朗() feud.Barony {
	return c.普莱欧卡斯特朗Palaeokastron
}
    
func (c *尼西亚NikaeaCounty) BYalova亚洛瓦() feud.Barony {
	return c.亚洛瓦Yalova
}
    
var CNikaea尼西亚 NikaeaCounty = &尼西亚NikaeaCounty{}

func init() {
	f := CNikaea尼西亚.(*尼西亚NikaeaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "750",
		Title:     "nikaea",
		TitleName: "尼西亚",
		TitleCode: "c_nikaea",
		Baronies:  map[string]feud.Barony{},
	}

	f.基奥斯Kios = BKios基奥斯
	f.基奥斯Kios.SetParent(f)
	
	f.科泰严Kotyaion = BKotyaion科泰严
	f.科泰严Kotyaion.SetParent(f)
	
	f.莫德里尼Modrene = BModrene莫德里尼
	f.莫德里尼Modrene.SetParent(f)
	
	f.尼西亚Nikaea = BNikaea尼西亚
	f.尼西亚Nikaea.SetParent(f)
	
	f.奥普提马同Optimatum = BOptimatum奥普提马同
	f.奥普提马同Optimatum.SetParent(f)
	
	f.普莱欧卡斯特朗Palaeokastron = BPalaeokastron普莱欧卡斯特朗
	f.普莱欧卡斯特朗Palaeokastron.SetParent(f)
	
	f.亚洛瓦Yalova = BYalova亚洛瓦
	f.亚洛瓦Yalova.SetParent(f)
	
}
