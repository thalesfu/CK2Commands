package attaleia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type AttaleiaCounty interface {
    feud.County
    BAttaleia阿塔利亚() 	feud.Barony
    BCibyra基比拉() 	feud.Barony
    BGalanauros贾拉纳罗斯() 	feud.Barony
    BPanemotichus帕尼莫提库斯() 	feud.Barony
    BSagalassos萨加拉索斯() 	feud.Barony
    BSide西德() 	feud.Barony
    BSillyon希梁() 	feud.Barony
    BSlege斯莱治() 	feud.Barony
}

type 阿塔利亚AttaleiaCounty struct {
	feud.BaseCounty
	阿塔利亚Attaleia 	feud.Barony
	基比拉Cibyra 	feud.Barony
	贾拉纳罗斯Galanauros 	feud.Barony
	帕尼莫提库斯Panemotichus 	feud.Barony
	萨加拉索斯Sagalassos 	feud.Barony
	西德Side 	feud.Barony
	希梁Sillyon 	feud.Barony
	斯莱治Slege 	feud.Barony
}

func (c *阿塔利亚AttaleiaCounty) BAttaleia阿塔利亚() feud.Barony {
	return c.阿塔利亚Attaleia
}
    
func (c *阿塔利亚AttaleiaCounty) BCibyra基比拉() feud.Barony {
	return c.基比拉Cibyra
}
    
func (c *阿塔利亚AttaleiaCounty) BGalanauros贾拉纳罗斯() feud.Barony {
	return c.贾拉纳罗斯Galanauros
}
    
func (c *阿塔利亚AttaleiaCounty) BPanemotichus帕尼莫提库斯() feud.Barony {
	return c.帕尼莫提库斯Panemotichus
}
    
func (c *阿塔利亚AttaleiaCounty) BSagalassos萨加拉索斯() feud.Barony {
	return c.萨加拉索斯Sagalassos
}
    
func (c *阿塔利亚AttaleiaCounty) BSide西德() feud.Barony {
	return c.西德Side
}
    
func (c *阿塔利亚AttaleiaCounty) BSillyon希梁() feud.Barony {
	return c.希梁Sillyon
}
    
func (c *阿塔利亚AttaleiaCounty) BSlege斯莱治() feud.Barony {
	return c.斯莱治Slege
}
    
var CAttaleia阿塔利亚 AttaleiaCounty = &阿塔利亚AttaleiaCounty{}

func init() {
	f := CAttaleia阿塔利亚.(*阿塔利亚AttaleiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "755",
		Title:     "attaleia",
		TitleName: "阿塔利亚",
		TitleCode: "c_attaleia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿塔利亚Attaleia = BAttaleia阿塔利亚
	f.阿塔利亚Attaleia.SetParent(f)
	
	f.基比拉Cibyra = BCibyra基比拉
	f.基比拉Cibyra.SetParent(f)
	
	f.贾拉纳罗斯Galanauros = BGalanauros贾拉纳罗斯
	f.贾拉纳罗斯Galanauros.SetParent(f)
	
	f.帕尼莫提库斯Panemotichus = BPanemotichus帕尼莫提库斯
	f.帕尼莫提库斯Panemotichus.SetParent(f)
	
	f.萨加拉索斯Sagalassos = BSagalassos萨加拉索斯
	f.萨加拉索斯Sagalassos.SetParent(f)
	
	f.西德Side = BSide西德
	f.西德Side.SetParent(f)
	
	f.希梁Sillyon = BSillyon希梁
	f.希梁Sillyon.SetParent(f)
	
	f.斯莱治Slege = BSlege斯莱治
	f.斯莱治Slege.SetParent(f)
	
}
