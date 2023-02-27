package olvia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type OlviaCounty interface {
    feud.County
    BKobleve科布列韦() 	feud.Barony
    BOchakiv奥恰科夫() 	feud.Barony
    BOlvia奥利维亚() 	feud.Barony
    BRybakivka雷巴基夫卡() 	feud.Barony
    BStavky斯塔夫基() 	feud.Barony
    BVeselynove韦谢利诺韦() 	feud.Barony
    BVoznesensk沃兹涅先斯克() 	feud.Barony
}

type 奥利维亚OlviaCounty struct {
	feud.BaseCounty
	科布列韦Kobleve 	feud.Barony
	奥恰科夫Ochakiv 	feud.Barony
	奥利维亚Olvia 	feud.Barony
	雷巴基夫卡Rybakivka 	feud.Barony
	斯塔夫基Stavky 	feud.Barony
	韦谢利诺韦Veselynove 	feud.Barony
	沃兹涅先斯克Voznesensk 	feud.Barony
}

func (c *奥利维亚OlviaCounty) BKobleve科布列韦() feud.Barony {
	return c.科布列韦Kobleve
}
    
func (c *奥利维亚OlviaCounty) BOchakiv奥恰科夫() feud.Barony {
	return c.奥恰科夫Ochakiv
}
    
func (c *奥利维亚OlviaCounty) BOlvia奥利维亚() feud.Barony {
	return c.奥利维亚Olvia
}
    
func (c *奥利维亚OlviaCounty) BRybakivka雷巴基夫卡() feud.Barony {
	return c.雷巴基夫卡Rybakivka
}
    
func (c *奥利维亚OlviaCounty) BStavky斯塔夫基() feud.Barony {
	return c.斯塔夫基Stavky
}
    
func (c *奥利维亚OlviaCounty) BVeselynove韦谢利诺韦() feud.Barony {
	return c.韦谢利诺韦Veselynove
}
    
func (c *奥利维亚OlviaCounty) BVoznesensk沃兹涅先斯克() feud.Barony {
	return c.沃兹涅先斯克Voznesensk
}
    
var COlvia奥利维亚 OlviaCounty = &奥利维亚OlviaCounty{}

func init() {
	f := COlvia奥利维亚.(*奥利维亚OlviaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "542",
		Title:     "olvia",
		TitleName: "奥利维亚",
		TitleCode: "c_olvia",
		Baronies:  map[string]feud.Barony{},
	}

	f.科布列韦Kobleve = BKobleve科布列韦
	f.科布列韦Kobleve.SetParent(f)
	
	f.奥恰科夫Ochakiv = BOchakiv奥恰科夫
	f.奥恰科夫Ochakiv.SetParent(f)
	
	f.奥利维亚Olvia = BOlvia奥利维亚
	f.奥利维亚Olvia.SetParent(f)
	
	f.雷巴基夫卡Rybakivka = BRybakivka雷巴基夫卡
	f.雷巴基夫卡Rybakivka.SetParent(f)
	
	f.斯塔夫基Stavky = BStavky斯塔夫基
	f.斯塔夫基Stavky.SetParent(f)
	
	f.韦谢利诺韦Veselynove = BVeselynove韦谢利诺韦
	f.韦谢利诺韦Veselynove.SetParent(f)
	
	f.沃兹涅先斯克Voznesensk = BVoznesensk沃兹涅先斯克
	f.沃兹涅先斯克Voznesensk.SetParent(f)
	
}
