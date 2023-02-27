package severin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SeverinCounty interface {
    feud.County
    BCaransebes卡兰塞贝什() 	feud.Barony
    BDrobeta德罗贝塔() 	feud.Barony
    BGradet格拉德茨() 	feud.Barony
    BMehadia梅哈迪亚() 	feud.Barony
    BOrsova奥尔绍瓦() 	feud.Barony
    BSeverin塞韦林() 	feud.Barony
    BVodita沃迪察() 	feud.Barony
}

type 塞韦林SeverinCounty struct {
	feud.BaseCounty
	卡兰塞贝什Caransebes 	feud.Barony
	德罗贝塔Drobeta 	feud.Barony
	格拉德茨Gradet 	feud.Barony
	梅哈迪亚Mehadia 	feud.Barony
	奥尔绍瓦Orsova 	feud.Barony
	塞韦林Severin 	feud.Barony
	沃迪察Vodita 	feud.Barony
}

func (c *塞韦林SeverinCounty) BCaransebes卡兰塞贝什() feud.Barony {
	return c.卡兰塞贝什Caransebes
}
    
func (c *塞韦林SeverinCounty) BDrobeta德罗贝塔() feud.Barony {
	return c.德罗贝塔Drobeta
}
    
func (c *塞韦林SeverinCounty) BGradet格拉德茨() feud.Barony {
	return c.格拉德茨Gradet
}
    
func (c *塞韦林SeverinCounty) BMehadia梅哈迪亚() feud.Barony {
	return c.梅哈迪亚Mehadia
}
    
func (c *塞韦林SeverinCounty) BOrsova奥尔绍瓦() feud.Barony {
	return c.奥尔绍瓦Orsova
}
    
func (c *塞韦林SeverinCounty) BSeverin塞韦林() feud.Barony {
	return c.塞韦林Severin
}
    
func (c *塞韦林SeverinCounty) BVodita沃迪察() feud.Barony {
	return c.沃迪察Vodita
}
    
var CSeverin塞韦林 SeverinCounty = &塞韦林SeverinCounty{}

func init() {
	f := CSeverin塞韦林.(*塞韦林SeverinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "516",
		Title:     "severin",
		TitleName: "塞韦林",
		TitleCode: "c_severin",
		Baronies:  map[string]feud.Barony{},
	}

	f.卡兰塞贝什Caransebes = BCaransebes卡兰塞贝什
	f.卡兰塞贝什Caransebes.SetParent(f)
	
	f.德罗贝塔Drobeta = BDrobeta德罗贝塔
	f.德罗贝塔Drobeta.SetParent(f)
	
	f.格拉德茨Gradet = BGradet格拉德茨
	f.格拉德茨Gradet.SetParent(f)
	
	f.梅哈迪亚Mehadia = BMehadia梅哈迪亚
	f.梅哈迪亚Mehadia.SetParent(f)
	
	f.奥尔绍瓦Orsova = BOrsova奥尔绍瓦
	f.奥尔绍瓦Orsova.SetParent(f)
	
	f.塞韦林Severin = BSeverin塞韦林
	f.塞韦林Severin.SetParent(f)
	
	f.沃迪察Vodita = BVodita沃迪察
	f.沃迪察Vodita.SetParent(f)
	
}
