package vastergotland

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type VastergotlandCounty interface {
    feud.County
    BAlvsborg埃尔夫斯堡() 	feud.Barony
    BFalkoping法尔雪平() 	feud.Barony
    BLacko莱雪() 	feud.Barony
    BLena莱纳() 	feud.Barony
    BLindholmen林德霍尔门() 	feud.Barony
    BLodose勒德瑟() 	feud.Barony
    BSkalunda斯卡伦达() 	feud.Barony
}

type 西约特兰VastergotlandCounty struct {
	feud.BaseCounty
	埃尔夫斯堡Alvsborg 	feud.Barony
	法尔雪平Falkoping 	feud.Barony
	莱雪Lacko 	feud.Barony
	莱纳Lena 	feud.Barony
	林德霍尔门Lindholmen 	feud.Barony
	勒德瑟Lodose 	feud.Barony
	斯卡伦达Skalunda 	feud.Barony
}

func (c *西约特兰VastergotlandCounty) BAlvsborg埃尔夫斯堡() feud.Barony {
	return c.埃尔夫斯堡Alvsborg
}
    
func (c *西约特兰VastergotlandCounty) BFalkoping法尔雪平() feud.Barony {
	return c.法尔雪平Falkoping
}
    
func (c *西约特兰VastergotlandCounty) BLacko莱雪() feud.Barony {
	return c.莱雪Lacko
}
    
func (c *西约特兰VastergotlandCounty) BLena莱纳() feud.Barony {
	return c.莱纳Lena
}
    
func (c *西约特兰VastergotlandCounty) BLindholmen林德霍尔门() feud.Barony {
	return c.林德霍尔门Lindholmen
}
    
func (c *西约特兰VastergotlandCounty) BLodose勒德瑟() feud.Barony {
	return c.勒德瑟Lodose
}
    
func (c *西约特兰VastergotlandCounty) BSkalunda斯卡伦达() feud.Barony {
	return c.斯卡伦达Skalunda
}
    
var CVastergotland西约特兰 VastergotlandCounty = &西约特兰VastergotlandCounty{}

func init() {
	f := CVastergotland西约特兰.(*西约特兰VastergotlandCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "297",
		Title:     "vastergotland",
		TitleName: "西约特兰",
		TitleCode: "c_vastergotland",
		Baronies:  map[string]feud.Barony{},
	}

	f.埃尔夫斯堡Alvsborg = BAlvsborg埃尔夫斯堡
	f.埃尔夫斯堡Alvsborg.SetParent(f)
	
	f.法尔雪平Falkoping = BFalkoping法尔雪平
	f.法尔雪平Falkoping.SetParent(f)
	
	f.莱雪Lacko = BLacko莱雪
	f.莱雪Lacko.SetParent(f)
	
	f.莱纳Lena = BLena莱纳
	f.莱纳Lena.SetParent(f)
	
	f.林德霍尔门Lindholmen = BLindholmen林德霍尔门
	f.林德霍尔门Lindholmen.SetParent(f)
	
	f.勒德瑟Lodose = BLodose勒德瑟
	f.勒德瑟Lodose.SetParent(f)
	
	f.斯卡伦达Skalunda = BSkalunda斯卡伦达
	f.斯卡伦达Skalunda.SetParent(f)
	
}
