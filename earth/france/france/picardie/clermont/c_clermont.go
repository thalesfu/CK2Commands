package clermont

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type ClermontCounty interface {
    feud.County
    BBeauvais博韦() 	feud.Barony
    BChelles谢勒() 	feud.Barony
    BGisors日索尔() 	feud.Barony
    BNoailles诺阿耶() 	feud.Barony
    BNogent_sur_oise诺让() 	feud.Barony
    BTille蒂耶() 	feud.Barony
}

type 克莱蒙ClermontCounty struct {
	feud.BaseCounty
	博韦Beauvais 	feud.Barony
	谢勒Chelles 	feud.Barony
	日索尔Gisors 	feud.Barony
	诺阿耶Noailles 	feud.Barony
	诺让Nogent_sur_oise 	feud.Barony
	蒂耶Tille 	feud.Barony
}

func (c *克莱蒙ClermontCounty) BBeauvais博韦() feud.Barony {
	return c.博韦Beauvais
}
    
func (c *克莱蒙ClermontCounty) BChelles谢勒() feud.Barony {
	return c.谢勒Chelles
}
    
func (c *克莱蒙ClermontCounty) BGisors日索尔() feud.Barony {
	return c.日索尔Gisors
}
    
func (c *克莱蒙ClermontCounty) BNoailles诺阿耶() feud.Barony {
	return c.诺阿耶Noailles
}
    
func (c *克莱蒙ClermontCounty) BNogent_sur_oise诺让() feud.Barony {
	return c.诺让Nogent_sur_oise
}
    
func (c *克莱蒙ClermontCounty) BTille蒂耶() feud.Barony {
	return c.蒂耶Tille
}
    
var CClermont克莱蒙 ClermontCounty = &克莱蒙ClermontCounty{}

func init() {
	f := CClermont克莱蒙.(*克莱蒙ClermontCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1961",
		Title:     "clermont",
		TitleName: "克莱蒙",
		TitleCode: "c_clermont",
		Baronies:  map[string]feud.Barony{},
	}

	f.博韦Beauvais = BBeauvais博韦
	f.博韦Beauvais.SetParent(f)
	
	f.谢勒Chelles = BChelles谢勒
	f.谢勒Chelles.SetParent(f)
	
	f.日索尔Gisors = BGisors日索尔
	f.日索尔Gisors.SetParent(f)
	
	f.诺阿耶Noailles = BNoailles诺阿耶
	f.诺阿耶Noailles.SetParent(f)
	
	f.诺让Nogent_sur_oise = BNogent_sur_oise诺让
	f.诺让Nogent_sur_oise.SetParent(f)
	
	f.蒂耶Tille = BTille蒂耶
	f.蒂耶Tille.SetParent(f)
	
}
