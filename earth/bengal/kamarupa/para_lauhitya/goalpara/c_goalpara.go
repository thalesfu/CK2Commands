package goalpara

import (
	"github.com/thalesfu/paradoxtools/CK2/feud"
)

type GoalparaCounty interface {
    feud.County
    BAmauli阿旄梨() 	feud.Barony
    BAmla阿姆拉() 	feud.Barony
    BDhubri菟婆利() 	feud.Barony
    BDigha提格诃() 	feud.Barony
    BGoalpara乔阿罗波罗() 	feud.Barony
    BKamakhya迦摩佉耶() 	feud.Barony
    BSri_surya_pahar室利苏利耶钵罗诃罗() 	feud.Barony
}

type 乔阿罗波罗GoalparaCounty struct {
	feud.BaseCounty
	阿旄梨Amauli 	feud.Barony
	阿姆拉Amla 	feud.Barony
	菟婆利Dhubri 	feud.Barony
	提格诃Digha 	feud.Barony
	乔阿罗波罗Goalpara 	feud.Barony
	迦摩佉耶Kamakhya 	feud.Barony
	室利苏利耶钵罗诃罗Sri_surya_pahar 	feud.Barony
}

func (c *乔阿罗波罗GoalparaCounty) BAmauli阿旄梨() feud.Barony {
	return c.阿旄梨Amauli
}
    
func (c *乔阿罗波罗GoalparaCounty) BAmla阿姆拉() feud.Barony {
	return c.阿姆拉Amla
}
    
func (c *乔阿罗波罗GoalparaCounty) BDhubri菟婆利() feud.Barony {
	return c.菟婆利Dhubri
}
    
func (c *乔阿罗波罗GoalparaCounty) BDigha提格诃() feud.Barony {
	return c.提格诃Digha
}
    
func (c *乔阿罗波罗GoalparaCounty) BGoalpara乔阿罗波罗() feud.Barony {
	return c.乔阿罗波罗Goalpara
}
    
func (c *乔阿罗波罗GoalparaCounty) BKamakhya迦摩佉耶() feud.Barony {
	return c.迦摩佉耶Kamakhya
}
    
func (c *乔阿罗波罗GoalparaCounty) BSri_surya_pahar室利苏利耶钵罗诃罗() feud.Barony {
	return c.室利苏利耶钵罗诃罗Sri_surya_pahar
}
    
var CGoalpara乔阿罗波罗 GoalparaCounty = &乔阿罗波罗GoalparaCounty{}

func init() {
	f := CGoalpara乔阿罗波罗.(*乔阿罗波罗GoalparaCounty)
	f.BaseCounty = feud.BaseCounty{
		ID:        "1246",
		Title:     "goalpara",
		TitleName: "乔阿罗波罗",
		TitleCode: "c_goalpara",
		Baronies:  map[string]feud.Barony{},
	}

	f.阿旄梨Amauli = BAmauli阿旄梨
	f.阿旄梨Amauli.SetParent(f)
	
	f.阿姆拉Amla = BAmla阿姆拉
	f.阿姆拉Amla.SetParent(f)
	
	f.菟婆利Dhubri = BDhubri菟婆利
	f.菟婆利Dhubri.SetParent(f)
	
	f.提格诃Digha = BDigha提格诃
	f.提格诃Digha.SetParent(f)
	
	f.乔阿罗波罗Goalpara = BGoalpara乔阿罗波罗
	f.乔阿罗波罗Goalpara.SetParent(f)
	
	f.迦摩佉耶Kamakhya = BKamakhya迦摩佉耶
	f.迦摩佉耶Kamakhya.SetParent(f)
	
	f.室利苏利耶钵罗诃罗Sri_surya_pahar = BSri_surya_pahar室利苏利耶钵罗诃罗
	f.室利苏利耶钵罗诃罗Sri_surya_pahar.SetParent(f)
	
}
