package padova

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type PadovaCounty interface {
    feud.County
    BAdria阿德里亚() 	feud.Barony
    BChioggia基奥贾() 	feud.Barony
    BEste埃斯泰() 	feud.Barony
    BMontagnana蒙塔尼亚纳() 	feud.Barony
    BPadova帕多瓦() 	feud.Barony
    BPolesine波莱西内() 	feud.Barony
    BRovigo罗维戈() 	feud.Barony
    BVigonza维贡扎() 	feud.Barony
}

type 帕多维亚PadovaCounty struct {
	feud.BaseCounty
	阿德里亚Adria 	feud.Barony
	基奥贾Chioggia 	feud.Barony
	埃斯泰Este 	feud.Barony
	蒙塔尼亚纳Montagnana 	feud.Barony
	帕多瓦Padova 	feud.Barony
	波莱西内Polesine 	feud.Barony
	罗维戈Rovigo 	feud.Barony
	维贡扎Vigonza 	feud.Barony
}

func (c *帕多维亚PadovaCounty) BAdria阿德里亚() feud.Barony {
	return c.阿德里亚Adria
}
    
func (c *帕多维亚PadovaCounty) BChioggia基奥贾() feud.Barony {
	return c.基奥贾Chioggia
}
    
func (c *帕多维亚PadovaCounty) BEste埃斯泰() feud.Barony {
	return c.埃斯泰Este
}
    
func (c *帕多维亚PadovaCounty) BMontagnana蒙塔尼亚纳() feud.Barony {
	return c.蒙塔尼亚纳Montagnana
}
    
func (c *帕多维亚PadovaCounty) BPadova帕多瓦() feud.Barony {
	return c.帕多瓦Padova
}
    
func (c *帕多维亚PadovaCounty) BPolesine波莱西内() feud.Barony {
	return c.波莱西内Polesine
}
    
func (c *帕多维亚PadovaCounty) BRovigo罗维戈() feud.Barony {
	return c.罗维戈Rovigo
}
    
func (c *帕多维亚PadovaCounty) BVigonza维贡扎() feud.Barony {
	return c.维贡扎Vigonza
}
    
var CPadova帕多维亚 PadovaCounty = &帕多维亚PadovaCounty{}

func init() {
	f := CPadova帕多维亚.(*帕多维亚PadovaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "355",
		Title:     "padova",
		TitleName: "帕多维亚",
		TitleCode: "c_padova",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿德里亚Adria = BAdria阿德里亚
	f.阿德里亚Adria.SetParent(f)
	
	f.基奥贾Chioggia = BChioggia基奥贾
	f.基奥贾Chioggia.SetParent(f)
	
	f.埃斯泰Este = BEste埃斯泰
	f.埃斯泰Este.SetParent(f)
	
	f.蒙塔尼亚纳Montagnana = BMontagnana蒙塔尼亚纳
	f.蒙塔尼亚纳Montagnana.SetParent(f)
	
	f.帕多瓦Padova = BPadova帕多瓦
	f.帕多瓦Padova.SetParent(f)
	
	f.波莱西内Polesine = BPolesine波莱西内
	f.波莱西内Polesine.SetParent(f)
	
	f.罗维戈Rovigo = BRovigo罗维戈
	f.罗维戈Rovigo.SetParent(f)
	
	f.维贡扎Vigonza = BVigonza维贡扎
	f.维贡扎Vigonza.SetParent(f)
	
}
