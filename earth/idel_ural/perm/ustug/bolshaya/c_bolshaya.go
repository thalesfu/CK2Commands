package bolshaya

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type BolshayaCounty interface {
    feud.County
    BBolshaya博利沙亚() 	feud.Barony
    BLena_bolshaya列那() 	feud.Barony
    BLitvino利特维诺() 	feud.Barony
    BSorovo索罗沃() 	feud.Barony
    BSoyga索伊加() 	feud.Barony
    BTyva特瓦() 	feud.Barony
    BUrdoma乌尔多马() 	feud.Barony
}

type 维列季BolshayaCounty struct {
	feud.BaseCounty
	博利沙亚Bolshaya 	feud.Barony
	列那Lena_bolshaya 	feud.Barony
	利特维诺Litvino 	feud.Barony
	索罗沃Sorovo 	feud.Barony
	索伊加Soyga 	feud.Barony
	特瓦Tyva 	feud.Barony
	乌尔多马Urdoma 	feud.Barony
}

func (c *维列季BolshayaCounty) BBolshaya博利沙亚() feud.Barony {
	return c.博利沙亚Bolshaya
}
    
func (c *维列季BolshayaCounty) BLena_bolshaya列那() feud.Barony {
	return c.列那Lena_bolshaya
}
    
func (c *维列季BolshayaCounty) BLitvino利特维诺() feud.Barony {
	return c.利特维诺Litvino
}
    
func (c *维列季BolshayaCounty) BSorovo索罗沃() feud.Barony {
	return c.索罗沃Sorovo
}
    
func (c *维列季BolshayaCounty) BSoyga索伊加() feud.Barony {
	return c.索伊加Soyga
}
    
func (c *维列季BolshayaCounty) BTyva特瓦() feud.Barony {
	return c.特瓦Tyva
}
    
func (c *维列季BolshayaCounty) BUrdoma乌尔多马() feud.Barony {
	return c.乌尔多马Urdoma
}
    
var CBolshaya维列季 BolshayaCounty = &维列季BolshayaCounty{}

func init() {
	f := CBolshaya维列季.(*维列季BolshayaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1818",
		Title:     "bolshaya",
		TitleName: "维列季",
		TitleCode: "c_bolshaya",
		Baronies:  map[string]feud.Barony{},
	}

	f.博利沙亚Bolshaya = BBolshaya博利沙亚
	f.博利沙亚Bolshaya.SetParent(f)
	
	f.列那Lena_bolshaya = BLena_bolshaya列那
	f.列那Lena_bolshaya.SetParent(f)
	
	f.利特维诺Litvino = BLitvino利特维诺
	f.利特维诺Litvino.SetParent(f)
	
	f.索罗沃Sorovo = BSorovo索罗沃
	f.索罗沃Sorovo.SetParent(f)
	
	f.索伊加Soyga = BSoyga索伊加
	f.索伊加Soyga.SetParent(f)
	
	f.特瓦Tyva = BTyva特瓦
	f.特瓦Tyva.SetParent(f)
	
	f.乌尔多马Urdoma = BUrdoma乌尔多马
	f.乌尔多马Urdoma.SetParent(f)
	
}
