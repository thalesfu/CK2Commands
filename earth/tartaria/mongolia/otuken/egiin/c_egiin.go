package egiin

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type EgiinCounty interface {
    feud.County
    BEgiin额金() 	feud.Barony
    BHatgal哈特嘎勒() 	feud.Barony
    BMandal曼达勒() 	feud.Barony
    BMurun木伦() 	feud.Barony
    BSaridag萨尔德克() 	feud.Barony
    BTugul图古勒() 	feud.Barony
    BZuulun卓伦() 	feud.Barony
}

type 额金EgiinCounty struct {
	feud.BaseCounty
	额金Egiin 	feud.Barony
	哈特嘎勒Hatgal 	feud.Barony
	曼达勒Mandal 	feud.Barony
	木伦Murun 	feud.Barony
	萨尔德克Saridag 	feud.Barony
	图古勒Tugul 	feud.Barony
	卓伦Zuulun 	feud.Barony
}

func (c *额金EgiinCounty) BEgiin额金() feud.Barony {
	return c.额金Egiin
}
    
func (c *额金EgiinCounty) BHatgal哈特嘎勒() feud.Barony {
	return c.哈特嘎勒Hatgal
}
    
func (c *额金EgiinCounty) BMandal曼达勒() feud.Barony {
	return c.曼达勒Mandal
}
    
func (c *额金EgiinCounty) BMurun木伦() feud.Barony {
	return c.木伦Murun
}
    
func (c *额金EgiinCounty) BSaridag萨尔德克() feud.Barony {
	return c.萨尔德克Saridag
}
    
func (c *额金EgiinCounty) BTugul图古勒() feud.Barony {
	return c.图古勒Tugul
}
    
func (c *额金EgiinCounty) BZuulun卓伦() feud.Barony {
	return c.卓伦Zuulun
}
    
var CEgiin额金 EgiinCounty = &额金EgiinCounty{}

func init() {
	f := CEgiin额金.(*额金EgiinCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1904",
		Title:     "egiin",
		TitleName: "额金",
		TitleCode: "c_egiin",
		Baronies:  map[string]feud.Barony{},
	}

	f.额金Egiin = BEgiin额金
	f.额金Egiin.SetParent(f)
	
	f.哈特嘎勒Hatgal = BHatgal哈特嘎勒
	f.哈特嘎勒Hatgal.SetParent(f)
	
	f.曼达勒Mandal = BMandal曼达勒
	f.曼达勒Mandal.SetParent(f)
	
	f.木伦Murun = BMurun木伦
	f.木伦Murun.SetParent(f)
	
	f.萨尔德克Saridag = BSaridag萨尔德克
	f.萨尔德克Saridag.SetParent(f)
	
	f.图古勒Tugul = BTugul图古勒
	f.图古勒Tugul.SetParent(f)
	
	f.卓伦Zuulun = BZuulun卓伦
	f.卓伦Zuulun.SetParent(f)
	
}
