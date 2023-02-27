package nubia

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type NubiaCounty interface {
    feud.County
    BAnag阿纳格() 	feud.Barony
    BEirayab艾拉雅布() 	feud.Barony
    BFanoidig法努瓦迪格() 	feud.Barony
    BKomotit科莫蒂特() 	feud.Barony
    BNafab拿法布() 	feud.Barony
    BNubia努比亚() 	feud.Barony
    BSalala塞拉莱() 	feud.Barony
    BTahtani塔赫塔尼() 	feud.Barony
}

type 努比亚NubiaCounty struct {
	feud.BaseCounty
	阿纳格Anag 	feud.Barony
	艾拉雅布Eirayab 	feud.Barony
	法努瓦迪格Fanoidig 	feud.Barony
	科莫蒂特Komotit 	feud.Barony
	拿法布Nafab 	feud.Barony
	努比亚Nubia 	feud.Barony
	塞拉莱Salala 	feud.Barony
	塔赫塔尼Tahtani 	feud.Barony
}

func (c *努比亚NubiaCounty) BAnag阿纳格() feud.Barony {
	return c.阿纳格Anag
}
    
func (c *努比亚NubiaCounty) BEirayab艾拉雅布() feud.Barony {
	return c.艾拉雅布Eirayab
}
    
func (c *努比亚NubiaCounty) BFanoidig法努瓦迪格() feud.Barony {
	return c.法努瓦迪格Fanoidig
}
    
func (c *努比亚NubiaCounty) BKomotit科莫蒂特() feud.Barony {
	return c.科莫蒂特Komotit
}
    
func (c *努比亚NubiaCounty) BNafab拿法布() feud.Barony {
	return c.拿法布Nafab
}
    
func (c *努比亚NubiaCounty) BNubia努比亚() feud.Barony {
	return c.努比亚Nubia
}
    
func (c *努比亚NubiaCounty) BSalala塞拉莱() feud.Barony {
	return c.塞拉莱Salala
}
    
func (c *努比亚NubiaCounty) BTahtani塔赫塔尼() feud.Barony {
	return c.塔赫塔尼Tahtani
}
    
var CNubia努比亚 NubiaCounty = &努比亚NubiaCounty{}

func init() {
	f := CNubia努比亚.(*努比亚NubiaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "792",
		Title:     "nubia",
		TitleName: "努比亚",
		TitleCode: "c_nubia",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿纳格Anag = BAnag阿纳格
	f.阿纳格Anag.SetParent(f)
	
	f.艾拉雅布Eirayab = BEirayab艾拉雅布
	f.艾拉雅布Eirayab.SetParent(f)
	
	f.法努瓦迪格Fanoidig = BFanoidig法努瓦迪格
	f.法努瓦迪格Fanoidig.SetParent(f)
	
	f.科莫蒂特Komotit = BKomotit科莫蒂特
	f.科莫蒂特Komotit.SetParent(f)
	
	f.拿法布Nafab = BNafab拿法布
	f.拿法布Nafab.SetParent(f)
	
	f.努比亚Nubia = BNubia努比亚
	f.努比亚Nubia.SetParent(f)
	
	f.塞拉莱Salala = BSalala塞拉莱
	f.塞拉莱Salala.SetParent(f)
	
	f.塔赫塔尼Tahtani = BTahtani塔赫塔尼
	f.塔赫塔尼Tahtani.SetParent(f)
	
}
