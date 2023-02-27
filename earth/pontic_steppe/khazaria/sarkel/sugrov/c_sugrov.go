package sugrov

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type SugrovCounty interface {
    feud.County
    BFatezh法捷日() 	feud.Barony
    BKhazar可萨() 	feud.Barony
    BKhratayak赫拉塔亚克() 	feud.Barony
    BKhursa胡尔萨() 	feud.Barony
    BOboyan奥博扬() 	feud.Barony
    BSugrov苏格罗夫() 	feud.Barony
    BTim托木() 	feud.Barony
    BYauchy雅乌赤() 	feud.Barony
}

type 苏格罗夫SugrovCounty struct {
	feud.BaseCounty
	法捷日Fatezh 	feud.Barony
	可萨Khazar 	feud.Barony
	赫拉塔亚克Khratayak 	feud.Barony
	胡尔萨Khursa 	feud.Barony
	奥博扬Oboyan 	feud.Barony
	苏格罗夫Sugrov 	feud.Barony
	托木Tim 	feud.Barony
	雅乌赤Yauchy 	feud.Barony
}

func (c *苏格罗夫SugrovCounty) BFatezh法捷日() feud.Barony {
	return c.法捷日Fatezh
}
    
func (c *苏格罗夫SugrovCounty) BKhazar可萨() feud.Barony {
	return c.可萨Khazar
}
    
func (c *苏格罗夫SugrovCounty) BKhratayak赫拉塔亚克() feud.Barony {
	return c.赫拉塔亚克Khratayak
}
    
func (c *苏格罗夫SugrovCounty) BKhursa胡尔萨() feud.Barony {
	return c.胡尔萨Khursa
}
    
func (c *苏格罗夫SugrovCounty) BOboyan奥博扬() feud.Barony {
	return c.奥博扬Oboyan
}
    
func (c *苏格罗夫SugrovCounty) BSugrov苏格罗夫() feud.Barony {
	return c.苏格罗夫Sugrov
}
    
func (c *苏格罗夫SugrovCounty) BTim托木() feud.Barony {
	return c.托木Tim
}
    
func (c *苏格罗夫SugrovCounty) BYauchy雅乌赤() feud.Barony {
	return c.雅乌赤Yauchy
}
    
var CSugrov苏格罗夫 SugrovCounty = &苏格罗夫SugrovCounty{}

func init() {
	f := CSugrov苏格罗夫.(*苏格罗夫SugrovCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "566",
		Title:     "sugrov",
		TitleName: "苏格罗夫",
		TitleCode: "c_sugrov",
		Baronies:  map[string]feud.Barony{},
	}

	f.法捷日Fatezh = BFatezh法捷日
	f.法捷日Fatezh.SetParent(f)
	
	f.可萨Khazar = BKhazar可萨
	f.可萨Khazar.SetParent(f)
	
	f.赫拉塔亚克Khratayak = BKhratayak赫拉塔亚克
	f.赫拉塔亚克Khratayak.SetParent(f)
	
	f.胡尔萨Khursa = BKhursa胡尔萨
	f.胡尔萨Khursa.SetParent(f)
	
	f.奥博扬Oboyan = BOboyan奥博扬
	f.奥博扬Oboyan.SetParent(f)
	
	f.苏格罗夫Sugrov = BSugrov苏格罗夫
	f.苏格罗夫Sugrov.SetParent(f)
	
	f.托木Tim = BTim托木
	f.托木Tim.SetParent(f)
	
	f.雅乌赤Yauchy = BYauchy雅乌赤
	f.雅乌赤Yauchy.SetParent(f)
	
}
