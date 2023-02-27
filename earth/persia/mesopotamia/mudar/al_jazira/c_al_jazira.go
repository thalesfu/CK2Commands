package al_jazira

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type Al_jaziraCounty interface {
    feud.County
    BAmuda阿穆达() 	feud.Barony
    BDakhiliyah代希利耶() 	feud.Barony
    BDarbasiyah达尔巴斯亚() 	feud.Barony
    BDayrik代里克() 	feud.Barony
    BHamoukar哈默卡尔() 	feud.Barony
    BHasakah哈塞克() 	feud.Barony
    BQamishhlo卡米什利() 	feud.Barony
    BTeltamer特达梅尔() 	feud.Barony
}

type 哈塞克Al_jaziraCounty struct {
	feud.BaseCounty
	阿穆达Amuda 	feud.Barony
	代希利耶Dakhiliyah 	feud.Barony
	达尔巴斯亚Darbasiyah 	feud.Barony
	代里克Dayrik 	feud.Barony
	哈默卡尔Hamoukar 	feud.Barony
	哈塞克Hasakah 	feud.Barony
	卡米什利Qamishhlo 	feud.Barony
	特达梅尔Teltamer 	feud.Barony
}

func (c *哈塞克Al_jaziraCounty) BAmuda阿穆达() feud.Barony {
	return c.阿穆达Amuda
}
    
func (c *哈塞克Al_jaziraCounty) BDakhiliyah代希利耶() feud.Barony {
	return c.代希利耶Dakhiliyah
}
    
func (c *哈塞克Al_jaziraCounty) BDarbasiyah达尔巴斯亚() feud.Barony {
	return c.达尔巴斯亚Darbasiyah
}
    
func (c *哈塞克Al_jaziraCounty) BDayrik代里克() feud.Barony {
	return c.代里克Dayrik
}
    
func (c *哈塞克Al_jaziraCounty) BHamoukar哈默卡尔() feud.Barony {
	return c.哈默卡尔Hamoukar
}
    
func (c *哈塞克Al_jaziraCounty) BHasakah哈塞克() feud.Barony {
	return c.哈塞克Hasakah
}
    
func (c *哈塞克Al_jaziraCounty) BQamishhlo卡米什利() feud.Barony {
	return c.卡米什利Qamishhlo
}
    
func (c *哈塞克Al_jaziraCounty) BTeltamer特达梅尔() feud.Barony {
	return c.特达梅尔Teltamer
}
    
var CAl_jazira哈塞克 Al_jaziraCounty = &哈塞克Al_jaziraCounty{}

func init() {
	f := CAl_jazira哈塞克.(*哈塞克Al_jaziraCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "698",
		Title:     "al_jazira",
		TitleName: "哈塞克",
		TitleCode: "c_al_jazira",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿穆达Amuda = BAmuda阿穆达
	f.阿穆达Amuda.SetParent(f)
	
	f.代希利耶Dakhiliyah = BDakhiliyah代希利耶
	f.代希利耶Dakhiliyah.SetParent(f)
	
	f.达尔巴斯亚Darbasiyah = BDarbasiyah达尔巴斯亚
	f.达尔巴斯亚Darbasiyah.SetParent(f)
	
	f.代里克Dayrik = BDayrik代里克
	f.代里克Dayrik.SetParent(f)
	
	f.哈默卡尔Hamoukar = BHamoukar哈默卡尔
	f.哈默卡尔Hamoukar.SetParent(f)
	
	f.哈塞克Hasakah = BHasakah哈塞克
	f.哈塞克Hasakah.SetParent(f)
	
	f.卡米什利Qamishhlo = BQamishhlo卡米什利
	f.卡米什利Qamishhlo.SetParent(f)
	
	f.特达梅尔Teltamer = BTeltamer特达梅尔
	f.特达梅尔Teltamer.SetParent(f)
	
}
