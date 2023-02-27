package roma

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type RomaCounty interface {
    feud.County
    BAragni阿纳尼() 	feud.Barony
    BOstia奥斯蒂亚() 	feud.Barony
    BRoma罗马() 	feud.Barony
    BSutri苏特里() 	feud.Barony
    BTerracina泰拉奇纳() 	feud.Barony
    BTivoli蒂沃利() 	feud.Barony
    BTusculum图斯库卢姆() 	feud.Barony
    BViterbo维泰博() 	feud.Barony
}

type 罗马RomaCounty struct {
	feud.BaseCounty
	阿纳尼Aragni 	feud.Barony
	奥斯蒂亚Ostia 	feud.Barony
	罗马Roma 	feud.Barony
	苏特里Sutri 	feud.Barony
	泰拉奇纳Terracina 	feud.Barony
	蒂沃利Tivoli 	feud.Barony
	图斯库卢姆Tusculum 	feud.Barony
	维泰博Viterbo 	feud.Barony
}

func (c *罗马RomaCounty) BAragni阿纳尼() feud.Barony {
	return c.阿纳尼Aragni
}
    
func (c *罗马RomaCounty) BOstia奥斯蒂亚() feud.Barony {
	return c.奥斯蒂亚Ostia
}
    
func (c *罗马RomaCounty) BRoma罗马() feud.Barony {
	return c.罗马Roma
}
    
func (c *罗马RomaCounty) BSutri苏特里() feud.Barony {
	return c.苏特里Sutri
}
    
func (c *罗马RomaCounty) BTerracina泰拉奇纳() feud.Barony {
	return c.泰拉奇纳Terracina
}
    
func (c *罗马RomaCounty) BTivoli蒂沃利() feud.Barony {
	return c.蒂沃利Tivoli
}
    
func (c *罗马RomaCounty) BTusculum图斯库卢姆() feud.Barony {
	return c.图斯库卢姆Tusculum
}
    
func (c *罗马RomaCounty) BViterbo维泰博() feud.Barony {
	return c.维泰博Viterbo
}
    
var CRoma罗马 RomaCounty = &罗马RomaCounty{}

func init() {
	f := CRoma罗马.(*罗马RomaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "333",
		Title:     "roma",
		TitleName: "罗马",
		TitleCode: "c_roma",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿纳尼Aragni = BAragni阿纳尼
	f.阿纳尼Aragni.SetParent(f)
	
	f.奥斯蒂亚Ostia = BOstia奥斯蒂亚
	f.奥斯蒂亚Ostia.SetParent(f)
	
	f.罗马Roma = BRoma罗马
	f.罗马Roma.SetParent(f)
	
	f.苏特里Sutri = BSutri苏特里
	f.苏特里Sutri.SetParent(f)
	
	f.泰拉奇纳Terracina = BTerracina泰拉奇纳
	f.泰拉奇纳Terracina.SetParent(f)
	
	f.蒂沃利Tivoli = BTivoli蒂沃利
	f.蒂沃利Tivoli.SetParent(f)
	
	f.图斯库卢姆Tusculum = BTusculum图斯库卢姆
	f.图斯库卢姆Tusculum.SetParent(f)
	
	f.维泰博Viterbo = BViterbo维泰博
	f.维泰博Viterbo.SetParent(f)
	
}
