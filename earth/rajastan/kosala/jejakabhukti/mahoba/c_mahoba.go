package mahoba

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type MahobaCounty interface {
    feud.County
    BAjaigarh阿阇耶姞利呬() 	feud.Barony
    BChitrakuta质多罗矩吒() 	feud.Barony
    BHamirpur诃弥罗城() 	feud.Barony
    BJhammanpur旬曼补() 	feud.Barony
    BKamodh伽牟陀() 	feud.Barony
    BMahoba摩呼婆() 	feud.Barony
    BRath罗提() 	feud.Barony
}

type 摩呼婆MahobaCounty struct {
	feud.BaseCounty
	阿阇耶姞利呬Ajaigarh 	feud.Barony
	质多罗矩吒Chitrakuta 	feud.Barony
	诃弥罗城Hamirpur 	feud.Barony
	旬曼补Jhammanpur 	feud.Barony
	伽牟陀Kamodh 	feud.Barony
	摩呼婆Mahoba 	feud.Barony
	罗提Rath 	feud.Barony
}

func (c *摩呼婆MahobaCounty) BAjaigarh阿阇耶姞利呬() feud.Barony {
	return c.阿阇耶姞利呬Ajaigarh
}
    
func (c *摩呼婆MahobaCounty) BChitrakuta质多罗矩吒() feud.Barony {
	return c.质多罗矩吒Chitrakuta
}
    
func (c *摩呼婆MahobaCounty) BHamirpur诃弥罗城() feud.Barony {
	return c.诃弥罗城Hamirpur
}
    
func (c *摩呼婆MahobaCounty) BJhammanpur旬曼补() feud.Barony {
	return c.旬曼补Jhammanpur
}
    
func (c *摩呼婆MahobaCounty) BKamodh伽牟陀() feud.Barony {
	return c.伽牟陀Kamodh
}
    
func (c *摩呼婆MahobaCounty) BMahoba摩呼婆() feud.Barony {
	return c.摩呼婆Mahoba
}
    
func (c *摩呼婆MahobaCounty) BRath罗提() feud.Barony {
	return c.罗提Rath
}
    
var CMahoba摩呼婆 MahobaCounty = &摩呼婆MahobaCounty{}

func init() {
	f := CMahoba摩呼婆.(*摩呼婆MahobaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1301",
		Title:     "mahoba",
		TitleName: "摩呼婆",
		TitleCode: "c_mahoba",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿阇耶姞利呬Ajaigarh = BAjaigarh阿阇耶姞利呬
	f.阿阇耶姞利呬Ajaigarh.SetParent(f)
	
	f.质多罗矩吒Chitrakuta = BChitrakuta质多罗矩吒
	f.质多罗矩吒Chitrakuta.SetParent(f)
	
	f.诃弥罗城Hamirpur = BHamirpur诃弥罗城
	f.诃弥罗城Hamirpur.SetParent(f)
	
	f.旬曼补Jhammanpur = BJhammanpur旬曼补
	f.旬曼补Jhammanpur.SetParent(f)
	
	f.伽牟陀Kamodh = BKamodh伽牟陀
	f.伽牟陀Kamodh.SetParent(f)
	
	f.摩呼婆Mahoba = BMahoba摩呼婆
	f.摩呼婆Mahoba.SetParent(f)
	
	f.罗提Rath = BRath罗提
	f.罗提Rath.SetParent(f)
	
}
